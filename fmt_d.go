package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtDayGregorian(locale language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case cldr.BS:
		if script, _ := locale.Script(); script == cldr.Cyrl {
			break
		}

		suffix = "."
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		suffix = "."
	case cldr.JA, cldr.YUE, cldr.ZH:
		suffix = "日"
	case cldr.KO:
		suffix = "일"
	case cldr.LT:
		opt = Day2Digit
	case cldr.II:
		suffix = "ꑍ"
	}

	dayDigits := convertDayDigits(digits, opt)

	return func(t timeReader) string { return dayDigits(t) + suffix }
}

func fmtDayBuddhist(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	dayDigits := convertDayDigits(digits, opt)
	return func(t timeReader) string { return dayDigits(t) }
}

func fmtDayPersian(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	return convertDayDigits(digits, opt)
}
