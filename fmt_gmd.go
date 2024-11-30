package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	case af, as, ia, ky, mi, rm, tg, wo:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit)
		}
	case sd:
		if script == deva {
			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		}

		fallthrough
	case bgc, bho, bo, ce, ckb, csw, eo, gv, ie, ii, kl, ksh, kw, lij, lkt, lmo, mgo, mt, nds, nnh, ne, nqo, oc, prg, ps,
		qu, raj, rw, sah, sat, sn, szl, tok, vmw, yi, za, zu:
		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
		}
	case kxv:
		if script == deva || script == orya || script == telu {
			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case dz, si:
		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case lt:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case nl:
		if region == regionBE {
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		fallthrough
	case fy, kok, ug:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case or:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case ms:
		separator := "/"
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "-"
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
		}
	case se:
		if region == regionFI {
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
		}
	case kn, mr, vi:
		separator := "-"
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "/"
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
		}
	case ti:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
		}
	case ff:
		separator := "/"
		if script == adlm {
			separator = "-"
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
		}
	case bn, ccp, gu, ta, te:
		if opts.Month == Month2Digit || opts.Day == Day2Digit {
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case az, cv, fo, hy, kk, ku, os, tk, tt, uk:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit)
		}
	case sq:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric)
		}
	case bg:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit)
		}
	case pl:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit)
		}
	case be, da, et, ka:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case mk:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + "." + fmtMonth(m, opts.Month)
		}
	case nb, nn, no:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "."
		}
	case lv:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
		}
	case sr:
		separator := "."
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = ". "
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month) + "."
		}
	case cs, sk, sl:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case hr:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case ro, ru:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case de, dsb, fi, gsw, hsb, lb, is, smn:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "."
		}
	case he, jgo:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case hu, ko:
		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + ". " + fmtDay(d, opts.Day) + "."
		}
	case wae:
		fmtMonth = fmtMonthName(locale.String(), "format", "abbreviated")

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month)
		}
	case bs:
		if script == cyrl {
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
			}
		}

		separator := ". "
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			separator = "."
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + separator + fmtMonth(m, MonthNumeric) + "."
		}
	case om:
		if opts.Month == Month2Digit || opts.Day == Day2Digit {
			break
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
		}
	case ks:
		if script == deva {
			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}

		fallthrough
	case ak, am, asa, bem, blo, bez, brx, ceb, cgg, chr, dav, ebu, ee, eu, fil, guz, ha, kam, kde, kln, teo, vai, ja, jmc,
		ki, ksb, lag, lg, luo, luy, mas, mer, naq, nd, nyn, rof, rwk, saq, sbp, so, tzm, vun, xh, xog, yue:
		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case mn:
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")

		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + "/" + fmtDay(d, Day2Digit)
		}
	case zh:
		separator := "/"

		switch region {
		case regionSG:
			separator = "-"
		case regionHK, regionMO:
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day)
		}
	case fr:
		switch region {
		case regionCA:
			if opts.Month == Month2Digit || opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
			}
		case regionCH:
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return func(m time.Month, d int) string {
					return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "."
				}
			}

			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
			}
		}

		fallthrough
	case br, ga, it, jv, kkj, sc, syr, vec, uz:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit)
		}
	case pcm:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + " /" + fmtMonth(m, opts.Month)
		}
	case sv:
		separator := "/"
		if region == regionFI {
			separator = "."
		}

		if opts.Month == Month2Digit && opts.Day == DayNumeric {
			opts.Month = MonthNumeric
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
		}
	case ti:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case kea, pt:
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case hi:
		if script != latn {
			break
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + " " + era
		}
	case ar:
		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "\u200f/" + fmtMonth(m, opts.Month)
		}
	case lrc:
		if region == regionIQ {
			return func(m time.Month, d int) string {
				return era + " " + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}
		}
	case en:
		separator := "/"

		switch region {
		case regionAS, regionBI, regionPH, regionPR, regionUM, regionUS, regionVI:
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + " " + era
			}
		case regionAU, regionBE, regionIE, regionNZ, regionZW:
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + " " + era
			}
		case regionCA:
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day) + " " + era
			}
		case regionCH:
			separator = "."
		case regionZA:
			if opts.Month == Month2Digit && opts.Day == Day2Digit {
				break
			}

			return func(m time.Month, d int) string {
				return fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit) + " " + era
			}
		case regionGU, regionMH, regionMP, regionZZ:
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + " " + era
			}
		}

		if script == shaw || script == dsrt {
			return func(m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + " " + era
			}
		}

		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month) + " " + era
		}
	case es:
		switch region {
		case regionCL:
			separator := "-"
			if opts.Month == Month2Digit {
				separator = "/"
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + separator + fmtMonth(m, opts.Month)
			}
		case regionMX, regionUS:
			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		case regionPA, regionPR:
			if opts.Month == MonthNumeric {
				return func(m time.Month, d int) string {
					return era + " " + fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit)
				}
			}

			if opts.Month == Month2Digit {
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return func(m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return era + " " + fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
		}
	}

	return func(m time.Month, d int) string {
		// g M/d   119
		// g d/M   166
		// g dd/MM 73
		// g MM/dd 2
		// g M-d   7
		// g d-M   28
		// g MM-dd 99
		// g dd-MM 19
		return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
	}
}

func fmtEraMonthDayPersian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)
	separator := "-"

	switch lang {
	case fa, ps:
		separator = "/"
	}

	return func(m time.Month, d int) string {
		return era + " " + fmtMonth(m, opts.Month) + separator + fmtDay(d, opts.Day)
	}
}

func fmtEraMonthDayBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	return func(m time.Month, d int) string {
		return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
	}
}
