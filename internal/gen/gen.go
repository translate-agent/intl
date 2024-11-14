package main

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"iter"
	"log/slog"
	"maps"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"
)

//go:embed cldr.go.tmpl
var datetimeTemplate string

type Generator struct {
	cldr *cldr.CLDR
}

type Conf struct {
	cldrDir    string
	out        string
	saveMerged bool
}

func Gen(conf Conf) error {
	var g Generator

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	if err := g.Load(conf.cldrDir, log); err != nil {
		return err
	}

	if conf.saveMerged {
		if err := g.saveMerged(conf.out); err != nil {
			return err
		}
	}

	if err := g.Write(conf.out, log); err != nil {
		return err
	}

	return nil
}

func (g *Generator) Load(dir string, log *slog.Logger) error {
	var (
		d   cldr.Decoder
		err error
	)

	d.SetDirFilter("main", "supplemental")

	g.cldr, err = d.DecodePath(dir)
	if err != nil {
		return fmt.Errorf(`decode CLDR at path "%s": %w`, dir, err)
	}

	g.filterApproved()
	g.merge(log)

	return nil
}

func (g *Generator) Write(out string, log *slog.Logger) error {
	tpl, err := template.New("datetime").Funcs(template.FuncMap{
		"join":     strings.Join,
		"contains": strings.Contains,
		"title":    title,
		"sub":      func(a, b int) int { return a - b },
	}).Parse(datetimeTemplate)
	if err != nil {
		return fmt.Errorf("parse datetime template: %w", err)
	}

	defaultNumberingSystems := g.defaultNumberingSystems()
	calendarPreferences := g.calendarPreferences()

	data := TemplateData{
		CalendarPreferences:     calendarPreferences,
		DateTimeFormats:         g.dateTimeFormats(calendarPreferences, log),
		NumberingSystems:        g.numberingSystems(defaultNumberingSystems),
		NumberingSystemIota:     g.numberingSystemsIota(defaultNumberingSystems),
		DefaultNumberingSystems: defaultNumberingSystems,
		Months:                  g.months(),
	}

	name := path.Join(out, "cldr.go")

	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create %s: %w", name, err)
	}

	defer f.Close()

	if err := tpl.Execute(f, data); err != nil {
		return fmt.Errorf("execute datetime template: %w", err)
	}

	return nil
}

func (g *Generator) saveMerged(out string) error {
	for _, locale := range g.cldr.Locales() {
		name := path.Join(out, ".cldr_merged", locale+".xml")

		f, err := os.Create(name)
		if err != nil {
			return fmt.Errorf("create %s: %w", name, err)
		}

		defer f.Close()

		ldml := g.cldr.RawLDML(locale)
		enc := xml.NewEncoder(f)
		enc.Indent("", "\t")

		if err = enc.Encode(ldml); err != nil {
			return fmt.Errorf("save to %s: %w", name, err)
		}
	}

	return nil
}

//nolint:gocognit
func (g *Generator) filterApproved() {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Numbers != nil {
			var defaultNumberingSystem []*cldr.Common

			for _, v := range ldml.Numbers.DefaultNumberingSystem {
				if isContributedOrApproved(v.Draft) {
					defaultNumberingSystem = append(defaultNumberingSystem, v)
				}
			}

			ldml.Numbers.DefaultNumberingSystem = defaultNumberingSystem
		}

		if ldml.Dates != nil && ldml.Dates.Calendars != nil {
			for _, calendar := range ldml.Dates.Calendars.Calendar {
				if calendar.Months == nil {
					continue
				}

				// calendar.Months.MonthContext.MonthWidth.Month
				for _, monthContext := range calendar.Months.MonthContext {
					for _, monthWidth := range monthContext.MonthWidth {
						var months []*struct {
							cldr.Common
							Yeartype string `xml:"yeartype,attr"`
						}

						for _, month := range monthWidth.Month {
							if isContributedOrApproved(month.Draft) {
								months = append(months, month)
							}
						}

						monthWidth.Month = months
					}
				}

				// calendar.DateTimeFormats.AvailableFormats
				if calendar.DateTimeFormats == nil {
					continue
				}

				for _, dateTimeFormat := range calendar.DateTimeFormats.AvailableFormats {
					var dateFormatItems []*struct {
						cldr.Common
						Id    string `xml:"id,attr"` //nolint:revive,stylecheck
						Count string `xml:"count,attr"`
					}

					for _, dateFormatItem := range dateTimeFormat.DateFormatItem {
						if isContributedOrApproved(dateFormatItem.Draft) {
							dateFormatItems = append(dateFormatItems, dateFormatItem)
						}
					}

					dateTimeFormat.DateFormatItem = dateFormatItems
				}
			}
		}
	}
}

func (g *Generator) merge(log *slog.Logger) {
	g.mergeAliases()
	g.mergeLocal(log)

	root := g.cldr.RawLDML("root")

	// merge parent to child
	for _, parentLocale := range g.cldr.Supplemental().ParentLocales.ParentLocale {
		// ignore, cldr package does NOT have the attribute "component"
		// <parentLocales component="collations">
		// 	<parentLocale parent="sr_ME" locales="sr_Cyrl_ME"/>
		// 	<parentLocale parent="zh_Hant" locales="yue yue_Hant"/>
		// 	<parentLocale parent="zh_Hans" locales="yue_CN yue_Hans yue_Hans_CN"/>
		// </parentLocales>
		if slices.Contains([]string{"sr_ME", "zh_Hant", "zh_Hans"}, parentLocale.Parent) {
			continue
		}

		parent := g.cldr.RawLDML(parentLocale.Parent)

		// merge root to parent
		merge(parent, root, log)

		for _, locale := range strings.Split(parentLocale.Locales, " ") {
			child := g.cldr.RawLDML(locale)

			if child == nil {
				continue
			}

			merge(child, parent, log)
		}
	}

	// merge root to language

	for _, locale := range g.cldr.Locales() {
		if strings.ContainsRune(locale, '_') {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		merge(ldml, root, log)
	}

	// merge language to territory
	for _, locale := range g.cldr.Locales() {
		if !strings.ContainsRune(locale, '_') {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		parts := strings.Split(locale, "_")
		parts = parts[:len(parts)-1]

		fallback := g.cldr.RawLDML(strings.Join(parts, "_"))

		merge(ldml, fallback, log)
	}
}

func (g *Generator) mergeAliases() {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for i, calendar := range ldml.Dates.Calendars.Calendar {
			if calendar.Alias != nil {
				calendarType := strings.Split(calendar.Alias.Path, "'")[1]
				calendar = findCalendar(ldml, calendarType)
				ldml.Dates.Calendars.Calendar[i] = calendar

				continue
			}

			if calendar.DateTimeFormats == nil || calendar.DateTimeFormats.Alias == nil {
				continue
			}

			// example: ../../calendar[@type='generic']/dateTimeFormats
			calendarType := strings.Split(calendar.DateTimeFormats.Alias.Path, "'")[1]

			calendar.DateTimeFormats = findCalendar(ldml, calendarType).DateTimeFormats
		}
	}
}

func (g *Generator) mergeLocal(log *slog.Logger) {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Identity.Language.Type == "root" || ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		generic := findCalendar(ldml, "generic")

		if generic == nil || generic.DateTimeFormats == nil {
			continue
		}

		for _, calendar := range ldml.Dates.Calendars.Calendar {
			if !slices.Contains([]string{"persian", "buddhist"}, calendar.Type) {
				continue
			}

			if calendar.DateTimeFormats == nil {
				calendar.DateTimeFormats = deepCopy(generic.DateTimeFormats)
			}

			mergeCalendar(calendar, generic, log.With("locale", locale))
		}
	}
}

// findCalendar returns *cldr.Calendar by its type if found. Otherwise, returns nil.
func findCalendar(ldml *cldr.LDML, calendarType string) *cldr.Calendar {
	for _, v := range ldml.Dates.Calendars.Calendar {
		if v.Type == calendarType {
			return v
		}
	}

	return nil
}

// containsDateFormatItem returns true if calendar contains dateFormatItem with given id.
func containsDateFormatItem(calendar *cldr.Calendar, id string) bool {
	for _, v := range calendar.DateTimeFormats.AvailableFormats[0].DateFormatItem {
		if v.Id == id {
			return true
		}
	}

	return false
}

var supportedCalendarTypes = []string{"gregorian", "persian", "buddhist"}

func supportedCalendars(calendars []*cldr.Calendar) iter.Seq[*cldr.Calendar] {
	return func(yield func(*cldr.Calendar) bool) {
		for _, v := range calendars {
			if slices.Contains(supportedCalendarTypes, v.Type) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// merge copies particular src values to dst.
func merge(dst, src *cldr.LDML, log *slog.Logger) {
	if src.Dates == nil || src.Dates.Calendars == nil {
		return
	}

	if dst.Dates == nil {
		dst.Dates = deepCopy(src.Dates)
	}

	if dst.Dates.Calendars == nil {
		dst.Dates.Calendars = deepCopy(src.Dates.Calendars)
	}

	if len(dst.Dates.Calendars.Calendar) == 0 {
		dst.Dates.Calendars.Calendar = deepCopy(src.Dates.Calendars.Calendar)
	}

	for parentCalendar := range supportedCalendars(src.Dates.Calendars.Calendar) {
		calendar := findCalendar(dst, parentCalendar.Type)
		if calendar == nil {
			calendar = deepCopy(parentCalendar)
			dst.Dates.Calendars.Calendar = append(dst.Dates.Calendars.Calendar, calendar)

			continue
		}

		mergeCalendar(calendar, parentCalendar, log)
	}
}

func mergeCalendar(dst, src *cldr.Calendar, log *slog.Logger) {
	log.Debug("merge calendars", "dst", dst.Type, "src", src.Type)

	if dst.DateTimeFormats == nil {
		dst.DateTimeFormats = deepCopy(src.DateTimeFormats)
	} else {
		if dst.DateTimeFormats.AvailableFormats == nil && src.DateTimeFormats != nil {
			dst.DateTimeFormats.AvailableFormats = deepCopy(src.DateTimeFormats.AvailableFormats)
		}

		if src.DateTimeFormats != nil {
			for _, availableFormats := range src.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					if containsDateFormatItem(dst, dateFormatItem.Id) {
						continue
					}

					dst.DateTimeFormats.AvailableFormats[0].DateFormatItem = append(
						dst.DateTimeFormats.AvailableFormats[0].DateFormatItem,
						deepCopy(dateFormatItem))
				}
			}
		}
	}

	// months
	if src.Months == nil {
		return
	}

	if dst.Months == nil {
		dst.Months = deepCopy(src.Months)
	}

	if dst.Months != nil && len(dst.Months.MonthContext) == 0 && len(src.Months.MonthContext) > 0 {
		dst.Months.MonthContext = deepCopy(src.Months.MonthContext)
	}
}

func (g *Generator) calendarPreferences() []CalendarPreference {
	calendarPreferences := g.cldr.Supplemental().CalendarPreferenceData.CalendarPreference
	preferences := make([]CalendarPreference, 0, len(calendarPreferences))

	// calendar preferences
	for _, v := range calendarPreferences {
		preferences = append(preferences, CalendarPreference{
			Regions:   strings.Split(v.Territories, " "),
			Calendars: strings.Split(v.Ordering, " "),
		})
	}

	return preferences
}

func (g *Generator) defaultNumberingSystems() DefaultNumberingSystems {
	defaultNumberingSystems := make(DefaultNumberingSystems)

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Numbers == nil || locale == "root" {
			continue
		}

		for _, v := range ldml.Numbers.DefaultNumberingSystem {
			if v.Alt != "" {
				continue
			}

			defaultNumberingSystems[v.CharData] = append(defaultNumberingSystems[v.CharData],
				strings.ReplaceAll(locale, "_", "-"))
		}
	}

	return defaultNumberingSystems
}

func (g *Generator) dateTimeFormats(calendarPreferences []CalendarPreference, log *slog.Logger) DateTimeFormats {
	dateTimeFormats := make(DateTimeFormats, len(supportedCalendarTypes))

	for _, calendarType := range supportedCalendarTypes {
		formats := NewCalendarDateTimeFormats()

		formats.Y.Default = g.findDateFormatItem("root", calendarType, "y")
		formats.YM.Default = cmp.Or(
			g.findDateFormatItem("root", calendarType, "yM"),
			g.findDateFormatItem("root", calendarType, "yMM"),
			g.findDateFormatItem("root", calendarType, "yyyyM"),
		)
		formats.M.Default = g.findDateFormatItem("root", calendarType, "M")
		formats.MD.Default = BuildFmtMD(
			g.findDateFormatItem("root", calendarType, "Md"),
			g.findDateFormatItem("root", calendarType, "MMd"),
			g.findDateFormatItem("root", calendarType, "Mdd"),
			g.findDateFormatItem("root", calendarType, "MMd"),
			log,
		)
		formats.D.Default = g.findDateFormatItem("root", calendarType, "d")

		dateTimeFormats[calendarType] = formats
	}

	for _, locale := range g.cldr.Locales() {
		// Ignore duplicate formatting for "y".
		// Locales containing "_" have the same "y" formatting, skip them for now.
		if strings.Contains(locale, "_") || locale == "root" {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		region, _ := language.MustParse(locale).Region()

		i := slices.IndexFunc(calendarPreferences, func(v CalendarPreference) bool {
			return slices.Contains(v.Regions, region.String())
		})

		preferedCalendar := "gregorian"
		if i >= 0 {
			preferedCalendar = calendarPreferences[i].Calendars[0]
		}

		calendar := findCalendar(ldml, preferedCalendar)
		if calendar == nil {
			continue
		}

		formats := dateTimeFormats[calendar.Type]

		if calendar.DateTimeFormats == nil {
			continue
		}

		localeLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})).
			With("locale", locale)

		g.addFormatY(locale, calendar, formats, localeLog)
		g.addFormatYM(locale, calendar, formats, localeLog)
		g.addFormatM(locale, calendar, formats, localeLog)
		g.addFormatMD(locale, calendar, formats, localeLog)
		g.addFormatD(locale, calendar, formats, localeLog)
	}

	for calendarType, formats := range dateTimeFormats {
		formats.Y.Default = strings.NewReplacer("G ", `"AP "+`, "y", "v").Replace(formats.Y.Default)
		formats.YM.Default = g.buildFmtYM(formats.YM.Default, "", "", log)
		dateTimeFormats[calendarType] = formats
	}

	return dateTimeFormats
}

func (g *Generator) addFormatY(
	locale string,
	calendar *cldr.Calendar,
	formats CalendarDateTimeFormats,
	log *slog.Logger,
) {
	y := findDateFormatItem(calendar, "y")

	log.Debug("add Y format", "y", y)

	if y == "" {
		return
	}

	var sb strings.Builder

	for i, v := range parseDatePattern(y) {
		if i > 0 {
			sb.WriteRune('+')
		}

		switch {
		default:
			sb.WriteString(`"` + v.Value + `"`)
		case v.Value == "y":
			sb.WriteString("v")
		}
	}

	formats.Y.Fmt[sb.String()] = append(formats.Y.Fmt[sb.String()], locale)
}

func (g *Generator) addFormatYM(
	locale string,
	calendar *cldr.Calendar,
	formats CalendarDateTimeFormats,
	log *slog.Logger,
) {
	yM := findDateFormatItem(calendar, "yM")
	yMM := findDateFormatItem(calendar, "yMM")
	yyyyM := findDateFormatItem(calendar, "yyyyM")

	log.Debug("add YM format", "yM", yM, "yMM", yMM, "yyyyM", yyyyM)

	if yM == formats.YM.Default &&
		(yMM == "" || yMM == formats.YM.Default) &&
		(yyyyM == "" || yyyyM == formats.YM.Default) ||
		yM == "" && yMM == "" && yyyyM == "" {
		return
	}

	s := g.buildFmtYM(yM, yMM, yyyyM, log)

	formats.YM.Fmt[s] = append(formats.YM.Fmt[s], locale)
}

func (g *Generator) addFormatM(
	locale string,
	calendar *cldr.Calendar,
	formats CalendarDateTimeFormats,
	log *slog.Logger,
) {
	// "L" and "M" have the same meaning - numeric with minimum digits
	m := cmp.Or(findDateFormatItem(calendar, "L"), findDateFormatItem(calendar, "M"))

	log.Debug("add M format", "M", m)

	if m == "" || m == formats.M.Default || m == "M" {
		return
	}

	var sb strings.Builder

	for i, v := range parseDatePattern(m) {
		if i > 0 {
			sb.WriteRune('+')
		}

		if v.Literal {
			sb.WriteString(`"` + v.Value + `"`)
			continue
		}

		switch v.Value {
		default:
			sb.WriteString("fmt(v, opt)")
		case "LL", "MM":
			sb.WriteString(`fmt(v, Month2Digit)`)
		case "LLL":
			sb.WriteString(`fmtMonthName(locale.String(), "stand-alone", "abbreviated")`)
		case "MMM":
			sb.WriteString(`fmtMonthName(locale.String(), "format", "abbreviated")`)
		case "LLLL":
			sb.WriteString(`fmtMonthName(locale.String(), "stand-alone", "wide")`)
		case "MMMM":
			sb.WriteString(`fmtMonthName(locale.String(), "format", "wide")`)
		case "LLLLL":
			sb.WriteString(`fmtMonthName(locale.String(), "stand-alone", "narrow")`)
		case "MMMMM":
			sb.WriteString(`fmtMonthName(locale.String(), "format", "narrow")`)
		}
	}

	s := sb.String()

	if strings.Contains(s, "fmtMonthName") {
		s = "return " + s
	} else {
		s = `fmt := fmtMonth(digits); return func(v time.Month, opt Month) string { return ` + s + ` }`
	}

	formats.M.Fmt[s] = append(formats.M.Fmt[s], locale)
}

func (g *Generator) addFormatMD(
	locale string,
	calendar *cldr.Calendar,
	formats CalendarDateTimeFormats,
	log *slog.Logger,
) {
	md := findDateFormatItem(calendar, "Md")
	mmd := findDateFormatItem(calendar, "MMd")
	mdd := findDateFormatItem(calendar, "Mdd")
	mmdd := findDateFormatItem(calendar, "MMdd")

	log.Debug("add MD format", "Md", md, "MMd", mmd, "Mdd", mdd, "MMdd", mmdd)

	if md == "" && mmd == "" && mdd == "" && mmdd == "" {
		return
	}

	s := BuildFmtMD(md, mmd, mdd, mmdd, log)

	if s == formats.MD.Default {
		return
	}

	formats.MD.Fmt[s] = append(formats.MD.Fmt[s], locale)
}

func (g *Generator) addFormatD(
	locale string,
	calendar *cldr.Calendar,
	formats CalendarDateTimeFormats,
	log *slog.Logger,
) {
	d := findDateFormatItem(calendar, "d")

	log.Debug("add D format", "d", d)

	if d == "" || d == formats.D.Default {
		return
	}

	var sb strings.Builder

	for i, v := range parseDatePattern(d) {
		if i > 0 {
			sb.WriteRune('+')
		}

		switch {
		default:
			sb.WriteString("fmt(v, opt)")
		case v.Literal:
			sb.WriteString(`"` + v.Value + `"`)
		case v.Value == "dd":
			sb.WriteString(`fmt(v, Day2Digit)`)
		}
	}

	formats.D.Fmt[sb.String()] = append(formats.D.Fmt[sb.String()], locale)
}

func (g *Generator) findDateFormatItem(locale, calendarType string, id string) string {
	calendar := findCalendar(g.cldr.RawLDML(locale), calendarType)

	return findDateFormatItem(calendar, id)
}

func findDateFormatItem(calendar *cldr.Calendar, id string) string {
	for _, availableFormats := range calendar.DateTimeFormats.AvailableFormats {
		for _, dateFormatItem := range availableFormats.DateFormatItem {
			if dateFormatItem.Id != id {
				continue
			}

			return dateFormatItem.CharData
		}
	}

	return ""
}

func (g *Generator) months() Months { //nolint:gocognit
	months := NewMonths()

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		locale = strings.ReplaceAll(locale, "_", "-")

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for calendar := range supportedCalendars(ldml.Dates.Calendars.Calendar) {
			if calendar.Months == nil {
				continue
			}

			for _, monthContext := range calendar.Months.MonthContext {
				for _, monthWidth := range monthContext.MonthWidth {
					if len(monthWidth.Month) == 0 {
						continue
					}

					month := monthWidth.Month[0]

					// skip months with the same digits
					if month.Type == month.CharData && month.CharData == "1" {
						continue
					}

					var monthNames MonthNames

					for _, month = range monthWidth.Month {
						i, err := strconv.Atoi(month.Type)
						if err != nil {
							panic(err)
						}

						i--

						monthNames[i] = month.CharData
					}

					// skip empty names
					if monthNames[0] == "" {
						continue
					}

					i := slices.IndexFunc(months.List, func(names MonthNames) bool {
						for i, v := range names {
							if v != monthNames[i] {
								return false
							}
						}

						return true
					})

					if i == -1 {
						months.List = append(months.List, monthNames)
						i = len(months.List) - 1
					}

					indexes := months.Lookup[locale]
					indexes.Set(calendar.Type, monthWidth.Type, monthContext.Type, i)

					// NOTE: fallback "format" context when "stand-alone" not defined
					if monthContext.Type == "format" {
						indexes.Set(calendar.Type, monthWidth.Type, "stand-alone", i)
					}

					months.Lookup[locale] = indexes
				}
			}
		}
	}

	return months
}

// CLDRDateFormatItem is a copy of CLDR DateFormatItem.
type CLDRDateFormatItem struct {
	cldr.Common
	Id    string //nolint:revive,stylecheck
	Count string
}

func (g *Generator) numberingSystems(defaultNumberingSystems DefaultNumberingSystems) []NumberingSystem {
	numberingSystems := make([]NumberingSystem, 0, 12) //nolint:mnd

	ids := slices.Collect(maps.Keys(defaultNumberingSystems))

	for _, v := range g.cldr.Supplemental().NumberingSystems.NumberingSystem {
		if v.Type != "numeric" || !slices.Contains(ids, v.Id) {
			continue
		}

		numberingSystem := NumberingSystem{ID: v.Id}

		var i int

		for _, digit := range v.Digits {
			numberingSystem.Digits[i] = digit
			i++
		}

		numberingSystems = append(numberingSystems, numberingSystem)
	}

	return numberingSystems
}

func (g *Generator) numberingSystemsIota(defaultNumberingSystems DefaultNumberingSystems) []string {
	ids := slices.Collect(maps.Keys(defaultNumberingSystems))

	slices.Sort(ids)

	return slices.Compact(ids)
}

type NumberingSystem struct {
	ID     string
	Digits [10]rune
}

type TemplateData struct {
	Months                  Months
	DefaultNumberingSystems DefaultNumberingSystems
	NumberingSystemIota     []string
	CalendarPreferences     []CalendarPreference
	DateTimeFormats         DateTimeFormats
	NumberingSystems        []NumberingSystem
}

// value - locales.
type Months struct {
	// key is locale, value is 18 indexes from [List].
	Lookup map[string]MonthIndexes
	List   []MonthNames
}

func NewMonths() Months {
	return Months{
		Lookup: make(map[string]MonthIndexes),
	}
}

type MonthKey struct {
	Locale       string
	CalendarType string // gregorian, persian or buddhist
	Width        string // wide, narrow, abbreviated
	Context      string // format or stand-alone
}

type MonthIndexes [18]int

func (m *MonthIndexes) Set(calendarType, width, context string, i int) {
	widthsCount := 3
	contextCount := 2

	var t, w, c int

	// the order MUST be the same as const of [intl.calendarType]
	switch calendarType {
	case "gregorian":
		t = 0
	case "buddhist":
		t = 1
	case "persian":
		t = 2
	}

	switch width {
	case "abbreviated":
		w = 0
	case "wide":
		w = 1
	case "narrow":
		w = 2
	}

	switch context {
	case "format":
		c = 0
	case "stand-alone":
		c = 1
	}

	index := t*widthsCount*contextCount + w*contextCount + c

	m[index] = i
}

type MonthNames [12]string

func (n MonthNames) String() string {
	return `{"` + strings.Join(n[:], `", "`) + `"}`
}

// key - calendar type.
type DateTimeFormats map[string]CalendarDateTimeFormats

type CalendarDateTimeFormats struct {
	Y, YM, M, MD, D CalendarDateTimeFormat
}

func NewCalendarDateTimeFormats() CalendarDateTimeFormats {
	return CalendarDateTimeFormats{
		Y:  NewCalendarDateTimeFormat(),
		YM: NewCalendarDateTimeFormat(),
		M:  NewCalendarDateTimeFormat(),
		MD: NewCalendarDateTimeFormat(),
		D:  NewCalendarDateTimeFormat(),
	}
}

type CalendarDateTimeFormat struct {
	// key - expr (format), value - languages.
	Fmt     map[string][]string
	Default string
}

func NewCalendarDateTimeFormat() CalendarDateTimeFormat {
	return CalendarDateTimeFormat{Fmt: make(map[string][]string)}
}

type CalendarPreference struct {
	Regions   []string
	Calendars []string
}

func Locale(ldml *cldr.LDML) string {
	lang := ldml.Identity.Language.Type

	if ldml.Identity.Script != nil && ldml.Identity.Script.Type != "" {
		lang += "-" + ldml.Identity.Script.Type
	}

	if ldml.Identity.Territory != nil && ldml.Identity.Territory.Type != "" {
		lang += "-" + ldml.Identity.Territory.Type
	}

	return lang
}

type DefaultNumberingSystems map[string][]string // key - numbering system, value - locales

func deepCopy[T any](v T) T { //nolint:ireturn
	var r T

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &r); err != nil {
		panic(err)
	}

	return r
}

func isContributedOrApproved(draft string) bool {
	return draft == "" || draft == "contributed"
}

func title(s string) string {
	var r string

	for _, v := range strings.Split(s, " ") {
		r += cases.Title(language.English).String(v)
	}

	return strings.ReplaceAll(r, "-", "") // e.g. "islamic - umalqura"
}

func (g *Generator) buildFmtYM(yM, yMM, yyyyM string, log *slog.Logger) string {
	yMPattern, yMMPattern, _ := yearMonthPatterns(yM, yMM, yyyyM)

	log.Debug("infer YM patterns", "yM", yMPattern.String(), "yMM", yMMPattern.String())

	switch {
	default: // yM == yMM
		var sb strings.Builder

		sb.WriteString("return ")

		for i, v := range yMPattern {
			if i > 0 {
				sb.WriteRune('+')
			}

			switch v.Value {
			default:
				sb.WriteString(`"` + v.Value + `"`)
			case "L", "M":
				if yMMmonth := yMMPattern.Month(); yMM != "" && (yMMmonth == "M" || yMMmonth == "L") { // TODO: use len() instead
					sb.WriteString(`fmtMonth(m, MonthNumeric)`)
				} else {
					sb.WriteString(`fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))`)
				}
			case "LL", "MM":
				if yMMPattern.Month() == v.Value {
					sb.WriteString(`fmtMonth(m, Month2Digit)`)
				} else {
					sb.WriteString(`fmtMonth(m, cmp.Or(opts.Month, Month2Digit))`)
				}
			case "MMMMM":
				sb.WriteString(`fmtMonthName(locale.String(), "stand-alone", "narrow")(m, opts.Month)`)
			case "y", "Y":
				sb.WriteString("fmtYear(y, cmp.Or(opts.Year, YearNumeric))")
			}
		}

		return sb.String()
	case yM == "y/M" && yMM == "y年M月":
		return `
	ys := fmtYear(y, cmp.Or(opts.Year, YearNumeric))
	ms := fmtMonth(m, MonthNumeric)
	if opts.Month == MonthNumeric {
		return ys+"/"+ms
	}
	return ys+"年"+ms+"月"`
	case yM == "y-MM" && yMM == "MM/y":
		return `
	ys := fmtYear(y, cmp.Or(opts.Year, YearNumeric))
	ms := fmtMonth(m, Month2Digit)
	if opts.Month == MonthNumeric {
		return ys+"-"+ms
	}
	return ms+"/"+ys`
	case len(yMPattern) > 0 && len(yMMPattern) > 0 && yMPattern[1] != yMMPattern[1]:
		return fmt.Sprintf(
			`if (opts.Month == MonthNumeric) { %s }; %s`,
			g.buildFmtYM(yM, "", "", log), g.buildFmtYM(yMM, "", "", log))
	case len(yMPattern) > 0 && yMPattern[0].Value == "MM" && yMMPattern[0].Value == "M":
		return `
	if opts.Month == MonthNumeric {
		return fmtMonth(m, Month2Digit)+"/"+fmtYear(y, cmp.Or(opts.Year, YearNumeric))
	}
	return fmtMonth(m, MonthNumeric)+"/"+fmtYear(y, cmp.Or(opts.Year, YearNumeric))`
	}
}
