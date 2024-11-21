package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits digits) func(m time.Month, opt Month) string {
	switch lang, _ := locale.Base(); lang {
	default:
		return fmtMonth(digits)
	case br, fo, ga, lt, uk, uz:
		fmt := fmtMonth(digits)
		return func(m time.Month, _ Month) string { return fmt(m, Month2Digit) }
	case hr, nb, nn, no, sk:
		fmt := fmtMonth(digits)
		return func(m time.Month, opt Month) string { return fmt(m, opt) + "." }
	case ja, yue, zh:
		fmt := fmtMonth(digits)
		return func(m time.Month, opt Month) string { return fmt(m, opt) + "月" }
	case ko:
		fmt := fmtMonth(digits)
		return func(m time.Month, opt Month) string { return fmt(m, opt) + "월" }
	case mn:
		return fmtMonthName(locale.String(), "stand-alone", "narrow")
	case wae:
		return fmtMonthName(locale.String(), "stand-alone", "abbreviated")
	}
}

func fmtMonthBuddhist(_ language.Tag, digits digits) func(m time.Month, opt Month) string {
	return fmtMonth(digits)
}

func fmtMonthPersian(_ language.Tag, digits digits) func(m time.Month, opt Month) string {
	return fmtMonth(digits)
}
