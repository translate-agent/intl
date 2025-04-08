package intl

import (
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtMonthDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()

	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutMonthDay
	middle := "-"
	suffix := ""

	switch lang {
	default:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case en:
		switch region {
		default:
			middle = "/"
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
			middle = "/"

			if script == shaw {
				break
			}

			layout = layoutDayMonth

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionAU, regionBE, regionIE, regionNZ, regionZW:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case regionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-2
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionCH:
			// month=numeric,day=numeric,out=02.01
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			layout = layoutDayMonth
			middle = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionZA:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=01/02
			// month=2-digit,day=2-digit,out=02/01
			middle = "/"

			if opts.Month.twoDigit() && opts.Day.twoDigit() {
				layout = layoutDayMonth
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}
	case af, as, ia, ky, mi, rm, tg, wo:
		layout = layoutDayMonth
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case hi:
		if script == latn && opts.Month.numeric() && opts.Day.numeric() {
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case am, agq, ast, bas, bm, ca, cy, dje, doi, dua, dyo, el, ewo, fur, gd, gl, haw, id, ig, kab, kgp, khq, km, ksf, kxv,
		ln, lo, lu, mai, mfe, mg, mgh, ml, mni, mua, my, nmg, nus, pa, rn, sa, seh, ses, sg, shi, su, sw, to, tr, twq, ur,
		xnr, yav, yo, yrl, zgh:
		layout = layoutDayMonth
		middle = "/"
	case br, ga, it, jv, kkj, sc, syr, uz, vec:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutDayMonth
		middle = "/"
	case ti:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "/"
	case kea, pt:
		layout = layoutDayMonth
		middle = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case ak, asa, bem, bez, blo, brx, ceb, cgg, chr, dav, ebu, ee, eu, fil, guz, ha, ja, jmc, kaa, kam, kde, ki, kln, ksb,
		lag, lg, luo, luy, mas, mer, mhn, naq, nd, nyn, rof, rwk, saq, sbp, so, teo, tzm, vai, vun, xh, xog, yue:
		middle = "/"
	case ks:
		if script == deva {
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=01-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			middle = "/"
		}
	case ar:
		layout = layoutDayMonth
		middle = "\u200f/"
	case az, cv, fo, hy, kk, ku, os, tk, tt, uk:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutDayMonth
		middle = "."
	case be, da, et, he, ie, jgo, ka:
		layout = layoutDayMonth
		middle = "."
	case mk:
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "."
	case bg, pl:
		opts.Month = Month2Digit
		layout = layoutDayMonth
		middle = "."
	case lv:
		opts.Month = Month2Digit
		opts.Day = Day2Digit

		fallthrough
	case de, dsb, fi, gsw, hsb, is, lb, smn:
		layout = layoutDayMonth
		middle = "."
		suffix = "."
	case nb, nn, no:
		suffix = "."
		fallthrough
	case sq:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "."
	case ro, ru:
		layout = layoutDayMonth
		middle = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case sr:
		layout = layoutDayMonth
		suffix = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			middle = ". "
		} else {
			middle = "."
		}
	case bn, ccp, gu, kn, mr, ta, te, vi:
		layout = layoutDayMonth

		if opts.Month.numeric() && opts.Day.numeric() {
			middle = "/"
		}
	case bs:
		layout = layoutDayMonth
		suffix = "."

		if script == cyrl {
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.01.
			// month=2-digit,day=numeric,out=02.01.
			// month=2-digit,day=2-digit,out=02.01.
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			middle = "."
		} else {
			// month=numeric,day=numeric,out=2. 1.
			// month=numeric,day=2-digit,out=2. 1.
			// month=2-digit,day=numeric,out=2. 1.
			// month=2-digit,day=2-digit,out=2. 1.
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
			middle = ". "
		}
	case hr:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cs, sk, sl:
		layout = layoutDayMonth

		fallthrough
	case hu, ko:
		middle = ". "
		suffix = "."
	case wae:
		month = fmtMonthName(locale.String(), "stand-alone", "abbreviated")
		layout = layoutDayMonth
		middle = ". "
	case dz, si: // noop
	case es:
		switch region {
		default:
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
			layout = layoutDayMonth
			middle = "/"
		case regionCL:
			// month=numeric,day=numeric,out=02-01
			// month=numeric,day=2-digit,out=02-01
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			layout = layoutDayMonth

			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				middle = "/"
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			}
		case regionMX, regionUS:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case regionPA, regionPR:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			middle = "/"

			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				layout = layoutDayMonth
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			}
		}
	case ff:
		layout = layoutDayMonth

		if script != adlm {
			middle = "/"
		}
	case fr:
		switch region {
		default:
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonth
			middle = "/"
		case regionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month.numeric() && opts.Day.twoDigit() {
				opts.Month = MonthNumeric
			} else {
				opts.Month = Month2Digit
			}

			opts.Day = Day2Digit
		case regionCH:
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			layout = layoutDayMonth
			middle = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				suffix = "."
			}
		}
	case nl:
		layout = layoutDayMonth

		if region == regionBE {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			middle = "/"
		}
	case fy, ug:
		layout = layoutDayMonth
	case iu:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		middle = "/"
	case lt:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
		}
	case mn:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
		opts.Day = Day2Digit
		middle = "/"
	case ms:
		layout = layoutDayMonth

		if !opts.Month.numeric() || !opts.Day.numeric() {
			middle = "/"
		}
	case om:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonth
			middle = "/"
		}
	case or:
		if opts.Month.numeric() && opts.Day.numeric() {
			middle = "/"
		} else {
			layout = layoutDayMonth
		}
	case pcm:
		layout = layoutDayMonth
		middle = " /"
	case sd:
		if script == deva {
			// month=numeric,day=numeric,out=1/2
			// month=numeric,day=2-digit,out=1/02
			// month=2-digit,day=numeric,out=01/2
			// month=2-digit,day=2-digit,out=01/02
			middle = "/"
		} else {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case se:
		if region == regionFI {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		} else {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case sv:
		layout = layoutDayMonth

		if region == regionFI {
			// month=numeric,day=numeric,out=2.1
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.1
			// month=2-digit,day=2-digit,out=02.01
			middle = "."

			if opts.Day.numeric() {
				opts.Month = MonthNumeric
			}

			break
		}

		middle = "/"

		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
		}
	case zh:
		switch region {
		default:
			middle = "/"
		case regionHK, regionMO:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case regionSG: // noop
		}
	case ii:
		// month=numeric,day=numeric,out=01ꆪ-02ꑍ
		// month=numeric,day=2-digit,out=01ꆪ-02ꑍ
		// month=2-digit,day=numeric,out=01ꆪ-02ꑍ
		// month=2-digit,day=2-digit,out=01ꆪ-02ꑍ
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		middle = "ꆪ-"
		suffix = "ꑍ"
	case kok:
		layout = layoutDayMonth

		if script == latn {
			middle = "/"
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t timeReader) string {
			return dayDigits(t) + middle + month(t) + suffix
		}
	}

	return func(t timeReader) string {
		return month(t) + middle + dayDigits(t) + suffix
	}
}

func fmtMonthDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutMonthDay

	if lang, _ := locale.Base(); lang == th {
		layout = layoutDayMonth
	} else {
		opts.Day = Day2Digit
	}

	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t timeReader) string {
			return dayDigits(t) + "/" + monthDigits(t)
		}
	}

	return func(t timeReader) string { return monthDigits(t) + "-" + dayDigits(t) }
}

func fmtMonthDayPersian(locale language.Tag, digits digits, opts Options) fmtFunc {
	middle := "-"

	switch lang, _ := locale.Base(); lang {
	default:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case fa, ps:
		middle = "/"
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v timeReader) string {
		return month(v) + middle + dayDigits(v)
	}
}
