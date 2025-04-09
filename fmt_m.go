package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtMonthGregorian(locale language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	fmt := cldr.Fmt{convertMonthDigitsFmt(digits, opt)}

	switch lang, _ := locale.Base(); lang {
	default:
		return fmt.Format
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		return cldr.Fmt{convertMonthDigitsFmt(digits, Month2Digit)}.Format
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		return append(fmt, cldr.Text(".")).Format
	case cldr.JA, cldr.YUE, cldr.ZH, cldr.KO:
		return append(fmt, cldr.Text(cldr.UnitName(locale).Month)).Format
	case cldr.MN:
		return cldr.Fmt{cldr.Month(cldr.MonthNames(locale.String(), "stand-alone", "narrow"))}.Format
	case cldr.WAE:
		return cldr.Fmt{cldr.Month(cldr.MonthNames(locale.String(), "stand-alone", "abbreviated"))}.Format
	}
}

func fmtMonthBuddhist(_ language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	return cldr.Fmt{convertMonthDigitsFmt(digits, opt)}.Format
}

func fmtMonthPersian(_ language.Tag, digits cldr.Digits, opt Month) fmtFunc {
	return convertMonthDigits(digits, opt)
}
