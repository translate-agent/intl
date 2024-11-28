package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	default:
		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + " " + fmtDay(d, opts.Day)
		}
	case lv:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
		}
	}
}
