package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtDayGregorian(locale language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	fmt := cldr.Fmt{convertDayDigitsFmt(digits, opt)}

	switch lang, _ := locale.Base(); lang {
	default:
		return fmt.Format
	case cldr.BS:
		if script, _ := locale.Script(); script == cldr.Cyrl {
			return fmt.Format
		}

		return append(fmt, cldr.Text(".")).Format
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		return append(fmt, cldr.Text(".")).Format
	case cldr.JA, cldr.YUE, cldr.ZH:
		return append(fmt, cldr.Text("日")).Format
	case cldr.KO:
		return append(fmt, cldr.Text("일")).Format
	case cldr.LT:
		return cldr.Fmt{convertDayDigitsFmt(digits, Day2Digit)}.Format
	case cldr.II:
		return append(fmt, cldr.Text("ꑍ")).Format
	}
}

func fmtDayBuddhist(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	return cldr.Fmt{convertDayDigitsFmt(digits, opt)}.Format
}

func fmtDayPersian(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	return cldr.Fmt{convertDayDigitsFmt(digits, opt)}.Format
}
