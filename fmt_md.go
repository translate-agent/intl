package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, script, region := locale.Raw()

	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	default:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case af, as, ia, ky, mi, rm, tg, wo:
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) }
	case hi:
		if script == latn && opts.Month == MonthNumeric && opts.Day == DayNumeric {
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case agq, ast, bas, bm, ca, cy, dje, doi, dua, dyo, el, ewo, fur, gd, gl, haw, id, ig, kab, kgp, khq, km, ksf, ln,
		lo, lu, mai, mfe, mg, mgh, ml, mni, mua, my, nmg, nus, pa, rn, sa, seh, ses, sg, shi, su, sw, to, tr, twq, ur,
		xnr, yav, yo, yrl, zgh:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case kxv:
		switch script {
		default:
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		case deva, orya, telu:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=01-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			return func(m time.Month, d int) string {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}
	case en:
		switch region {
		case region001, region150, regionAE, regionAG, regionAI, regionAT, regionBB, regionBM, regionBS, regionBW,
			regionBZ, regionCC, regionCK, regionCM, regionCX, regionCY, regionDE, regionDG, regionDK, regionDM, regionER,
			regionFI, regionFJ, regionFK, regionFM, regionGB, regionGD, regionGG, regionGH, regionGI, regionGM, regionGY,
			regionHK, regionID, regionIL, regionIM, regionIN, regionIO, regionJE, regionJM, regionKE, regionKI, regionKN,
			regionKY, regionLC, regionLR, regionLS, regionMG, regionMO, regionMS, regionMT, regionMU, regionMV, regionMW,
			regionMY, regionNA, regionNF, regionNG, regionNL, regionNR, regionNU, regionPG, regionPK, regionPN, regionPW,
			regionRW, regionSB, regionSC, regionSD, regionSE, regionSG, regionSH, regionSI, regionSL, regionSS, regionSX,
			regionSZ, regionTC, regionTK, regionTO, regionTT, regionTV, regionTZ, regionUG, regionVC, regionVG, regionVU,
			regionWS, regionZM:
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			if script == shaw {
				break
			}

			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string { return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) }
		case regionAU, regionBE, regionIE, regionNZ, regionZW:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		case regionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-2
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
			}
		case regionCH:
			// month=numeric,day=numeric,out=02.01
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string { return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) }
		case regionZA:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=01/02
			// month=2-digit,day=2-digit,out=02/01
			if opts.Month == Month2Digit && opts.Day == Day2Digit {
				return func(m time.Month, d int) string { return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) }
			}

			return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit) }
		}

		fallthrough
	case ak, am, asa, bem, bez, blo, brx, ceb, cgg, chr, dav, ebu, ee, eu, fil, guz, ha, ja, jmc, kam, kde, ki, kln,
		ksb, lag, lg, luo, luy, mas, mer, naq, nd, nyn, rof, rwk, saq, sbp, so, teo, tzm, vai, vun, xh, xog, yue:
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case ks:
		if script == deva {
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=01-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			return func(m time.Month, d int) string {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}

		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case ar:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "\u200f/" + fmtMonth(m, opts.Month)
		}
	case az, cv, fo, hy, kk, ku, os, tk, tt, uk:
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) }
	case be, da, et, he, jgo, ka:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case bg, pl:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit)
		}
	case bn, ccp, gu, kn, mr, ta, te, vi:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case br, ga, it, jv, kkj, sc, syr, uz, vec:
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) }
	case bs:
		if script == cyrl {
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.01.
			// month=2-digit,day=numeric,out=02.01.
			// month=2-digit,day=2-digit,out=02.01.
			return func(m time.Month, d int) string {
				return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
			}
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(m time.Month, d int) string {
				return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "."
			}
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, DayNumeric) + ". " + fmtMonth(m, MonthNumeric) + "."
		}
	case cs, sk, sl:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case de, dsb, fi, gsw, hsb, is, lb, smn:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "."
		}
	case dz, si:
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case es:
		switch region {
		case regionCL:
			// month=numeric,day=numeric,out=02-01
			// month=numeric,day=2-digit,out=02-01
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			separator := "/"
			if opts.Month == MonthNumeric {
				separator = "-"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			}

			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
			}
		case regionMX, regionUS:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return func(m time.Month, d int) string { return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) }
		case regionPA, regionPR:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			if opts.Month == MonthNumeric {
				return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit) }
			}

			return func(m time.Month, d int) string {
				return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
			}
		}

		fallthrough
	case ti:
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric) }
	case ff:
		if script == adlm {
			// month=numeric,day=numeric,out=𞥒-𞥑
			// month=numeric,day=2-digit,out=𞥐𞥒-𞥑
			// month=2-digit,day=numeric,out=𞥒-𞥐𞥑
			// month=2-digit,day=2-digit,out=𞥐𞥒-𞥐𞥑
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case fr:
		switch region {
		default:
			return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) }
		case regionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month == MonthNumeric && opts.Day == Day2Digit {
				opts.Month = MonthNumeric
			} else {
				opts.Month = Month2Digit
			}

			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "-" + fmtDay(d, Day2Digit)
			}
		case regionCH:
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(m time.Month, d int) string {
					return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
				}
			}

			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
			}
		}
	case nl:
		if region == regionBE {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		fallthrough
	case fy, kok, ug:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case hr:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + ". " + fmtMonth(m, Month2Digit) + "."
			}

			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case hu, ko:
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + ". " + fmtDay(d, opts.Day) + "."
		}
	case iu:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit) }
	case kea, pt:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case lt:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, DayNumeric)
			}

			return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case lv:
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." }
	case mk:
		return func(m time.Month, d int) string {
			return fmtDay(d, DayNumeric) + "." + fmtMonth(m, opts.Month)
		}
	case mn:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")
		return func(m time.Month, d int) string { return fmtMonth(m, opts.Month) + "/" + fmtDay(d, Day2Digit) }
	case ms:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + "-" + fmtMonth(m, MonthNumeric)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case nb, nn, no:
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "." }
	case om:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case or:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, MonthNumeric) + "/" + fmtDay(d, DayNumeric)
			}

			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case pcm:
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + " /" + fmtMonth(m, opts.Month)
		}
	case ro, ru:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit)
			}

			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case sd:
		if script == deva {
			// month=numeric,day=numeric,out=1/2
			// month=numeric,day=2-digit,out=1/02
			// month=2-digit,day=numeric,out=01/2
			// month=2-digit,day=2-digit,out=01/02
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		}

		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case se:
		if region == regionFI {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case sq:
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) }
	case sr:
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + ". " + fmtMonth(m, MonthNumeric) + "."
			}

			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "."
		}
	case sv:
		if region == regionFI {
			// month=numeric,day=numeric,out=2.1
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.1
			// month=2-digit,day=2-digit,out=02.01
			if opts.Day == DayNumeric {
				opts.Month = MonthNumeric
			}

			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			if opts.Month == Month2Digit && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case wae:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "abbreviated")

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month)
		}
	case zh:
		switch region {
		default:
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		case regionHK, regionMO:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		case regionSG:
			// month=numeric,day=numeric,out=1-2
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-2
			// month=2-digit,day=2-digit,out=01-02
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
			}
		}
	}
}

func fmtMonthDayBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	if lang, _ := locale.Base(); lang == th {
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	}

	return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
}

func fmtMonthDayPersian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang, _ := locale.Base(); lang {
	default:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case fa, ps:
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	}
}
