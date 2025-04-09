package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop
func fmtEraYearMonthGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearGregorian(locale, digits, opts.Year)
	monthName := cldr.UnitName(locale).Month

	const (
		// eraYearMonth includes "era year month" and "year month era".
		eraYearMonth = iota
		// eraMonthYear includes "era month year" and "month year era".
		eraMonthYear
	)

	layout := eraMonthYear
	prefix := ""
	middle := " "
	suffix := " " + era

	switch lang {
	case cldr.AZ, cldr.QU, cldr.TE, cldr.TK, cldr.TR:
		prefix = era + " "
		suffix = ""
	case cldr.BE, cldr.RU:
		suffix = " г. " + era
	case cldr.BG:
		middle = "."
		suffix = " " + era

		if opts.Month.numeric() {
			opts.Month = Month2Digit
		}
	case cldr.CY, cldr.MK:
		middle = " "
		suffix = " " + era
	case cldr.CV:
		suffix = " ҫ. " + era
	case cldr.HI:
		if script != cldr.Latn {
			middle = " " + era + " "
			suffix = ""

			break
		}

		fallthrough
	case cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BGC, cldr.BHO, cldr.BM, cldr.BO,
		cldr.BRX, cldr.CE, cldr.CGG, cldr.CKB, cldr.CSW, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EBU,
		cldr.EO, cldr.EWO, cldr.FUR, cldr.GAA, cldr.GSW, cldr.GUZ, cldr.GV, cldr.HA, cldr.HU, cldr.II, cldr.JGO, cldr.JMC,
		cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN, cldr.KN, cldr.KO, cldr.KSB, cldr.KSF, cldr.KSH,
		cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.LN, cldr.LRC, cldr.LU, cldr.LUO, cldr.LUY, cldr.LV,
		cldr.MAS, cldr.MG, cldr.MER, cldr.MFE, cldr.MGH, cldr.ML, cldr.MGO, cldr.MNI, cldr.MUA, cldr.MY, cldr.NAQ, cldr.ND,
		cldr.NDS, cldr.NE, cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OS, cldr.PCM,
		cldr.PRG, cldr.PS, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RW, cldr.RWK, cldr.SAH, cldr.SAQ, cldr.SAT, cldr.SBP,
		cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SI, cldr.SN, cldr.ST, cldr.SZL, cldr.TA, cldr.TEO, cldr.TN, cldr.TOK,
		cldr.TWQ, cldr.TZM, cldr.VAI, cldr.VMW, cldr.VUN, cldr.WAE, cldr.XOG, cldr.YAV, cldr.YI, cldr.YO, cldr.ZA,
		cldr.ZGH, cldr.ZU:
		layout = eraYearMonth
		prefix = era + " "
		suffix = ""
	case cldr.SE:
		if region != cldr.RegionFI {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case cldr.SD:
		if script != cldr.Deva {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case cldr.KS:
		if script == cldr.Deva {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case cldr.IG, cldr.KXV, cldr.MAI, cldr.MR, cldr.SA, cldr.XNR:
		middle = " " + era + " "
		suffix = ""
	case cldr.PA:
		if script == cldr.Arab {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""

			break
		}

		fallthrough
	case cldr.GU, cldr.LO, cldr.UZ:
		middle = ", " + era + " "
		suffix = ""
	case cldr.KGP, cldr.WO:
		middle = ", "
		suffix = " " + era
	case cldr.TT:
		layout = eraYearMonth
		prefix = era + " "
		middle = " ел, "
		suffix = ""
	case cldr.ES:
		if region != cldr.RegionCO {
			break
		}

		fallthrough
	case cldr.GL, cldr.PT:
		middle = " de "
	case cldr.YUE, cldr.ZH:
		opts.Month = MonthNumeric
		layout = eraYearMonth
		prefix = era
		middle = ""
		suffix = monthName
	case cldr.DZ:
		layout = eraYearMonth
		prefix = era + " "
		middle = " སྤྱི་ཟླ་"
		suffix = ""
	case cldr.JA:
		layout = eraYearMonth
		prefix = era
		middle = ""
		opts.Month = MonthNumeric
		suffix = "月"
	case cldr.EU:
		layout = eraYearMonth
		prefix = era + " "
		middle = ". urteko "
		suffix = ""
	case cldr.FF:
		if script != cldr.Adlm {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	case cldr.HY:
		layout = eraYearMonth
		prefix = era + " "
		middle = " թ. "
		suffix = ""
	case cldr.KK:
		layout = eraYearMonth
		prefix = era + " "
		middle = " ж. "
		suffix = ""
	case cldr.KA:
		middle = ". "
	case cldr.KU:
		prefix = era + " "
		middle = "a "
		suffix = "an"
	case cldr.KY:
		layout = eraYearMonth
		prefix = era + " "
		middle = "-ж. "
		suffix = ""
	case cldr.LT:
		opts.Month = Month2Digit
		layout = eraYearMonth
		middle = "-"
	case cldr.MN:
		layout = eraYearMonth
		prefix = era + " "
		middle = " оны "
		suffix = ""
	case cldr.SL:
		month = fmtMonthName(locale.String(), "format", "abbreviated")
	case cldr.UG:
		layout = eraYearMonth
	case cldr.UK:
		suffix = " р. " + era
	case cldr.KOK:
		if script != cldr.Latn {
			layout = eraYearMonth
			prefix = era + " "
			suffix = ""
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	switch layout {
	default: // eraYearMonth
		return func(t cldr.TimeReader) string {
			return prefix + year(t) + middle + month(t) + suffix
		}
	case eraMonthYear:
		return func(t cldr.TimeReader) string {
			return prefix + month(t) + middle + year(t) + suffix
		}
	}
}

func fmtEraYearMonthPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		// eraYearMonth includes "era year month" and "year month era".
		eraYearMonth = iota
		// eraMonthYear includes "era month year" and "month year era".
		eraMonthYear
	)

	layout := eraYearMonth
	prefix := era + " "
	middle := " "
	suffix := ""

	switch lang {
	case cldr.FA:
		layout = eraMonthYear
		prefix = ""
		middle = " "
		suffix = " " + era
	case cldr.CKB, cldr.UZ:
		if region != cldr.RegionAF {
			prefix = ""
		}
	case cldr.LRC, cldr.MZN, cldr.PS:
		prefix = ""
	}

	month := convertMonthDigits(digits, opts.Month)

	switch layout {
	default: // eraYearMonth
		return func(v cldr.TimeReader) string {
			return prefix + year(yearDigits(v)) + middle + month(v) + suffix
		}
	case eraMonthYear:
		return func(v cldr.TimeReader) string {
			return prefix + month(v) + middle + year(yearDigits(v)) + suffix
		}
	}
}

func fmtEraYearMonthBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)

	return func(t cldr.TimeReader) string {
		return monthDigits(t) + " " + year(t)
	}
}
