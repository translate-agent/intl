package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraMonthGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Month == Month2Digit
	monthName := unitName(locale).Month

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	switch lang {
	case en:
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = ""
			suffix = " " + era
		}
	case bg, cy:
		prefix = era + " "
		if withName {
			prefix = era + " (" + monthName + ": "
		}
	case br, fo, ga, lt, uk, uz:
		opts.Month = Month2Digit
	case hr, nb, nn, no, sk:
		suffix = "."
		if withName {
			suffix = ".)"
		}
	case hi:
		if script != latn {
			break
		}

		if opts.Era == EraLong && opts.Month == MonthNumeric || opts.Era == EraNarrow {
			prefix = ""
			suffix = " " + era
		}
	case mn:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")
	case wae:
		fmtMonth = fmtMonthName(locale.String(), "format", "abbreviated")
	case ja, ko, zh, yue:
		suffix = monthName

		if withName {
			prefix = era + " (" + monthName + ": "
			suffix = monthName + ")"
		}
	}

	return func(m time.Month) string {
		return prefix + fmtMonth(m, opts.Month) + suffix
	}
}

func fmtEraMonthPersian(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	monthName := unitName(locale).Month
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Month == Month2Digit

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	if lang == fa {
		prefix = era + " "

		if withName {
			prefix = era + " (" + monthName + ": "
		}
	}

	return func(m time.Month) string {
		return prefix + fmtMonth(m, opts.Month) + suffix
	}
}

func fmtEraMonthBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	monthName := unitName(locale).Month
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Month == Month2Digit

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	return func(m time.Month) string {
		return prefix + fmtMonth(m, opts.Month) + suffix
	}
}
