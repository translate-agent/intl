package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func seqMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, _ := locale.Region()
	seq := symbols.NewSeq(locale)
	month := opts.Month.symbolFormat()
	day := opts.Day.symbol()

	switch lang {
	case cldr.TH:
		seq.Add(day, '/', month)
	case cldr.FA, cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(month, '/', day)
		}
	case cldr.EN:
		switch region {
		default:
			seq.Add(month, '/', day)
		case cldr.Region001, cldr.Region150, cldr.RegionAE, cldr.RegionAG, cldr.RegionAI, cldr.RegionAT, cldr.RegionBB,
			cldr.RegionBM, cldr.RegionBS, cldr.RegionBW, cldr.RegionBZ, cldr.RegionCC, cldr.RegionCK, cldr.RegionCM,
			cldr.RegionCX, cldr.RegionCY, cldr.RegionCZ, cldr.RegionDE, cldr.RegionDG, cldr.RegionDK, cldr.RegionDM,
			cldr.RegionER, cldr.RegionFI, cldr.RegionFJ, cldr.RegionFK, cldr.RegionFM, cldr.RegionGB, cldr.RegionGD,
			cldr.RegionGG, cldr.RegionGH, cldr.RegionGI, cldr.RegionGM, cldr.RegionGY, cldr.RegionHK, cldr.RegionID,
			cldr.RegionIL, cldr.RegionIM, cldr.RegionIN, cldr.RegionIO, cldr.RegionJE, cldr.RegionJM, cldr.RegionKE,
			cldr.RegionKI, cldr.RegionKN, cldr.RegionKY, cldr.RegionLC, cldr.RegionLR, cldr.RegionLS, cldr.RegionMG,
			cldr.RegionMO, cldr.RegionMS, cldr.RegionMT, cldr.RegionMU, cldr.RegionMV, cldr.RegionMW, cldr.RegionMY,
			cldr.RegionNA, cldr.RegionNF, cldr.RegionNG, cldr.RegionNL, cldr.RegionNR, cldr.RegionNU, cldr.RegionPG,
			cldr.RegionPK, cldr.RegionPN, cldr.RegionPW, cldr.RegionRW, cldr.RegionSB, cldr.RegionSC, cldr.RegionSD,
			cldr.RegionSE, cldr.RegionSG, cldr.RegionSH, cldr.RegionSI, cldr.RegionSL, cldr.RegionSS, cldr.RegionSX,
			cldr.RegionSZ, cldr.RegionTC, cldr.RegionTK, cldr.RegionTO, cldr.RegionTT, cldr.RegionTV, cldr.RegionTZ,
			cldr.RegionUG, cldr.RegionVC, cldr.RegionVG, cldr.RegionVU, cldr.RegionWS, cldr.RegionZM:
			switch {
			case script == cldr.Shaw:
				seq.Add(month, '/', day)
			case opts.Month.numeric() && opts.Day.numeric():
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
			default:
				seq.Add(day, '/', month)
			}
		case cldr.RegionAU, cldr.RegionBE, cldr.RegionIE, cldr.RegionNZ, cldr.RegionZW:
			seq.Add(day, '/', month)
		case cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
			} else {
				seq.Add(month, '-', day)
			}
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
			} else {
				seq.Add(day, '.', month)
			}
		case cldr.RegionZA:
			if opts.Month.twoDigit() && opts.Day.twoDigit() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
			} else {
				seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
			}
		case cldr.RegionEE, cldr.RegionES, cldr.RegionFR, cldr.RegionGE, cldr.RegionGS, cldr.RegionHU, cldr.RegionIT,
			cldr.RegionLT, cldr.RegionLV, cldr.RegionNO, cldr.RegionPL, cldr.RegionPT, cldr.RegionRO, cldr.RegionSK,
			cldr.RegionUA:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
			} else {
				seq.Add(day, '/', month)
			}
		}
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.TG, cldr.WO, cldr.RW:
		seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM)
	case cldr.HI:
		if script == cldr.Latn && opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)

			break
		}

		fallthrough
	case cldr.AM, cldr.AGQ, cldr.AST, cldr.BAS, cldr.BM, cldr.CA, cldr.CY, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EL,
		cldr.EWO, cldr.FUR, cldr.GD, cldr.GL, cldr.HAW, cldr.ID, cldr.IG, cldr.KAB, cldr.KGP, cldr.KHQ, cldr.KM, cldr.KOK,
		cldr.KSF, cldr.KXV, cldr.LN, cldr.LO, cldr.LU, cldr.MAI, cldr.MFE, cldr.MG, cldr.MGH, cldr.ML, cldr.MNI, cldr.MUA,
		cldr.MY, cldr.NMG, cldr.NUS, cldr.PA, cldr.RN, cldr.SA, cldr.SCN, cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SU,
		cldr.SW, cldr.TO, cldr.TR, cldr.TWQ, cldr.UR, cldr.XNR, cldr.YAV, cldr.YO, cldr.YRL, cldr.ZGH:
		seq.Add(day, '/', month)
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.VEC:
		seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
	case cldr.UZ:
		if region == cldr.RegionAF {
			seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		}
	case cldr.TI:
		seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
	case cldr.KEA, cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		} else {
			seq.Add(day, '/', month)
		}
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BEZ, cldr.BLO, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.JA, cldr.JMC, cldr.KAA, cldr.KAM, cldr.KDE, cldr.KI, cldr.KLN,
		cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.MHN, cldr.NAQ, cldr.ND, cldr.NYN,
		cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TEO, cldr.TZM, cldr.VAI, cldr.VUN, cldr.XOG,
		cldr.YUE:
		seq.Add(month, '/', day)
	case cldr.KS:
		if script == cldr.Deva {
			seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(month, '/', day)
		}
	case cldr.AR:
		seq.Add(day, symbols.Txt02, month)
	case cldr.BA, cldr.AZ, cldr.FO, cldr.HY, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(day, '-', month)
		} else {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
		}
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		seq.Add(day, '.', month)
	case cldr.MK:
		seq.Add(symbols.Symbol_d, '.', month)
	case cldr.BG, cldr.PL:
		seq.Add(day, '.', symbols.Symbol_MM)
	case cldr.LV:
		seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
	case cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.IS, cldr.LB, cldr.SMN:
		seq.Add(day, '.', month, '.')
	case cldr.DE:
		if opts.Month.twoDigit() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', month, '.')
		} else {
			seq.Add(day, '.', month, '.')
		}
	case cldr.NB, cldr.NN, cldr.NO: // d.M.
		seq.Add(symbols.Symbol_d, '.', symbols.Symbol_M, '.')
	case cldr.SQ:
		seq.Add(symbols.Symbol_d, '.', symbols.Symbol_M)
	case cldr.RO, cldr.RU:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
		} else {
			seq.Add(day, '.', month)
		}
	case cldr.CV:
		seq.Add(symbols.Symbol_MM, '.', symbols.Symbol_dd)
	case cldr.SR:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '.', ' ', month, '.')
		} else {
			seq.Add(day, '.', month, '.')
		}
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.TA, cldr.TE, cldr.VI:
		var sep symbols.Symbol = '-'
		if opts.Month.numeric() && opts.Day.numeric() {
			sep = '/'
		}

		seq.Add(day, sep, month)
	case cldr.BS:
		if script == cldr.Cyrl {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
		} else {
			seq.Add(symbols.Symbol_d, '.', ' ', symbols.Symbol_M, '.')
		}
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			month = symbols.Symbol_MM
			day = symbols.Symbol_dd
		}

		seq.Add(day, '.', ' ', month, '.')
	case cldr.CS, cldr.SK, cldr.SL:
		seq.Add(day, '.', ' ', month, '.')
	case cldr.HU, cldr.KO:
		seq.Add(month, '.', ' ', day, '.')
	case cldr.WAE:
		seq.Add(day, '.', ' ', symbols.Symbol_LLL)
	case cldr.DZ, cldr.SI: // noop
		seq.Add(month, '-', day)
	case cldr.ES:
		switch region {
		default:
			seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
		case cldr.RegionCL:
			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit

				seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM)
			} else {
				seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
			}
		case cldr.RegionMX, cldr.RegionUS:
			seq.Add(day, '/', month)
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() {
				seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
			} else {
				seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
			}
		}
	case cldr.FF:
		if script == cldr.Adlm {
			seq.Add(day, '-', month)
		} else {
			seq.Add(day, '/', month)
		}
	case cldr.FR:
		switch region {
		default:
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		case cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.twoDigit() {
				seq.Add(symbols.Symbol_M, '-', day)
			} else {
				seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
			} else {
				seq.Add(day, '.', month)
			}
		}
	case cldr.NL:
		if region == cldr.RegionBE {
			seq.Add(day, '/', month)
		} else {
			seq.Add(day, '-', month)
		}
	case cldr.FY, cldr.UG:
		seq.Add(day, '-', month)
	case cldr.IU:
		seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_MM, '-', day)
		} else {
			seq.Add(month, '-', day)
		}
	case cldr.MN:
		seq.Add(symbols.Symbol_LLLLL, '/', symbols.Symbol_dd)
	case cldr.MS:
		if !opts.Month.numeric() || !opts.Day.numeric() {
			seq.Add(day, '/', month)
		} else {
			seq.Add(day, '-', month)
		}
	case cldr.OM:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '/', month)
		}
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(month, '/', day)
		} else {
			seq.Add(day, '-', month)
		}
	case cldr.PCM:
		seq.Add(day, ' ', '/', month)
	case cldr.SD:
		if script == cldr.Deva {
			seq.Add(month, '/', day)
		} else {
			seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			seq.Add(day, '/', month)
		} else {
			seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.SV:
		if region == cldr.RegionFI {
			if opts.Day.numeric() {
				month = symbols.Symbol_M
			}

			seq.Add(day, '.', month)
		} else {
			if opts.Month.twoDigit() && opts.Day.numeric() {
				month = symbols.Symbol_M
				day = symbols.Symbol_d
			}

			seq.Add(day, '/', month)
		}
	case cldr.ZH:
		switch region {
		default:
			seq.Add(month, '/', day)
		case cldr.RegionHK, cldr.RegionMO:
			seq.Add(day, '/', month)
		case cldr.RegionSG:
			seq.Add(month, '-', day)
		}
	case cldr.II:
		seq.Add(symbols.Symbol_MM, symbols.Txt03, symbols.Symbol_dd, symbols.Txtꑍ)

	}

	if seq.Empty() {
		seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
	}

	return seq
}
