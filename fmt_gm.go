package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtEraMonthGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()
	monthName := cldr.UnitName(locale).Month

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	switch lang {
	case cldr.EN, cldr.KAA, cldr.MHN:
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = ""
			suffix = " " + era
		}
	case cldr.BG, cldr.CY, cldr.MK:
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = era + " "
		}
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		opts.Month = Month2Digit
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		if withName {
			suffix = ".)"
		} else {
			suffix = "."
		}
	case cldr.HI:
		if script != cldr.Latn {
			break
		}

		if opts.Era.long() && opts.Month.numeric() || opts.Era.narrow() {
			prefix = ""
			suffix = " " + era
		}
	case cldr.MN:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
	case cldr.WAE:
		month = fmtMonthName(locale.String(), "format", "abbreviated")
	case cldr.JA, cldr.KO, cldr.ZH, cldr.YUE:
		if withName {
			prefix = era + " (" + monthName + ": "
			suffix = monthName + ")"
		} else {
			suffix = monthName
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	return func(t timeReader) string { return prefix + month(t) + suffix }
}

func fmtEraMonthPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	monthName := cldr.UnitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	if lang == cldr.FA {
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = era + " "
		}
	}

	month := convertMonthDigits(digits, opts.Month)

	return func(v timeReader) string { return prefix + month(v) + suffix }
}

func fmtEraMonthBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	monthDigits := convertMonthDigits(digits, opts.Month)
	monthName := cldr.UnitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	return func(t timeReader) string { return prefix + monthDigits(t) + suffix }
}
