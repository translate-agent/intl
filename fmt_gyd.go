package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraYearDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := seqYear(locale, opts.Year)
	day := seqDay(locale, opts.Day)

	switch lang {
	case cldr.BE, cldr.RU:
		return seq.AddSeq(year).Add(' ', symbols.Txt00, ' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.CV:
		return seq.
			AddSeq(year).Add(' ', symbols.Txtҫ, '.', ' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.KK:
		return seq.
			Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtж, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.KY:
		return seq.
			Add(era, ' ').AddSeq(year).Add('-', symbols.Txtж, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.HY:
		return seq.
			Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtթ, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.TT:
		return seq.
			Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt05, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.SAH:
		return seq.
			AddSeq(year).Add(' ', symbols.Txtс, '.', ' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.LT:
		return seq.
			AddSeq(year).Add(' ', symbols.Txtm, '.', ' ', era, ' ', '(', symbols.DayUnit, ':', ' ', Day2Digit.symbol(), ')')
	case cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AZ, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BGC, cldr.BHO, cldr.BM,
		cldr.BO, cldr.CE, cldr.CGG, cldr.CKB, cldr.CSW, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.DZ,
		cldr.EBU, cldr.EO, cldr.EU, cldr.EWO, cldr.FUR, cldr.FY, cldr.GAA, cldr.GSW, cldr.GU, cldr.GUZ, cldr.GV, cldr.HA,
		cldr.HU, cldr.IG, cldr.II, cldr.JGO, cldr.JMC, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN,
		cldr.KN, cldr.KO, cldr.KSB, cldr.KSF, cldr.KSH, cldr.KU, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO,
		cldr.LN, cldr.LO, cldr.LRC, cldr.LU, cldr.LUO, cldr.LUY, cldr.LV, cldr.MAS, cldr.MER, cldr.MFE, cldr.MG, cldr.MGH,
		cldr.MGO, cldr.ML, cldr.MN, cldr.MNI, cldr.MR, cldr.MT, cldr.MUA, cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NE,
		cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OM, cldr.OS, cldr.PA, cldr.PCM, cldr.PRG,
		cldr.PS, cldr.QU, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RW, cldr.RWK, cldr.SAQ, cldr.SAT, cldr.SBP, cldr.SEH, cldr.SES,
		cldr.SG, cldr.SHI, cldr.SI, cldr.SN, cldr.ST, cldr.SZL, cldr.TA, cldr.TE, cldr.TEO, cldr.TK, cldr.TN, cldr.TOK,
		cldr.TR, cldr.TWQ, cldr.TZM, cldr.VAI, cldr.VMW, cldr.VUN, cldr.WAE, cldr.XOG, cldr.YAV, cldr.YI, cldr.YO, cldr.ZA,
		cldr.ZGH, cldr.ZU:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.UZ:
		if script != cldr.Arab {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.FF:
		if script != cldr.Adlm {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.HI:
		if script == cldr.Latn {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.KOK:
		if script != cldr.Latn {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.SD:
		if script != cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.BRX, cldr.JA, cldr.YUE, cldr.ZH:
		return seq.Add(era).AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.SE:
		if region != cldr.RegionFI {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	}

	return seq.AddSeq(year).Add(' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
}

func seqEraYearDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := seqYearPersian(locale, opts.Year)
	day := opts.Day.symbol()

	switch lang {
	case cldr.UZ:
		if region == cldr.RegionAF {
			seq.Add(era, ' ')
		}
	case cldr.FA:
		return seq.AddSeq(year).Add(' ', era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
	}

	return seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ', day, ')')
}

func seqEraYearDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	year := seqYearBuddhist(locale, opts)
	day := opts.Day.symbol()

	return symbols.NewSeq(locale).AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ', day, ')')
}
