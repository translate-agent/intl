package intl

import "golang.org/x/text/language"

func fmtEraYearGregorian(locale language.Tag, digits digits, opts Options) func(y int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearGregorian(locale)

	switch lang {
	default:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + era
		}
	case agq, ak, as, asa, az, bas, bem, bez, bgc, bho, bm, bo, ce, cgg, ckb, csw, dav, dje, doi, dua, dz, dyo, ebu, eo,
		eu, ewo, fur, fy, gsw, gu, guz, gv, ha, hu, ie, ig, ii, jmc, jgo, kab, kam, kde, khq, ki, kl, kln, kn, ko, kok,
		ksb, ksf, ksh, ku, kw, lag, lg, lij, lkt, lmo, ln, lo, lrc, lv, lu, luo, luy, mas, mer, mfe, mg, mgh, mgo, ml, mn,
		mni, mr, mt, mua, my, naq, nd, nds, ne, nmg, nnh, nqo, nus, nyn, om, os, pa, pcm, prg, ps, qu, raj, rn, rof, rw,
		rwk, saq, sbp, seh, ses, sg, shi, si, sn, szl, ta, te, teo, tk, tok, tr, twq, tzm, uz, vai, vmw, vun, wae, xog,
		yav, yi, yo, za, zgh, zu, sat:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year))
		}
	case kxv:
		if script == deva || script == orya || script == telu {
			return func(y int) string {
				return era + " " + layoutYear(fmtYear(y, opts.Year))
			}
		}

		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + era
		}
	case hi:
		if script == latn {
			return func(y int) string {
				return era + " " + layoutYear(fmtYear(y, opts.Year))
			}
		}

		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + era
		}
	case ks:
		if script == deva {
			return func(y int) string {
				return era + " " + layoutYear(fmtYear(y, opts.Year))
			}
		}

		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + era
		}
	case oc:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year))
		}
	case sd:
		if script == deva {
			return func(y int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + era
			}
		}

		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year))
		}
	case ff:
		if script == adlm {
			return func(y int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + era
			}
		}

		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year))
		}
	case se:
		if region == regionFI {
			return func(y int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + era
			}
		}

		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year))
		}
	case brx, ja, yue, zh:
		return func(y int) string {
			return era + layoutYear(fmtYear(y, opts.Year))
		}
	case be, ru:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " г. " + era
		}
	case bg:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + era
		}
	case cv:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " ҫ. " + era
		}
	case kk:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + " ж."
		}
	case hy:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + " թ."
		}
	case ky:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + "-ж."
		}
	case lt:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " m. " + era
		}
	case tt:
		return func(y int) string {
			return era + " " + layoutYear(fmtYear(y, opts.Year)) + " ел"
		}
	case sah:
		return func(y int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " с. " + era
		}
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
