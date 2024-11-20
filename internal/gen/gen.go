package main

import (
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
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"
)

//go:embed cldr_data.go.tmpl
var cldrDataTemplate string

// LocaleLookup maps a shared property (the key) to a list of locales that share that property.
type LocaleLookup map[string][]string

type Generator struct {
	cldr *cldr.CLDR
}

type Conf struct {
	cldrDir    string
	out        string
	saveMerged bool
}

func Gen(conf Conf, log *slog.Logger) error {
	var g Generator

	if err := g.load(conf.cldrDir, log); err != nil {
		return err
	}

	if conf.saveMerged {
		if err := g.saveMerged(conf.out); err != nil {
			return err
		}
	}

	if err := g.write(conf.out); err != nil {
		return err
	}

	return nil
}

func (g *Generator) load(dir string, log *slog.Logger) error {
	var (
		d   cldr.Decoder
		err error
	)

	now := time.Now()

	defer func() {
		log.Debug("loading", "duration", time.Since(now))
	}()

	d.SetDirFilter("main", "supplemental")

	g.cldr, err = d.DecodePath(dir)
	if err != nil {
		return fmt.Errorf(`decode CLDR at path "%s": %w`, dir, err)
	}

	g.filterApproved()
	g.merge(log)

	return nil
}

func (g *Generator) write(out string) error {
	defaultNumberingSystems := g.defaultNumberingSystems()
	calendarPreferences := g.calendarPreferences()

	data := TemplateData{
		CalendarPreferences:     calendarPreferences,
		NumberingSystems:        g.numberingSystems(defaultNumberingSystems),
		NumberingSystemIota:     g.numberingSystemsIota(defaultNumberingSystems),
		DefaultNumberingSystems: defaultNumberingSystems,
		Months:                  g.months(),
	}

	dataTpl, err := template.New("cldr_data").Funcs(template.FuncMap{
		"join":  strings.Join,
		"sub":   func(a, b int) int { return a - b },
		"title": title,
	}).Parse(cldrDataTemplate)
	if err != nil {
		return fmt.Errorf("parse cldr_data: %w", err)
	}

	name := path.Join(out, "cldr_data.go")

	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file %s: %w", name, err)
	}

	defer f.Close()

	if err := dataTpl.Execute(f, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
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

	switch dst.DateTimeFormats {
	default:
		if dst.DateTimeFormats.AvailableFormats == nil && src.DateTimeFormats != nil {
			dst.DateTimeFormats.AvailableFormats = deepCopy(src.DateTimeFormats.AvailableFormats)
		}

		if src.DateTimeFormats != nil {
			for _, availableFormats := range src.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					if containsDateFormatItem(dst, dateFormatItem.Id) {
						continue
					}

					// NOTE(jhorsts): Why the first AvailableFormats? I don't remember.
					dst.DateTimeFormats.AvailableFormats[0].DateFormatItem = append(
						dst.DateTimeFormats.AvailableFormats[0].DateFormatItem,
						deepCopy(dateFormatItem))
				}
			}
		}
	case nil:
		dst.DateTimeFormats = deepCopy(src.DateTimeFormats)
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

func (g *Generator) defaultNumberingSystems() LocaleLookup {
	defaultNumberingSystems := make(LocaleLookup)

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

func (g *Generator) months() Months { //nolint:gocognit
	months := NewMonths()

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		locale = strings.ReplaceAll(locale, "_", "-")

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for calendar := range supportedCalendars(ldml.Dates.Calendars.Calendar) {
			// month names are available only in gregorian calendars (default)
			if calendar.Months == nil || calendar.Type != "gregorian" {
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
					indexes.Set(monthWidth.Type, monthContext.Type, i)

					// NOTE: fallback "format" context when "stand-alone" not defined
					if monthContext.Type == "format" {
						indexes.Set(monthWidth.Type, "stand-alone", i)
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

func (g *Generator) numberingSystems(defaultNumberingSystems LocaleLookup) []NumberingSystem {
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

func (g *Generator) numberingSystemsIota(defaultNumberingSystems LocaleLookup) []string {
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
	DefaultNumberingSystems LocaleLookup
	NumberingSystemIota     []string
	CalendarPreferences     []CalendarPreference
	NumberingSystems        []NumberingSystem
}

// value - locales.
type Months struct {
	// key is locale, value is 6 indexes from [List].
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

// MonthIndexes contains indexes for month names in [Months.List]:
//
//	0 - abbreviated, format
//	1 - abbreviated, stand-alone
//	2 - wide, format
//	3 - wide, stand-alone
//	4 - narrow, format
//	5 - narrow, stand-alone
type MonthIndexes [6]int

func (m *MonthIndexes) Set(width, context string, i int) {
	contextCount := 2

	var w, c int

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

	index := w*contextCount + c

	m[index] = i
}

type MonthNames [12]string

func (n MonthNames) String() string {
	return `{"` + strings.Join(n[:], `", "`) + `"}`
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

//nolint:ireturn
func deepCopy[T any](v T) T {
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
