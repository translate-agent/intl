package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtEraMonthDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)

	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutDayMonth

	prefix := era + " "
	suffix := ""
	separator := "/"

	switch lang {
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "-"
	case cldr.SD:
		if script == cldr.Deva {
			layout = layoutMonthDay
			break
		}

		fallthrough
	case cldr.BGC, cldr.BHO, cldr.BO, cldr.CE, cldr.CKB, cldr.CSW, cldr.EO, cldr.GAA, cldr.GV, cldr.KL, cldr.KSH,
		cldr.KW, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.MGO, cldr.MT, cldr.NDS, cldr.NNH, cldr.NE, cldr.NQO, cldr.NSO, cldr.OC,
		cldr.PRG, cldr.PS, cldr.QU, cldr.RAJ, cldr.RW, cldr.SAH, cldr.SAT, cldr.SN, cldr.ST, cldr.SZL, cldr.TN, cldr.TOK,
		cldr.VMW, cldr.YI, cldr.ZA, cldr.ZU:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "-"
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
		}

		fallthrough
	case cldr.DZ, cldr.SI:
		layout = layoutMonthDay
		separator = "-"
	case cldr.NL:
		if region == cldr.RegionBE {
			break
		}

		fallthrough
	case cldr.FY, cldr.UG:
		separator = "-"
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDay
			break
		}

		separator = "-"
	case cldr.MS:
		if opts.Month.numeric() && opts.Day.numeric() {
			separator = "-"
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			break
		}

		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "-"
	case cldr.KN, cldr.MR, cldr.VI:
		if !opts.Month.numeric() || !opts.Day.numeric() {
			separator = "-"
		}
	case cldr.TI:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
	case cldr.FF:
		if script == cldr.Adlm {
			separator = "-"
		}
	case cldr.BN, cldr.CCP, cldr.GU, cldr.TA, cldr.TE:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			separator = "-"
			break
		}
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "."
	case cldr.SQ:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = "."
	case cldr.BG:
		opts.Month = Month2Digit
		prefix = era + " "
		separator = "."
	case cldr.CY:
		prefix = era + " "
	case cldr.PL:
		opts.Month = Month2Digit
		separator = "."
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		separator = "."
	case cldr.MK:
		opts.Day = DayNumeric
		prefix = era + " "
		separator = "."
	case cldr.NB, cldr.NN, cldr.NO:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = "."
		suffix = "."
	case cldr.LV:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "."
		suffix = "."
	case cldr.SR:
		separator = "."
		suffix = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			separator = ". "
		}
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cldr.CS, cldr.SK, cldr.SL:
		separator = ". "
		suffix = "."
	case cldr.RO, cldr.RU:
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.DE, cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.LB, cldr.IS, cldr.SMN:
		separator = "."
		suffix = "."
	case cldr.HU, cldr.KO:
		layout = layoutMonthDay
		separator = ". "
		suffix = "."
	case cldr.WAE:
		month = fmtMonthName(locale.String(), "format", "abbreviated")
		separator = ". "
	case cldr.BS:
		suffix = "."

		if script == cldr.Cyrl {
			separator = "."
			opts.Month = Month2Digit
			opts.Day = Day2Digit

			break
		}

		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		separator = ". "
	case cldr.OM:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			break
		}

		layout = layoutMonthDay
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		separator = "-"
	case cldr.KS:
		if script == cldr.Deva {
			layout = layoutMonthDay
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			separator = "-"

			break
		}

		fallthrough
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BLO, cldr.BEZ, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.KAM, cldr.KDE, cldr.KLN, cldr.TEO, cldr.VAI, cldr.JA,
		cldr.JMC, cldr.KI, cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.NAQ, cldr.ND,
		cldr.NYN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TZM, cldr.VUN, cldr.XH, cldr.XOG, cldr.YUE:
		layout = layoutMonthDay
	case cldr.MN:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
		opts.Day = Day2Digit
		layout = layoutMonthDay
	case cldr.ZH:
		if region == cldr.RegionHK || region == cldr.RegionMO {
			break
		}

		layout = layoutMonthDay

		if region == cldr.RegionSG {
			separator = "-"
		}
	case cldr.FR:
		if region == cldr.RegionCA {
			layout = layoutMonthDay
			separator = "-"

			if opts.Month.twoDigit() || opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			break
		}

		if region == cldr.RegionCH {
			separator = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				suffix = "."
			}

			break
		}

		fallthrough
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.VEC, cldr.UZ:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case cldr.PCM:
		separator = " /"
	case cldr.SV:
		if region == cldr.RegionFI {
			separator = "."
		}

		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
		}
	case cldr.KEA, cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.HI:
		if script != cldr.Latn {
			break
		}

		prefix = ""
		suffix = " " + era

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.AR:
		separator = "\u200f/"
	case cldr.LRC:
		if region == cldr.RegionIQ {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutMonthDay
			separator = "-"
		}
	case cldr.EN:
		prefix = ""
		suffix = " " + era

		switch region {
		case cldr.RegionUS, cldr.RegionAS, cldr.RegionBI, cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionVI:
			layout = layoutMonthDay
			goto breakEN
		case cldr.RegionAU, cldr.RegionBE, cldr.RegionIE, cldr.RegionNZ, cldr.RegionZW:
			goto breakEN
		case cldr.RegionGU, cldr.RegionMH, cldr.RegionMP, cldr.RegionZZ:
			layout = layoutMonthDay
			goto breakEN
		case cldr.RegionCA:
			layout = layoutMonthDay
			separator = "-"
		case cldr.RegionCH:
			separator = "."
		case cldr.RegionZA:
			if !opts.Month.twoDigit() || !opts.Day.twoDigit() {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}

		if script == cldr.Shaw || script == cldr.Dsrt {
			layout = layoutMonthDay
			break
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	breakEN:
		break
	case cldr.ES:
		switch region {
		case cldr.RegionUS, cldr.RegionMX:
			goto breakES
		case cldr.RegionCL:
			if opts.Month.twoDigit() {
				separator = "/"
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				separator = "-"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			goto breakES
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() {
				layout = layoutMonthDay
				opts.Month = Month2Digit
				opts.Day = Day2Digit

				goto breakES
			}

			if opts.Month.twoDigit() {
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			goto breakES
		}

		opts.Month = MonthNumeric
		opts.Day = DayNumeric
	breakES:
		break
	case cldr.II:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutMonthDay
		separator = "ꆪ-"
		suffix = "ꑍ"
	case cldr.KOK:
		if script != cldr.Latn {
			separator = "-"
		}
	case cldr.KAA, cldr.MHN:
		layout = layoutMonthDay
		prefix = ""
		suffix = " " + era
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t timeReader) string {
			return prefix + dayDigits(t) + separator + month(t) + suffix
		}
	}

	return func(t timeReader) string {
		return prefix + month(t) + separator + dayDigits(t) + suffix
	}
}

func fmtEraMonthDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	prefix := era + " "
	separator := "-"

	switch lang {
	case cldr.FA, cldr.PS:
		prefix = era + " "
		separator = "/"
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v timeReader) string { return prefix + month(v) + separator + dayDigits(v) }
}

func fmtEraMonthDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	prefix := era + " "
	month := fmtMonthDayBuddhist(locale, digits, opts)

	return func(t timeReader) string { return prefix + month(t) }
}
