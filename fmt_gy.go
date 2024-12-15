package intl

import "golang.org/x/text/language"

func fmtEraYearGregorian(locale language.Tag, digits digits, opts Options) func(y int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearGregorian(locale)

	prefix := ""
	suffix := " " + era

	switch lang {
	case kok:
		if script == latn {
			break
		}

		fallthrough
	case agq, ak, as, asa, az, bas, bem, bez, bgc, bho, bm, bo, ce, cgg, ckb, csw, dav, dje, doi, dua, dz, dyo, ebu, eo,
		eu, ewo, fur, fy, gaa, gsw, gu, guz, gv, ha, hu, ig, ii, jmc, jgo, kab, kam, kde, khq, ki, kl, kln, kn, ko, ksb, ksf,
		ksh, ku, kw, lag, lg, lij, lkt, lmo, ln, lo, lrc, lv, lu, luo, luy, mas, mer, mfe, mg, mgh, mgo, ml, mn, mni, mr, mt,
		mua, my, naq, nd, nds, ne, nmg, nnh, nqo, nso, nus, nyn, oc, om, os, pa, pcm, prg, ps, qu, raj, rn, rof, rw, rwk, saq,
		sat, sbp, seh, ses, sg, shi, si, sn, st, szl, ta, te, teo, tk, tn, tok, tr, twq, tzm, uz, vai, vmw, vun, wae, xog,
		yav, yi, yo, za, zgh, zu:
		prefix = era + " "
		suffix = ""
	case ks:
		if script == deva {
			prefix = era + " "
			suffix = ""
		}
	case hi:
		if script == latn {
			prefix = era + " "
			suffix = ""
		}
	case sd:
		if script != deva {
			prefix = era + " "
			suffix = ""
		}
	case ff:
		if script != adlm {
			prefix = era + " "
			suffix = ""
		}
	case se:
		if region != regionFI {
			prefix = era + " "
			suffix = ""
		}
	case be, ru:
		suffix = " г. " + era
	case bg, cy, mk:
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
		prefix = era + " "
		suffix = " ел"
	case sah:
		suffix = " с. " + era
	case ja, brx, yue, zh:
		prefix = era
		suffix = ""
	}

	return func(y int) string { return prefix + layoutYear(fmtYear(y, opts.Year)) + suffix }
}

func fmtEraYearPersian(locale language.Tag, digits digits, opts Options) func(y int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)

	prefix := ""
	suffix := " " + era

	switch lang {
	case ckb, lrc, mzn, ps, uz:
		prefix = era + " "
		suffix = ""
	case fa:
		suffix = " " + era
	}

	return func(y int) string {
		return prefix + fmtYear(y, opts.Year) + suffix
	}
}

func fmtEraYearBuddhist(locale language.Tag, digits digits, opts Options) func(y int) string {
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)

	return func(y int) string {
		return era + " " + fmtYear(y, opts.Year)
	}
}
