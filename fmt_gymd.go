package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraYearMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month, d int) string {
	lang, _, _ := locale.Raw()

	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	default:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case lv:
		// date always formatted as dd-MM-yyyy or dd-MM-yy
		// era=long,...,out=mūsu ērā 02-01-2024
		// era=short,...,out=m.ē. 02-01-2024
		// era=narrow,...,out=m.ē. 02-01-2024
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtYear(y, opts.Year)
		}
	}
}
