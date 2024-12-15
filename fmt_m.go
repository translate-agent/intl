package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits digits, opt Month) func(m time.Month) string {
	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case br, fo, ga, lt, uk, uz:
		opt = Month2Digit
	case hr, nb, nn, no, sk:
		suffix = "."
	case ja, yue, zh, ko:
		suffix = unitName(locale).Month
	case mn:
		return fmtMonthName(locale.String(), "stand-alone", "narrow")
	case wae:
		return fmtMonthName(locale.String(), "stand-alone", "abbreviated")
	}

	month := fmtMonth(digits, opt)

	return func(m time.Month) string { return month(m) + suffix }
}

func fmtMonthBuddhist(_ language.Tag, digits digits, opt Month) func(m time.Month) string {
	return fmtMonth(digits, opt)
}

func fmtMonthPersian(_ language.Tag, digits digits, opt Month) func(m time.Month) string {
	return fmtMonth(digits, opt)
}
