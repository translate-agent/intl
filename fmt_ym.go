package intl

import (
	"cmp"
	"time"

	"golang.org/x/text/language"
)

//nolint:cyclop
func fmtYearMonthGregorian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch language, _ := locale.Base(); language.String() {
	default:
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "-" + fmtMonth(m, Month2Digit)
		}
	case "af", "as", "ia", "jv", "mi", "rm", "tg", "wo":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "-" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "agq", "ak", "am", "asa", "ast", "bas", "bem", "bez", "blo", "bm", "brx", "ca", "ceb", "cgg", "chr", "ckb",
		"cs", "cy", "dav", "dje", "doi", "dua", "dyo", "ebu", "ee", "el", "en", "ewo", "ff", "fil", "fur", "gd", "gl", "guz",
		"ha", "haw", "hi", "id", "ig", "jmc", "kab", "kam", "kde", "khq", "ki", "kln", "km", "ks", "ksb", "ksf", "kxv",
		"lag", "lg", "ln", "lo", "lu", "luo", "luy", "mai", "mas", "mer", "mfe", "mg", "mgh", "mni", "mua", "my", "naq",
		"nd", "nmg", "nus", "nyn", "pa", "pcm", "rn", "rof", "rwk", "sa", "saq", "sbp", "ses", "sg", "shi", "sk", "sl",
		"so", "su", "sw", "teo", "twq", "tzm", "ur", "vai", "vun", "xh", "xnr", "xog", "yav", "yo", "zgh":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "ar":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "\u200f/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "az", "cv", "fo", "hy", "kk", "ku", "os", "pl", "ro", "ru", "tk", "tt", "uk", "uz":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "be", "da", "dsb", "et", "hsb", "ka", "lb", "mk", "nb", "nn", "no", "smn", "sq":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "bg":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + " г."
		}
	case "bn", "ccp", "gu", "kn", "mr", "or", "ta", "te", "to":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			}

			return fmtMonth(m, Month2Digit) + "-" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "br", "fr", "ga", "it", "iu", "kea", "kgp", "pt", "sc", "seh", "syr", "vec", "yrl":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "bs":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			}

			return fmtMonth(m, MonthNumeric) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "de":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			}

			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "dz", "si":
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "-" + fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))
		}
	case "es", "ti":
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "eu", "ja", "yue":
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "/" + fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))
		}
	case "fi", "he":
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "fy", "kok", "ms", "nl", "ug":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "-" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "gsw":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "-" + fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))
			}

			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "hr":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + ". " + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "."
		}
	case "hu":
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + ". " + fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "."
		}
	case "is":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + ". " + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "kkj":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + " " + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "ko":
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + ". " + fmtMonth(m, MonthNumeric) + "."
		}
	case "lv":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "."
		}
	case "mn":
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")

		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + " " + fmtMonth(m, opts.Month)
		}
	case "om", "yi":
		return func(y int, m time.Month) string {
			ys := fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			ms := fmtMonth(m, Month2Digit)

			if opts.Month == MonthNumeric {
				return ys + "-" + ms
			}

			return ms + "/" + ys
		}
	case "sr":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + ". " + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "."
			}

			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "."
		}
	case "tr":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			}

			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "vi":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			}

			return "tháng " + fmtMonth(m, Month2Digit) + ", " + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	case "zh":
		return func(y int, m time.Month) string {
			ys := fmtYear(y, cmp.Or(opts.Year, YearNumeric))
			ms := fmtMonth(m, MonthNumeric)

			if opts.Month == MonthNumeric {
				return ys + "/" + ms
			}

			return ys + "年" + ms + "月"
		}
	}
}

func fmtYearMonthBuddhist(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch language, _ := locale.Base(); language.String() {
	default:
		return func(y int, m time.Month) string {
			return fmtEra(locale) + " " + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "-" + fmtMonth(m, Month2Digit)
		}
	case "th":
		return func(y int, m time.Month) string {
			return fmtMonth(m, cmp.Or(opts.Month, MonthNumeric)) + "/" + fmtYear(y, cmp.Or(opts.Year, YearNumeric))
		}
	}
}

func fmtYearMonthPersian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch language, _ := locale.Base(); language.String() {
	default:
		return func(y int, m time.Month) string {
			return fmtEra(locale) + " " + fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "-" + fmtMonth(m, Month2Digit)
		}
	case "fa":
		return func(y int, m time.Month) string {
			return fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "/" + fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))
		}
	case "ps":
		return func(y int, m time.Month) string {
			return fmtEra(locale) + " " +
				fmtYear(y, cmp.Or(opts.Year, YearNumeric)) + "/" +
				fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))
		}
	}
}
