package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits digits) func(m time.Month, opt Month) string {
	fmt := fmtMonth(digits)
	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case br, fo, ga, lt, uk, uz:
		return func(m time.Month, _ Month) string { return fmt(m, Month2Digit) }
	case hr, nb, nn, no, sk:
		suffix = "."
	case ja, yue, zh, ko:
		suffix = unitName(locale).Month
	case mn:
		return fmtMonthName(locale.String(), "stand-alone", "narrow")
	case wae:
		return fmtMonthName(locale.String(), "stand-alone", "abbreviated")
	}

	return func(m time.Month, opt Month) string { return fmt(m, opt) + suffix }
}

func fmtMonthBuddhist(_ language.Tag, digits digits) func(m time.Month, opt Month) string {
	return fmtMonth(digits)
}

func fmtMonthPersian(_ language.Tag, digits digits) func(m time.Month, opt Month) string {
	return fmtMonth(digits)
}
