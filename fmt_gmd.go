package intl

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraMonthDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)

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
	case bgc, bho, bo, ce, ckb, csw, eo, gaa, gv, kl, ksh, kw, lij, lkt, lmo, mgo, mt, nds, nnh, ne, nqo, nso, oc, prg,
		ps, qu, raj, rw, sah, sat, sn, st, szl, tn, tok, vmw, yi, za, zu:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "-"
	case lt:
		if opts.Month.numeric() && opts.Day.numeric() {
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
	case fy, ug:
		separator = "-"
	case or:
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDay
			break
		}

		separator = "-"
	case ms:
		if opts.Month.numeric() && opts.Day.numeric() {
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
		if !opts.Month.numeric() || !opts.Day.numeric() {
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
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
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

		if opts.Month.numeric() && opts.Day.numeric() {
			separator = ". "
		}
	case hr:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cs, sk, sl:
		separator = ". "
		suffix = "."
	case ro, ru:
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
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
		month = fmtMonthName(locale.String(), "format", "abbreviated")
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
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
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
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
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

			if opts.Month.twoDigit() || opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			break
		}

		if region == regionCH {
			separator = "."

			if opts.Month.numeric() && opts.Day.numeric() {
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

		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
		}
	case kea, pt:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case hi:
		if script != latn {
			break
		}

		prefix = ""
		suffix = " " + era

		if opts.Month.numeric() && opts.Day.numeric() {
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
			if !opts.Month.twoDigit() || !opts.Day.twoDigit() {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}

		if script == shaw || script == dsrt {
			layout = layoutMonthDay
			break
		}

		if opts.Month.numeric() && opts.Day.numeric() {
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
			if opts.Month.twoDigit() {
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
			if opts.Month.numeric() {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit

				goto breakES
			}

			if opts.Month.twoDigit() {
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
	case kok:
		if script != latn {
			separator = "-"
		}
	case kaa, mhn:
		layout = layoutMonthDay
		prefix = ""
		suffix = " " + era
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t time.Time) string {
			return prefix + dayDigits(t) + separator + month(t) + suffix
		}
	}

	return func(t time.Time) string {
		return prefix + month(t) + separator + dayDigits(t) + suffix
	}
}

func fmtEraMonthDayPersian(locale language.Tag, digits digits, opts Options) fmtPersianFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	prefix := era + " "
	separator := "-"

	switch lang {
	case fa, ps:
		prefix = era + " "
		separator = "/"
	}

	month := convertMonthDigitsPersian(digits, opts.Month)
	dayDigits := convertDayDigitsPersian(digits, opts.Day)

	return func(v ptime.Time) string { return prefix + month(v) + separator + dayDigits(v) }
}

func fmtEraMonthDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	prefix := era + " "
	month := fmtMonthDayBuddhist(locale, digits, opts)

	return func(t time.Time) string { return prefix + month(t) }
}
