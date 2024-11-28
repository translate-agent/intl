package intl

import "golang.org/x/text/language"

func fmtEraYearDayGregorian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)

	switch lang {
	default:
		return func(y, d int) string {
			return era + fmtDay(d, opts.Day)
		}
	case lv:
		return func(y, d int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + " (" + name + ": " + fmtDay(d, opts.Day) + ")"
		}
	}
}
