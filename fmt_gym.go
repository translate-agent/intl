package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop
func seqEraYearMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := seqYear(locale, opts.Year)
	month := opts.Month.symbolFormat()

	switch lang {
	case cldr.AZ, cldr.QU, cldr.TE, cldr.TK, cldr.TR:
		return seq.Add(era, ' ', month, ' ').AddSeq(year)
	case cldr.BE, cldr.RU:
		return seq.Add(month, ' ').AddSeq(year).Add(' ', symbols.Txt00, ' ', era)
	case cldr.BG:
		if opts.Month.numeric() {
			opts.Month = Month2Digit
		}

		return seq.Add(opts.Month.symbolFormat(), '.').AddSeq(year).Add(' ', era)
	case cldr.CV:
		return seq.Add(month, ' ').AddSeq(year).Add(' ', symbols.Txtҫ, '.', ' ', era)
	case cldr.HI:
		if script != cldr.Latn {
			return seq.Add(month, ' ', era, ' ').AddSeq(year)
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
		return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
	case cldr.SE:
		if region != cldr.RegionFI {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}
	case cldr.SD:
		if script != cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}
	case cldr.IG, cldr.KXV, cldr.MAI, cldr.MR, cldr.SA, cldr.XNR:
		return seq.Add(month, ' ', era, ' ').AddSeq(year)
	case cldr.FF:
		if script != cldr.Adlm {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}
	case cldr.PA:
		if script == cldr.Arab {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}

		fallthrough
	case cldr.GU, cldr.LO, cldr.UZ:
		return seq.Add(month, ',', ' ', era, ' ').AddSeq(year)
	case cldr.KGP, cldr.WO:
		return seq.Add(month, ',', ' ').AddSeq(year).Add(' ', era)
	case cldr.TT:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt05, ',', ' ', month)
	case cldr.ES:
		if region == cldr.RegionCO {
			return seq.Add(month, ' ', symbols.Txt07, ' ').AddSeq(year).Add(' ', era)
		}
	case cldr.GL, cldr.PT:
		return seq.Add(month, ' ', symbols.Txt07, ' ').AddSeq(year).Add(' ', era)
	case cldr.JA, cldr.YUE, cldr.ZH:
		return seq.Add(era).AddSeq(year).Add(symbols.Symbol_M, symbols.MonthUnit)
	case cldr.DZ:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt06, month)
	case cldr.EU:
		return seq.Add(era, ' ').AddSeq(year).Add('.', ' ', symbols.Txt08, ' ', month)
	case cldr.HY:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtթ, '.', ' ', month)
	case cldr.KK:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtж, '.', ' ', month)
	case cldr.KA:
		return seq.Add(month, '.', ' ').AddSeq(year).Add(' ', era)
	case cldr.KU:
		return seq.Add(era, ' ', month, 'a', ' ').AddSeq(year).Add(symbols.Txt09)
	case cldr.KY:
		return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Txtж, '.', ' ', month)
	case cldr.LT:
		return seq.AddSeq(year).Add('-', symbols.Symbol_MM, ' ', era)
	case cldr.MN:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt10, ' ', month)
	case cldr.SL:
		return seq.Add(symbols.Symbol_MMM, ' ').AddSeq(year).Add(' ', era)
	case cldr.UG:
		return seq.AddSeq(year).Add(' ', month, ' ', era)
	case cldr.UK:
		return seq.Add(month, ' ').AddSeq(year).Add(' ', symbols.Txtр, '.', ' ', era)
	case cldr.KOK:
		if script != cldr.Latn {
			return seq.Add(era, ' ').AddSeq(year).Add(' ', month)
		}
	}

	return seq.Add(month, ' ').AddSeq(year).Add(' ', era)
}

func seqEraYearMonthPersian(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	year := seqYearPersian(locale, opts.Year)
	month := opts.Month.symbolFormat()

	switch lang {
	case cldr.FA:
		return seq.Add(month, ' ').AddSeq(year).Add(' ', opts.Era.symbol())
	case cldr.CKB, cldr.UZ:
		if region != cldr.RegionAF {
			return seq.AddSeq(year).Add(' ', month)
		}
	case cldr.LRC, cldr.MZN, cldr.PS:
		return seq.AddSeq(year).Add(' ', month)
	}

	return seq.Add(opts.Era.symbol(), ' ').AddSeq(year).Add(' ', month)
}

func seqEraYearMonthBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Month.symbolFormat(), ' ').AddSeq(seqYearBuddhist(locale, opts))
}
