package intl

import (
	"time"

	"golang.org/x/text/language"
)

type EraYearMonthDay int

const (
	eraYearMonthDay EraYearMonthDay = iota
	eraMonthDayYear
	eraDayMonthYear
	dayMonthEraYear
)

//nolint:cyclop,gocognit
func fmtEraYearMonthDayGregorian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, script, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	month, day := Month2Digit, Day2Digit
	layout := eraYearMonthDay
	prefix := era + " "
	suffix := ""
	separator := "-"

	switch lang {
	case en:
		switch region {
		default:
			month, day = opts.Month, opts.Day
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
			month, day = opts.Month, opts.Day
			layout = eraMonthDayYear
			separator = "/"
			prefix = ""
			suffix = " " + era
		case regionCH:
			month, day = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "."
			prefix = ""
			suffix = " " + era
		case regionGB:
			separator = "/"
			prefix = ""
			suffix = " " + era

			if script == shaw {
				month, day = opts.Month, opts.Day
				layout = eraMonthDayYear
			} else {
				layout = eraDayMonthYear
			}
		}
	case brx, lv, mni:
		layout = eraDayMonthYear
	case da, dsb, hsb, ie, ka, sq:
		month, day = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case mk:
		month = opts.Month
		day = opts.Day
		layout = eraDayMonthYear
		prefix = ""
		suffix = " г. " + era
		separator = "."
	case et, pl:
		day = opts.Day
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
		day = opts.Day
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
		month, day = opts.Month, opts.Day
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
		month, day = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case ga, it, kea, pt, sc, syr, vec:
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case ceb, chr, blo, fil, ml, ne, or, ps, sd, so, ti, xh, zu:
		month, day = opts.Month, opts.Day
		layout = eraMonthDayYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case cy:
		month, day = opts.Month, opts.Day
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
			month, day = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = ". "
			prefix = ""
			suffix = ". " + era
		}
	case ff:
		if script == adlm {
			month, day = opts.Month, opts.Day
			layout = eraDayMonthYear
			prefix = ""
			suffix = " " + era
		}
	case ks:
		if script != deva {
			month, day = opts.Month, opts.Day
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
			fmtMonth = fmtMonthName(locale.String(), "format", "abbreviated")
			month, day = opts.Month, opts.Day
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
		month, day = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = " " + era
	case hr:
		month, day = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = ". " + era
	case hi:
		month, day = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"

		if script == latn {
			prefix = ""
			suffix = " " + era
		}
	case zh:
		if script == hant {
			month, day = opts.Month, opts.Day
			separator = "/"
		}
	case kxv:
		if script != deva && script != orya && script != telu {
			month, day = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "/"
		}
	case ja:
		month, day = opts.Month, opts.Day
		separator = "/"
		prefix = era
	case ko, my:
		month, day = opts.Month, opts.Day
		separator = "/"
	case mr, qu:
		month, day = opts.Month, opts.Day
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
		month, day = opts.Month, opts.Day
		layout = dayMonthEraYear
		separator = "/"
	case pa:
		if script != arab {
			month, day = opts.Month, opts.Day
			layout = dayMonthEraYear
			separator = "/"
		}
	}

	switch layout {
	default: // eraYearMonthDay
		return func(y int, m time.Month, d int) string {
			return prefix + fmtYear(y, opts.Year) + separator + fmtMonth(m, month) + separator + fmtDay(d, day) + suffix
		}
	case eraMonthDayYear:
		return func(y int, m time.Month, d int) string {
			return prefix + fmtMonth(m, month) + separator + fmtDay(d, day) + separator + fmtYear(y, opts.Year) + suffix
		}
	case eraDayMonthYear:
		return func(y int, m time.Month, d int) string {
			return prefix + fmtDay(d, day) + separator + fmtMonth(m, month) + separator + fmtYear(y, opts.Year) + suffix
		}
	case dayMonthEraYear:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, day) + separator + fmtMonth(m, month) + separator + era + " " + fmtYear(y, opts.Year)
		}
	}
}

func fmtEraYearMonthDayPersian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, _, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)
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

		if opts.Era != EraNarrow {
			separator = "-"
		}
	case fa:
		suffix = " " + era
	}

	switch layout {
	default: // eraMonthDayYear
		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day) + separator + fmtYear(y, opts.Year) + suffix
		}
	case eraYearMonthDay:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + separator + fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day)
		}
	}
}

func fmtEraYearMonthDayBuddhist(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	return func(y int, m time.Month, d int) string {
		return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + era + " " + fmtYear(y, opts.Year)
	}
}
