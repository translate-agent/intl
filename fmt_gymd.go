package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraYearMonthDayGregorian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, script, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	case ff:
		if script == adlm {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month) + "-" + fmtYear(y, opts.Year) + " " + era
			}
		}
	case en:
		switch region {
		default:
			if script == dsrt || script == shaw || region == regionZZ {
				return func(y int, m time.Month, d int) string {
					return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year) + " " + era
			}
		case regionAE, regionAS, regionBI, regionCA, regionGU, regionMH, regionMP, regionPH, regionPR, regionUM, regionUS,
			regionVI:
			return func(y int, m time.Month, d int) string {
				return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
			}
		case regionCH:
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year) + " " + era
			}
		case regionGB:
			if script == shaw {
				return func(y int, m time.Month, d int) string {
					return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
				}
			}

			return func(y int, m time.Month, d int) string {
				return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year) + " " + era
			}
		}
	case brx, lv, mni:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtYear(y, opts.Year)
		}
	case da, dsb, hsb, ka, mk, sq:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year) + " " + era
		}
	case be, cv, de, fo, hy, nb, nn, no, ro, ru:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " " + era
		}
	case et, pl:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " " + era
		}
	case sr:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + ". " + era
		}
	case bg:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " г. " + era
		}
	case uz:
		if script != cyrl {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " " + era
			}
		}
	case fi:
		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "." + fmtDay(d, opts.Day) + "." + fmtYear(y, opts.Year) + " " + era
		}
	case fr:
		if region == regionCA {
			return func(y int, m time.Month, d int) string {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) + " " + era
			}
		}

		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year) + " " + era
		}
	case ga, it, kea, pt, sc, syr, vec:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year) + " " + era
		}
	case as, es, gd, gl, he, el, id, is, jv, nl, su, sw, ta, ti, tr, xnr, ur, vi, yo:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year) + " " + era
		}
	case bs:
		if script != cyrl {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year) + ". " + era
			}
		}
	case am, ceb, chr, cy, blo, fil, ml, ne, ps, sd, so, xh, zu:
		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
		}
	case ks:
		if script == deva {
			break
		}

		return func(y int, m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
		}
	case ar, ia, bn, ca, mai, rm, uk, wo:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtYear(y, opts.Year) + " " + era
		}
	case kk:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) + "-" + era + " " + fmtYear(y, opts.Year)
		}
	case lt, sv:
		return func(y int, m time.Month, d int) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) + " " + era
		}
	case az:
		if script != cyrl {
			fmtMonth = fmtMonthName(locale.String(), "format", "abbreviated")

			return func(y int, m time.Month, d int) string {
				return era + " " + fmtDay(d, opts.Day) + " " + fmtMonth(m, opts.Month) + " " + fmtYear(y, opts.Year)
			}
		}
	case pa:
		if script != arab {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + era + " " + fmtYear(y, opts.Year)
			}
		}
	case ku, tk:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
		}
	case hu:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + ". " + fmtMonth(m, Month2Digit) + ". " + fmtDay(d, Day2Digit) + "."
		}
	case cs, sk, sl:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + ". " + fmtYear(y, opts.Year) + " " + era
		}
	case hr:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + ". " + fmtYear(y, opts.Year) + ". " + era
		}
	case hi:
		if script == latn {
			return func(y int, m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year) + " " + era
			}
		}

		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case zh:
		if script == hant {
			return func(y int, m time.Month, d int) string {
				return era + " " + fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		}
	case kxv:
		if script == deva || script == orya || script == telu {
			break
		}

		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case ja:
		return func(y int, m time.Month, d int) string {
			return era + fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case ko, my:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case mr, qu:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case lo:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + era + " " + fmtYear(y, opts.Year)
		}
	case to:
		return func(y int, m time.Month, d int) string {
			return fmtDay(d, Day2Digit) + " " + fmtMonth(m, Month2Digit) + " " + fmtYear(y, opts.Year) + " " + era
		}
	}

	return func(y int, m time.Month, d int) string {
		return era + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
	}
}

func fmtEraYearMonthDayPersian(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	lang, _, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang {
	case ckb:
		if region != regionIR {
			break
		}

		fallthrough
	case lrc, mzn, uz:
		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case ps:
		if opts.Era == EraNarrow {
			return func(y int, m time.Month, d int) string {
				return era + " " + fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
			}
		}

		return func(y int, m time.Month, d int) string {
			return era + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	}

	return func(y int, m time.Month, d int) string {
		return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day) + "/" + fmtYear(y, opts.Year) + " " + era
	}
}

func fmtEraYearMonthDayBuddhist(
	locale language.Tag,
	digits digits,
	opts Options,
) func(y int, m time.Month, d int) string {
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	return func(y int, m time.Month, d int) string {
		return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month) + "/" + era + " " + fmtYear(y, opts.Year)
	}
}
