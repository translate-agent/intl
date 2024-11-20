package intl

import (
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
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}
	case "af", "as", "ia", "jv", "mi", "rm", "tg", "wo":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "-" + fmtYear(y, opts.Year)
		}
	case "agq", "ak", "am", "asa", "ast", "bas", "bem", "bez", "blo", "bm", "brx", "ca", "ceb", "cgg", "chr", "ckb",
		"cs", "cy", "dav", "dje", "doi", "dua", "dyo", "ebu", "ee", "el", "en", "ewo", "ff", "fil", "fur", "gd", "gl", "guz",
		"ha", "haw", "hi", "id", "ig", "jmc", "kab", "kam", "kde", "khq", "ki", "kln", "km", "ks", "ksb", "ksf", "kxv",
		"lag", "lg", "ln", "lo", "lu", "luo", "luy", "mai", "mas", "mer", "mfe", "mg", "mgh", "mni", "mua", "my", "naq",
		"nd", "nmg", "nus", "nyn", "pa", "pcm", "rn", "rof", "rwk", "sa", "saq", "sbp", "ses", "sg", "shi", "sk", "sl",
		"so", "su", "sw", "teo", "twq", "tzm", "ur", "vai", "vun", "xh", "xnr", "xog", "yav", "yo", "zgh":
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case "ar":
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "\u200f/" + fmtYear(y, opts.Year)
		}
	case "az", "cv", "fo", "hy", "kk", "ku", "os", "pl", "ro", "ru", "tk", "tt", "uk", "uz":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year)
		}
	case "be", "da", "dsb", "et", "hsb", "ka", "lb", "mk", "nb", "nn", "no", "smn", "sq":
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year)
		}
	case "bg":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + " г."
		}
	case "bn", "ccp", "gu", "kn", "mr", "or", "ta", "te", "to":
		separator := "-"
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year)
		}
	case "br", "fr", "ga", "it", "iu", "kea", "kgp", "pt", "sc", "seh", "syr", "vec", "yrl":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "/" + fmtYear(y, opts.Year)
		}
	case "bs":
		if opts.Month == MonthNumeric {
			opts.Month = Month2Digit
		} else {
			opts.Month = MonthNumeric
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	case "de":
		separator := "."
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year)
		}
	case "dz", "si":
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month)
		}
	case "es", "ti":
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "/" + fmtYear(y, opts.Year)
		}
	case "eu", "ja", "yue":
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month)
		}
	case "fi", "he":
		return func(y int, m time.Month) string {
			return fmtMonth(m, MonthNumeric) + "." + fmtYear(y, opts.Year)
		}
	case "fy", "kok", "ms", "nl", "ug":
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "-" + fmtYear(y, opts.Year)
		}
	case "gsw":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtYear(y, opts.Year) + "-" + fmtMonth(m, opts.Month)
			}

			return fmtMonth(m, opts.Month) + "." + fmtYear(y, opts.Year)
		}
	case "hr":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + ". " + fmtYear(y, opts.Year) + "."
		}
	case "hu":
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case "is":
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + ". " + fmtYear(y, opts.Year)
		}
	case "kkj":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + " " + fmtYear(y, opts.Year)
		}
	case "ko":
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + ". " + fmtMonth(m, MonthNumeric) + "."
		}
	case "lv":
		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + "." + fmtYear(y, opts.Year) + "."
		}
	case "mn":
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")

		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + " " + fmtMonth(m, opts.Month)
		}
	case "om", "yi":
		return func(y int, m time.Month) string {
			ys := fmtYear(y, opts.Year)
			ms := fmtMonth(m, Month2Digit)

			if opts.Month == MonthNumeric {
				return ys + "-" + ms
			}

			return ms + "/" + ys
		}
	case "sr":
		separator := "."
		if opts.Month == MonthNumeric {
			separator = ". "
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + separator + fmtYear(y, opts.Year) + "."
		}
	case "tr":
		separator := "."
		if opts.Month == MonthNumeric {
			separator = "/"
		}

		return func(y int, m time.Month) string {
			return fmtMonth(m, Month2Digit) + separator + fmtYear(y, opts.Year)
		}
	case "vi":
		return func(y int, m time.Month) string {
			if opts.Month == MonthNumeric {
				return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
			}

			return "tháng " + fmtMonth(m, Month2Digit) + ", " + fmtYear(y, opts.Year)
		}
	case "zh":
		return func(y int, m time.Month) string {
			ys := fmtYear(y, opts.Year)
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

	if language, _ := locale.Base(); language.String() == "th" {
		return func(y int, m time.Month) string {
			return fmtMonth(m, opts.Month) + "/" + fmtYear(y, opts.Year)
		}
	}

	return func(y int, m time.Month) string {
		return fmtEra(locale) + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
	}
}

func fmtYearMonthPersian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	fmtMonth := fmtMonth(digits)

	switch language, _ := locale.Base(); language.String() {
	default:
		return func(y int, m time.Month) string {
			return fmtEra(locale) + " " + fmtYear(y, opts.Year) + "-" + fmtMonth(m, Month2Digit)
		}
	case "fa":
		return func(y int, m time.Month) string {
			return fmtYear(y, opts.Year) + "/" + fmtMonth(m, opts.Month)
		}
	case "ps":
		return func(y int, m time.Month) string {
			return fmtEra(locale) + " " +
				fmtYear(y, opts.Year) + "/" +
				fmtMonth(m, opts.Month)
		}
	}
}
