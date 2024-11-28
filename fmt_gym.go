package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraYearMonth(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearGregorian(locale)
	fmtMonth := fmtMonth(digits)

	switch lang {
	default:
		return func(y int, m time.Month) string {
			return era
		}
	case lv:
		return func(y int, m time.Month) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + " " + fmtMonth(m, opts.Month)
		}
	}
}
