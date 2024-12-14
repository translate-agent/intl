package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutDayMonth

	prefix := era + " "
	suffix := ""
	separator := "/"

	switch lang {
	case af, as, ia, ky, mi, rm, tg, wo:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "-"
	case sd:
		if script == deva {
			layout = layoutMonthDay
			break
		}

		fallthrough
	case bgc, bho, bo, ce, ckb, csw, eo, gv, kl, ksh, kw, lij, lkt, lmo, mgo, mt, nds, nnh, ne, nqo, oc, prg, ps, qu, raj,
		rw, sah, sat, sn, szl, tok, vmw, yi, za, zu:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "-"
	case lt:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
		}

		fallthrough
	case dz, si:
		layout = layoutMonthDay
		separator = "-"
	case nl:
		if region == regionBE {
			break
		}

		fallthrough
	case fy, kok, ug:
		separator = "-"
	case or:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			layout = layoutMonthDay
			break
		}

		separator = "-"
	case ms:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "-"
		}
	case se:
		if region == regionFI {
			break
		}

		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "-"
	case kn, mr, vi:
		if opts.Month != MonthNumeric || opts.Day != DayNumeric {
			separator = "-"
		}
	case ti:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
	case ff:
		if script == adlm {
			separator = "-"
		}
	case bn, ccp, gu, ta, te:
		if opts.Month == Month2Digit || opts.Day == Day2Digit {
			separator = "-"
			break
		}
	case az, cv, fo, hy, kk, ku, os, tk, tt, uk:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "."
	case sq:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = "."
	case bg:
		opts.Month = Month2Digit
		prefix = era + " "
		separator = "."
	case cy:
		prefix = era + " "
	case pl:
		opts.Month = Month2Digit
		separator = "."
	case be, da, et, he, ie, jgo, ka:
		separator = "."
	case mk:
		opts.Day = DayNumeric
		prefix = era + " "
		separator = "."
	case nb, nn, no:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = "."
		suffix = "."
	case lv:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "."
		suffix = "."
	case sr:
		separator = "."
		suffix = "."

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = ". "
		}
	case hr:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cs, sk, sl:
		separator = ". "
		suffix = "."
	case ro, ru:
		separator = "."

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case de, dsb, fi, gsw, hsb, lb, is, smn:
		separator = "."
		suffix = "."
	case hu, ko:
		layout = layoutMonthDay
		separator = ". "
		suffix = "."
	case wae:
		fmtMonth = fmtMonthName(locale.String(), "format", "abbreviated")
		separator = ". "
	case bs:
		suffix = "."

		if script == cyrl {
			separator = "."
			opts.Month = Month2Digit
			opts.Day = Day2Digit

			break
		}

		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = ". "
	case om:
		if opts.Month == Month2Digit || opts.Day == Day2Digit {
			break
		}

		layout = layoutMonthDay
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "-"
	case ks:
		if script == deva {
			layout = layoutMonthDay
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			separator = "-"

			break
		}

		fallthrough
	case ak, asa, bem, blo, bez, brx, ceb, cgg, chr, dav, ebu, ee, eu, fil, guz, ha, kam, kde, kln, teo, vai, ja, jmc, ki,
		ksb, lag, lg, luo, luy, mas, mer, naq, nd, nyn, rof, rwk, saq, sbp, so, tzm, vun, xh, xog, yue:
		layout = layoutMonthDay
	case mn:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")
		opts.Day = Day2Digit
		layout = layoutMonthDay
	case zh:
		if region == regionHK || region == regionMO {
			break
		}

		layout = layoutMonthDay

		if region == regionSG {
			separator = "-"
		}
	case fr:
		if region == regionCA {
			layout = layoutMonthDay
			separator = "-"

			if opts.Month == Month2Digit || opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			break
		}

		if region == regionCH {
			separator = "."

			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				suffix = "."
			}

			break
		}

		fallthrough
	case br, ga, it, jv, kkj, sc, syr, vec, uz:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case pcm:
		separator = " /"
	case sv:
		if region == regionFI {
			separator = "."
		}

		if opts.Month == Month2Digit && opts.Day == DayNumeric {
			opts.Month = MonthNumeric
		}
	case kea, pt:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case hi:
		if script != latn {
			break
		}

		prefix = ""
		suffix = " " + era

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case ar:
		separator = "\u200f/"
	case lrc:
		if region == regionIQ {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutMonthDay
			separator = "-"
		}
	case en:
		prefix = ""
		suffix = " " + era

		switch region {
		case regionUS, regionAS, regionBI, regionPH, regionPR, regionUM, regionVI:
			layout = layoutMonthDay
			goto breakEN
		case regionAU, regionBE, regionIE, regionNZ, regionZW:
			goto breakEN
		case regionGU, regionMH, regionMP, regionZZ:
			layout = layoutMonthDay
			goto breakEN
		case regionCA:
			layout = layoutMonthDay
			separator = "-"
		case regionCH:
			separator = "."
		case regionZA:
			if opts.Month != Month2Digit || opts.Day != Day2Digit {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}

		if script == shaw || script == dsrt {
			layout = layoutMonthDay
			break
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	breakEN:
		break
	case es:
		switch region {
		case regionUS, regionMX:
			goto breakES
		case regionCL:
			if opts.Month == Month2Digit {
				separator = "/"
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				separator = "-"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			goto breakES
		case regionPA, regionPR:
			if opts.Month == MonthNumeric {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit

				goto breakES
			}

			if opts.Month == Month2Digit {
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			goto breakES
		}

		opts.Month = MonthNumeric
		opts.Day = DayNumeric
	breakES:
		break
	case ii:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "ꆪ-"
		suffix = "ꑍ"
	}

	if layout == layoutDayMonth {
		return func(m time.Month, d int) string {
			return prefix + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month) + suffix
		}
	}

	return func(m time.Month, d int) string {
		return prefix + fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day) + suffix
	}
}

func fmtEraMonthDayPersian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)
	prefix := era + " "
	separator := "-"

	switch lang {
	case fa, ps:
		prefix = era + " "
		separator = "/"
	}

	return func(m time.Month, d int) string {
		return prefix + fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day)
	}
}

func fmtEraMonthDayBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	return func(m time.Month, d int) string {
		return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
	}
}
