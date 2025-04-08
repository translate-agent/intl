package intl

import (
	"golang.org/x/text/language"
)

//nolint:cyclop
func fmtEraYearDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearGregorian(locale, digits, opts.Year)
	dayName := unitName(locale).Day

	const (
		eraYearDay = iota
		eraDayYear
	)

	layout := eraYearDay
	prefix := ""
	middle := " " + era + " (" + dayName + ": "
	suffix := ")"

	switch lang {
	case be, ru:
		middle = " г. " + era + " (" + dayName + ": "
	case cv:
		middle = " ҫ. " + era + " (" + dayName + ": "
	case kk:
		prefix = era + " "
		middle = " ж. (" + dayName + ": "
	case ky:
		prefix = era + " "
		middle = "-ж. (" + dayName + ": "
	case hy:
		prefix = era + " "
		middle = " թ. (" + dayName + ": "
	case tt:
		prefix = era + " "
		middle = " ел (" + dayName + ": "
	case sah:
		middle = " с. " + era + " (" + dayName + ": "
	case lt:
		opts.Day = Day2Digit
		middle = " m. " + era + " (" + dayName + ": "
	case bg, cy, mk:
		middle = " " + era + " (" + dayName + ": "
	case bs:
		if script != cyrl {
			suffix = ".)"
		}
	case agq, ak, as, asa, az, bas, bem, bez, bgc, bho, bm, bo, ce, cgg, ckb, csw, dav, dje, doi, dua, dyo, dz, ebu, eo,
		eu, ewo, fur, fy, gaa, gsw, gu, guz, gv, ha, hu, ig, jgo, jmc, kab, kam, kde, khq, ki, kl, kln, kn, ksb, ksf, ksh, ku,
		kw, lag, lg, lij, lkt, lmo, ln, lo, lrc, lu, luo, luy, lv, mas, mer, mfe, mg, mgh, mgo, ml, mn, mni, mr, mt, mua, my,
		naq, nd, nds, ne, nmg, nnh, nqo, nso, nus, nyn, oc, om, os, pa, pcm, prg, ps, qu, raj, rn, rof, rw, rwk, saq, sat,
		sbp, seh, ses, sg, shi, si, sn, st, szl, ta, te, teo, tk, tn, tok, tr, twq, tzm, vai, vmw, vun, wae, xog, yav, yi, yo,
		za, zgh, zu:
		prefix = era + " "
		middle = " (" + dayName + ": "
	case uz:
		if script != arab {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case brx:
		prefix = era
		middle = " (" + dayName + ": "
	case cs, da, dsb, fo, hr, hsb, ie, nb, nn, no, sk, sl:
		suffix = ".)"
	case ff:
		if script != adlm {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case hi:
		if script == latn {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case ks:
		if script == deva {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case sd:
		if script != deva {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case ja, yue, zh:
		prefix = era
		middle = " (" + dayName + ": "
		suffix = dayName + ")"
	case ko:
		prefix = era + " "
		middle = " (" + dayName + ": "
		suffix = dayName + ")"
	case ii:
		prefix = era + " "
		middle = " (" + dayName + ": "
		suffix = "ꑍ)"
	case kxv:
		if script == deva || script == orya || script == telu {
			prefix = ""
			middle = " " + era + " (" + dayName + ": "
		}
	case se:
		if region != regionFI {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case kok:
		if script != latn {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // eraYearDay
		return func(t timeReader) string {
			return prefix + year(t) + middle + dayDigits(t) + suffix
		}
	case eraDayYear:
		return func(t timeReader) string {
			return prefix + dayDigits(t) + middle + year(t) + suffix
		}
	}
}

func fmtEraYearDayPersian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigits(digits, opts.Year)
	dayName := unitName(locale).Day

	prefix := ""
	middle := " (" + dayName + ": "
	suffix := ")"

	switch lang {
	case uz:
		if region == regionAF {
			era := fmtEra(locale, opts.Era)
			prefix = era + " "
		}
	case fa:
		era := fmtEra(locale, opts.Era)
		middle = " " + era + " (" + dayName + ": "
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v timeReader) string {
		return prefix + year(yearDigits(v)) + middle + dayDigits(v) + suffix
	}
}

func fmtEraYearDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	dayDigits := convertDayDigits(digits, opts.Day)
	dayName := unitName(locale).Day
	middle, suffix := " ("+dayName+": ", ")"

	return func(t timeReader) string {
		return year(t) + middle + dayDigits(t) + suffix
	}
}
