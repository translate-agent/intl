package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtYearMonthDayGregorian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	var month func(int) string

	lang, script, region := locale.Raw()
	year := fmtYear(digits, opts.Year)

	const (
		layoutYearMonthDay = iota
		layoutDayMonthYear
		layoutMonthDayYear
		layoutYearDayMonth
	)

	layout := layoutYearMonthDay
	prefix := ""
	separator := "-"
	suffix := ""

	switch lang {
	default:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case es:
		switch region {
		default:
			layout = layoutDayMonthYear
			separator = "/"
		case regionCL:
			// year=numeric,month=numeric,day=numeric,out=02-01-2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02-01-24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			layout = layoutDayMonthYear

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionPA, regionPR:
			// year=numeric,month=numeric,day=numeric,out=01/02/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=01/02/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			layout = layoutMonthDayYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}
	case agq, am, asa, ast, bas, bem, bez, bm, bn, ca, ccp, cgg, cy, dav, dje, doi, dua, dyo, ebu, el, ewo, gd, gl, gu,
		haw, hi, id, ig, km, kn, ksf, kxv, ln, lo, lu, mai, mgh, ml, mni, mr, ms, mua, my, nmg, nnh, nus, pcm, rn, sa, su,
		sw, ta, to, twq, ur, vi, xnr, yav:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear
		separator = "/"
	case pa:
		if script == arab && opts.Month.numeric() && opts.Day.numeric() {
			// year=numeric,month=numeric,day=numeric,out=۲۰۲۴-۰۱-۰۲
			// year=numeric,month=numeric,day=2-digit,out=۰۲/۱/۲۰۲۴
			// year=numeric,month=2-digit,day=numeric,out=۲/۰۱/۲۰۲۴
			// year=numeric,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۰۲۴
			// year=2-digit,month=numeric,day=numeric,out=۲۴-۰۱-۰۲
			// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱/۲۴
			// year=2-digit,month=2-digit,day=numeric,out=۲/۰۱/۲۴
			// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۴
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
		}
	case ak:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		separator = "/"

		if opts.Year.twoDigit() {
			layout = layoutMonthDayYear
		}
	case eu, ja, yue:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		separator = "/"
	case ar:
		// year=numeric,month=numeric,day=numeric,out=٢‏/١‏/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٠٢٤
		// year=2-digit,month=numeric,day=numeric,out=٢‏/١‏/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٤
		// year=2-digit,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٤
		// year=2-digit,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٤
		layout = layoutDayMonthYear
		separator = "\u200f/"
	case az, hy, kk, uk:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=02.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Year.numeric() ||
			opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case be, da, de, dsb, et, fi, he, hsb, ie, is, ka, lb, nb, nn, no, smn, sq:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."
	case bg:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.01.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		opts.Month = Month2Digit
		fallthrough
	case mk:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.1.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		layout = layoutDayMonthYear
		separator = "."
		suffix = " г."
	case en:
		switch region {
		default:
			layout = layoutMonthDayYear
			separator = "/"
		case region001, region150, regionAE, regionAG, regionAI, regionAT, regionBB, regionBM, regionBS, regionCC,
			regionCK, regionCM, regionCX, regionCY, regionDE, regionDG, regionDK, regionDM, regionER, regionFI, regionFJ,
			regionFK, regionFM, regionGB, regionGD, regionGG, regionGH, regionGI, regionGM, regionGY, regionID, regionIL,
			regionIM, regionIO, regionJE, regionJM, regionKE, regionKI, regionKN, regionKY, regionLC, regionLR, regionLS,
			regionMG, regionMO, regionMS, regionMT, regionMU, regionMW, regionMY, regionNA, regionNF, regionNG, regionNL,
			regionNR, regionNU, regionPG, regionPK, regionPN, regionPW, regionRW, regionSB, regionSC, regionSD, regionSH,
			regionSI, regionSL, regionSS, regionSX, regionSZ, regionTC, regionTK, regionTO, regionTT, regionTV, regionTZ,
			regionUG, regionVC, regionVG, regionVU, regionWS, regionZM:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			separator = "/"

			if script == shaw {
				layout = layoutMonthDayYear
			} else {
				layout = layoutDayMonthYear

				if opts.Month.numeric() && opts.Day.numeric() {
					opts.Month = Month2Digit
					opts.Day = Day2Digit
				}
			}
		case regionAU, regionSG:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionBE, regionHK, regionIE, regionIN, regionZW:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
		case regionBW, regionBZ:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionCA, regionSE:
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-1-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-2
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-1-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-2
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear
			separator = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionMV:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			layout = layoutDayMonthYear

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				separator = "/"
			}
		case regionNZ:
			// year=numeric,month=numeric,day=numeric,out=2/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
			}
		case regionZA:
			// year=numeric,month=numeric,day=numeric,out=2024/01/02
			// year=numeric,month=numeric,day=2-digit,out=2024/1/02
			// year=numeric,month=2-digit,day=numeric,out=2024/01/2
			// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
			// year=2-digit,month=numeric,day=numeric,out=24/01/02
			// year=2-digit,month=numeric,day=2-digit,out=24/1/02
			// year=2-digit,month=2-digit,day=numeric,out=24/01/2
			// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}
	case blo, ceb, chr, ee, fil, kaa, mhn, om, or, ti, xh:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		layout = layoutMonthDayYear
		separator = "/"
	case ks:
		separator = "/"

		if script == deva && opts.Year.twoDigit() {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
		} else {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			layout = layoutMonthDayYear
		}
	case br, ga, kea, kgp, pt, sc, yrl:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear
		separator = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case bs:
		layout = layoutDayMonthYear
		suffix = "."

		if script == cyrl {
			// year=numeric,month=numeric,day=numeric,out=02.01.2024.
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024.
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
			// year=2-digit,month=numeric,day=numeric,out=02.01.24.
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24.
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24.
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
			day := opts.Day
			separator = "."

			if opts.Month.numeric() {
				opts.Day = Day2Digit
			}

			if day.numeric() {
				opts.Month = Month2Digit
			}
		} else {
			// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
			// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			separator = ". "
		}
	case ckb:
		// year=numeric,month=numeric,day=numeric,out=٢/١/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٢٠٢٤-١-٠٢
		// year=numeric,month=2-digit,day=numeric,out=٢٠٢٤-٠١-٢
		// year=numeric,month=2-digit,day=2-digit,out=٢٠٢٤-٠١-٠٢
		// year=2-digit,month=numeric,day=numeric,out=٢/١/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٢٤-١-٠٢
		// year=2-digit,month=2-digit,day=numeric,out=٢٤-٠١-٢
		// year=2-digit,month=2-digit,day=2-digit,out=٢٤-٠١-٠٢
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case cs, sk, sl:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024
		// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24
		// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24
		layout = layoutDayMonthYear
		separator = ". "
	case cv, fo, ku, ro, ru, tk, tt:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case dz, si: // noop
		// year=numeric,month=numeric,day=numeric,out=2024-1-2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-1-2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
	case eo:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		switch {
		case opts.Year.numeric():
			opts.Year = YearNumeric
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month.numeric() && opts.Day.numeric():
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case kab, khq, ksh, mfe, zgh, ps, seh, ses, sg, shi:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-01-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-02
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case ff:
		if script == adlm {
			// year=numeric,month=numeric,day=numeric,out=𞥒-𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=numeric,day=2-digit,out=𞥐𞥒-𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=2-digit,day=numeric,out=𞥒-𞥐𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=2-digit,day=2-digit,out=𞥐𞥒-𞥐𞥑-𞥒𞥐𞥒𞥔
			// year=2-digit,month=numeric,day=numeric,out=𞥒-𞥑-𞥒𞥔
			// year=2-digit,month=numeric,day=2-digit,out=𞥐𞥒-𞥑-𞥒𞥔
			// year=2-digit,month=2-digit,day=numeric,out=𞥒-𞥐𞥑-𞥒𞥔
			// year=2-digit,month=2-digit,day=2-digit,out=𞥐𞥒-𞥐𞥑-𞥒𞥔
			layout = layoutDayMonthYear
		} else {
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-01-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-02
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-01-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-02
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case fr:
		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionCA:
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-1-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-2
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-1-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-2
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.01.2024
			// year=numeric,month=2-digit,day=numeric,out=02.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear
			separator = "."

			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionBE:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() {
				if opts.Month.numeric() && opts.Day.numeric() {
					opts.Day = Day2Digit
				}

				opts.Month = Month2Digit
			}
		}
	case vai:
		if script == latn {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				layout = layoutMonthDayYear
			} else {
				layout = layoutDayMonthYear
			}

			break
		}

		fallthrough
	case fur, guz, jmc, kam, kde, ki, kln, ksb, lag, lg, luo, luy, mas, mer, naq, nd, nyn, rof, rwk, saq, teo, tzm, vun,
		xog:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case nl:
		layout = layoutDayMonthYear

		if region == regionBE {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			separator = "/"
		}
	case fy, kok:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear
	case gsw:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "."
		}
	case ha, sat:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case hr:
		layout = layoutDayMonthYear
		separator = ". "
		suffix = "."

		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=02. 01. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
			// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=02. 01. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case regionBA:
			// year=numeric,month=numeric,day=numeric,out=02. 01. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 01. 2024.
			// year=numeric,month=2-digit,day=numeric,out=02. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			if opts.Year.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}
	case hu:
		// year=numeric,month=numeric,day=numeric,out=2024. 01. 02.
		// year=numeric,month=numeric,day=2-digit,out=2024. 1. 02.
		// year=numeric,month=2-digit,day=numeric,out=2024. 01. 2.
		// year=numeric,month=2-digit,day=2-digit,out=2024. 01. 02.
		// year=2-digit,month=numeric,day=numeric,out=24. 01. 02.
		// year=2-digit,month=numeric,day=2-digit,out=24. 1. 02.
		// year=2-digit,month=2-digit,day=numeric,out=24. 01. 2.
		// year=2-digit,month=2-digit,day=2-digit,out=24. 01. 02.
		separator = ". "
		suffix = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case nds, prg:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutDayMonthYear
			separator = "."
		}
	case it:
		if region == regionCH {
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear

			if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
				opts.Month.numeric() && opts.Day.numeric() {
				separator = "/"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				separator = "."
			}

			break
		}

		fallthrough
	case vec, uz:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=02/01/2024
		// year=numeric,month=2-digit,day=numeric,out=02/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear
		separator = "/"

		switch {
		case opts.Year.numeric():
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month.numeric() && opts.Day.numeric():
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case jgo:
		// year=numeric,month=numeric,day=numeric,out=1.2.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1.2.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
			separator = "."
		}
	case kkj:
		// year=numeric,month=numeric,day=numeric,out=02/01 2024
		// year=numeric,month=numeric,day=2-digit,out=02/1 2024
		// year=numeric,month=2-digit,day=numeric,out=2/01 2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01 2024
		// year=2-digit,month=numeric,day=numeric,out=02/01 24
		// year=2-digit,month=numeric,day=2-digit,out=02/1 24
		// year=2-digit,month=2-digit,day=numeric,out=2/01 24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01 24
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		month = fmtMonth(digits, opts.Month)
		day := fmtDay(digits, opts.Day)

		return func(y int, m time.Month, d int) string {
			return day(d) + "/" + month(int(m)) + " " + year(y)
		}
	case ko:
		// year=numeric,month=numeric,day=numeric,out=2024. 1. 2.
		// year=numeric,month=numeric,day=2-digit,out=2024. 1. 02.
		// year=numeric,month=2-digit,day=numeric,out=2024. 01. 2.
		// year=numeric,month=2-digit,day=2-digit,out=2024. 01. 02.
		// year=2-digit,month=numeric,day=numeric,out=24. 1. 2.
		// year=2-digit,month=numeric,day=2-digit,out=24. 1. 02.
		// year=2-digit,month=2-digit,day=numeric,out=24. 01. 2.
		// year=2-digit,month=2-digit,day=2-digit,out=24. 01. 02.
		separator = ". "
		suffix = "."
	case ky:
		// year=numeric,month=numeric,day=numeric,out=2024-02-01
		// year=numeric,month=numeric,day=2-digit,out=2024-02-01
		// year=numeric,month=2-digit,day=numeric,out=2024-02-01
		// year=numeric,month=2-digit,day=2-digit,out=2024-02-01
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutYearDayMonth
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case lij, vmw:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case lkt, zu:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutMonthDayYear
			separator = "/"
		}
	case lv:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Year.twoDigit() && opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			suffix = "."
		}
	case as, brx, ia, jv, mi, rm, wo:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case rw:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonthYear
		}
	case mn:
		// year=numeric,month=numeric,day=numeric,out=2024.01.02
		// year=numeric,month=numeric,day=2-digit,out=2024.1.02
		// year=numeric,month=2-digit,day=numeric,out=2024.01.2
		// year=numeric,month=2-digit,day=2-digit,out=2024.01.02
		// year=2-digit,month=numeric,day=numeric,out=24.01.02
		// year=2-digit,month=numeric,day=2-digit,out=24.1.02
		// year=2-digit,month=2-digit,day=numeric,out=24.01.2
		// year=2-digit,month=2-digit,day=2-digit,out=24.01.02
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case mt, sbp:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		separator = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
		} else {
			layout = layoutDayMonthYear
		}
	case ne:
		// year=numeric,month=numeric,day=numeric,out=२०२४-०१-०२
		// year=numeric,month=numeric,day=2-digit,out=२०२४-०१-०२
		// year=numeric,month=2-digit,day=numeric,out=२०२४-०१-०२
		// year=numeric,month=2-digit,day=2-digit,out=२०२४-०१-०२
		// year=2-digit,month=numeric,day=numeric,out=२४/१/२
		// year=2-digit,month=numeric,day=2-digit,out=२४/१/०२
		// year=2-digit,month=2-digit,day=numeric,out=२४/०१/२
		// year=2-digit,month=2-digit,day=2-digit,out=२४/०१/०२
		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
	case nqo:
		// year=numeric,month=numeric,day=numeric,out=߂߀߂߄ / ߀߂ / ߀߁
		// year=numeric,month=numeric,day=2-digit,out=߂߀߂߄-߁-߀߂
		// year=numeric,month=2-digit,day=numeric,out=߂߀߂߄-߀߁-߂
		// year=numeric,month=2-digit,day=2-digit,out=߂߀߂߄-߀߁-߀߂
		// year=2-digit,month=numeric,day=numeric,out=߂߄ / ߀߂ / ߀߁
		// year=2-digit,month=numeric,day=2-digit,out=߂߄-߁-߀߂
		// year=2-digit,month=2-digit,day=numeric,out=߂߄-߀߁-߂
		// year=2-digit,month=2-digit,day=2-digit,out=߂߄-߀߁-߀߂
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutYearDayMonth
			separator = " / "
		}
	case oc:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutDayMonthYear
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			separator = "/"
		}
	case os:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "."
		}
	case pl:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		opts.Month = Month2Digit
		layout = layoutDayMonthYear
		separator = "."
	case qu:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-01-2024
		// year=numeric,month=2-digit,day=numeric,out=02-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear

		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
	case sah:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
	case sd:
		if script == deva {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			layout = layoutMonthDayYear
			separator = "/"
		} else if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case se:
		if region == regionFI {
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear
			separator = "."
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case so:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		separator = "/"

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
		} else {
			layout = layoutDayMonthYear
		}
	case sr:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
		layout = layoutDayMonthYear
		suffix = "."

		if opts.Month.twoDigit() && opts.Day.twoDigit() {
			separator = "."
		} else {
			separator = ". "
		}
	case syr:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Month.numeric() {
			separator = "/"
		}
	case szl:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonthYear
			separator = "."
		}
	case te:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			separator = "/"
		}
	case tok:
		// year=numeric,month=numeric,day=numeric,out=#2024)#1)#2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=#24)#1)#2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			prefix = "#"
			separator = ")#"
		}
	case tr:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Day = Day2Digit
		}

		opts.Month = Month2Digit
	case ug:
		// year=numeric,month=numeric,day=numeric,out=2024-2-1
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-2-1
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutYearDayMonth
		}
	case yi:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear

		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			separator = "/"
		}
	case yo:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2 01 2024
		// year=numeric,month=2-digit,day=2-digit,out=02 01 2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2 01 24
		// year=2-digit,month=2-digit,day=2-digit,out=02 01 24
		layout = layoutDayMonthYear

		if opts.Month.twoDigit() {
			separator = " "
		} else {
			separator = "/"
		}
	case za:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			separator = "/"
		}
	case zh:
		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=2024/1/2
			// year=numeric,month=numeric,day=2-digit,out=2024/1/02
			// year=numeric,month=2-digit,day=numeric,out=2024/01/2
			// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
			// year=2-digit,month=numeric,day=numeric,out=24/1/2
			// year=2-digit,month=numeric,day=2-digit,out=24/1/02
			// year=2-digit,month=2-digit,day=numeric,out=24/01/2
			// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
			separator = "/"
		case regionMO, regionSG:
			if script == hans {
				// year=numeric,month=numeric,day=numeric,out=2024年1月2日
				// year=numeric,month=numeric,day=2-digit,out=2024年1月02日
				// year=numeric,month=2-digit,day=numeric,out=2024年01月2日
				// year=numeric,month=2-digit,day=2-digit,out=2024年01月02日
				// year=2-digit,month=numeric,day=numeric,out=24年1月2日
				// year=2-digit,month=numeric,day=2-digit,out=24年1月02日
				// year=2-digit,month=2-digit,day=numeric,out=24年01月2日
				// year=2-digit,month=2-digit,day=2-digit,out=24年01月02日
				month = fmtMonth(digits, opts.Month)
				day := fmtDay(digits, opts.Day)

				return func(y int, m time.Month, d int) string {
					return year(y) + "年" + month(int(m)) + "月" + day(d) + "日"
				}
			}

			fallthrough
		case regionHK:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
		}
	case tg:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear

		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			separator = "/"
		} else {
			separator = "."
		}
	case gaa:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
			separator = "/"
		}
	}

	if month == nil {
		month = fmtMonth(digits, opts.Month)
	}

	day := fmtDay(digits, opts.Day)

	switch layout {
	default: // layoutYearMonthDay
		return func(y int, m time.Month, d int) string {
			return prefix + year(y) + separator + month(int(m)) + separator + day(d) + suffix
		}
	case layoutDayMonthYear:
		return func(y int, m time.Month, d int) string {
			return day(d) + separator + month(int(m)) + separator + year(y) + suffix
		}
	case layoutMonthDayYear:
		return func(y int, m time.Month, d int) string {
			return month(int(m)) + separator + day(d) + separator + year(y) + suffix
		}
	case layoutYearDayMonth:
		return func(y int, m time.Month, d int) string {
			return year(y) + separator + day(d) + separator + month(int(m)) + suffix
		}
	}
}

func fmtYearMonthDayPersian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, _, region := locale.Raw()

	year := fmtYear(digits, opts.Year)

	const (
		layoutYearMonthDay = iota
		layoutDayMonthYear
	)

	layout := layoutYearMonthDay

	// "lrc", "mzn", "ps", "uz"
	// year=numeric,month=numeric,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=numeric,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=2-digit,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=2-digit,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
	// year=2-digit,month=numeric,day=numeric,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=numeric,day=2-digit,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=2-digit,day=numeric,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=2-digit,day=2-digit,out=AP ۰۲-۱۰-۱۲
	opts.Month = Month2Digit
	opts.Day = Day2Digit
	prefix := ""
	separator := "-"

	switch lang {
	default:
		prefix = "AP "
	case ckb: // ckb-IR
		// year=numeric,month=numeric,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=numeric,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=2-digit,month=numeric,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=numeric,day=2-digit,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=2-digit,out=١٢/١٠/٠٢
		layout = layoutDayMonthYear
		separator = "/"
	case fa: // fa-IR
		// year=numeric,month=numeric,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=numeric,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۱۰/۱۲
		separator = "/"
	case uz:
		if region != regionAF {
			prefix = "AP "
		}
	}

	month := fmtMonth(digits, opts.Month)
	day := fmtDay(digits, opts.Day)

	if layout == layoutDayMonthYear {
		return func(y int, m time.Month, d int) string {
			return day(d) + "/" + month(int(m)) + "/" + year(y)
		}
	}

	return func(y int, m time.Month, d int) string {
		return prefix + year(y) + separator + month(int(m)) + separator + day(d)
	}
}

func fmtYearMonthDayBuddhist(
	_ language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	year := fmtYear(digits, opts.Year)
	month := fmtMonth(digits, opts.Month)
	day := fmtDay(digits, opts.Day)

	// th-TH
	// year=numeric,month=numeric,day=numeric,out=2/1/2024
	// year=numeric,month=numeric,day=2-digit,out=02/1/2024
	// year=numeric,month=2-digit,day=numeric,out=2/01/2024
	// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
	// year=2-digit,month=numeric,day=numeric,out=2/1/24
	// year=2-digit,month=numeric,day=2-digit,out=02/1/24
	// year=2-digit,month=2-digit,day=numeric,out=2/01/24
	// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
	return func(y int, m time.Month, d int) string {
		return day(d) + "/" + month(int(m)) + "/" + year(y)
	}
}
