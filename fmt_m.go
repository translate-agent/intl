package intl

import (
	"cmp"
	"time"

	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits digits) func(v time.Month, opt Month) string {
	lang, _ := locale.Base()

	switch lang.String() {
	default:
		return fmtMonth(digits)
	case "br", "fo", "ga", "lt", "uk", "uz":
		fmt := fmtMonth(digits)
		return func(v time.Month, _ Month) string { return fmt(v, Month2Digit) }
	case "hr", "nb", "nn", "no", "sk":
		fmt := fmtMonth(digits)
		return func(v time.Month, opt Month) string { return fmt(v, opt) + "." }
	case "ja", "yue", "zh":
		fmt := fmtMonth(digits)
		return func(v time.Month, opt Month) string { return fmt(v, opt) + "月" }
	case "ko":
		fmt := fmtMonth(digits)
		return func(v time.Month, opt Month) string { return fmt(v, opt) + "월" }
	case "mn":
		return fmtMonthName(locale.String(), "stand-alone", "narrow")
	case "wae":
		return fmtMonthName(locale.String(), "stand-alone", "abbreviated")
	}
}

func fmtMonthBuddhist(_ language.Tag, digits digits) func(v time.Month, opt Month) string {
	return fmtMonth(digits)
}

func fmtMonthPersian(_ language.Tag, digits digits) func(v time.Month, opt Month) string {
	return fmtMonth(digits)
}

func fmtMonthDayPersian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, _ := locale.Base()

	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang.String() {
	default:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case "fa", "ps":
		return func(m time.Month, d int) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtDay(d, cmp.Or(opts.Day, DayNumeric))
		}
	}
}
