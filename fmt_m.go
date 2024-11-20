package intl

import (
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
