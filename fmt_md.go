package intl

import (
	"time"

	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtMonthDayGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang, _ := locale.Base(); lang.String() {
	default:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case "af", "as", "ia", "ky", "mi", "rm", "tg", "wo":
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "-" + fmtMonth(m, Month2Digit) }
	case "agq", "ast", "bas", "bm", "ca", "cy", "dje", "doi", "dua", "dyo", "el", "ewo", "ff", "fur", "gd", "gl", "haw",
		"hi", "id", "ig", "kab", "kgp", "khq", "km", "ksf", "kxv", "ln", "lo", "lu", "mai", "mfe", "mg", "mgh", "ml",
		"mni", "mua", "my", "nmg", "nus", "pa", "rn", "sa", "seh", "ses", "sg", "shi", "su", "sw", "to", "tr", "twq", "ur",
		"xnr", "yav", "yo", "yrl", "zgh":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case "ak", "am", "asa", "bem", "bez", "blo", "brx", "ceb", "cgg", "chr", "dav", "ebu", "ee", "en", "eu", "fil", "guz",
		"ha", "ja", "jmc", "kam", "kde", "ki", "kln", "ks", "ksb", "lag", "lg", "luo", "luy", "mas", "mer", "naq", "nd",
		"nyn", "rof", "rwk", "saq", "sbp", "so", "teo", "tzm", "vai", "vun", "xh", "xog", "yue", "zh":
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	case "ar":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "\u200f/" + fmtMonth(m, opts.Month)
		}
	case "az", "cv", "fo", "hy", "kk", "ku", "os", "tk", "tt", "uk":
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) }
	case "be", "da", "et", "he", "jgo", "ka":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case "bg", "pl":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, Month2Digit)
		}
	case "bn", "ccp", "gu", "kn", "mr", "ta", "te", "vi":
		if opts.Month == MonthNumeric && opts.Day == DayNumeric {
			return func(m time.Month, d int) string {
				return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
			}
		}

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case "br", "fr", "ga", "it", "jv", "kkj", "sc", "syr", "uz", "vec":
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit) }
	case "bs":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "."
			}

			return fmtDay(d, DayNumeric) + ". " + fmtMonth(m, MonthNumeric) + "."
		}
	case "cs", "sk", "sl":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case "de", "dsb", "fi", "gsw", "hsb", "is", "lb", "smn":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "."
		}
	case "dz", "si":
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case "es", "ti":
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric) }
	case "fy", "kok", "nl", "ug":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case "hr":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + ". " + fmtMonth(m, Month2Digit) + "."
			}

			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month) + "."
		}
	case "hu", "ko":
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + ". " + fmtDay(d, opts.Day) + "."
		}
	case "iu":
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "/" + fmtDay(d, Day2Digit) }
	case "kea", "pt":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + "/" + fmtMonth(m, Month2Digit)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case "lt":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, DayNumeric)
			}

			return fmtMonth(m, opts.Month) + "-" + fmtDay(d, opts.Day)
		}
	case "lv":
		return func(m time.Month, d int) string { return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit) + "." }
	case "mk":
		return func(m time.Month, d int) string {
			return fmtDay(d, DayNumeric) + "." + fmtMonth(m, opts.Month)
		}
	case "mn":
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")
		return func(m time.Month, d int) string { return fmtMonth(m, opts.Month) + "/" + fmtDay(d, Day2Digit) }
	case "ms":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + "-" + fmtMonth(m, MonthNumeric)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case "nb", "nn", "no":
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "." }
	case "om":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case "or":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtMonth(m, MonthNumeric) + "/" + fmtDay(d, DayNumeric)
			}

			return fmtDay(d, opts.Day) + "-" + fmtMonth(m, opts.Month)
		}
	case "pcm":
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + " /" + fmtMonth(m, opts.Month)
		}
	case "ro", "ru":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, Day2Digit) + "." + fmtMonth(m, Month2Digit)
			}

			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month)
		}
	case "sq":
		return func(m time.Month, d int) string { return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) }
	case "sr":
		return func(m time.Month, d int) string {
			if opts.Month == MonthNumeric && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + ". " + fmtMonth(m, MonthNumeric) + "."
			}

			return fmtDay(d, opts.Day) + "." + fmtMonth(m, opts.Month) + "."
		}
	case "sv":
		return func(m time.Month, d int) string {
			if opts.Month == Month2Digit && opts.Day == DayNumeric {
				return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
			}

			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	case "wae":
		fmtMonth = fmtMonthName(locale.String(), "stand-alone", "abbreviated")

		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + ". " + fmtMonth(m, opts.Month)
		}
	}
}

func fmtMonthDayBuddhist(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	if lang, _ := locale.Base(); lang.String() == "th" {
		return func(m time.Month, d int) string {
			return fmtDay(d, opts.Day) + "/" + fmtMonth(m, opts.Month)
		}
	}

	return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
}

func fmtMonthDayPersian(locale language.Tag, digits digits, opts Options) func(m time.Month, d int) string {
	fmtMonth := fmtMonth(digits)
	fmtDay := fmtDay(digits)

	switch lang, _ := locale.Base(); lang.String() {
	default:
		return func(m time.Month, d int) string { return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit) }
	case "fa", "ps":
		return func(m time.Month, d int) string {
			return fmtMonth(m, opts.Month) + "/" + fmtDay(d, opts.Day)
		}
	}
}