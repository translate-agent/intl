package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop
func fmtEraYearDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearGregorian(locale, digits, opts.Year)
	dayName := cldr.UnitName(locale).Day

	const (
		eraYearDay = iota
		eraDayYear
	)

	layout := eraYearDay
	prefix := ""
	middle := " " + era + " (" + dayName + ": "
	suffix := ")"

	switch lang {
	case cldr.BE, cldr.RU:
		middle = " г. " + era + " (" + dayName + ": "
	case cldr.CV:
		middle = " ҫ. " + era + " (" + dayName + ": "
	case cldr.KK:
		prefix = era + " "
		middle = " ж. (" + dayName + ": "
	case cldr.KY:
		prefix = era + " "
		middle = "-ж. (" + dayName + ": "
	case cldr.HY:
		prefix = era + " "
		middle = " թ. (" + dayName + ": "
	case cldr.TT:
		prefix = era + " "
		middle = " ел (" + dayName + ": "
	case cldr.SAH:
		middle = " с. " + era + " (" + dayName + ": "
	case cldr.LT:
		opts.Day = Day2Digit
		middle = " m. " + era + " (" + dayName + ": "
	case cldr.BG, cldr.CY, cldr.MK:
		middle = " " + era + " (" + dayName + ": "
	case cldr.BS:
		if script != cldr.Cyrl {
			suffix = ".)"
		}
	case cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AZ, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BGC, cldr.BHO, cldr.BM,
		cldr.BO, cldr.CE, cldr.CGG, cldr.CKB, cldr.CSW, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.DZ,
		cldr.EBU, cldr.EO, cldr.EU, cldr.EWO, cldr.FUR, cldr.FY, cldr.GAA, cldr.GSW, cldr.GU, cldr.GUZ, cldr.GV, cldr.HA,
		cldr.HU, cldr.IG, cldr.JGO, cldr.JMC, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN, cldr.KN,
		cldr.KSB, cldr.KSF, cldr.KSH, cldr.KU, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.LN, cldr.LO,
		cldr.LRC, cldr.LU, cldr.LUO, cldr.LUY, cldr.LV, cldr.MAS, cldr.MER, cldr.MFE, cldr.MG, cldr.MGH, cldr.MGO, cldr.ML,
		cldr.MN, cldr.MNI, cldr.MR, cldr.MT, cldr.MUA, cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NE, cldr.NMG, cldr.NNH,
		cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OM, cldr.OS, cldr.PA, cldr.PCM, cldr.PRG, cldr.PS, cldr.QU,
		cldr.RAJ, cldr.RN, cldr.ROF, cldr.RW, cldr.RWK, cldr.SAQ, cldr.SAT, cldr.SBP, cldr.SEH, cldr.SES, cldr.SG,
		cldr.SHI, cldr.SI, cldr.SN, cldr.ST, cldr.SZL, cldr.TA, cldr.TE, cldr.TEO, cldr.TK, cldr.TN, cldr.TOK, cldr.TR,
		cldr.TWQ, cldr.TZM, cldr.VAI, cldr.VMW, cldr.VUN, cldr.WAE, cldr.XOG, cldr.YAV, cldr.YI, cldr.YO, cldr.ZA,
		cldr.ZGH, cldr.ZU:
		prefix = era + " "
		middle = " (" + dayName + ": "
	case cldr.UZ:
		if script != cldr.Arab {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.BRX:
		prefix = era
		middle = " (" + dayName + ": "
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		suffix = ".)"
	case cldr.FF:
		if script != cldr.Adlm {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.HI:
		if script == cldr.Latn {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.KS:
		if script == cldr.Deva {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.SD:
		if script != cldr.Deva {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.JA, cldr.YUE, cldr.ZH:
		prefix = era
		middle = " (" + dayName + ": "
		suffix = dayName + ")"
	case cldr.KO:
		prefix = era + " "
		middle = " (" + dayName + ": "
		suffix = dayName + ")"
	case cldr.II:
		prefix = era + " "
		middle = " (" + dayName + ": "
		suffix = "ꑍ)"
	case cldr.KXV:
		if script == cldr.Deva || script == cldr.Orya || script == cldr.Telu {
			prefix = ""
			middle = " " + era + " (" + dayName + ": "
		}
	case cldr.SE:
		if region != cldr.RegionFI {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	case cldr.KOK:
		if script != cldr.Latn {
			prefix = era + " "
			middle = " (" + dayName + ": "
		}
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // eraYearDay
		return func(t cldr.TimeReader) string {
			return prefix + year(t) + middle + dayDigits(t) + suffix
		}
	case eraDayYear:
		return func(t cldr.TimeReader) string {
			return prefix + dayDigits(t) + middle + year(t) + suffix
		}
	}
}

func fmtEraYearDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigits(digits, opts.Year)
	dayName := cldr.UnitName(locale).Day

	prefix := ""
	middle := " (" + dayName + ": "
	suffix := ")"

	switch lang {
	case cldr.UZ:
		if region == cldr.RegionAF {
			era := fmtEra(locale, opts.Era)
			prefix = era + " "
		}
	case cldr.FA:
		era := fmtEra(locale, opts.Era)
		middle = " " + era + " (" + dayName + ": "
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v cldr.TimeReader) string {
		return prefix + year(yearDigits(v)) + middle + dayDigits(v) + suffix
	}
}

func fmtEraYearDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	dayDigits := convertDayDigits(digits, opts.Day)
	dayName := cldr.UnitName(locale).Day
	middle, suffix := " ("+dayName+": ", ")"

	return func(t cldr.TimeReader) string {
		return year(t) + middle + dayDigits(t) + suffix
	}
}
