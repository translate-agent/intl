package intl

import (
	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits digits, opt Month) fmtFunc {
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

	monthDigits := convertMonthDigits(digits, opt)

	return func(t timeReader) string { return monthDigits(t) + suffix }
}

func fmtMonthBuddhist(_ language.Tag, digits digits, opt Month) fmtFunc {
	monthDigits := convertMonthDigits(digits, opt)

	return func(t timeReader) string { return monthDigits(t) }
}

func fmtMonthPersian(_ language.Tag, digits digits, opt Month) fmtFunc {
	return convertMonthDigits(digits, opt)
}
