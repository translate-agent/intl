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
	lang, script, region := locale.Raw()

	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	default:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
		}
	case es:
		switch region {
		case regionCL:
			// year=numeric,month=numeric,day=numeric,out=02-01-2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02-01-24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		fallthrough
	case agq, am, asa, ast, bas, bem, bez, bm, bn, ca, ccp, cgg, cy, dav, dje, doi, dua, dyo, ebu, el, ewo, gd, gl, gu,
		haw, hi, id, ig, km, kn, ksf, ln, lo, lu, mai, mgh, ml, mni, mr, ms, mua, my, nmg, nnh, nus, pcm, rn, sa, su, sw,
		ta, tg, ti, to, twq, ur, vi, xnr, yav:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
		}
	case kxv:
		switch script {
		default:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
		case deva, orya, telu:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=2024-1-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-2
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=24-1-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-2
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(y int, m time.Month, d int) string {
					return fmtDay(d, opts.Day) + "/" +
						fmtMonth(m, opts.Month) + "/" +
						fmtYear(y, opts.Year)
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtDay(d, opts.Day)
			}
		}
	case pa:
		if script == arab {
			// year=numeric,month=numeric,day=numeric,out=۲۰۲۴-۰۱-۰۲
			// year=numeric,month=numeric,day=2-digit,out=۰۲/۱/۲۰۲۴
			// year=numeric,month=2-digit,day=numeric,out=۲/۰۱/۲۰۲۴
			// year=numeric,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۰۲۴
			// year=2-digit,month=numeric,day=numeric,out=۲۴-۰۱-۰۲
			// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱/۲۴
			// year=2-digit,month=2-digit,day=numeric,out=۲/۰۱/۲۴
			// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۴
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(y int, m time.Month, d int) string {
					return fmtYear(y, opts.Year) + "-" +
						fmtMonth(m, Month2Digit) + "-" +
						fmtDay(d, Day2Digit)
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
		}
	case ak, eu, ja, yue:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtDay(d, opts.Day)
		}
	case ar:
		// year=numeric,month=numeric,day=numeric,out=٢‏/١‏/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٠٢٤
		// year=2-digit,month=numeric,day=numeric,out=٢‏/١‏/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٤
		// year=2-digit,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٤
		// year=2-digit,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٤
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "\u200f/" +
				fmtMonth(m, opts.Month) + "\u200f/" +
				fmtYear(y, opts.Year)
		}
	case as, brx, ia, jv:
		// year=numeric,month=numeric,day=numeric,out=০২-০১-২০২৪
		// year=numeric,month=numeric,day=2-digit,out=০২-১-২০২৪
		// year=numeric,month=2-digit,day=numeric,out=২-০১-২০২৪
		// year=numeric,month=2-digit,day=2-digit,out=০২-০১-২০২৪
		// year=2-digit,month=numeric,day=numeric,out=০২-০১-২৪
		// year=2-digit,month=numeric,day=2-digit,out=০২-১-২৪
		// year=2-digit,month=2-digit,day=numeric,out=২-০১-২৪
		// year=2-digit,month=2-digit,day=2-digit,out=০২-০১-২৪
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtYear(y, opts.Year)
		}
	case az, hy, kk, uk:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=02.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Year == YearNumeric ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year)
		}
	case be, da, de, dsb, et, fi, he, hsb, is, ka, lb, mk, nb, nn, no, smn, sq:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year)
		}
	case bg:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.01.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, Month2Digit) + "." +
				fmtYear(y, opts.Year) + " г."
		}
	case en:
		switch region {
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
			if script == shaw {
				break
			}

			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			if opts.Year == YearNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
		case regionBW, regionBZ:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Year == YearNumeric || opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtDay(d, opts.Day)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." +
					fmtMonth(m, opts.Month) + "." +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(y int, m time.Month, d int) string {
					return fmtDay(d, Day2Digit) + "/" +
						fmtMonth(m, Month2Digit) + "/" +
						fmtYear(y, opts.Year)
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day)
			}
		}

		fallthrough
	case blo, ceb, chr, ee, fil, or, xh:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" +
				fmtDay(d, opts.Day) + "/" +
				fmtYear(y, opts.Year)
		}
	case ks:
		if script == deva {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Year == Year2Digit {
				return func(y int, m time.Month, d int) string {
					return fmtDay(d, opts.Day) + "/" +
						fmtMonth(m, opts.Month) + "/" +
						fmtYear(y, opts.Year)
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" +
				fmtDay(d, opts.Day) + "/" +
				fmtYear(y, opts.Year)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
		}
	case bs:
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

			if opts.Month == MonthNumeric {
				opts.Day = Day2Digit
			}

			if day == DayNumeric {
				opts.Month = Month2Digit
			}
		}

		// year=numeric,month=numeric,day=numeric,out=2.1.2024.
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
		// year=2-digit,month=numeric,day=numeric,out=2.1.24.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year) + "."
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, DayNumeric) + "/" +
					fmtMonth(m, MonthNumeric) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " +
				fmtMonth(m, opts.Month) + ". " +
				fmtYear(y, opts.Year)
		}
	case cv, fo, ku, ro, ru, tk, tt:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year)
		}
	case dz, si:
		// year=numeric,month=numeric,day=numeric,out=2024-1-2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-1-2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
		}
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
		case opts.Year == YearNumeric:
			opts.Year = YearNumeric
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month == MonthNumeric && opts.Day == DayNumeric:
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, Month2Digit) + "-" +
				fmtDay(d, Day2Digit)
		}
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
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtYear(y, opts.Year)
			}
		}

		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-01-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-02
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, Month2Digit) + "-" +
				fmtDay(d, Day2Digit)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, opts.Month) + "-" +
					fmtDay(d, opts.Day)
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
			if opts.Year == YearNumeric || opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." +
					fmtMonth(m, opts.Month) + "." +
					fmtYear(y, opts.Year)
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
			if opts.Year == YearNumeric {
				if opts.Month == MonthNumeric && opts.Day == DayNumeric {
					opts.Day = Day2Digit
				}

				opts.Month = Month2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(y int, m time.Month, d int) string {
					return fmtMonth(m, opts.Month) + "/" +
						fmtDay(d, opts.Day) + "/" +
						fmtYear(y, opts.Year)
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
		}
	case nl:
		if region == regionBE {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		fallthrough
	case fy, kok:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtYear(y, opts.Year)
		}
	case gsw:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year)
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
		if opts.Year == YearNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, YearNumeric) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, Year2Digit)
		}
	case hr:
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
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
			if opts.Year == YearNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " +
				fmtMonth(m, opts.Month) + ". " +
				fmtYear(y, opts.Year) + "."
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + ". " +
				fmtMonth(m, opts.Month) + ". " +
				fmtDay(d, opts.Day) + "."
		}
	case ie, nds, prg:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, DayNumeric) + "." +
					fmtMonth(m, MonthNumeric) + "." +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
			separator := "."
			if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
				opts.Month == MonthNumeric && opts.Day == DayNumeric {
				separator = "/"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + separator +
					fmtMonth(m, opts.Month) + separator +
					fmtYear(y, opts.Year)
			}
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
		switch {
		case opts.Year == YearNumeric:
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month == MonthNumeric && opts.Day == DayNumeric:
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, MonthNumeric) + "." +
					fmtDay(d, DayNumeric) + "." +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + " " +
				fmtYear(y, opts.Year)
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
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + ". " +
				fmtMonth(m, opts.Month) + ". " +
				fmtDay(d, opts.Day) + "."
		}
	case ky:
		// year=numeric,month=numeric,day=numeric,out=2024-02-01
		// year=numeric,month=numeric,day=2-digit,out=2024-02-01
		// year=numeric,month=2-digit,day=numeric,out=2024-02-01
		// year=numeric,month=2-digit,day=2-digit,out=2024-02-01
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year == YearNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, YearNumeric) + "-" +
					fmtDay(d, Day2Digit) + "-" +
					fmtMonth(m, Month2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, Year2Digit)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, DayNumeric) + "/" +
					fmtMonth(m, MonthNumeric) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		if opts.Year == YearNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, YearNumeric) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" +
				fmtDay(d, opts.Day) + "/" +
				fmtYear(y, Year2Digit)
		}
	case lv:
		if opts.Year == Year2Digit && (opts.Month == Month2Digit || opts.Day == Day2Digit) ||
			(opts.Month == Month2Digit && opts.Day == Day2Digit) {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." +
					fmtMonth(m, opts.Month) + "." +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, Month2Digit) + "." +
				fmtYear(y, opts.Year) + "."
		}
	case mi, rm, wo:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtYear(y, opts.Year)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtDay(d, opts.Day)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, MonthNumeric) + "/" +
					fmtDay(d, DayNumeric) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
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
		separator := "/"

		if opts.Year == YearNumeric {
			separator = "-"
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtDay(d, opts.Day)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + " / " +
					fmtDay(d, Day2Digit) + " / " +
					fmtMonth(m, Month2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, Day2Digit) + "/" +
					fmtMonth(m, Month2Digit) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
		}
	case om:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
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
		if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, opts.Month) + "." +
				fmtYear(y, opts.Year)
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
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, Month2Digit) + "." +
				fmtYear(y, opts.Year)
		}
	case qu:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-01-2024
		// year=numeric,month=2-digit,day=numeric,out=02-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		separator := "/"
		if opts.Year == YearNumeric {
			separator = "-"
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year)
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
		separator := "/"

		if opts.Year == YearNumeric {
			separator = "-"
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtDay(d, opts.Day)
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
			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." +
					fmtMonth(m, opts.Month) + "." +
					fmtYear(y, opts.Year)
			}
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day) + "/" +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" +
				fmtMonth(m, opts.Month) + "/" +
				fmtYear(y, opts.Year)
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
		separator := ". "
		if opts.Month == Month2Digit && opts.Day == Day2Digit {
			separator = "."
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year) + "."
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
		separator := "-"
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, Day2Digit) + "." +
					fmtMonth(m, Month2Digit) + "." +
					fmtYear(y, opts.Year)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		separator := "-"
		if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "/"
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return "#" + fmtYear(y, opts.Year) + ")#" +
					fmtMonth(m, opts.Month) + ")#" +
					fmtDay(d, opts.Day)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Day = Day2Digit
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." +
				fmtMonth(m, Month2Digit) + "." +
				fmtYear(y, opts.Year)
		}
	case ug:
		// year=numeric,month=numeric,day=numeric,out=2024-2-1
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-2-1
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtDay(d, opts.Day) + "-" +
					fmtMonth(m, opts.Month)
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, opts.Month) + "-" +
				fmtDay(d, opts.Day)
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
		separator := "/"
		if opts.Year == YearNumeric && !(opts.Month == Month2Digit && opts.Day == Day2Digit) ||
			opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "-"
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year)
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
		separator := "/"
		if opts.Month == Month2Digit {
			separator = " "
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtYear(y, opts.Year)
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
		separator := "-"
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "/"
		}

		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + separator +
				fmtMonth(m, opts.Month) + separator +
				fmtDay(d, opts.Day)
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
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtDay(d, opts.Day)
			}
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
				return func(y int, m time.Month, d int) string {
					return fmtYear(y, opts.Year) + "年" +
						fmtMonth(m, opts.Month) + "月" +
						fmtDay(d, opts.Day) + "日"
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
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" +
					fmtMonth(m, opts.Month) + "/" +
					fmtYear(y, opts.Year)
			}
		}
	}
}

func fmtYearMonthDayPersian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, _, region := locale.Raw()

	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	case ckb: // ckb-IR
		// year=numeric,month=numeric,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=numeric,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=2-digit,month=numeric,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=numeric,day=2-digit,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=2-digit,out=١٢/١٠/٠٢
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "/" +
				fmtMonth(m, Month2Digit) + "/" +
				fmtYear(y, opts.Year)
		}
	case fa: // fa-IR
		// year=numeric,month=numeric,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=numeric,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۱۰/۱۲
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "/" +
				fmtMonth(m, Month2Digit) + "/" +
				fmtDay(d, Day2Digit)
		}
	case uz:
		if region == regionAF {
			// year=numeric,month=numeric,day=numeric,out=۱۴۰۲-۱۰-۱۲
			// year=numeric,month=numeric,day=2-digit,out=۱۴۰۲-۱۰-۱۲
			// year=numeric,month=2-digit,day=numeric,out=۱۴۰۲-۱۰-۱۲
			// year=numeric,month=2-digit,day=2-digit,out=۱۴۰۲-۱۰-۱۲
			// year=2-digit,month=numeric,day=numeric,out=۰۲-۱۰-۱۲
			// year=2-digit,month=numeric,day=2-digit,out=۰۲-۱۰-۱۲
			// year=2-digit,month=2-digit,day=numeric,out=۰۲-۱۰-۱۲
			// year=2-digit,month=2-digit,day=2-digit,out=۰۲-۱۰-۱۲
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" +
					fmtMonth(m, Month2Digit) + "-" +
					fmtDay(d, Day2Digit)
			}
		}

		fallthrough
	default: // "lrc", "mzn", "ps", "uz"
		// year=numeric,month=numeric,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
		// year=numeric,month=numeric,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
		// year=numeric,month=2-digit,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
		// year=numeric,month=2-digit,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
		// year=2-digit,month=numeric,day=numeric,out=AP ۰۲-۱۰-۱۲
		// year=2-digit,month=numeric,day=2-digit,out=AP ۰۲-۱۰-۱۲
		// year=2-digit,month=2-digit,day=numeric,out=AP ۰۲-۱۰-۱۲
		// year=2-digit,month=2-digit,day=2-digit,out=AP ۰۲-۱۰-۱۲
		return func(y int, m time.Month, d int) string {
			return "AP " +
				fmtYear(y, opts.Year) + "-" +
				fmtMonth(m, Month2Digit) + "-" +
				fmtDay(d, Day2Digit)
		}
	}
}

func fmtYearMonthDayBuddhist(
	_ language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

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
		return fmtDay(d, opts.Day) + "/" +
			fmtMonth(m, opts.Month) + "/" +
			fmtYear(y, opts.Year)
	}
}
