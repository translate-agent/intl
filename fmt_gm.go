package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraMonthGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	var month func(time.Month) string

	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()
	monthName := unitName(locale).Month

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	switch lang {
	case en, kaa, mhn:
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = ""
			suffix = " " + era
		}
	case bg, cy, mk:
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = era + " "
		}
	case br, fo, ga, lt, uk, uz:
		opts.Month = Month2Digit
	case hr, nb, nn, no, sk:
		if withName {
			suffix = ".)"
		} else {
			suffix = "."
		}
	case hi:
		if script != latn {
			break
		}

		if opts.Era.long() && opts.Month.numeric() || opts.Era.narrow() {
			prefix = ""
			suffix = " " + era
		}
	case mn:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
	case wae:
		month = fmtMonthName(locale.String(), "format", "abbreviated")
	case ja, ko, zh, yue:
		if withName {
			prefix = era + " (" + monthName + ": "
			suffix = monthName + ")"
		} else {
			suffix = monthName
		}
	}

	if month == nil {
		month = fmtMonth(digits, opts.Month)
	}

	return func(m time.Month) string { return prefix + month(m) + suffix }
}

func fmtEraMonthPersian(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	monthName := unitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	if lang == fa {
		if withName {
			prefix = era + " (" + monthName + ": "
		} else {
			prefix = era + " "
		}
	}

	month := fmtMonth(digits, opts.Month)

	return func(m time.Month) string { return prefix + month(m) + suffix }
}

func fmtEraMonthBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	era := fmtEra(locale, opts.Era)
	monthName := unitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	month := fmtMonth(digits, opts.Month)

	return func(m time.Month) string { return prefix + month(m) + suffix }
}
