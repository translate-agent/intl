package intl

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/text/language"
)

func fmtEraDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	dayName := unitName(locale).Day
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + dayName + ": "
		suffix = ")"
	}

	switch lang {
	case hi:
		if script != latn {
			break
		}

		fallthrough
	case kaa, en, mhn:
		if !withName {
			prefix = ""
			suffix = " " + era
		}
	case bg, cy, mk:
		if withName {
			prefix = era + " (" + dayName + ": "
		} else {
			prefix = era + " "
		}
	case bs:
		if script == cyrl {
			break
		}

		fallthrough
	case cs, da, dsb, fo, hr, hsb, ie, nb, nn, no, sk, sl:
		if withName {
			suffix = ".)"
		} else {
			suffix = "."
		}
	case ja, ko, yue, zh:
		if withName {
			suffix = dayName + ")"
		} else {
			suffix = dayName
		}
	case lt:
		opts.Day = Day2Digit
	case ii:
		if withName {
			suffix = "ꑍ)"
		} else {
			suffix = "ꑍ"
		}
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t time.Time) string { return prefix + dayDigits(t) + suffix }
}

func fmtEraDayPersian(locale language.Tag, digits digits, opts Options) fmtPersianFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + unitName(locale).Day + ": "
		suffix = ")"
	}

	if lang == fa {
		if withName {
			prefix = era + " (" + unitName(locale).Day + ": "
		} else {
			prefix = era + " "
		}
	}

	dayDigits := convertDayDigitsPersian(digits, opts.Day)

	return func(v ptime.Time) string { return prefix + dayDigits(v) + suffix }
}

func fmtEraDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	prefix, suffix := era+" ", ""
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		prefix, suffix = era+" ("+unitName(locale).Day+": ", ")"
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t time.Time) string { return prefix + dayDigits(t) + suffix }
}
