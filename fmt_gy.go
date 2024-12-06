package intl

import "golang.org/x/text/language"

func fmtEraYearGregorian(locale language.Tag, digits digits, opts Options) func(y int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearGregorian(locale)
	prefix, suffix := "", " "+era

	switch lang {
	case agq, ak, as, asa, az, bas, bem, bez, bgc, bho, bm, bo, brx, ce, cgg, ckb, csw, dav, dje, doi, dua, dz, dyo, ebu,
		eo, eu, ewo, fur, fy, gsw, gu, guz, gv, ha, hu, ie, ig, ii, ja, jmc, jgo, kab, kam, kde, khq, ki, kl, kln, kn, ko,
		kok, ksb, ksf, ksh, ku, kw, lag, lg, lij, lkt, lmo, ln, lo, lrc, lv, lu, luo, luy, mas, mer, mfe, mg, mgh, mgo, ml,
		mn, mni, mr, mt, mua, my, naq, nd, nds, ne, nmg, nnh, nqo, nus, nyn, oc, om, os, pa, pcm, prg, ps, qu, raj, rn, rof,
		rw, rwk, saq, sat, sbp, seh, ses, sg, shi, si, sn, szl, ta, te, teo, tk, tok, tr, twq, tzm, uz, vai, vmw, vun, wae,
		xog, yav, yi, yo, yue, za, zgh, zh, zu:
		prefix = era + " "
		suffix = ""
	case hi, ks, kxv:
		// hi-latn
		// ks-deva
		// kxv-deva, kxv-orya, kxv-telu
		if script == deva || script == orya || script == telu {
			prefix = era + " "
			suffix = ""

			break
		}
	case sd:
		if script == deva {
			break
		}

		prefix = era + " "
		suffix = ""
	case ff:
		if script == adlm {
			break
		}

		prefix = era + " "
		suffix = ""
	case se:
		if region == regionFI {
			break
		}

		prefix = era + " "
		suffix = ""
	case be, ru:
		suffix = " г. " + era
	case bg, cy:
		suffix = " " + era
	case cv:
		suffix = " ҫ. " + era
	case kk:
		prefix = era + " "
		suffix = " ж."
	case hy:
		prefix = era + " "
		suffix = " թ."
	case ky:
		prefix = era + " "
		suffix = "-ж."
	case lt:
		suffix = " m. " + era
	case tt:
		suffix = era + " "
		prefix = " ел"
	case sah:
		suffix = " с. " + era
	}

	return func(y int) string {
		return prefix + layoutYear(fmtYear(y, opts.Year)) + suffix
	}
}

func fmtEraYearPersian(locale language.Tag, digits digits, opts Options) func(y int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)

	switch lang {
	default:
		return func(y int) string {
			return fmtYear(y, opts.Year) + " " + era
		}
	case ckb, lrc, mzn, ps, uz:
		return func(y int) string {
			return era + " " + fmtYear(y, opts.Year)
		}
	}
}

func fmtEraYearBuddhist(locale language.Tag, digits digits, opts Options) func(y int) string {
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)

	return func(y int) string {
		return era + " " + fmtYear(y, opts.Year)
	}
}
