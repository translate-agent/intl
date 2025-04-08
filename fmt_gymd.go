package intl

import (
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraYearMonthDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		eraYearMonthDay = iota
		eraMonthDayYear
		eraDayMonthYear
		dayMonthEraYear
	)

	monthOpt, monthDay := Month2Digit, Day2Digit
	layout := eraYearMonthDay
	prefix := era + " "
	suffix := ""
	separator := "-"

	switch lang {
	case en:
		switch region {
		default:
			monthOpt, monthDay = opts.Month, opts.Day
			separator = "/"
			prefix = ""
			suffix = " " + era

			if script == dsrt || script == shaw || region == regionZZ {
				layout = eraMonthDayYear
			} else {
				layout = eraDayMonthYear
			}
		case regionAE, regionAS, regionBI, regionCA, regionGU, regionMH, regionMP, regionPH, regionPR, regionUM, regionUS,
			regionVI:
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraMonthDayYear
			separator = "/"
			prefix = ""
			suffix = " " + era
		case regionCH:
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "."
			prefix = ""
			suffix = " " + era
		case regionGB:
			separator = "/"
			prefix = ""
			suffix = " " + era

			if script == shaw {
				monthOpt, monthDay = opts.Month, opts.Day
				layout = eraMonthDayYear
			} else {
				layout = eraDayMonthYear
			}
		}
	case brx, lv, mni:
		layout = eraDayMonthYear
	case da, dsb, hsb, ie, ka, sq:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case mk:
		monthOpt = opts.Month
		monthDay = opts.Day
		layout = eraDayMonthYear
		prefix = ""
		suffix = " г. " + era
		separator = "."
	case et, pl:
		monthDay = opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case be, cv, de, fo, hy, nb, nn, no, ro, ru:
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case sr:
		monthDay = opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = ". " + era
	case bg:
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " г. " + era
	case fi:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case fr:
		prefix = ""
		suffix = " " + era

		if region != regionCA {
			layout = eraDayMonthYear
			separator = "/"
		}
	case am, as, es, gd, gl, he, el, id, is, jv, nl, su, sw, ta, xnr, ur, vi, yo:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case ga, it, kea, pt, sc, syr, vec:
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case ceb, chr, blo, fil, kaa, mhn, ml, ne, or, ps, sd, so, ti, xh, zu:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case cy:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		prefix = ""
		suffix = " " + era
		separator = "/"
	case ar, ia, bn, ca, mai, rm, uk, wo:
		layout = eraDayMonthYear
		prefix = ""
		suffix = " " + era
	case lt, sv:
		prefix = ""
		suffix = " " + era
	case bs:
		if script != cyrl {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = ". "
			prefix = ""
			suffix = ". " + era
		}
	case ff:
		if script == adlm {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			prefix = ""
			suffix = " " + era
		}
	case ks:
		if script != deva {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraMonthDayYear
			separator = "/"
			prefix = ""
			suffix = " " + era
		}
	case uz:
		if script != cyrl {
			layout = eraDayMonthYear
			separator = "."
			prefix = ""
			suffix = " " + era
		}
	case az:
		if script != cyrl {
			month = fmtMonthName(locale.String(), "format", "abbreviated")
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = " "
		}
	case ku, tk, tr:
		layout = eraDayMonthYear
		separator = "."
	case hu:
		separator = ". "
		suffix = "."
	case cs, sk, sl:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = " " + era
	case hr:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = ". " + era
	case hi:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"

		if script == latn {
			prefix = ""
			suffix = " " + era
		}
	case zh:
		if script == hant {
			monthOpt, monthDay = opts.Month, opts.Day
			separator = "/"
		}
	case kxv:
		if script != deva && script != orya && script != telu {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "/"
		}
	case ja:
		monthOpt, monthDay = opts.Month, opts.Day
		separator = "/"
		prefix = era
	case ko, my:
		monthOpt, monthDay = opts.Month, opts.Day
		separator = "/"
	case mr, qu:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"
	case to:
		layout = eraDayMonthYear
		separator = " "
		prefix = ""
		suffix = " " + era
	case kk:
		layout = dayMonthEraYear
	case lo:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = dayMonthEraYear
		separator = "/"
	case pa:
		if script != arab {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = dayMonthEraYear
			separator = "/"
		}
	case kok:
		if script == latn {
			monthOpt = opts.Month
			monthDay = opts.Day
			layout = eraDayMonthYear
			prefix = ""
			suffix = " " + era
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, monthOpt)
	}

	dayDigits := convertDayDigits(digits, monthDay)

	switch layout {
	default: // eraYearMonthDay
		return func(t timeReader) string {
			return prefix + yearDigits(t) + separator + month(t) + separator + dayDigits(t) + suffix
		}
	case eraMonthDayYear:
		return func(t timeReader) string {
			return prefix + month(t) + separator + dayDigits(t) + separator + yearDigits(t) + suffix
		}
	case eraDayMonthYear:
		return func(t timeReader) string {
			return prefix + dayDigits(t) + separator + month(t) + separator + yearDigits(t) + suffix
		}
	case dayMonthEraYear:
		return func(t timeReader) string {
			return dayDigits(t) + separator + month(t) + separator + era + " " + yearDigits(t)
		}
	}
}

func fmtEraYearMonthDayPersian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		eraYearMonthDay = iota
		eraMonthDayYear
	)

	layout := eraMonthDayYear
	separator := "/"
	suffix := " " + era

	switch lang {
	case ckb:
		if region == regionIR {
			layout = eraYearMonthDay
			separator = "-"
		}
	case lrc, mzn, uz:
		layout = eraYearMonthDay
		separator = "-"
	case ps:
		layout = eraYearMonthDay

		if !opts.Era.narrow() {
			separator = "-"
		}
	case fa:
		suffix = " " + era
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // eraMonthDayYear
		return func(v timeReader) string {
			return month(v) + separator + dayDigits(v) + separator + yearDigits(v) + suffix
		}
	case eraYearMonthDay:
		prefix := era + " "

		return func(v timeReader) string {
			return prefix + yearDigits(v) + separator + month(v) + separator + dayDigits(v)
		}
	}
}

func fmtEraYearMonthDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t timeReader) string {
		return dayDigits(t) + "/" + monthDigits(t) + "/" + year(t)
	}
}
