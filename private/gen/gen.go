package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/unicode/cldr"
)

//go:embed datetime.tmpl
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
	var preferences []CalendarPreference

	// calendar preferences
	for _, v := range g.cldr.Supplemental().CalendarPreferenceData.CalendarPreference {
		preferences = append(preferences, CalendarPreference{
			Regions:   strings.Split(v.Territories, " "),
			Calendars: strings.Split(v.Ordering, " "),
		})
	}

	return preferences
}

func (g *Generator) defaultNumberingSystems() []DefaultNumberingSystem {
	var defaultNumberingSystems []DefaultNumberingSystem

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Numbers == nil {
			continue
		}

		for _, v := range ldml.Numbers.DefaultNumberingSystem {
			if v.Alt != "" || v.Draft != "" {
				continue
			}

			defaultNumberingSystems = append(defaultNumberingSystems, DefaultNumberingSystem{
				Locale: strings.ReplaceAll(locale, "_", "-"),
				ID:     v.CharData,
			})
		}
	}

	return defaultNumberingSystems
}

func (g *Generator) numerals() []Numerals {
	var numerals []Numerals

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Characters != nil {
			for _, characters := range ldml.Characters.ExemplarCharacters {
				entry := Numerals{Locale: Locale(ldml)}

				if characters.Type == "numbers" && strings.Contains(characters.CharData, " 0") && !strings.Contains(characters.CharData, " 0 ") {
					entry.Characters = numeralCharacters(characters.CharData)
				}

				numerals = append(numerals, entry)
			}
		}
	}

	return numerals
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

		if ldml.Dates != nil && ldml.Dates.Calendars != nil {
			for _, calendar := range ldml.Dates.Calendars.Calendar {
				if calendar.Type != "gregorian" {
					continue
				}

				if calendar.DateTimeFormats != nil {
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
		}
	}

	return dateTimeFormats
}

func (g *Generator) numberingSystems() []NumberingSystem {
	var numberingSystems []NumberingSystem

	for _, v := range g.cldr.Supplemental().NumberingSystems.NumberingSystem {
		if v.Type != "numeric" {
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

func (g *Generator) Write() error {
	tpl, err := template.New("datetime").Funcs(template.FuncMap{
		"join":     strings.Join,
		"contains": strings.Contains,
	}).Parse(datetimeTemplate)
	if err != nil {
		return fmt.Errorf("parse datetime template: %w", err)
	}

	data := TemplateData{
		Numerals:                g.numerals(),
		CalendarPreferences:     g.calendarPreferences(),
		DateTimeFormats:         g.dateTimeFormats(),
		NumberingSystems:        g.numberingSystems(),
		DefaultNumberingSystems: g.defaultNumberingSystems(),
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

type NumeralCharacters [10]rune

func numeralCharacters(text string) *NumeralCharacters {
	var numerals NumeralCharacters

	for _, r := range strings.Split(text, " ") {
		if r[0] >= '0' && r[0] <= '9' {
			var n int
			i := int(r[0] - '0')
			numerals[i], n = utf8.DecodeRune([]byte(r)[1:])
			if n == 0 {
				panic("missing number rune")
			}
		}
	}

	return &numerals
}

type TemplateData struct {
	Numerals                []Numerals
	CalendarPreferences     []CalendarPreference
	DateTimeFormats         DateTimeFormats
	NumberingSystems        []NumberingSystem
	DefaultNumberingSystems []DefaultNumberingSystem
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

type Numerals struct {
	Locale     string
	Characters *NumeralCharacters
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

type DefaultNumberingSystem struct {
	Locale string
	ID     string
}
