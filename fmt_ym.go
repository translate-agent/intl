package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtYearMonthGregorian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	var month func(time.Month) string

	lang, script, region := locale.Raw()
	year := fmtYear(digits, opts.Year)

	const (
		layoutYearMonth = iota
		layoutMonthYear
	)

	layout := layoutYearMonth
	prefix := ""
	middle := "-"
	suffix := ""

	switch lang {
	default:
		opts.Month = Month2Digit
	case af, as, ia, jv, mi, rm, tg, wo:
		opts.Month = Month2Digit
		layout = layoutMonthYear
	case en:
		switch region {
		default:
			layout = layoutMonthYear
			middle = "/"
		case region001, region150, regionAE, regionAG, regionAI, regionAT, regionAU, regionBB, regionBE, regionBM, regionBS,
			regionBW, regionBZ, regionCC, regionCK, regionCM, regionCX, regionCY, regionDE, regionDG, regionDK, regionDM,
			regionER, regionFI, regionFJ, regionFK, regionFM, regionGB, regionGD, regionGG, regionGH, regionGI, regionGM,
			regionGY, regionHK, regionID, regionIE, regionIL, regionIM, regionIN, regionIO, regionJE, regionJM, regionKE,
			regionKI, regionKN, regionKY, regionLC, regionLR, regionLS, regionMG, regionMO, regionMS, regionMT, regionMU,
			regionMV, regionMW, regionMY, regionNA, regionNF, regionNG, regionNL, regionNR, regionNU, regionNZ, regionPG,
			regionPK, regionPN, regionPW, regionRW, regionSB, regionSC, regionSD, regionSG, regionSH, regionSI, regionSL,
			regionSS, regionSX, regionSZ, regionTC, regionTK, regionTO, regionTT, regionTV, regionTZ, regionUG, regionVC,
			regionVG, regionVU, regionWS, regionZA, regionZM, regionZW:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"

			if script != shaw {
				opts.Month = Month2Digit
			}
		case regionCA, regionSE:
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			opts.Month = Month2Digit
		case regionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			opts.Month = Month2Digit
			layout = layoutMonthYear
			middle = "."
		}
	case agq, ak, am, asa, ast, bas, bem, bez, blo, bm, brx, ca, ceb, cgg, chr, ckb, cs, cy, dav, dje, doi, dua, dyo, ebu,
		ee, el, ewo, fil, fur, gd, gl, guz, ha, haw, id, ig, jmc, kaa, kab, kam, kde, khq, ki, kln, km, ksb, ksf, kxv, lag,
		lg, ln, lo, lu, luo, luy, mai, mas, mer, mfe, mg, mgh, mhn, mni, mua, naq, nd, nmg, nus, nyn, om, pcm, rn, rof, rwk,
		sa, saq, sbp, ses, sg, shi, sk, sl, so, su, sw, teo, twq, tzm, ur, vai, vun, xh, xnr, xog, yav, yo, zgh:
		layout = layoutMonthYear
		middle = "/"
	case pa:
		if script == arab {
			// year=numeric,month=numeric,out=۲۰۲۴-۰۱
			// year=numeric,month=2-digit,out=۲۰۲۴-۰۱
			// year=2-digit,month=numeric,out=۲۴-۰۱
			// year=2-digit,month=2-digit,out=۲۴-۰۱
			opts.Month = Month2Digit
		} else {
			layout = layoutMonthYear
			middle = "/"
		}
	case ks:
		if script == deva {
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			opts.Month = Month2Digit
		} else {
			layout = layoutMonthYear
			middle = "/"
		}
	case hi:
		layout = layoutMonthYear
		middle = "/"

		if script == latn {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			opts.Month = Month2Digit
		}
	case ar:
		layout = layoutMonthYear
		middle = "\u200f/"
	case az, cv, fo, hy, kk, ku, os, pl, ro, ru, tk, tt, uk:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
	case uz:
		opts.Month = Month2Digit
		layout = layoutMonthYear

		if script == cyrl {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			middle = "/"
		} else {
			middle = "."
		}
	case be, da, dsb, et, hsb, ie, ka, lb, nb, nn, no, smn, sq:
		layout = layoutMonthYear
		middle = "."
	case bg:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
		suffix = " г."
	case mk:
		layout = layoutMonthYear
		middle = "."
		suffix = " г."
	case bn, ccp, gu, kn, mr, or, ta, te, to:
		layout = layoutMonthYear

		if opts.Month.numeric() {
			middle = "/"
		}
	case br, ga, it, iu, kea, kgp, pt, sc, seh, syr, vec, yrl:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "/"
	case bs:
		layout = layoutMonthYear

		if script == cyrl {
			opts.Month = Month2Digit
			middle = "."
			suffix = "."

			break
		}

		middle = "/"

		if !opts.Month.numeric() {
			middle = ". "
			suffix = "."
		}

		if opts.Month.numeric() {
			opts.Month = Month2Digit
		} else {
			opts.Month = MonthNumeric
		}
	case de:
		layout = layoutMonthYear
		middle = "."

		if opts.Month.numeric() {
			middle = "/"
		}
	case dz, si: // noop
	case es:
		switch region {
		default:
			opts.Month = MonthNumeric
			layout = layoutMonthYear
			middle = "/"
		case regionAR:
			// year=numeric,month=numeric,out=1-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=1-24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear

			if !opts.Month.numeric() {
				middle = "/"
			}

			opts.Month = MonthNumeric
		case regionCL:
			// year=numeric,month=numeric,out=01-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01-24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear

			if opts.Month.numeric() {
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
				middle = "/"
			}
		case regionMX, regionUS:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"
		case regionPA, regionPR:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear
			middle = "/"

			if opts.Month.numeric() {
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
			}
		}
	case ti:
		opts.Month = MonthNumeric
		layout = layoutMonthYear
		middle = "/"
	case yue:
		if script == hans {
			// year=numeric,month=numeric,out=2024年1月
			// year=numeric,month=2-digit,out=2024年1月
			// year=2-digit,month=numeric,out=24年1月
			// year=2-digit,month=2-digit,out=24年1月
			opts.Month = MonthNumeric
			middle = "年"
			suffix = "月"
		} else {
			middle = "/"
		}
	case eu, ja:
		middle = "/"
	case fi, he:
		opts.Month = MonthNumeric
		layout = layoutMonthYear
		middle = "."
	case ff:
		layout = layoutMonthYear

		if script != adlm {
			middle = "/"
		}
	case fr:
		opts.Month = Month2Digit

		switch region {
		default:
			layout = layoutMonthYear
			middle = "/"
		case regionCA: // noop
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
		case regionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			layout = layoutMonthYear
			middle = "."
		}
	case nl:
		layout = layoutMonthYear

		if region == regionBE {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			middle = "/"
		}
	case fy, kok, ms, ug:
		layout = layoutMonthYear
	case gsw:
		if !opts.Month.numeric() {
			layout = layoutMonthYear
			middle = "."
		}
	case hr:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = ". "
		suffix = "."
	case hu:
		middle = ". "
		suffix = "."
	case is:
		layout = layoutMonthYear
		middle = ". "
	case kkj:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = " "
	case ko:
		opts.Month = MonthNumeric
		middle = ". "
		suffix = "."
	case lv:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
		suffix = "."
	case mn:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
		middle = " "
	case yi:
		if !opts.Month.numeric() {
			layout = layoutMonthYear
			middle = "/"
		}

		opts.Month = Month2Digit
	case sd:
		if script == deva {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"
		} else {
			opts.Month = Month2Digit
		}
	case se:
		opts.Month = Month2Digit

		if region == regionFI {
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			layout = layoutMonthYear
			middle = "."
		}
	case sr:
		layout = layoutMonthYear
		suffix = "."

		if opts.Month.numeric() {
			middle = ". "
		} else {
			middle = "."
		}
	case tr:
		layout = layoutMonthYear
		middle = "."

		if opts.Month.numeric() {
			middle = "/"
		}

		opts.Month = Month2Digit
	case vi:
		layout = layoutMonthYear

		if opts.Month.numeric() {
			middle = "/"
		} else {
			prefix = "tháng "
			middle = ", "
		}
	case zh:
		middle = "/"

		switch script {
		case hant:
			switch region {
			default:
				// year=numeric,month=numeric,out=2024/1
				// year=numeric,month=2-digit,out=2024/01
				// year=2-digit,month=numeric,out=24/1
				// year=2-digit,month=2-digit,out=24/01
			case regionHK, regionMO:
				// year=numeric,month=numeric,out=1/2024
				// year=numeric,month=2-digit,out=01/2024
				// year=2-digit,month=numeric,out=1/24
				// year=2-digit,month=2-digit,out=01/24
				layout = layoutMonthYear
			}
		case hans:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			if region == regionHK {
				layout = layoutMonthYear
				break
			}

			fallthrough
		default:
			if !opts.Month.numeric() {
				opts.Month = MonthNumeric
				middle = "年"
				suffix = "月"
			}
		}
	}

	if month == nil {
		month = fmtMonth(digits, opts.Month)
	}

	if layout == layoutMonthYear {
		return func(y int, m time.Month) string {
			return prefix + month(m) + middle + year(y) + suffix
		}
	}

	return func(y int, m time.Month) string {
		return prefix + year(y) + middle + month(m) + suffix
	}
}

func fmtYearMonthBuddhist(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	year := fmtYear(digits, opts.Year)

	if lang, _ := locale.Base(); lang == th {
		month := fmtMonth(digits, opts.Month)

		return func(y int, m time.Month) string {
			return month(m) + "/" + year(y)
		}
	}

	prefix := fmtEra(locale, EraNarrow) + " "
	month := fmtMonth(digits, Month2Digit)

	return func(y int, m time.Month) string {
		return prefix + year(y) + "-" + month(m)
	}
}

func fmtYearMonthPersian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	lang, _, region := locale.Raw()
	year := fmtYear(digits, opts.Year)
	month := fmtMonth(digits, Month2Digit)

	prefix := ""
	separator := "-"

	switch lang {
	case ckb: // ckb-IR
		// year=numeric,month=numeric,out=١٠/١٤٠٢
		// year=numeric,month=2-digit,out=١٠/١٤٠٢
		// year=2-digit,month=numeric,out=١٠/٠٢
		// year=2-digit,month=2-digit,out=١٠/٠٢
		return func(y int, m time.Month) string {
			return month(m) + "/" + year(y)
		}
	case fa:
		separator = "/"
	case ps:
		prefix = fmtEra(locale, EraNarrow) + " "
		separator = "/"
	case uz:
		if region == regionAF {
			// year=numeric,month=numeric,out=۱۴۰۲-۱۰
			// year=numeric,month=2-digit,out=۱۴۰۲-۱۰
			// year=2-digit,month=numeric,out=۰۲-۱۰
			// year=2-digit,month=2-digit,out=۰۲-۱۰
			break
		}

		fallthrough
	default:
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(y int, m time.Month) string {
		return prefix + year(y) + separator + month(m)
	}
}
