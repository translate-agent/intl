package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraYearMonthDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, script, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	year := convertYearDigitsFmt(digits, opts.Year)
	month := convertMonthDigitsFmt(digits, Month2Digit)

	separator := cldr.Text("-")

	switch lang {
	default:
		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.EN:
		switch region {
		default:
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text("/")
			suffix := cldr.Text(" " + era)

			if script == cldr.Dsrt || script == cldr.Shaw || region == cldr.RegionZZ {
				return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
			}

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		case cldr.RegionAE, cldr.RegionAS, cldr.RegionBI, cldr.RegionCA, cldr.RegionGU, cldr.RegionMH, cldr.RegionMP,
			cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionUS,
			cldr.RegionVI:
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text("/")
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
		case cldr.RegionCH:
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text(".")
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		case cldr.RegionGB:
			separator = cldr.Text("/")
			suffix := cldr.Text(" " + era)

			if script == cldr.Shaw {
				month = convertMonthDigitsFmt(digits, opts.Month)
				day := convertDayDigitsFmt(digits, opts.Day)
				return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
			}

			day := convertDayDigitsFmt(digits, Day2Digit)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}
	case cldr.BRX, cldr.LV, cldr.MNI:
		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
	case cldr.DA, cldr.DSB, cldr.HSB, cldr.IE, cldr.KA, cldr.SQ:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text(".")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.MK:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		suffix := cldr.Text(" г. " + era)
		separator = "."

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.ET, cldr.PL:
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text(".")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.BE, cldr.CV, cldr.DE, cldr.FO, cldr.HY, cldr.NB, cldr.NN, cldr.NO, cldr.RO, cldr.RU:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text(".")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.SR:
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text(".")
		suffix := cldr.Text(". " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.BG:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text(".")
		suffix := cldr.Text(" г. " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.FI:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text(".")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
	case cldr.FR:
		day := convertDayDigitsFmt(digits, Day2Digit)
		suffix := cldr.Text(" " + era)

		if region != cldr.RegionCA {
			separator = cldr.Text("/")

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}

		return cldr.Fmt{year, separator, month, separator, day, suffix}.Format
	case cldr.AM, cldr.AS, cldr.ES, cldr.GD, cldr.GL, cldr.HE, cldr.EL, cldr.ID, cldr.IS, cldr.JV, cldr.NL, cldr.SU,
		cldr.SW, cldr.TA, cldr.XNR, cldr.UR, cldr.VI, cldr.YO:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text("/")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.GA, cldr.IT, cldr.KEA, cldr.PT, cldr.SC, cldr.SYR, cldr.VEC:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text("/")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.CEB, cldr.CHR, cldr.BLO, cldr.FIL, cldr.KAA, cldr.MHN, cldr.ML, cldr.NE, cldr.OR, cldr.PS, cldr.SD,
		cldr.SO, cldr.TI, cldr.XH, cldr.ZU:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text("/")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
	case cldr.CY:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		suffix := cldr.Text(" " + era)
		separator = "/"

		return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
	case cldr.AR, cldr.IA, cldr.BN, cldr.CA, cldr.MAI, cldr.RM, cldr.UK, cldr.WO:
		day := convertDayDigitsFmt(digits, Day2Digit)
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.LT, cldr.SV:
		day := convertDayDigitsFmt(digits, Day2Digit)
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{year, separator, month, separator, day, suffix}.Format
	case cldr.BS:
		if script != cldr.Cyrl {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text(". ")
			suffix := cldr.Text(". " + era)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.FF:
		if script == cldr.Adlm {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.KS:
		if script != cldr.Deva {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text("/")
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{month, separator, day, separator, year, suffix}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.UZ:
		day := convertDayDigitsFmt(digits, Day2Digit)
		if script != cldr.Cyrl {
			separator = cldr.Text(".")
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}

		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.AZ:
		if script != cldr.Cyrl {
			month = cldr.Month(cldr.MonthNames(locale.String(), "format", "abbreviated"))
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text(" ")

			prefix := cldr.Text(era + " ")

			return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.KU, cldr.TK, cldr.TR:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text(".")
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
	case cldr.HU:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text(". ")
		suffix := cldr.Text(".")
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day, suffix}.Format
	case cldr.CS, cldr.SK, cldr.SL:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = ". "
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.HR:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = ". "
		suffix := cldr.Text(". " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.HI:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = "/"

		if script != cldr.Latn {
			prefix := cldr.Text(era + " ")

			return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
		}

		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.ZH:
		dayOpt := Day2Digit

		if script == cldr.Hant {
			month = convertMonthDigitsFmt(digits, opts.Month)
			dayOpt = opts.Day
			separator = cldr.Text("/")
		}

		day := convertDayDigitsFmt(digits, dayOpt)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.KXV:
		prefix := cldr.Text(era + " ")

		if script != cldr.Deva && script != cldr.Orya && script != cldr.Telu {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text("/")

			return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.JA:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = "/"
		prefix := cldr.Text(era)

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.KO, cldr.MY:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = "/"
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.MR, cldr.QU:
		month = convertMonthDigitsFmt(digits, opts.Month)
		day := convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text("/")
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, day, separator, month, separator, year}.Format
	case cldr.TO:
		day := convertDayDigitsFmt(digits, Day2Digit)
		separator = cldr.Text(" ")
		suffix := cldr.Text(" " + era)

		return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
	case cldr.KK:
		day := convertDayDigitsFmt(digits, Day2Digit)
		return cldr.Fmt{day, separator, month, separator, cldr.Text(era + " "), year}.Format
	case cldr.LO:
		day := convertDayDigitsFmt(digits, Day2Digit)
		month = convertMonthDigitsFmt(digits, opts.Month)
		day = convertDayDigitsFmt(digits, opts.Day)
		separator = cldr.Text("/")

		return cldr.Fmt{day, separator, month, separator, cldr.Text(era + " "), year}.Format
	case cldr.PA:
		if script != cldr.Arab {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			separator = cldr.Text("/")

			return cldr.Fmt{day, separator, month, separator, cldr.Text(era + " "), year}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	case cldr.KOK:
		if script == cldr.Latn {
			month = convertMonthDigitsFmt(digits, opts.Month)
			day := convertDayDigitsFmt(digits, opts.Day)
			suffix := cldr.Text(" " + era)

			return cldr.Fmt{day, separator, month, separator, year, suffix}.Format
		}

		day := convertDayDigitsFmt(digits, Day2Digit)
		prefix := cldr.Text(era + " ")

		return cldr.Fmt{prefix, year, separator, month, separator, day}.Format
	}
}

func fmtEraYearMonthDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		eraYearMonthDay = iota
		eraMonthDayYear
	)

	layout := eraMonthDayYear
	separator := "/"
	suffix := " " + era

	switch lang {
	case cldr.CKB:
		if region == cldr.RegionIR {
			layout = eraYearMonthDay
			separator = "-"
		}
	case cldr.LRC, cldr.MZN, cldr.UZ:
		layout = eraYearMonthDay
		separator = "-"
	case cldr.PS:
		layout = eraYearMonthDay

		if !opts.Era.narrow() {
			separator = "-"
		}
	case cldr.FA:
		suffix = " " + era
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // eraMonthDayYear
		return func(v cldr.TimeReader) string {
			return month(v) + separator + dayDigits(v) + separator + yearDigits(v) + suffix
		}
	case eraYearMonthDay:
		prefix := era + " "

		return func(v cldr.TimeReader) string {
			return prefix + yearDigits(v) + separator + month(v) + separator + dayDigits(v)
		}
	}
}

func fmtEraYearMonthDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t cldr.TimeReader) string {
		return dayDigits(t) + "/" + monthDigits(t) + "/" + year(t)
	}
}
