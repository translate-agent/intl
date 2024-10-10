package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
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

func Gen(cldrDir, out string) error {
	g := Generator{}

	if err := g.Load(cldrDir); err != nil {
		return err
	}

	if err := g.Write(out); err != nil {
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

	g.merge()

	return nil
}

func (g *Generator) Write(out string) error {
	tpl, err := template.New("datetime").Funcs(template.FuncMap{
		"join":     strings.Join,
		"contains": strings.Contains,
		"title":    cases.Title(language.English).String,
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

type CalendarTypes []string

func (t CalendarTypes) Skip(s string) bool {
	return !slices.Contains(t, s)
}

var calendarTypes = CalendarTypes{"gregorian", "persian", "buddhist"}

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

	for _, calendarType := range calendarTypes {
		parentCalendar := findCalendar(fallback, calendarType)
		// skip if parent calendar not found
		if parentCalendar == nil {
			continue
		}

		if parentCalendar.Alias != nil &&
			parentCalendar.Alias.Path == "../../calendar[@type='generic']/dateTimeFormats" {
			parentCalendar.DateTimeFormats = deepCopy(findCalendar(fallback, "generic").DateTimeFormats)
		}

		calendar := findCalendar(dst, calendarType)
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
			if v.Alt != "" || !isContributedOrApproved(v.Draft) {
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

		for _, calendar := range ldml.Dates.Calendars.Calendar {
			if calendarTypes.Skip(calendar.Type) {
				continue
			}

			formats, ok := dateTimeFormats[calendar.Type]
			if !ok {
				formats = NewCalendarDateTimeFormats()

				formats.Y.Default = g.defaultDateFormatItem(calendar.Type, "y")
				formats.D.Default = g.defaultDateFormatItem(calendar.Type, "d")

				dateTimeFormats[calendar.Type] = formats
			}

			for _, availableFormats := range calendar.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					g.addDateFormatItem(calendar.Type, formats, (*CLDRDateFormatItem)(dateFormatItem), locale)
				}
			}
		}
	}

	return dateTimeFormats
}

func (g *Generator) defaultDateFormatItem(calendarType string, id string) string {
	calendars := g.cldr.RawLDML("root").Dates.Calendars.Calendar

	// TODO(jhorsts): use findCalendar()
	i := slices.IndexFunc(g.cldr.RawLDML("root").Dates.Calendars.Calendar, func(calendar *cldr.Calendar) bool {
		return calendar.Type == calendarType
	})

	calendar := calendars[i]

	if calendar.DateTimeFormats.Alias != nil {
		switch {
		case strings.Contains(calendar.DateTimeFormats.Alias.Path, "gregorian"):
			return g.defaultDateFormatItem("gregorian", id)
		case strings.Contains(calendar.DateTimeFormats.Alias.Path, "generic"):
			return g.defaultDateFormatItem("generic", id)
		}
	}

	for _, availableFormats := range calendar.DateTimeFormats.AvailableFormats {
		for _, dateFormatItem := range availableFormats.DateFormatItem {
			if dateFormatItem.Id != id {
				continue
			}

			if id != "y" && calendarType != "persian" {
				return dateFormatItem.CharData
			}

			return strings.Replace(dateFormatItem.CharData, "G ", `"AP "+`, 1)
		}
	}

	return ""
}

func (g *Generator) months() Months {
	months := NewMonths()

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)
		loc := strings.ReplaceAll(locale, "_", "-")

		for _, calendar := range ldml.Dates.Calendars.Calendar {
			if calendarTypes.Skip(calendar.Type) {
				continue
			}

			for _, monthContext := range calendar.Months.MonthContext {
				for _, monthWidth := range monthContext.MonthWidth {
					var monthNames MonthNames

				monthLoop:
					for _, month := range monthWidth.Month {
						// skip draft and months with the same digits
						if !isContributedOrApproved(month.Draft) ||
							month.Type == month.CharData && month.CharData == "1" {
							break monthLoop
						}

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

					widthsCount := 3
					contextCount := 2

					var t, w, c int

					switch calendar.Type {
					case "gregorian":
						t = 0
					case "buddhist":
						t = 1
					case "persian":
						t = 2
					}

					switch monthWidth.Type {
					case "abbreviated":
						w = 0
					case "wide":
						w = 1
					case "narrow":
						w = 2
					}

					switch monthContext.Type {
					case "format":
						c = 0
					case "stand-alone":
						c = 1
					}

					index := t*widthsCount*contextCount + w*contextCount + c

					indexes := months.Lookup[loc]
					if monthContext.Type == "format" {
						indexes[index+1] = i // NOTE: fallback "format" context when "stand-alone" not defined
					}

					indexes[index] = i
					months.Lookup[loc] = indexes
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

//nolint:gocognit
func (g *Generator) addDateFormatItem(
	calendarType string,
	dateTimeFormats CalendarDateTimeFormats,
	dateFormatItem *CLDRDateFormatItem,
	locale string,
) {
	if !isContributedOrApproved(dateFormatItem.Draft) {
		return
	}

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
				sb.WriteString("y")
			}
		}

		dateTimeFormats.Y.Fmt[sb.String()] = append(dateTimeFormats.Y.Fmt[sb.String()], locale)
	case "M", "L":
		if dateFormatItem.CharData == dateTimeFormats.M.Default {
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

			switch v.value {
			default:
				sb.WriteString("fmt(m, f)")
			case "LL", "MM":
				sb.WriteString(`fmt(m, "01")`)
			case "LLL":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "stand-alone", "abbreviated")`, calendarType))
			case "MMM":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "format", "abbreviated")`, calendarType))
			case "LLLL":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "stand-alone", "wide")`, calendarType))
			case "MMMM":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "format", "wide")`, calendarType))
			case "LLLLL":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "stand-alone", "narrow")`, calendarType))
			case "MMMMM":
				sb.WriteString(fmt.Sprintf(`fmtMonth(locale.String(), "%s", "format", "narrow")`, calendarType))
			}
		}

		s := sb.String()

		if !strings.Contains(s, "fmtMonth") {
			s = `func(m int, f string) string { return ` + s + ` }`
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
				sb.WriteString("fmt(d, f)")
			case v.literal:
				sb.WriteString(`"` + v.value + `"`)
			case v.value == "dd":
				sb.WriteString(`fmt(d, "02")`)
			}
		}

		dateTimeFormats.D.Fmt[sb.String()] = append(dateTimeFormats.D.Fmt[sb.String()], locale)
	}
}

func (g *Generator) numberingSystems(defaultNumberingSystems DefaultNumberingSystems) []NumberingSystem {
	numberingSystems := make([]NumberingSystem, 0, 12) //nolint:mnd

	ids := slices.Collect(maps.Keys(defaultNumberingSystems))

	for _, v := range g.cldr.Supplemental().NumberingSystems.NumberingSystem {
		// only use default numbering systems
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
	DefaultNumberingSystems DefaultNumberingSystems
	NumberingSystemIota     []string
	CalendarPreferences     []CalendarPreference
	DateTimeFormats         DateTimeFormats
	NumberingSystems        []NumberingSystem
	Months                  Months
}

// value - locales.
type Months struct {
	List []MonthNames
	// key is locale, value is 18 indexes from [List].
	Lookup map[string][18]int
}

func NewMonths() Months {
	return Months{
		Lookup: make(map[string][18]int, 18), //nolint:mnd
	}
}

type MonthKey struct {
	Locale       string
	CalendarType string // gregorian, persian or buddhist
	Width        string // wide, narrow, abbreviated
	Context      string // format or stand-alone
}

type MonthNames [12]string

func (n MonthNames) String() string {
	return `{"` + strings.Join(n[:], `", "`) + `"}`
}

// key - calendar type.
type DateTimeFormats map[string]CalendarDateTimeFormats

type CalendarDateTimeFormats struct {
	Y CalendarDateTimeFormat
	M CalendarDateTimeFormat
	D CalendarDateTimeFormat
}

func NewCalendarDateTimeFormats() CalendarDateTimeFormats {
	return CalendarDateTimeFormats{
		Y: NewCalendarDateTimeFormat(),
		M: NewCalendarDateTimeFormat(),
		D: NewCalendarDateTimeFormat(),
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

func splitDatePattern(pattern string) []datePatternElement {
	var last rune

	elements := make([]datePatternElement, 0)
	elem := new(strings.Builder)
	literal := false
	quoted := false

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

func isContributedOrApproved(s string) bool {
	return s == "" || s == "contributed"
}
