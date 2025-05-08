package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraYear(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := symbolEra(opts.Era)
	year := seqYear(locale, opts.Year)

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
		return seq.Add(era, ' ').AddSeq(year)
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year)
		}
	case cldr.HI:
		if script == cldr.Latn {
			return seq.Add(era, ' ').AddSeq(year)
		}
	case cldr.SD:
		if script != cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year)
		}
	case cldr.FF:
		if script != cldr.Adlm {
			return seq.Add(era, ' ').AddSeq(year)
		}
	case cldr.SE:
		if region != cldr.RegionFI {
			return seq.Add(era, ' ').AddSeq(year)
		}
	case cldr.BE, cldr.RU:
		return seq.AddSeq(year).Add(' ', symbols.Txt00, ' ', era)
	case cldr.CV:
		return seq.AddSeq(year).Add(' ', symbols.Txtҫ, '.', ' ', era)
	case cldr.KK:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtж, '.')
	case cldr.HY:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtթ, '.')
	case cldr.KY:
		return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Txtж, '.')
	case cldr.LT:
		return seq.AddSeq(year).Add(' ', 'm', '.', ' ', era)
	case cldr.TT:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt05)
	case cldr.SAH:
		return seq.AddSeq(year).Add(' ', symbols.Txtс, '.', ' ', era)
	case cldr.JA, cldr.BRX, cldr.YUE, cldr.ZH:
		return seq.Add(era).AddSeq(year)
	}

	return seq.AddSeq(year).Add(' ', era)
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
	}

	return func(v cldr.TimeReader) string {
		return prefix + yearDigits(v) + suffix
	}
}

func fmtEraYearBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	return fmtYearBuddhist(locale, digits, opts)
}
