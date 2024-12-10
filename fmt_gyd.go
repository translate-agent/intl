package intl

import "golang.org/x/text/language"

type EraYearDay int

const (
	eraYearDay EraYearDay = iota
	eraDayYear
)

//nolint:cyclop
func fmtEraYearDayGregorian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	prefix := ""
	middle := " " + era + " (" + name + ": "
	suffix := ")"
	layout := eraYearDay

	switch lang {
	case be, ru:
		prefix = ""
		middle = " г. " + era + " (" + name + ": "
	case cv:
		prefix = ""
		middle = " ҫ. " + era + " (" + name + ": "
	case kk:
		prefix = era + " "
		middle = " ж. (" + name + ": "
	case ky:
		prefix = era + " "
		middle = "-ж. (" + name + ": "
	case hy:
		prefix = era + " "
		middle = " թ. (" + name + ": "
	case tt:
		prefix = era + " "
		middle = " ел (" + name + ": "
	case sah:
		middle = " с. " + era + " (" + name + ": "
	case lt:
		opts.Day = Day2Digit
		middle = " m. " + era + " (" + name + ": "
	case bg, cy:
		prefix = ""
		middle = " " + era + " (" + name + ": "
	case bs:
		prefix = ""
		middle = " " + era + " (" + name + ": "

		if script != cyrl {
			suffix = ".)"
		}
	case agq, ak, as, asa, az, bas, bem, bez, bgc, bho, bm, bo, ce, cgg, ckb, csw, dav, dje, doi, dua, dyo, dz, ebu, eo,
		eu, ewo, fur, fy, gsw, gu, guz, gv, ha, hu, ie, ig, ii, jgo, jmc, kab, kam, kde, khq, ki, kl, kln, kn, kok, ksb, ksf,
		ksh, ku, kw, lag, lg, lij, lkt, lmo, ln, lo, lrc, lu, luo, luy, lv, mas, mer, mfe, mg, mgh, mgo, ml, mn, mni, mr, mt,
		mua, my, naq, nd, nds, ne, nmg, nnh, nqo, nus, nyn, oc, om, os, pa, pcm, prg, ps, qu, raj, rn, rof, rw, rwk, saq, sat,
		sbp, seh, ses, sg, shi, si, sn, szl, ta, te, teo, tk, tok, tr, twq, tzm, vai, vmw, vun, wae, xog, yav, yi, yo,
		za, zgh, zu:
		prefix = era + " "
		middle = " (" + name + ": "
	case uz:
		if script == arab {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case brx:
		prefix = era
		middle = " (" + name + ": "
	case cs, da, dsb, fo, hr, hsb, nb, nn, no, sk, sl:
		prefix = ""
		middle = " " + era + " (" + name + ": "
		suffix = ".)"
	case ff:
		if script == adlm {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case hi:
		if script != latn {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case ks:
		if script != deva {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case sd:
		if script == deva {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case ja, yue, zh:
		prefix = era
		middle = " (" + name + ": "
		suffix = name + ")"
	case ko:
		prefix = era + " "
		middle = " (" + name + ": "
		suffix = name + ")"
	case kxv:
		if script != deva && script != orya && script != telu {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	case se:
		if region == regionFI {
			break
		}

		prefix = era + " "
		middle = " (" + name + ": "
	}

	switch layout {
	default: // eraYearDay
		return func(y, d int) string {
			return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d, opts.Day) + suffix
		}
	case eraDayYear:
		return func(y, d int) string {
			return prefix + fmtDay(d, opts.Day) + middle + layoutYear(fmtYear(y, opts.Year)) + suffix
		}
	}
}

func fmtEraYearDayPersian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	prefix, middle, suffix := era+" ", " ("+name+": ", ")"

	if lang == fa {
		prefix = ""
		middle = " " + era + " (" + name + ": "
		suffix = ")"
	}

	return func(y, d int) string {
		return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d, opts.Day) + suffix
	}
}

func fmtEraYearDayBuddhist(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	prefix, middle, suffix := era+" ", " ("+name+": ", ")"

	return func(y, d int) string {
		return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d, opts.Day) + suffix
	}
}
