package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		opt = Month2Digit
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		suffix = "."
	case cldr.JA, cldr.YUE, cldr.ZH, cldr.KO:
		suffix = cldr.UnitName(locale).Month
	case cldr.MN:
		return fmtMonthName(locale.String(), "stand-alone", "narrow")
	case cldr.WAE:
		return fmtMonthName(locale.String(), "stand-alone", "abbreviated")
	}

	monthDigits := convertMonthDigits(digits, opt)

	return func(t timeReader) string { return monthDigits(t) + suffix }
}

func fmtMonthBuddhist(_ language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	monthDigits := convertMonthDigits(digits, opt)

	return func(t timeReader) string { return monthDigits(t) }
}

func fmtMonthPersian(_ language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	return convertMonthDigits(digits, opt)
}
