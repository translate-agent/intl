package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtEraDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	dayName := cldr.UnitName(locale).Day
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + dayName + ": "
		suffix = ")"
	}

	switch lang {
	case cldr.HI:
		if script != cldr.Latn {
			break
		}

		fallthrough
	case cldr.KAA, cldr.EN, cldr.MHN:
		if !withName {
			prefix = ""
			suffix = " " + era
		}
	case cldr.BG, cldr.CY, cldr.MK:
		if withName {
			prefix = era + " (" + dayName + ": "
		} else {
			prefix = era + " "
		}
	case cldr.BS:
		if script == cldr.Cyrl {
			break
		}

		fallthrough
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		if withName {
			suffix = ".)"
		} else {
			suffix = "."
		}
	case cldr.JA, cldr.KO, cldr.YUE, cldr.ZH:
		if withName {
			suffix = dayName + ")"
		} else {
			suffix = dayName
		}
	case cldr.LT:
		opts.Day = Day2Digit
	case cldr.II:
		if withName {
			suffix = "ꑍ)"
		} else {
			suffix = "ꑍ"
		}
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t cldr.TimeReader) string { return prefix + dayDigits(t) + suffix }
}

func fmtEraDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + cldr.UnitName(locale).Day + ": "
		suffix = ")"
	}

	if lang == cldr.FA {
		if withName {
			prefix = era + " (" + cldr.UnitName(locale).Day + ": "
		} else {
			prefix = era + " "
		}
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v cldr.TimeReader) string { return prefix + dayDigits(v) + suffix }
}

func fmtEraDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	prefix, suffix := era+" ", ""
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		prefix, suffix = era+" ("+cldr.UnitName(locale).Day+": ", ")"
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t cldr.TimeReader) string { return prefix + dayDigits(t) + suffix }
}
