package main

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"iter"
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
	g := Generator{}

	if err := g.Load(conf.cldrDir); err != nil {
		return err
	}

	if conf.saveMerged {
		if err := g.saveMerged(conf.out); err != nil {
			return err
		}
	}

	if err := g.Write(conf.out); err != nil {
		return err
	}

	return nil
}

func (g *Generator) Load(dir string) error {
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
	g.merge()

	return nil
}

func (g *Generator) Write(out string) error {
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

	data := TemplateData{
		CalendarPreferences:     g.calendarPreferences(),
		DateTimeFormats:         g.dateTimeFormats(),
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

func (g *Generator) merge() {
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
		merge(parent, root)

		for _, locale := range strings.Split(parentLocale.Locales, " ") {
			child := g.cldr.RawLDML(locale)

			if child == nil {
				continue
			}

			merge(child, parent)
		}
	}

	// merge root to language

	for _, locale := range g.cldr.Locales() {
		if strings.ContainsRune(locale, '_') {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		merge(ldml, root)
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

		merge(ldml, fallback)
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

func supportedCalendars(calendars []*cldr.Calendar) iter.Seq[*cldr.Calendar] {
	return func(yield func(*cldr.Calendar) bool) {
		for _, v := range calendars {
			if slices.Contains([]string{"gregorian", "persian", "buddhist"}, v.Type) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// merge copies particular fallback values to dst.
func merge(dst, fallback *cldr.LDML) {
	if dst.Dates == nil {
		dst.Dates = deepCopy(fallback.Dates)
	}

	if dst.Dates.Calendars == nil {
		dst.Dates.Calendars = deepCopy(fallback.Dates.Calendars)
	}

	if len(dst.Dates.Calendars.Calendar) == 0 {
		dst.Dates.Calendars.Calendar = deepCopy(fallback.Dates.Calendars.Calendar)
	}

	for parentCalendar := range supportedCalendars(fallback.Dates.Calendars.Calendar) {
		if parentCalendar.Alias != nil &&
			parentCalendar.Alias.Path == "../../calendar[@type='generic']/dateTimeFormats" {
			parentCalendar.DateTimeFormats = deepCopy(findCalendar(fallback, "generic").DateTimeFormats)
		}

		calendar := findCalendar(dst, parentCalendar.Type)
		if calendar == nil {
			calendar = deepCopy(parentCalendar)
		}

		// datetimeformat

		if calendar.DateTimeFormats == nil {
			calendar.DateTimeFormats = deepCopy(parentCalendar.DateTimeFormats)
		}

		if calendar.DateTimeFormats.AvailableFormats == nil {
			calendar.DateTimeFormats.AvailableFormats = deepCopy(parentCalendar.DateTimeFormats.AvailableFormats)
		}

		for _, availableFormats := range parentCalendar.DateTimeFormats.AvailableFormats {
			for _, dateFormatItem := range availableFormats.DateFormatItem {
				if containsDateFormatItem(calendar, dateFormatItem.Id) {
					continue
				}

				calendar.DateTimeFormats.AvailableFormats[0].DateFormatItem = append(
					calendar.DateTimeFormats.AvailableFormats[0].DateFormatItem,
					deepCopy(dateFormatItem))
			}
		}

		// months

		if calendar.Months == nil {
			calendar.Months = deepCopy(parentCalendar.Months)
		}

		if len(calendar.Months.MonthContext) == 0 && len(parentCalendar.Months.MonthContext) > 0 {
			calendar.Months.MonthContext = deepCopy(parentCalendar.Months.MonthContext)
		}
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

func (g *Generator) dateTimeFormats() DateTimeFormats {
	dateTimeFormats := make(DateTimeFormats)

	for _, locale := range g.cldr.Locales() {
		// Ignore duplicate formatting for "y".
		// Locales containing "_" have the same "y" formatting, skip them for now.
		if strings.Contains(locale, "_") {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for calendar := range supportedCalendars(ldml.Dates.Calendars.Calendar) {
			formats, ok := dateTimeFormats[calendar.Type]
			if !ok {
				formats = NewCalendarDateTimeFormats()

				formats.Y.Default = g.findDateFormatItem("root", calendar.Type, "y")
				formats.YM.Default = cmp.Or(
					g.findDateFormatItem("root", calendar.Type, "yM"),
					g.findDateFormatItem("root", calendar.Type, "yMM"),
					g.findDateFormatItem("root", calendar.Type, "yyyyM"),
				)
				formats.M.Default = g.findDateFormatItem("root", calendar.Type, "M")
				formats.D.Default = g.findDateFormatItem("root", calendar.Type, "d")

				dateTimeFormats[calendar.Type] = formats
			}

			for _, availableFormats := range calendar.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					g.addDateFormatItem(calendar.Type, formats, (*CLDRDateFormatItem)(dateFormatItem), locale)
				}
			}
		}
	}

	for calendarType, formats := range dateTimeFormats {
		formats.Y.Default = strings.NewReplacer("G ", `"AP "+`, "y", "v").Replace(formats.Y.Default)
		formats.YM.Default = buildFmtYm(formats.YM.Default, "", "")
		dateTimeFormats[calendarType] = formats
	}

	return dateTimeFormats
}

func (g *Generator) findDateFormatItem(locale, calendarType string, id string) string {
	calendar := findCalendar(g.cldr.RawLDML(locale), calendarType)

	if calendar.DateTimeFormats.Alias != nil {
		switch {
		case strings.Contains(calendar.DateTimeFormats.Alias.Path, "gregorian"):
			return g.findDateFormatItem("root", "gregorian", id)
		case strings.Contains(calendar.DateTimeFormats.Alias.Path, "generic"):
			return g.findDateFormatItem("root", "generic", id)
		}
	}

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

		for calendar := range supportedCalendars(ldml.Dates.Calendars.Calendar) {
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

//nolint:gocognit,cyclop
func (g *Generator) addDateFormatItem(
	calendarType string,
	dateTimeFormats CalendarDateTimeFormats,
	dateFormatItem *CLDRDateFormatItem,
	locale string,
) {
	switch dateFormatItem.Id {
	case "y":
		if dateFormatItem.CharData == dateTimeFormats.Y.Default {
			return
		}

		var sb strings.Builder

		for i, v := range splitDatePattern(dateFormatItem.CharData) {
			if i > 0 {
				sb.WriteRune('+')
			}

			switch {
			default:
				sb.WriteString(`"` + v.value + `"`)
			case v.value == "y":
				sb.WriteString("v")
			}
		}

		dateTimeFormats.Y.Fmt[sb.String()] = append(dateTimeFormats.Y.Fmt[sb.String()], locale)
	case "yM", "yyyyM":
		yMM := g.findDateFormatItem(locale, calendarType, "yMM")
		yyyyM := g.findDateFormatItem(locale, calendarType, "yyyyM")

		if dateFormatItem.CharData == dateTimeFormats.YM.Default &&
			(yMM == "" || yMM == dateTimeFormats.YM.Default) &&
			(yyyyM == "" || yyyyM == dateTimeFormats.YM.Default) {
			return
		}

		s := buildFmtYm(dateFormatItem.CharData, yMM, yyyyM)

		dateTimeFormats.YM.Fmt[s] = append(dateTimeFormats.YM.Fmt[s], locale)
	case "M", "L":
		// "L" and "M" have the same meaning - numeric with minimum digits
		if dateFormatItem.CharData == dateTimeFormats.M.Default ||
			dateFormatItem.CharData == "M" {
			return
		}

		var sb strings.Builder

		for i, v := range splitDatePattern(dateFormatItem.CharData) {
			if i > 0 {
				sb.WriteRune('+')
			}

			if v.literal {
				sb.WriteString(`"` + v.value + `"`)
				continue
			}

			f := func(s string) string {
				return fmt.Sprintf(s, title(calendarType))
			}

			switch v.value {
			default:
				sb.WriteString("fmt(v, opt)")
			case "LL", "MM":
				sb.WriteString(`fmt(v, Month2Digit)`)
			case "LLL":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "stand-alone", "abbreviated")`))
			case "MMM":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "format", "abbreviated")`))
			case "LLLL":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "stand-alone", "wide")`))
			case "MMMM":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "format", "wide")`))
			case "LLLLL":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "stand-alone", "narrow")`))
			case "MMMMM":
				sb.WriteString(f(`fmtMonthName(locale.String(), calendarType%s, "format", "narrow")`))
			}
		}

		s := sb.String()

		if strings.Contains(s, "fmtMonthName") {
			s = "return " + s
		} else {
			s = `fmt := fmtMonth(digits); return func(v time.Month, opt Month) string { return ` + s + ` }`
		}

		dateTimeFormats.M.Fmt[s] = append(dateTimeFormats.M.Fmt[s], locale)
	case "d":
		if dateFormatItem.CharData == dateTimeFormats.D.Default {
			return
		}

		var sb strings.Builder

		for i, v := range splitDatePattern(dateFormatItem.CharData) {
			if i > 0 {
				sb.WriteRune('+')
			}

			switch {
			default:
				sb.WriteString("fmt(v, opt)")
			case v.literal:
				sb.WriteString(`"` + v.value + `"`)
			case v.value == "dd":
				sb.WriteString(`fmt(v, Day2Digit)`)
			}
		}

		dateTimeFormats.D.Fmt[sb.String()] = append(dateTimeFormats.D.Fmt[sb.String()], locale)
	}
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
	Y, YM, M, D CalendarDateTimeFormat
}

func NewCalendarDateTimeFormats() CalendarDateTimeFormats {
	return CalendarDateTimeFormats{
		Y:  NewCalendarDateTimeFormat(),
		YM: NewCalendarDateTimeFormat(),
		M:  NewCalendarDateTimeFormat(),
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

type datePatternElement struct {
	value   string
	literal bool
}

type datePatternElements []datePatternElement

func (e datePatternElements) Month() string {
	for _, v := range e {
		if v.literal {
			continue
		}

		// all patterns starting with 'M' or 'L' are months
		if v.value[0] == 'M' || v.value[0] == 'L' {
			return v.value
		}
	}

	return ""
}

func splitDatePattern(pattern string) datePatternElements {
	var (
		last            rune
		literal, quoted bool
		elem            strings.Builder
		elements        = make(datePatternElements, 0, 1)
	)

	write := func(r rune, asLiteral bool) {
		if literal && asLiteral {
			elem.WriteRune(r)
			last = r

			return
		}

		if !asLiteral && r == last {
			elem.WriteRune(r)

			return
		}

		if elem.Len() > 0 {
			elements = append(elements, datePatternElement{value: elem.String(), literal: literal})
		}

		elem.Reset()
		elem.WriteRune(r)

		last = r
		literal = asLiteral
	}

	for i, r := range pattern {
		if i == 0 {
			last = r

			if r == '\'' && len(pattern) > 1 {
				quoted = true
				continue
			}

			elem.WriteRune(r)
			literal = !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')

			continue
		}

		switch {
		default:
			write(r, true)
		case r == '\'':
			quoted = !quoted

			if last != r {
				last = r
				continue
			}

			write(r, true)

			last = 0
		case !quoted && ('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'):
			write(r, false)
		}
	}

	if elem.Len() > 0 {
		elements = append(elements, datePatternElement{value: elem.String(), literal: literal})
	}

	return elements
}

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

func buildFmtYm(yM, yMM, yyyyM string) string {
	yMPattern := splitDatePattern(yM)
	yMMPattern := splitDatePattern(cmp.Or(yyyyM, yMM, yM))

	switch {
	default: // yM == yMM
		var sb strings.Builder

		sb.WriteString("return ")

		for i, v := range yMPattern {
			if i > 0 {
				sb.WriteRune('+')
			}

			switch v.value {
			default:
				sb.WriteString(`"` + v.value + `"`)
			case "L", "M":
				if yMMmonth := yMMPattern.Month(); yMM != "" && (yMMmonth == "M" || yMMmonth == "L") {
					sb.WriteString(`fmtMonth(m, MonthNumeric)`)
				} else {
					sb.WriteString(`fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))`)
				}
			case "LL", "MM":
				if yMMPattern.Month() == v.value {
					sb.WriteString(`fmtMonth(m, Month2Digit)`)
				} else {
					sb.WriteString(`fmtMonth(m, cmp.Or(opts.Month, Month2Digit))`)
				}
			case "MMMMM":
				sb.WriteString(`fmtMonthName(locale.String(), calendarTypeGregorian, "stand-alone", "narrow")(m, opts.Month)`)
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
	case yMPattern[1] != yMMPattern[1]:
		return fmt.Sprintf(
			`if (opts.Month == MonthNumeric) { %s }; %s`,
			buildFmtYm(yM, "", ""), buildFmtYm(yMM, "", ""))
	case yMPattern[0].value == "MM" && yMMPattern[0].value == "M":
		return `
	if opts.Month == MonthNumeric {
		return fmtMonth(m, Month2Digit)+"/"+fmtYear(y, cmp.Or(opts.Year, YearNumeric))
	}
	return fmtMonth(m, MonthNumeric)+"/"+fmtYear(y, cmp.Or(opts.Year, YearNumeric))`
	}
}
