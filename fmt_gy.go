package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtEraYearGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, script, region := locale.Raw()
	era := fmtEra(locale, opts.Era)
	year := fmtYearGregorian(locale, digits, opts.Year)

	prefix := ""
	suffix := " " + era

	switch lang {
	case cldr.KOK:
		if script == cldr.Latn {
			break
		}

		fallthrough
	case cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AZ, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BGC, cldr.BHO, cldr.BM,
		cldr.BO, cldr.CE, cldr.CGG, cldr.CKB, cldr.CSW, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DZ, cldr.DYO,
		cldr.EBU, cldr.EO, cldr.EU, cldr.EWO, cldr.FUR, cldr.FY, cldr.GAA, cldr.GSW, cldr.GU, cldr.GUZ, cldr.GV, cldr.HA,
		cldr.HU, cldr.IG, cldr.II, cldr.JMC, cldr.JGO, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN,
		cldr.KN, cldr.KO, cldr.KSB, cldr.KSF, cldr.KSH, cldr.KU, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO,
		cldr.LN, cldr.LO, cldr.LRC, cldr.LV, cldr.LU, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.MFE, cldr.MG, cldr.MGH,
		cldr.MGO, cldr.ML, cldr.MN, cldr.MNI, cldr.MR, cldr.MT, cldr.MUA, cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NE,
		cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OM, cldr.OS, cldr.PA, cldr.PCM, cldr.PRG,
		cldr.PS, cldr.QU, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RW, cldr.RWK, cldr.SAQ, cldr.SAT, cldr.SBP, cldr.SEH, cldr.SES,
		cldr.SG, cldr.SHI, cldr.SI, cldr.SN, cldr.ST, cldr.SZL, cldr.TA, cldr.TE, cldr.TEO, cldr.TK, cldr.TN, cldr.TOK,
		cldr.TR, cldr.TWQ, cldr.TZM, cldr.UZ, cldr.VAI, cldr.VMW, cldr.VUN, cldr.WAE, cldr.XOG, cldr.YAV, cldr.YI, cldr.YO,
		cldr.ZA, cldr.ZGH, cldr.ZU:
		prefix = era + " "
		suffix = ""
	case cldr.KS:
		if script == cldr.Deva {
			prefix = era + " "
			suffix = ""
		}
	case cldr.HI:
		if script == cldr.Latn {
			prefix = era + " "
			suffix = ""
		}
	case cldr.SD:
		if script != cldr.Deva {
			prefix = era + " "
			suffix = ""
		}
	case cldr.FF:
		if script != cldr.Adlm {
			prefix = era + " "
			suffix = ""
		}
	case cldr.SE:
		if region != cldr.RegionFI {
			prefix = era + " "
			suffix = ""
		}
	case cldr.BE, cldr.RU:
		suffix = " г. " + era
	case cldr.BG, cldr.CY, cldr.MK:
		suffix = " " + era
	case cldr.CV:
		suffix = " ҫ. " + era
	case cldr.KK:
		prefix = era + " "
		suffix = " ж."
	case cldr.HY:
		prefix = era + " "
		suffix = " թ."
	case cldr.KY:
		prefix = era + " "
		suffix = "-ж."
	case cldr.LT:
		suffix = " m. " + era
	case cldr.TT:
		prefix = era + " "
		suffix = " ел"
	case cldr.SAH:
		suffix = " с. " + era
	case cldr.JA, cldr.BRX, cldr.YUE, cldr.ZH:
		prefix = era
		suffix = ""
	}

	return func(t timeReader) string { return prefix + year(t) + suffix }
}

func fmtEraYearPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	prefix := ""
	suffix := " " + era

	switch lang {
	case cldr.CKB, cldr.LRC, cldr.MZN, cldr.PS, cldr.UZ:
		prefix = era + " "
		suffix = ""
	case cldr.FA:
		suffix = " " + era
	}

	return func(v timeReader) string {
		return prefix + yearDigits(v) + suffix
	}
}

func fmtEraYearBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	return fmtYearBuddhist(locale, digits, opts)
}
