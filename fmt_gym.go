package intl

import (
	"golang.org/x/text/language"
)

//nolint:cyclop
func fmtEraYearMonthGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearGregorian(locale, digits, opts.Year)
	monthName := unitName(locale).Month

	const (
		// eraYearMonth includes "era year month" and "year month era".
		eraYearMonth = iota
		// eraMonthYear includes "era month year" and "month year era".
		eraMonthYear
	)

	layout := eraMonthYear
	prefix := ""
	middle := " "
	suffix := " " + era

	switch lang {
	case az, qu, te, tk, tr:
		prefix = era + " "
		suffix = ""
	case be, ru:
		suffix = " г. " + era
	case bg:
		middle = "."
		suffix = " " + era

		if opts.Month.numeric() {
			opts.Month = Month2Digit
		}
	case cy, mk:
		middle = " "
		suffix = " " + era
	case cv:
		suffix = " ҫ. " + era
	case hi:
		if script != latn {
			middle = " " + era + " "
			suffix = ""

			break
		}

		fallthrough
	case agq, ak, as, asa, bas, bem, bez, bgc, bho, bm, bo, brx, ce, cgg, ckb, csw, dav, dje, doi, dua, dyo, ebu, eo, ewo,
		fur, gaa, gsw, guz, gv, ha, hu, ii, jgo, jmc, kab, kam, kde, khq, ki, kl, kln, kn, ko, ksb, ksf, ksh, kw, lag, lg,
		lij, lkt, lmo, ln, lrc, lu, luo, luy, lv, mas, mg, mer, mfe, mgh, ml, mgo, mni, mua, my, naq, nd, nds, ne, nmg, nnh,
		nqo, nso, nus, nyn, oc, os, pcm, prg, ps, raj, rn, rof, rw, rwk, sah, saq, sat, sbp, seh, ses, sg, shi, si, sn, st,
		szl, ta, teo, tn, tok, twq, tzm, vai, vmw, vun, wae, xog, yav, yi, yo, za, zgh, zu:
		layout = eraYearMonth
		prefix = era + " "
		suffix = ""
	case se:
		if region != regionFI {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case sd:
		if script != deva {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case ks:
		if script == deva {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case ig, kxv, mai, mr, sa, xnr:
		middle = " " + era + " "
		suffix = ""
	case pa:
		if script == arab {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""

			break
		}

		fallthrough
	case gu, lo, uz:
		middle = ", " + era + " "
		suffix = ""
	case kgp, wo:
		middle = ", "
		suffix = " " + era
	case tt:
		layout = eraYearMonth
		prefix = era + " "
		middle = " ел, "
		suffix = ""
	case es:
		if region != regionCO {
			break
		}

		fallthrough
	case gl, pt:
		middle = " de "
	case yue, zh:
		opts.Month = MonthNumeric
		layout = eraYearMonth
		prefix = era
		middle = ""
		suffix = monthName
	case dz:
		layout = eraYearMonth
		prefix = era + " "
		middle = " སྤྱི་ཟླ་"
		suffix = ""
	case ja:
		layout = eraYearMonth
		prefix = era
		middle = ""
		opts.Month = MonthNumeric
		suffix = "月"
	case eu:
		layout = eraYearMonth
		prefix = era + " "
		middle = ". urteko "
		suffix = ""
	case ff:
		if script != adlm {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case hy:
		layout = eraYearMonth
		prefix = era + " "
		middle = " թ. "
		suffix = ""
	case kk:
		layout = eraYearMonth
		prefix = era + " "
		middle = " ж. "
		suffix = ""
	case ka:
		middle = ". "
	case ku:
		prefix = era + " "
		middle = "a "
		suffix = "an"
	case ky:
		layout = eraYearMonth
		prefix = era + " "
		middle = "-ж. "
		suffix = ""
	case lt:
		opts.Month = Month2Digit
		layout = eraYearMonth
		middle = "-"
	case mn:
		layout = eraYearMonth
		prefix = era + " "
		middle = " оны "
		suffix = ""
	case sl:
		month = fmtMonthName(locale.String(), "format", "abbreviated")
	case ug:
		layout = eraYearMonth
	case uk:
		suffix = " р. " + era
	case kok:
		if script != latn {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	switch layout {
	default: // eraYearMonth
		return func(t timeReader) string {
			return prefix + year(t) + middle + month(t) + suffix
		}
	case eraMonthYear:
		return func(t timeReader) string {
			return prefix + month(t) + middle + year(t) + suffix
		}
	}
}

func fmtEraYearMonthPersian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		// eraYearMonth includes "era year month" and "year month era".
		eraYearMonth = iota
		// eraMonthYear includes "era month year" and "month year era".
		eraMonthYear
	)

	layout := eraYearMonth
	prefix := era + " "
	middle := " "
	suffix := ""

	switch lang {
	case fa:
		layout = eraMonthYear
		prefix = ""
		middle = " "
		suffix = " " + era
	case ckb, uz:
		if region != regionAF {
			prefix = ""
		}
	case lrc, mzn, ps:
		prefix = ""
	}

	month := convertMonthDigits(digits, opts.Month)

	switch layout {
	default: // eraYearMonth
		return func(v timeReader) string {
			return prefix + year(yearDigits(v)) + middle + month(v) + suffix
		}
	case eraMonthYear:
		return func(v timeReader) string {
			return prefix + month(v) + middle + year(yearDigits(v)) + suffix
		}
	}
}

func fmtEraYearMonthBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)

	return func(t timeReader) string {
		return monthDigits(t) + " " + year(t)
	}
}
