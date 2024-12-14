package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtYearMonthGregorian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	lang, script, region := locale.Raw()
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch lang {
	default:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}
	case af, as, ia, jv, mi, rm, tg, wo:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "-" + fmtYear(y, opts.Year)
		}
	case en:
		switch region {
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
			if script == shaw {
				break
			}

			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
			}
		case regionCA, regionSE:
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
			}
		case regionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
			}
		}

		fallthrough
	case agq, ak, am, asa, ast, bas, bem, bez, blo, bm, brx, ca, ceb, cgg, chr, ckb, cs, cy, dav, dje, doi, dua, dyo, ebu,
		ee, el, ewo, fil, fur, gd, gl, guz, ha, haw, id, ig, jmc, kab, kam, kde, khq, ki, kln, km, ksb, ksf, kxv, lag, lg,
		ln, lo, lu, luo, luy, mai, mas, mer, mfe, mg, mgh, mni, mua, naq, nd, nmg, nus, nyn, om, pcm, rn, rof, rwk, sa, saq,
		sbp, ses, sg, shi, sk, sl, so, su, sw, teo, twq, tzm, ur, vai, vun, xh, xnr, xog, yav, yo, zgh:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case pa:
		if script == arab {
			// year=numeric,month=numeric,out=۲۰۲۴-۰۱
			// year=numeric,month=2-digit,out=۲۰۲۴-۰۱
			// year=2-digit,month=numeric,out=۲۴-۰۱
			// year=2-digit,month=2-digit,out=۲۴-۰۱
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
			}
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case ks:
		if script == deva {
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
			}
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case hi:
		if script == latn {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			opts.Month = Month2Digit
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case ar:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "\u200f/" + fmtYear(y, opts.Year)
		}
	case az, cv, fo, hy, kk, ku, os, pl, ro, ru, tk, tt, uk:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
		}
	case uz:
		if script == cyrl {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
		}
	case be, da, dsb, et, hsb, ie, ka, lb, nb, nn, no, smn, sq:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year)
		}
	case bg:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " г."
		}
	case mk:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year) + " г."
		}
	case bn, ccp, gu, kn, mr, or, ta, te, to:
		separator := "-"
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year)
		}
	case br, ga, it, iu, kea, kgp, pt, sc, seh, syr, vec, yrl:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
		}
	case bs:
		if script == cyrl {
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + "."
			}
		}

		separator := ". "
		suffix := "."

		if opts.Month == MonthNumeric {
			separator = "/"
			suffix = ""
		}

		if opts.Month == MonthNumeric {
			opts.Month = Month2Digit
		} else {
			opts.Month = MonthNumeric
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year) + suffix
		}
	case de:
		separator := "."
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year)
		}
	case dz, si:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month)
		}
	case es:
		switch region {
		case regionAR:
			// year=numeric,month=numeric,out=1-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=1-24
			// year=2-digit,month=2-digit,out=1/24
			separator := "/"
			if opts.Month == MonthNumeric {
				separator = "-"
			}

			return func(y int, m time.Month) string {
				return fmtMonth(m, MonthNumeric) + separator + fmtYear(y, opts.Year)
			}
		case regionCL:
			// year=numeric,month=numeric,out=01-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01-24
			// year=2-digit,month=2-digit,out=1/24
			separator := "/"
			if opts.Month == MonthNumeric {
				separator = "-"
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
			}

			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year)
			}
		case regionMX, regionUS:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}
		case regionPA, regionPR:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=1/24
			if opts.Month == MonthNumeric {
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
			}

			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}
		}

		fallthrough
	case ti:
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "/" + fmtYear(y, opts.Year)
		}
	case yue:
		if script == hans {
			// year=numeric,month=numeric,out=2024年1月
			// year=numeric,month=2-digit,out=2024年1月
			// year=2-digit,month=numeric,out=24年1月
			// year=2-digit,month=2-digit,out=24年1月
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "年" + fmtMonth(m, MonthNumeric) + "月"
			}
		}

		fallthrough
	case eu, ja:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month)
		}
	case fi, he:
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "." + fmtYear(y, opts.Year)
		}
	case ff:
		if script == adlm {
			// year=numeric,month=numeric,out=𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=2-digit,out=𞥐𞥑-𞥒𞥐𞥒𞥔
			// year=2-digit,month=numeric,out=𞥑-𞥒𞥔
			// year=2-digit,month=2-digit,out=𞥐𞥑-𞥒𞥔
			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + "-" + fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case fr:
		switch region {
		default:
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
			}
		case regionCA:
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
			}
		case regionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
			}
		}
	case nl:
		if region == regionBE {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}
		}

		fallthrough
	case fy, kok, ms, ug:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "-" + fmtYear(y, opts.Year)
		}
	case gsw:
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month)
			}

			return fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year)
		}
	case hr:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + ". " + fmtYear(y, opts.Year) + "."
		}
	case hu:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case is:
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + ". " + fmtYear(y, opts.Year)
		}
	case kkj:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + " " + fmtYear(y, opts.Year)
		}
	case ko:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + ". " + fmtMonth(m, MonthNumeric) + "."
		}
	case lv:
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + "."
		}
	case mn:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")

		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + " " + fmtMonth(m, opts.Month)
		}
	case yi:
		return func(y int, m time.Month) string {
			ys := fmtYear(y, opts.Year)
			ms := fmtMonth(m, Month2Digit)

			if opts.Month == MonthNumeric {
				return ys + "-" + ms
			}

			return ms + "/" + ys
		}
	case sd:
		if script == deva {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return func(y int, m time.Month) string {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}

	case se:
		if region == regionFI {
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return func(y int, m time.Month) string {
				return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}
	case sr:
		separator := "."
		if opts.Month == MonthNumeric {
			separator = ". "
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year) + "."
		}
	case tr:
		separator := "."
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + separator + fmtYear(y, opts.Year)
		}
	case vi:
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}

			return "tháng " + fmtMonth(m, Month2Digit) + ", " + fmtYear(y, opts.Year)
		}
	case zh:
		switch script {
		case hant:
			switch region {
			default:
				// year=numeric,month=numeric,out=2024/1
				// year=numeric,month=2-digit,out=2024/01
				// year=2-digit,month=numeric,out=24/1
				// year=2-digit,month=2-digit,out=24/01
				return func(y int, m time.Month) string {
					return fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month)
				}
			case regionHK, regionMO:
				// year=numeric,month=numeric,out=1/2024
				// year=numeric,month=2-digit,out=01/2024
				// year=2-digit,month=numeric,out=1/24
				// year=2-digit,month=2-digit,out=01/24
				return func(y int, m time.Month) string {
					return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
				}
			}
		case hans:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			if region == regionHK {
				return func(y int, m time.Month) string {
					return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
				}
			}

			fallthrough
		default:
			if opts.Month == MonthNumeric {
				return func(y int, m time.Month) string {
					return fmtYear(y, opts.Year) + "/" + fmtMonth(m, MonthNumeric)
				}
			}

			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "年" + fmtMonth(m, MonthNumeric) + "月"
			}
		}
	}
}

func fmtYearMonthBuddhist(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	if lang, _ := locale.Base(); lang == th {
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	}

	return func(y int, m time.Month) string {
		return fmtEra(locale, EraNarrow) + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
	}
}

func fmtYearMonthPersian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	lang, _, region := locale.Raw()
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch lang {
	case ckb: // ckb-IR
		// year=numeric,month=numeric,out=١٠/١٤٠٢
		// year=numeric,month=2-digit,out=١٠/١٤٠٢
		// year=2-digit,month=numeric,out=١٠/٠٢
		// year=2-digit,month=2-digit,out=١٠/٠٢
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
		}
	case fa:
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month)
		}
	case ps:
		return func(y int, m time.Month) string {
			return fmtEra(locale, EraNarrow) + " " +
				fmtYear(y, opts.Year) + "/" +
				fmtMonth(m, opts.Month)
		}
	case uz:
		if region == regionAF {
			// year=numeric,month=numeric,out=۱۴۰۲-۱۰
			// year=numeric,month=2-digit,out=۱۴۰۲-۱۰
			// year=2-digit,month=numeric,out=۰۲-۱۰
			// year=2-digit,month=2-digit,out=۰۲-۱۰
			return func(y int, m time.Month) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
			}
		}

		fallthrough
	default:
		return func(y int, m time.Month) string {
			return fmtEra(locale, EraNarrow) + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}
	}
}
