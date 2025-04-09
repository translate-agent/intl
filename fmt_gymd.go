package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraYearMonthDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		eraYearMonthDay = iota
		eraMonthDayYear
		eraDayMonthYear
		dayMonthEraYear
	)

	monthOpt, monthDay := Month2Digit, Day2Digit
	layout := eraYearMonthDay
	prefix := era + " "
	suffix := ""
	separator := "-"

	switch lang {
	case cldr.EN:
		switch region {
		default:
			monthOpt, monthDay = opts.Month, opts.Day
			separator = "/"
			prefix = ""
			suffix = " " + era

			if script == cldr.Dsrt || script == cldr.Shaw || region == cldr.RegionZZ {
				layout = eraMonthDayYear
			} else {
				layout = eraDayMonthYear
			}
		case cldr.RegionAE, cldr.RegionAS, cldr.RegionBI, cldr.RegionCA, cldr.RegionGU, cldr.RegionMH, cldr.RegionMP,
			cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionUS,
			cldr.RegionVI:
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraMonthDayYear
			separator = "/"
			prefix = ""
			suffix = " " + era
		case cldr.RegionCH:
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "."
			prefix = ""
			suffix = " " + era
		case cldr.RegionGB:
			separator = "/"
			prefix = ""
			suffix = " " + era

			if script == cldr.Shaw {
				monthOpt, monthDay = opts.Month, opts.Day
				layout = eraMonthDayYear
			} else {
				layout = eraDayMonthYear
			}
		}
	case cldr.BRX, cldr.LV, cldr.MNI:
		layout = eraDayMonthYear
	case cldr.DA, cldr.DSB, cldr.HSB, cldr.IE, cldr.KA, cldr.SQ:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case cldr.MK:
		monthOpt = opts.Month
		monthDay = opts.Day
		layout = eraDayMonthYear
		prefix = ""
		suffix = " г. " + era
		separator = "."
	case cldr.ET, cldr.PL:
		monthDay = opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case cldr.BE, cldr.CV, cldr.DE, cldr.FO, cldr.HY, cldr.NB, cldr.NN, cldr.NO, cldr.RO, cldr.RU:
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case cldr.SR:
		monthDay = opts.Day
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = ". " + era
	case cldr.BG:
		layout = eraDayMonthYear
		separator = "."
		prefix = ""
		suffix = " г. " + era
	case cldr.FI:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		separator = "."
		prefix = ""
		suffix = " " + era
	case cldr.FR:
		prefix = ""
		suffix = " " + era

		if region != cldr.RegionCA {
			layout = eraDayMonthYear
			separator = "/"
		}
	case cldr.AM, cldr.AS, cldr.ES, cldr.GD, cldr.GL, cldr.HE, cldr.EL, cldr.ID, cldr.IS, cldr.JV, cldr.NL, cldr.SU,
		cldr.SW, cldr.TA, cldr.XNR, cldr.UR, cldr.VI, cldr.YO:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case cldr.GA, cldr.IT, cldr.KEA, cldr.PT, cldr.SC, cldr.SYR, cldr.VEC:
		layout = eraDayMonthYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case cldr.CEB, cldr.CHR, cldr.BLO, cldr.FIL, cldr.KAA, cldr.MHN, cldr.ML, cldr.NE, cldr.OR, cldr.PS, cldr.SD,
		cldr.SO, cldr.TI, cldr.XH, cldr.ZU:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		separator = "/"
		prefix = ""
		suffix = " " + era
	case cldr.CY:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraMonthDayYear
		prefix = ""
		suffix = " " + era
		separator = "/"
	case cldr.AR, cldr.IA, cldr.BN, cldr.CA, cldr.MAI, cldr.RM, cldr.UK, cldr.WO:
		layout = eraDayMonthYear
		prefix = ""
		suffix = " " + era
	case cldr.LT, cldr.SV:
		prefix = ""
		suffix = " " + era
	case cldr.BS:
		if script != cldr.Cyrl {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = ". "
			prefix = ""
			suffix = ". " + era
		}
	case cldr.FF:
		if script == cldr.Adlm {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			prefix = ""
			suffix = " " + era
		}
	case cldr.KS:
		if script != cldr.Deva {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraMonthDayYear
			separator = "/"
			prefix = ""
			suffix = " " + era
		}
	case cldr.UZ:
		if script != cldr.Cyrl {
			layout = eraDayMonthYear
			separator = "."
			prefix = ""
			suffix = " " + era
		}
	case cldr.AZ:
		if script != cldr.Cyrl {
			month = fmtMonthName(locale.String(), "format", "abbreviated")
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = " "
		}
	case cldr.KU, cldr.TK, cldr.TR:
		layout = eraDayMonthYear
		separator = "."
	case cldr.HU:
		separator = ". "
		suffix = "."
	case cldr.CS, cldr.SK, cldr.SL:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = " " + era
	case cldr.HR:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = ". "
		prefix = ""
		suffix = ". " + era
	case cldr.HI:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"

		if script == cldr.Latn {
			prefix = ""
			suffix = " " + era
		}
	case cldr.ZH:
		if script == cldr.Hant {
			monthOpt, monthDay = opts.Month, opts.Day
			separator = "/"
		}
	case cldr.KXV:
		if script != cldr.Deva && script != cldr.Orya && script != cldr.Telu {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = eraDayMonthYear
			separator = "/"
		}
	case cldr.JA:
		monthOpt, monthDay = opts.Month, opts.Day
		separator = "/"
		prefix = era
	case cldr.KO, cldr.MY:
		monthOpt, monthDay = opts.Month, opts.Day
		separator = "/"
	case cldr.MR, cldr.QU:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = eraDayMonthYear
		separator = "/"
	case cldr.TO:
		layout = eraDayMonthYear
		separator = " "
		prefix = ""
		suffix = " " + era
	case cldr.KK:
		layout = dayMonthEraYear
	case cldr.LO:
		monthOpt, monthDay = opts.Month, opts.Day
		layout = dayMonthEraYear
		separator = "/"
	case cldr.PA:
		if script != cldr.Arab {
			monthOpt, monthDay = opts.Month, opts.Day
			layout = dayMonthEraYear
			separator = "/"
		}
	case cldr.KOK:
		if script == cldr.Latn {
			monthOpt = opts.Month
			monthDay = opts.Day
			layout = eraDayMonthYear
			prefix = ""
			suffix = " " + era
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, monthOpt)
	}

	dayDigits := convertDayDigits(digits, monthDay)

	switch layout {
	default: // eraYearMonthDay
		return func(t timeReader) string {
			return prefix + yearDigits(t) + separator + month(t) + separator + dayDigits(t) + suffix
		}
	case eraMonthDayYear:
		return func(t timeReader) string {
			return prefix + month(t) + separator + dayDigits(t) + separator + yearDigits(t) + suffix
		}
	case eraDayMonthYear:
		return func(t timeReader) string {
			return prefix + dayDigits(t) + separator + month(t) + separator + yearDigits(t) + suffix
		}
	case dayMonthEraYear:
		return func(t timeReader) string {
			return dayDigits(t) + separator + month(t) + separator + era + " " + yearDigits(t)
		}
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
		return func(v timeReader) string {
			return month(v) + separator + dayDigits(v) + separator + yearDigits(v) + suffix
		}
	case eraYearMonthDay:
		prefix := era + " "

		return func(v timeReader) string {
			return prefix + yearDigits(v) + separator + month(v) + separator + dayDigits(v)
		}
	}
}

func fmtEraYearMonthDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t timeReader) string {
		return dayDigits(t) + "/" + monthDigits(t) + "/" + year(t)
	}
}
