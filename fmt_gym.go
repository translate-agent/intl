package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqEraYearMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, regionConfidence := locale.Region()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := seqYear(locale, opts)
	month := opts.Month.symbolFormat()

	switch lang {
	case cldr.TH:
		if region == cldr.RegionTH {
			seq.Add(month, '/').AddSeq(year)
		}
	case cldr.SHN:
		if region == cldr.RegionTH {
			seq.AddSeq(year).Add('-', symbols.Symbol_MM)
		} else {
			seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		}
	case cldr.FA:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.AddSeq(year).Add('/', month).Add(' ', era)
		}
	case cldr.LRC, cldr.MZN, cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.AddSeq(year).Add('-', month)
		} else {
			seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		}
	case cldr.CKB:
		if region == cldr.RegionIR {
			seq.Add(era, ' ', opts.Year.symbol(), '-', month)
		} else {
			seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		}
	case cldr.RU:
		seq.Add(symbols.Symbol_MM, '.').AddSeq(year).Add(' ', era)
	case cldr.BG:
		seq.Add(opts.Month.symbolFormat(), '.').AddSeq(year).Add(' ', era)
	case cldr.CV:
		seq.Add(era, ' ').AddSeq(year).Add('.', symbols.Symbol_MM)
	case cldr.HI:
		if script == cldr.Latn {
			seq.Add(month, '/').AddSeq(year).Add(' ', era)
		} else {
			seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		}
	case cldr.AF, cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AST, cldr.BAS, cldr.BE, cldr.BEM, cldr.BEZ, cldr.BGC,
		cldr.BHO, cldr.BM, cldr.BO, cldr.BRX, cldr.BS, cldr.BUA, cldr.CCP, cldr.CE, cldr.CEB, cldr.CGG, cldr.CSW,
		cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EBU, cldr.EO, cldr.EWO, cldr.FUR, cldr.GAA, cldr.GSW, cldr.GUZ,
		cldr.GV, cldr.II, cldr.JGO, cldr.JMC, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN, cldr.KN,
		cldr.KSB, cldr.KSF, cldr.KSH, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.LN, cldr.LU,
		cldr.LUO, cldr.LUY, cldr.LV, cldr.MAS, cldr.MG, cldr.MER, cldr.MFE, cldr.MGH, cldr.MGO, cldr.MNI, cldr.MR, cldr.MUA,
		cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OS,
		cldr.PCM, cldr.PRG, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RWK, cldr.SA, cldr.SAH, cldr.SAQ, cldr.SAT, cldr.SBP,
		cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SI, cldr.SK, cldr.SMN, cldr.SN, cldr.SR, cldr.ST, cldr.SU, cldr.SYR,
		cldr.SZL, cldr.TA, cldr.TEO, cldr.TI, cldr.TN, cldr.TO, cldr.TOK, cldr.TWQ, cldr.TYV, cldr.TZM, cldr.VAI, cldr.VMW,
		cldr.VUN, cldr.WAE, cldr.WO, cldr.XH, cldr.XNR, cldr.XOG, cldr.YAV, cldr.YI, cldr.YRL, cldr.ZA, cldr.ZGH, cldr.ZU:
		seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.AZ:
		if script == cldr.Cyrl {
			seq.Add(era, ' ').Add(opts.Year.symbol(), '-', symbols.Symbol_MM)
		} else {
			seq.Add(era, ' ', opts.Month.symbol("format"), ' ', opts.Year.symbol())
		}
	case cldr.DZ, cldr.KKJ, cldr.KS, cldr.KY, cldr.MAI, cldr.MN, cldr.SE, cldr.TT, cldr.UG:
		seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
	case cldr.SD:
		if script == cldr.Deva {
			seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		} else {
			seq.Add(symbols.Symbol_MM, '-').AddSeq(year).Add(' ', era)
		}
	case cldr.RM, cldr.RW, cldr.YO:
		seq.Add(symbols.Symbol_MM, '-').AddSeq(year).Add(' ', era)
	case cldr.ZH:
		if script == cldr.Hant {
			seq.Add(era, opts.Year.symbol(), '/', opts.Month.symbolFormat())
		} else {
			seq.Add(era).AddSeq(year).Add(month, symbols.MonthUnit)
		}
	case cldr.YUE:
		seq.Add(era, month, '/', opts.Year.symbol())
	case cldr.JA:
		seq.Add(era, opts.Year.symbol(), '/', month)
	case cldr.EU, cldr.ML:
		seq.Add(era, ' ').AddSeq(year).Add('/', month)
	case cldr.HY, cldr.TK:
		seq.Add(era, ' ', symbols.Symbol_MM, '.').AddSeq(year)
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		} else {
			seq.Add(symbols.Symbol_MM, '-', era, ' ').AddSeq(year)
		}
	case cldr.FI, cldr.MK, cldr.NB, cldr.NN, cldr.NO, cldr.SQ:
		seq.Add(month, '.').AddSeq(year).Add(' ', era)
	case cldr.HR:
		seq.Add(symbols.Symbol_MM, '.', ' ').AddSeq(year).Add(' ', era)
	case cldr.KU, cldr.TE:
		seq.Add(era, ' ', month, '/').AddSeq(year)
	case cldr.SV:
		if region == cldr.RegionFI {
			seq.Add(month, '.').AddSeq(year).Add(' ', era)
		} else {
			seq.AddSeq(year).Add('-', symbols.Symbol_MM, ' ', era)
		}
	case cldr.LT, cldr.TG:
		seq.AddSeq(year).Add('-', symbols.Symbol_MM, ' ', era)
	case cldr.SL:
		seq.Add(symbols.Symbol_MMM, ' ').AddSeq(year).Add(' ', era)
	case cldr.UK:
		seq.Add(symbols.Symbol_MM, ' ').AddSeq(year).Add(' ', era)
	case cldr.AM, cldr.BLO, cldr.CA, cldr.CHR, cldr.CS, cldr.CY, cldr.DA, cldr.DSB, cldr.EN, cldr.FIL, cldr.GA,
		cldr.GD, cldr.GL, cldr.GU, cldr.HA, cldr.HSB, cldr.ID, cldr.IS, cldr.IT, cldr.NE, cldr.NL, cldr.SO, cldr.UR:
		seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.ES:
		if region == cldr.RegionUS {
			seq.Add(symbols.Symbol_MM, '-', opts.Year.symbol(), ' ', era)
		} else {
			seq.Add(month, '/', opts.Year.symbol(), ' ', era)
		}
	case cldr.DE:
		if opts.Month.numeric() && region != cldr.RegionCH {
			seq.Add(symbols.Symbol_MM, '/', opts.Year.symbol(), ' ', era)
		} else {
			seq.Add(month, '/', opts.Year.symbol(), ' ', era)
		}
	case cldr.FR:
		if region == cldr.RegionCA {
			seq.Add(opts.Year.symbol(), '-', symbols.Symbol_MM, ' ', era)
		} else {
			seq.Add(symbols.Symbol_MM, '/', opts.Year.symbol(), ' ', era)
		}
	case cldr.SC, cldr.SCN, cldr.TR, cldr.VEC:
		seq.Add(symbols.Symbol_MM, '/', opts.Year.symbol(), ' ', era)
	case cldr.AR:
		seq.Add(symbols.Symbol_MM, symbols.TxtArabicComma, ' ', opts.Year.symbol(), ' ', era)
	case cldr.BA:
		seq.Add(era, ' ', symbols.Symbol_MM, '.', opts.Year.symbol())
	case cldr.FF, cldr.FY, cldr.HAW, cldr.HE, cldr.IE, cldr.IG, cldr.KXV, cldr.PMS, cldr.QU:
		seq.Add(era, ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.ET, cldr.FO, cldr.PL, cldr.RO:
		seq.Add(symbols.Symbol_MM, '.', opts.Year.symbol(), ' ', era)
	case cldr.UZ:
		switch {
		case region == cldr.RegionAF:
			if regionConfidence == language.Exact {
				seq.Add(era, ' ').AddSeq(year).Add('-', month)
			} else {
				seq.AddSeq(year).Add('-', month)
			}
		case script == cldr.Cyrl:
			seq.Add(era, ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
		default:
			seq.Add(symbols.Symbol_MM, '-', opts.Year.symbol(), ' ', era)
		}
	case cldr.BN, cldr.IA:
		seq.Add(symbols.Symbol_MM, '-', opts.Year.symbol(), ' ', era)
	case cldr.BR, cldr.PT:
		if opts.Month.numeric() {
			month = symbols.Symbol_MM
		}

		seq.Add(month, '/', opts.Year.symbol(), ' ', era)
	case cldr.HU:
		seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Symbol_MM, '.')
	case cldr.EE, cldr.EL, cldr.JV, cldr.KA, cldr.KEA, cldr.KGP, cldr.LB, cldr.MI, cldr.MT, cldr.OM, cldr.OR:
		seq.Add(opts.Era.symbol(), ' ', opts.Year.symbol(), '-', symbols.Symbol_MM)
	case cldr.KM, cldr.KOK, cldr.MS, cldr.VI:
		seq.Add(month, '/').AddSeq(year).Add(' ', era)
	case cldr.KO:
		seq.Add(era, ' ').Add(opts.Year.symbol(), '/', month)
	case cldr.PA:
		if script == cldr.Arab {
			seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Symbol_MM)
		} else {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '/').AddSeq(year)
		}
	case cldr.LO:
		seq.Add(era, ' ', opts.Month.symbolFormat(), '/').AddSeq(year)
	case cldr.SW:
		seq.Add(era, ' ', symbols.Symbol_MM, '-').AddSeq(year)

	}

	if seq.Empty() {
		seq.Add(month, ' ').AddSeq(year).Add(' ', era)
	}

	return seq
}
