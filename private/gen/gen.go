package main

import (
	_ "embed"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"
)

//go:embed cldr.tmpl
var datetimeTemplate string

type Generator struct {
	cldr *cldr.CLDR
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

	return nil
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

		if ldml.Numbers == nil {
			continue
		}

		for _, v := range ldml.Numbers.DefaultNumberingSystem {
			if v.Alt != "" || v.Draft != "" {
				continue
			}

			defaultNumberingSystems[v.CharData] = append(defaultNumberingSystems[v.CharData],
				strings.ReplaceAll(locale, "_", "-"))
		}
	}

	return defaultNumberingSystems
}

func (g *Generator) dateTimeFormats() DateTimeFormats {
	dateTimeFormats := DateTimeFormats{Y: make(map[string][]string)}

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
			if calendar.Type != "gregorian" || calendar.DateTimeFormats == nil {
				continue
			}

			for _, availableFormats := range calendar.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					// skip all but "y"
					if dateFormatItem.Id != "y" || dateFormatItem.CharData == "y" {
						continue
					}

					var fmt string

					switch {
					default:
						fmt = `"` + strings.NewReplacer("y", `"+y+"`, "'", "").Replace(dateFormatItem.CharData) + `"`
					case strings.HasPrefix(dateFormatItem.CharData, "y"):
						fmt = strings.NewReplacer("y", `y+"`, "'", "").Replace(dateFormatItem.CharData) + `"`
					case strings.HasSuffix(dateFormatItem.CharData, "y"):
						fmt = `"` + strings.NewReplacer("y", `"+y`, "'", "").Replace(dateFormatItem.CharData)
					}

					dateTimeFormats.Y[fmt] = append(dateTimeFormats.Y[fmt], locale)
				}
			}
		}
	}

	return dateTimeFormats
}

func (g *Generator) numberingSystems(defaultNumberingSystems DefaultNumberingSystems) []NumberingSystem {
	numberingSystems := make([]NumberingSystem, 0, 30) //nolint:mnd

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

func (g *Generator) Write() error {
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
	}

	if err := tpl.Execute(os.Stdout, data); err != nil {
		return fmt.Errorf("execute datetime template: %w", err)
	}

	return nil
}

func Gen(dir string) error {
	g := Generator{}

	if err := g.Load(dir); err != nil {
		return err
	}

	if err := g.Write(); err != nil {
		return err
	}

	return nil
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
}

type DateTimeFormats struct {
	Y map[string][]string // key - expr (format), value - languages
}

type DateTimeFormatItems struct {
	Y string
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

	for i, r := range pattern {
		if i == 0 {
			last = r
			elem.WriteRune(r)
			literal = !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')

			continue
		}

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

			elements = append(elements, datePatternElement{value: elem.String(), literal: literal})

			elem.Reset()
			elem.WriteRune(r)

			last = r
			literal = asLiteral
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
