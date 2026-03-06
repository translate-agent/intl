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
	case cldr.RU:
		return seq.Add(symbols.Symbol_MM, '.').AddSeq(year).Add(' ', era)
	case cldr.BG:
		return seq.Add(opts.Month.symbolFormat(), '.').AddSeq(year).Add(' ', era)
	case cldr.CV:
		return seq.Add(era, ' ').AddSeq(year).Add('.', symbols.Symbol_MM)
	case cldr.HI:
		if script == cldr.Latn {
			return seq.Add(month, '/').AddSeq(year).Add(' ', era)
		}

		fallthrough
	case cldr.AF, cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AST, cldr.BAS, cldr.BE, cldr.BEM, cldr.BEZ, cldr.BGC,
		cldr.BHO, cldr.BM, cldr.BO, cldr.BRX, cldr.BS, cldr.BUA, cldr.CCP, cldr.CE, cldr.CEB, cldr.CGG, cldr.CKB, cldr.CSW,
		cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EBU, cldr.EO, cldr.EWO, cldr.FUR, cldr.GAA, cldr.GSW, cldr.GUZ,
		cldr.GV, cldr.II, cldr.JGO, cldr.JMC, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN, cldr.KN,
		cldr.KSB, cldr.KSF, cldr.KSH, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.LN, cldr.LRC, cldr.LU,
		cldr.LUO, cldr.LUY, cldr.LV, cldr.MAS, cldr.MG, cldr.MER, cldr.MFE, cldr.MGH, cldr.MGO, cldr.MNI, cldr.MR, cldr.MUA,
		cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OS,
		cldr.PCM, cldr.PRG, cldr.PS, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RWK, cldr.SA, cldr.SAH, cldr.SAQ, cldr.SAT, cldr.SBP,
		cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SI, cldr.SK, cldr.SMN, cldr.SN, cldr.SR, cldr.ST, cldr.SU, cldr.SYR,
		cldr.SZL, cldr.TA, cldr.TEO, cldr.TI, cldr.TN, cldr.TO, cldr.TOK, cldr.TWQ, cldr.TYV, cldr.TZM, cldr.VAI, cldr.VMW,
		cldr.VUN, cldr.WAE, cldr.WO, cldr.XH, cldr.XNR, cldr.XOG, cldr.YAV, cldr.YI, cldr.YRL, cldr.ZA, cldr.ZGH, cldr.ZU:
		return seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.AZ:
		if script == cldr.Cyrl {
			return seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		}

		return seq.Add(era, ' ', opts.Month.symbol("format"), ' ', opts.Year.symbol())
	case cldr.DZ, cldr.KKJ, cldr.KS, cldr.KY, cldr.MAI, cldr.MN, cldr.SE, cldr.SHN, cldr.TT, cldr.UG:
		return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
	case cldr.SD:
		if script == cldr.Deva {
			return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		}

		fallthrough
	case cldr.RM, cldr.RW, cldr.YO:
		return seq.Add(symbols.Symbol_MM, '-').AddSeq(year).Add(' ', era)
	case cldr.ZH:
		if script == cldr.Hant {
			return seq.Add(era, opts.Year.symbol(), '/', opts.Month.symbolFormat())
		}

		return seq.Add(era).AddSeq(year).Add(month, symbols.MonthUnit)
	case cldr.YUE:
		return seq.Add(era, month, '/', opts.Year.symbol())
	case cldr.JA:
		return seq.Add(era, opts.Year.symbol(), '/', month)
	case cldr.EU, cldr.ML:
		return seq.Add(era, ' ').AddSeq(year).Add('/', month)
	case cldr.HY, cldr.TK:
		return seq.Add(era, ' ', symbols.Symbol_MM, '.').AddSeq(year)
	case cldr.KK:
		if script == cldr.Arab {
			return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		}

		return seq.Add(symbols.Symbol_MM, '-', era, ' ').AddSeq(year)
	case cldr.FI, cldr.MK, cldr.NB, cldr.NN, cldr.NO, cldr.SQ:
		return seq.Add(month, '.').AddSeq(year).Add(' ', era)
	case cldr.HR:
		return seq.Add(symbols.Symbol_MM, '.', ' ').AddSeq(year).Add(' ', era)
	case cldr.KU, cldr.TE:
		return seq.Add(era, ' ', month, '/').AddSeq(year)
	case cldr.SV:
		if region == cldr.RegionFI {
			return seq.Add(month, '.').AddSeq(year).Add(' ', era)
		}

		fallthrough
	case cldr.LT, cldr.TG:
		return seq.AddSeq(year).Add('-', symbols.Symbol_MM, ' ', era)
	case cldr.SL:
		return seq.Add(symbols.Symbol_MMM, ' ').AddSeq(year).Add(' ', era)
	case cldr.UK:
		return seq.Add(symbols.Symbol_MM, ' ').AddSeq(year).Add(' ', era)
	case cldr.AM, cldr.BLO, cldr.CA, cldr.CHR, cldr.CS, cldr.CY, cldr.DA, cldr.DSB, cldr.EN, cldr.FIL, cldr.GA,
		cldr.GD, cldr.GL, cldr.GU, cldr.HA, cldr.HSB, cldr.ID, cldr.IS, cldr.IT, cldr.NE, cldr.NL, cldr.SO, cldr.UR:
		return seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.ES:
		if region == cldr.RegionUS {
			return seq.Add(symbols.Symbol_MM, '-', opts.Year.symbol(), ' ', era)
		}

		return seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.DE:
		if opts.Month.numeric() && region != cldr.RegionCH {
			return seq.Add(symbols.Symbol_MM, '/', opts.Year.symbol(), ' ', era)
		}

		return seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.FR:
		if region == cldr.RegionCA {
			return seq.Add(opts.Year.symbol(), '-', symbols.Symbol_MM, ' ', era)
		}

		fallthrough
	case cldr.SC, cldr.SCN, cldr.TR, cldr.VEC:
		return seq.Add(symbols.Symbol_MM, '/', opts.Year.symbol(), ' ', era)
	case cldr.AR:
		return seq.Add(symbols.Symbol_MM, symbols.TxtArabicComma, ' ', opts.Year.symbol(), ' ', era)
	case cldr.BA:
		return seq.Add(era, ' ', symbols.Symbol_MM, '.', opts.Year.symbol())
	case cldr.FF, cldr.FY, cldr.HAW, cldr.HE, cldr.IE, cldr.IG, cldr.KXV, cldr.PMS, cldr.QU:
		return seq.Add(era, ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.ET, cldr.FO, cldr.PL, cldr.RO:
		return seq.Add(symbols.Symbol_MM, '.', opts.Year.symbol(), ' ', era)
	case cldr.UZ:
		if script == cldr.Cyrl {
			return seq.Add(era, ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
		}

		fallthrough
	case cldr.BN, cldr.IA:
		return seq.Add(symbols.Symbol_MM, '-', opts.Year.symbol(), ' ', era)
	case cldr.BR, cldr.PT:
		if opts.Month.numeric() {
			month = symbols.Symbol_MM
		}

		return seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.HU:
		return seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Symbol_MM, '.')
	case cldr.EE, cldr.EL, cldr.JV, cldr.KA, cldr.KEA, cldr.KGP, cldr.LB, cldr.MI, cldr.MT, cldr.OM, cldr.OR:
		return seq.Add(opts.Era.symbol(), ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.KM, cldr.KOK, cldr.MS, cldr.VI:
		return seq.Add(month, '/').AddSeq(year).Add(' ', era)
	case cldr.KO:
		return seq.Add(era, ' ').Add(opts.Year.symbol(), '/', month)
	case cldr.PA:
		if script == cldr.Arab {
			return seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		}

		fallthrough
	case cldr.LO:
		return seq.Add(era, ' ', opts.Month.symbolFormat(), '/').AddSeq(year)
	case cldr.SW:
		return seq.Add(era, ' ', symbols.Symbol_MM, '-').AddSeq(year)
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
		return seq.AddSeq(year).Add('/', month).Add(' ', opts.Era.symbol())
	case cldr.UZ:
		if region == cldr.RegionAF {
			return seq.Add(opts.Era.symbol(), ' ').AddSeq(year).Add('-', month)
		}

		fallthrough
	case cldr.LRC, cldr.MZN, cldr.PS:
		return seq.AddSeq(year).Add('-', month)
	case cldr.CKB:
		return seq.Add(opts.Era.symbol(), ' ', opts.Year.symbol(), '-', month)
	}

	return seq.Add(opts.Era.symbol(), ' ').AddSeq(year).Add(' ', month)
}

func seqEraYearMonthBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, _ := locale.Raw()
	seq := symbols.NewSeq(locale)
	year := seqYearBuddhist(locale, opts)

	switch lang {
	default:
		return seq.Add(opts.Month.symbolFormat(), '/').AddSeq(year)
	case cldr.SHN:
		return seq.AddSeq(year).Add('-', symbols.Symbol_MM)
	}
}
