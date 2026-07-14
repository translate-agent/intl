package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqYearMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, regionConfidence := locale.Region()
	seq := symbols.NewSeq(locale)
	year := opts.Year.symbol()
	month := opts.Month.symbolFormat()

	switch lang {
	case cldr.TH:
		if region == cldr.RegionTH {
			seq.Add(month, '/', year)
		}
	case cldr.SHN:
		if region == cldr.RegionTH {
			seq.Add(symbols.Symbol_G, ' ', year, '-', symbols.Symbol_MM)
		}
	case cldr.CKB:
		seq.Add(month, '/', year)
	case cldr.FA:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(year, '/', month)
		}
	case cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(symbols.Symbol_GGGGG, ' ', year, '/', month)
		}
	case cldr.LRC, cldr.MZN:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(symbols.Symbol_GGGGG, ' ', year, '-', month)
		}
	case cldr.AF, cldr.AS, cldr.IA, cldr.JV, cldr.MI, cldr.RM, cldr.RW, cldr.TG, cldr.WO:
		seq.Add(symbols.Symbol_MM, '-', year)
	case cldr.EN:
		switch region {
		default:
			seq.Add(month, '/', year)
		case cldr.Region001, cldr.Region150, cldr.RegionAE, cldr.RegionAG, cldr.RegionAI, cldr.RegionAT, cldr.RegionAU,
			cldr.RegionBB, cldr.RegionBE, cldr.RegionBM, cldr.RegionBS, cldr.RegionBW, cldr.RegionBZ, cldr.RegionCC,
			cldr.RegionCK, cldr.RegionCM, cldr.RegionCX, cldr.RegionCY, cldr.RegionCZ, cldr.RegionDE, cldr.RegionDG,
			cldr.RegionDK, cldr.RegionDM, cldr.RegionER, cldr.RegionFI, cldr.RegionFJ, cldr.RegionFK, cldr.RegionFM,
			cldr.RegionGB, cldr.RegionGD, cldr.RegionGG, cldr.RegionGH, cldr.RegionGI, cldr.RegionGM, cldr.RegionGY,
			cldr.RegionHK, cldr.RegionID, cldr.RegionIE, cldr.RegionIL, cldr.RegionIM, cldr.RegionIN, cldr.RegionIO,
			cldr.RegionJE, cldr.RegionJM, cldr.RegionKE, cldr.RegionKI, cldr.RegionKN, cldr.RegionKY, cldr.RegionLC,
			cldr.RegionLR, cldr.RegionLS, cldr.RegionMG, cldr.RegionMO, cldr.RegionMS, cldr.RegionMT, cldr.RegionMU,
			cldr.RegionMV, cldr.RegionMW, cldr.RegionMY, cldr.RegionNA, cldr.RegionNF, cldr.RegionNG, cldr.RegionNL,
			cldr.RegionNR, cldr.RegionNU, cldr.RegionNZ, cldr.RegionPG, cldr.RegionPK, cldr.RegionPN, cldr.RegionPW,
			cldr.RegionRW, cldr.RegionSB, cldr.RegionSC, cldr.RegionSD, cldr.RegionSG, cldr.RegionSH, cldr.RegionSI,
			cldr.RegionSL, cldr.RegionSS, cldr.RegionSX, cldr.RegionSZ, cldr.RegionTC, cldr.RegionTK, cldr.RegionTO,
			cldr.RegionTT, cldr.RegionTV, cldr.RegionTZ, cldr.RegionUG, cldr.RegionVC, cldr.RegionVG, cldr.RegionVU,
			cldr.RegionWS, cldr.RegionZA, cldr.RegionZM, cldr.RegionZW:
			if script == cldr.Shaw {
				seq.Add(month, '/', year)
			} else {
				seq.Add(symbols.Symbol_MM, '/', year)
			}
		case cldr.RegionEE, cldr.RegionES, cldr.RegionFR, cldr.RegionGE, cldr.RegionGS, cldr.RegionHU, cldr.RegionIT,
			cldr.RegionJP, cldr.RegionLT, cldr.RegionLV, cldr.RegionNO, cldr.RegionPL, cldr.RegionPT, cldr.RegionRO,
			cldr.RegionSK, cldr.RegionUA:
			seq.Add(symbols.Symbol_MM, '/', year)
		case cldr.RegionCA, cldr.RegionSE:
			seq.Add(year, '-', symbols.Symbol_MM)
		case cldr.RegionCH:
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.AGQ, cldr.AK, cldr.AM, cldr.ASA, cldr.AST, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BLO, cldr.BM, cldr.BRX,
		cldr.CA, cldr.CEB, cldr.CGG, cldr.CHR, cldr.CS, cldr.CY, cldr.DAV, cldr.DE, cldr.DJE, cldr.DOI, cldr.DUA,
		cldr.DYO, cldr.EBU, cldr.EE, cldr.EL, cldr.EWO, cldr.FIL, cldr.FUR, cldr.GD, cldr.GL, cldr.GUZ, cldr.HA, cldr.HAW,
		cldr.ID, cldr.IG, cldr.JMC, cldr.KAA, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KLN, cldr.KM, cldr.KSB,
		cldr.KSF, cldr.KXV, cldr.LAG, cldr.LG, cldr.LN, cldr.LO, cldr.LU, cldr.LUO, cldr.LUY, cldr.MAI, cldr.MAS, cldr.MER,
		cldr.MFE, cldr.MG, cldr.MGH, cldr.MHN, cldr.ML, cldr.MNI, cldr.MUA, cldr.NAQ, cldr.ND, cldr.NMG, cldr.NUS, cldr.NYN,
		cldr.OM, cldr.PCM, cldr.RN, cldr.ROF, cldr.RWK, cldr.SA, cldr.SAQ, cldr.SBP, cldr.SCN, cldr.SES, cldr.SG, cldr.SHI,
		cldr.SK, cldr.SL, cldr.SO, cldr.SU, cldr.SW, cldr.TEO, cldr.TWQ, cldr.TZM, cldr.UR, cldr.VAI, cldr.VUN, cldr.XH,
		cldr.XNR, cldr.XOG, cldr.YAV, cldr.YO, cldr.ZGH:
		seq.Add(month, '/', year)
	case cldr.PA:
		if script != cldr.Arab {
			seq.Add(month, '/', year)
		}
	case cldr.KS:
		if script != cldr.Deva {
			seq.Add(month, '/', year)
		}
	case cldr.HI:
		if script == cldr.Latn {
			seq.Add(symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(month, '/', year)
		}
	case cldr.AR:
		seq.Add(month, symbols.Txt02, year)
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(month, '-', year)
		} else {
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.AZ, cldr.FO, cldr.HY, cldr.KU, cldr.OS, cldr.PL, cldr.RO, cldr.RU, cldr.TK, cldr.TT,
		cldr.UK:
		seq.Add(symbols.Symbol_MM, '.', year)
	case cldr.UZ:
		switch {
		case region == cldr.RegionAF:
			if regionConfidence == language.Exact {
				seq.Add(year, '-', month)
			} else {
				seq.Add(symbols.Symbol_GGGGG, ' ', year, '-', month)
			}
		case script == cldr.Cyrl:
			seq.Add(symbols.Symbol_MM, '/', year)
		default:
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.BE, cldr.DA, cldr.DSB, cldr.ET, cldr.HSB, cldr.IE, cldr.KA, cldr.LB, cldr.NB, cldr.NN, cldr.NO, cldr.SMN,
		cldr.SQ:
		seq.Add(month, '.', year)
	case cldr.BG:
		seq.Add(symbols.Symbol_MM, '.', year, ' ', symbols.Txt00)
	case cldr.MK:
		seq.Add(month, '.', year, ' ', symbols.Txt00)
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.OR, cldr.TA, cldr.TE, cldr.TO:
		if opts.Month.numeric() {
			seq.Add(month, '/', year)
		} else {
			seq.Add(month, '-', year)
		}
	case cldr.BR, cldr.GA, cldr.IT, cldr.IU, cldr.KEA, cldr.KGP, cldr.PT, cldr.SC, cldr.SEH, cldr.SYR, cldr.VEC, cldr.YRL:
		seq.Add(symbols.Symbol_MM, '/', year)
	case cldr.BS:
		switch {
		case script == cldr.Cyrl:
			seq.Add(symbols.Symbol_MM, '.', year, '.')
		case !opts.Month.numeric():
			seq.Add(symbols.Symbol_M, '.', ' ', year, '.')
		default:
			seq.Add(symbols.Symbol_MM, '/', year)
		}
	case cldr.DZ, cldr.SI:
		seq.Add(year, '-', month)
	case cldr.ES:
		switch region {
		default:
			seq.Add(symbols.Symbol_M, '/', year)
		case cldr.RegionAR:
			if opts.Month.numeric() {
				seq.Add(symbols.Symbol_M, '-', year)
			} else {
				seq.Add(symbols.Symbol_M, '/', year)
			}
		case cldr.RegionCL:
			if opts.Month.numeric() {
				seq.Add(symbols.Symbol_MM, '-', year)
			} else {
				seq.Add(symbols.Symbol_M, '/', year)
			}
		case cldr.RegionMX, cldr.RegionUS:
			seq.Add(month, '/', year)
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() {
				seq.Add(symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(symbols.Symbol_M, '/', year)
			}
		}
	case cldr.TI:
		seq.Add(symbols.Symbol_M, '/', year)
	case cldr.YUE:
		seq.Add(year, '/', month)
	case cldr.EU, cldr.JA:
		seq.Add(year, '/', month)
	case cldr.FI, cldr.HE:
		seq.Add(symbols.Symbol_M, '.', year)
	case cldr.FF:
		if script == cldr.Adlm {
			seq.Add(month, '-', year)
		} else {
			seq.Add(month, '/', year)
		}
	case cldr.FR:
		switch region {
		default:
			seq.Add(symbols.Symbol_MM, '/', year)
		case cldr.RegionCA: // noop
			seq.Add(year, '-', symbols.Symbol_MM)
		case cldr.RegionCH:
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.NL:
		if region == cldr.RegionBE {
			seq.Add(month, '/', year)
		} else {
			seq.Add(month, '-', year)
		}
	case cldr.KOK:
		if script == cldr.Latn {
			seq.Add(month, '/', year)
		} else {
			seq.Add(month, '-', year)
		}
	case cldr.FY, cldr.MS, cldr.UG:
		seq.Add(month, '-', year)
	case cldr.GSW:
		if opts.Month.numeric() {
			seq.Add(year, '-', month)
		} else {
			seq.Add(month, '.', year)
		}
	case cldr.HR:
		seq.Add(symbols.Symbol_MM, '.', ' ', year, '.')
	case cldr.HU:
		seq.Add(year, '.', ' ', month, '.')
	case cldr.IS:
		seq.Add(month, '.', ' ', year)
	case cldr.KKJ:
		seq.Add(symbols.Symbol_MM, ' ', year)
	case cldr.KO:
		seq.Add(year, '.', ' ', symbols.Symbol_M, '.')
	case cldr.LV:
		seq.Add(symbols.Symbol_MM, '.', year, '.')
	case cldr.MN:
		seq.Add(year, ' ', symbols.Symbol_LLLLL)
	case cldr.YI:
		if !opts.Month.numeric() {
			seq.Add(month, '/', year)
		}
	case cldr.SD:
		if script == cldr.Deva {
			seq.Add(month, '/', year)
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.SR:
		if opts.Month.numeric() {
			seq.Add(month, '.', ' ', year, '.')
		} else {
			seq.Add(month, '.', year, '.')
		}
	case cldr.TR:
		if opts.Month.numeric() {
			seq.Add(symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.VI:
		if opts.Month.numeric() {
			seq.Add(month, '/', year)
		} else {
			seq.Add(symbols.Txt04, month, ',', ' ', year)
		}
	case cldr.ZH:
		switch script {
		case cldr.Hant:
			if region == cldr.RegionHK || region == cldr.RegionMO {
				seq.Add(month, '/', year)
			} else {
				seq.Add(year, '/', month)
			}
		case cldr.Hans:
			if region == cldr.RegionHK {
				seq.Add(month, '/', year)
			} else {
				if !opts.Month.numeric() {
					seq.Add(year, symbols.Txt年, symbols.Symbol_M, symbols.Txt月)
				} else {
					seq.Add(year, '/', month)
				}
			}
		default:
			if opts.Month.numeric() {
				seq.Add(year, '/', month)
			} else {
				seq.Add(year, symbols.Txt年, symbols.Symbol_M, symbols.Txt月)
			}
		}
	case cldr.CV:
		seq.Add(year, '.', symbols.Symbol_MM)
	case cldr.SV:
		if region == cldr.RegionFI {
			seq.Add(symbols.Symbol_M, '.', year)
		}

	}

	if seq.Empty() {
		seq.Add(year, '-', symbols.Symbol_MM)
	}

	return seq
}
