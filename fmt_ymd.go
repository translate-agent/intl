package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, _ := locale.Region()
	seq := symbols.NewSeq(locale)
	year := opts.Year.symbol()
	month := opts.Month.symbolFormat()
	day := opts.Day.symbol()

	switch lang {
	case cldr.TH:
		if region == cldr.RegionTH {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.SHN:
		if region == cldr.RegionTH {
			seq.Add(symbols.Symbol_G, ' ', year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.FA:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(year, '/', symbols.Symbol_MM, '/', symbols.Symbol_dd)
		}
	case cldr.LRC, cldr.MZN, cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(symbols.Symbol_G, ' ', year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else if lang == cldr.PS {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.CKB:
		if region == cldr.RegionIR || opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.ES:
		switch region {
		default:
			seq.Add(day, '/', month, '/', year)
		case cldr.RegionCL:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
			} else {
				seq.Add(day, '-', month, '-', year)
			}
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd, '/', year)
			} else {
				seq.Add(month, '/', day, '/', year)
			}
		}
	case cldr.AGQ, cldr.AM, cldr.ASA, cldr.AST, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BM, cldr.BN, cldr.CA, cldr.CCP,
		cldr.CGG, cldr.CY, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EBU, cldr.EL, cldr.EWO, cldr.GD, cldr.GL,
		cldr.GU, cldr.HAW, cldr.HI, cldr.ID, cldr.IG, cldr.KM, cldr.KN, cldr.KSF, cldr.KXV, cldr.LN, cldr.LO, cldr.LU,
		cldr.MAI, cldr.MGH, cldr.ML, cldr.MNI, cldr.MR, cldr.MS, cldr.MUA, cldr.MY, cldr.NMG, cldr.NNH, cldr.NUS, cldr.PCM,
		cldr.RN, cldr.SA, cldr.SCN, cldr.SU, cldr.SW, cldr.TA, cldr.TO, cldr.TWQ, cldr.UR, cldr.VI, cldr.XNR, cldr.YAV:
		seq.Add(day, '/', month, '/', year)
	case cldr.PA:
		if script == cldr.Arab && opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.AK:
		if opts.Year.numeric() {
			seq.Add(year, '/', month, '/', day)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.EU, cldr.JA, cldr.YUE:
		seq.Add(year, '/', month, '/', day)
	case cldr.AR:
		seq.Add(day, symbols.Txt02, month, symbols.Txt02, year)
	case cldr.KK:
		if script == cldr.Arab {
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(year, '-', day, '-', month)
			} else {
				seq.Add(day, '-', month, '-', year)
			}

			break
		}

		fallthrough
	case cldr.AZ, cldr.HY, cldr.UK:
		if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.BE, cldr.DA, cldr.DE, cldr.DSB, cldr.ET, cldr.FI, cldr.HE, cldr.HSB, cldr.IE, cldr.IS, cldr.KA, cldr.LB,
		cldr.NB, cldr.NN, cldr.NO, cldr.SMN, cldr.SQ:
		seq.Add(day, '.', month, '.', year)
	case cldr.BG:
		seq.Add(day, '.', symbols.Symbol_MM, '.', year, ' ', symbols.Txt00)
	case cldr.MK:
		seq.Add(day, '.', month, '.', year, ' ', symbols.Txt00)
	case cldr.EN:
		switch region {
		default:
			seq.Add(month, '/', day, '/', year)
		// TODO(jhorsts): split regions? Do all these regions have shaw script?
		case cldr.Region001, cldr.Region150, cldr.RegionAE, cldr.RegionAG, cldr.RegionAI, cldr.RegionAT, cldr.RegionBB,
			cldr.RegionBM, cldr.RegionBS, cldr.RegionCC, cldr.RegionCK, cldr.RegionCM, cldr.RegionCX, cldr.RegionCY,
			cldr.RegionDE, cldr.RegionDG, cldr.RegionDK, cldr.RegionDM, cldr.RegionER, cldr.RegionFI, cldr.RegionFJ,
			cldr.RegionFK, cldr.RegionFM, cldr.RegionGB, cldr.RegionGD, cldr.RegionGG, cldr.RegionGH, cldr.RegionGI,
			cldr.RegionGM, cldr.RegionGY, cldr.RegionID, cldr.RegionIL, cldr.RegionIM, cldr.RegionIO, cldr.RegionJE,
			cldr.RegionJM, cldr.RegionKE, cldr.RegionKI, cldr.RegionKN, cldr.RegionKY, cldr.RegionLC, cldr.RegionLR,
			cldr.RegionLS, cldr.RegionMG, cldr.RegionMO, cldr.RegionMS, cldr.RegionMT, cldr.RegionMU, cldr.RegionMW,
			cldr.RegionMY, cldr.RegionNA, cldr.RegionNF, cldr.RegionNG, cldr.RegionNL, cldr.RegionNR, cldr.RegionNU,
			cldr.RegionPG, cldr.RegionPK, cldr.RegionPN, cldr.RegionPW, cldr.RegionRW, cldr.RegionSB, cldr.RegionSC,
			cldr.RegionSD, cldr.RegionSH, cldr.RegionSI, cldr.RegionSL, cldr.RegionSS, cldr.RegionSX, cldr.RegionSZ,
			cldr.RegionTC, cldr.RegionTK, cldr.RegionTO, cldr.RegionTT, cldr.RegionTV, cldr.RegionTZ, cldr.RegionUG,
			cldr.RegionVC, cldr.RegionVG, cldr.RegionVU, cldr.RegionWS, cldr.RegionZM:
			switch {
			case script == cldr.Shaw:
				seq.Add(month, '/', day, '/', year)
			case opts.Month.numeric() && opts.Day.numeric():
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			default:
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionCZ, cldr.RegionEE, cldr.RegionES, cldr.RegionFR, cldr.RegionGE, cldr.RegionGS, cldr.RegionHU,
			cldr.RegionIT, cldr.RegionLT, cldr.RegionLV, cldr.RegionNO, cldr.RegionPL, cldr.RegionPT, cldr.RegionRO,
			cldr.RegionSK, cldr.RegionUA:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionAU, cldr.RegionSG:
			if opts.Year.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionBE, cldr.RegionHK, cldr.RegionIE, cldr.RegionIN, cldr.RegionZW:
			seq.Add(day, '/', month, '/', year)
		case cldr.RegionBW, cldr.RegionBZ:
			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionCA, cldr.RegionSE:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
			} else {
				seq.Add(day, '.', month, '.', year)
			}
		case cldr.RegionMV:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '-', month, '-', year)
			}
		case cldr.RegionNZ:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(day, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionJP, cldr.RegionZA:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(year, '/', symbols.Symbol_MM, '/', symbols.Symbol_dd)
			} else {
				seq.Add(year, '/', month, '/', day)
			}
		}
	case cldr.BLO, cldr.CEB, cldr.CHR, cldr.EE, cldr.FIL, cldr.KAA, cldr.MHN, cldr.OM, cldr.OR, cldr.TI, cldr.XH:
		seq.Add(month, '/', day, '/', year)
	case cldr.KS:
		if script == cldr.Deva && opts.Year.twoDigit() {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(month, '/', day, '/', year)
		}
	case cldr.YRL:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.GA:
		if !opts.Year.twoDigit() &&
			(opts.Month.twoDigit() || opts.Day.twoDigit() || opts.Month.numeric() && opts.Day.numeric()) {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.KEA, cldr.KGP:
		if opts.Month.numeric() && opts.Day.numeric() {
			day = symbols.Symbol_dd
			month = symbols.Symbol_MM
		}

		seq.Add(day, '/', month, '/', year)
	case cldr.BS:
		if script == cldr.Cyrl {
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year, '.')
			} else {
				seq.Add(day, '.', month, '.', year, '.')
			}
		} else {
			seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
		}
	case cldr.CS, cldr.SK, cldr.SL:
		seq.Add(day, '.', ' ', month, '.', ' ', year)
	case cldr.FO, cldr.KU, cldr.RO, cldr.RU, cldr.TK, cldr.TT:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.BUA:
		switch {
		default:
			seq.Add(day, '.', month, '.', year)
		case opts.Month.twoDigit() || opts.Day.twoDigit():
		case opts.Month.numeric() && opts.Day.numeric():
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}
	case cldr.DZ, cldr.SI:
		seq.Add(year, '-', month, '-', day)
	case cldr.EO:
		switch {
		case opts.Year.numeric():
			seq.Add(symbols.Symbol_y, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		case opts.Month.numeric() && opts.Day.numeric():
			seq.Add(symbols.Symbol_yy, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.SV:
		switch region {
		default:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}
		case cldr.RegionAX:
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		case cldr.RegionFI:
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.KAB, cldr.KHQ, cldr.KSH, cldr.MFE, cldr.ZGH, cldr.SEH, cldr.SES, cldr.SG, cldr.SHI:
		seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.FF:
		if script == cldr.Adlm {
			seq.Add(day, '-', month, '-', year)
		} else {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.FR:
		switch region {
		default:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}
		case cldr.RegionCH:
			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
			} else {
				seq.Add(day, '.', month, '.', year)
			}
		case cldr.RegionBE:
			if opts.Year.numeric() {
				if opts.Month.numeric() && opts.Day.numeric() {
					seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
				} else {
					seq.Add(day, '/', symbols.Symbol_MM, '/', year)
				}

				break
			}

			seq.Add(day, '/', month, '/', year)
		}
	case cldr.VAI:
		if script == cldr.Latn {
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(month, '/', day, '/', year)
			} else {
				seq.Add(day, '/', month, '/', year)
			}

			break
		}

		fallthrough
	case cldr.FUR, cldr.GUZ, cldr.JMC, cldr.KAM, cldr.KDE, cldr.KI, cldr.KLN, cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO,
		cldr.LUY, cldr.MAS, cldr.MER, cldr.NAQ, cldr.ND, cldr.NYN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.TEO, cldr.TZM,
		cldr.VUN, cldr.XOG:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.NL:
		if region == cldr.RegionBE {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(day, '-', month, '-', year)
		}
	case cldr.FY, cldr.KOK, cldr.RM:
		seq.Add(day, '-', month, '-', year)
	case cldr.GSW:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.HA, cldr.SAT:
		if opts.Year.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.HR:
		switch region {
		default:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', ' ', symbols.Symbol_MM, '.', ' ', year, '.')
			} else {
				seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
			}
		case cldr.RegionBA:
			if opts.Year.numeric() {
				seq.Add(symbols.Symbol_dd, '.', ' ', symbols.Symbol_MM, '.', ' ', year, '.')
			} else {
				seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
			}
		}
	case cldr.HU:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '.', ' ', symbols.Symbol_MM, '.', ' ', symbols.Symbol_dd, '.')
		} else {
			seq.Add(year, '.', ' ', month, '.', ' ', day, '.')
		}
	case cldr.NDS, cldr.PRG:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.IT:
		if region == cldr.RegionCH {
			if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
				opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			} else {
				seq.Add(day, '.', month, '.', year)
			}

			break
		}

		fallthrough
	case cldr.UZ:
		if region == cldr.RegionAF {
			_, _, rawRegion := locale.Raw()
			if rawRegion == cldr.RegionAF {
				seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			} else {
				seq.Add(symbols.Symbol_GGGGG, ' ', year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}

			break
		}

		fallthrough
	case cldr.VEC:
		switch {
		default:
			seq.Add(day, '/', month, '/', year)
		case opts.Year.numeric():
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		case opts.Month.numeric() && opts.Day.numeric():
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', symbols.Symbol_yy)
		}
	case cldr.JGO:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(month, '.', day, '.', year)
		}
	case cldr.KKJ:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, ' ', year)
		} else {
			seq.Add(day, '/', month, ' ', year)
		}
	case cldr.KO:
		seq.Add(year, '.', ' ', month, '.', ' ', day, '.')
	case cldr.KY:
		if opts.Year.numeric() {
			seq.Add(year, '-', symbols.Symbol_dd, '-', symbols.Symbol_MM)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.LIJ, cldr.VMW:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.LKT, cldr.ZU:
		if opts.Year.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(month, '/', day, '/', year)
		}
	case cldr.LV:
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Year.twoDigit() && opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '.', symbols.Symbol_MM, '.', year, '.')
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.AS, cldr.BRX, cldr.IA, cldr.JV, cldr.MI, cldr.WO:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		} else {
			seq.Add(day, '-', month, '-', year)
		}
	case cldr.RW:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		}
	case cldr.CV, cldr.MN:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '.', symbols.Symbol_MM, '.', symbols.Symbol_dd)
		} else {
			seq.Add(year, '.', month, '.', day)
		}
	case cldr.MT, cldr.SBP:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(month, '/', day, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.NQO:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, ' ', '/', ' ', symbols.Symbol_dd, ' ', '/', ' ', symbols.Symbol_MM)
		}
	case cldr.OC, cldr.PMS:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		}
	case cldr.OS:
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.PL:
		seq.Add(day, '.', symbols.Symbol_MM, '.', year)
	case cldr.QU:
		if opts.Year.numeric() {
			seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.NE, cldr.SAH:
		if opts.Year.numeric() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else {
			seq.Add(year, '/', month, '/', day)
		}
	case cldr.SD:
		if script == cldr.Deva {
			seq.Add(month, '/', day, '/', year)
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)

				break
			}

			seq.Add(day, '.', month, '.', year)
		}
	case cldr.SO:
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(month, '/', day, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.SR:
		if opts.Month.twoDigit() && opts.Day.twoDigit() {
			seq.Add(day, '.', month, '.', year, '.')
		} else {
			seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
		}
	case cldr.SYR:
		if opts.Month.numeric() {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(day, '-', month, '-', year)
		}
	case cldr.SZL:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}
	case cldr.TE:
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(day, '-', month, '-', year)
		}
	case cldr.TOK:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add('#', year, ')', '#', month, ')', '#', day)
		}
	case cldr.TR:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		} else {
			seq.Add(day, '.', symbols.Symbol_MM, '.', year)
		}
	case cldr.UG:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '-', day, '-', month)
		}
	case cldr.YI:
		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(day, '-', month, '-', year)
		}
	case cldr.YO:
		if opts.Month.twoDigit() {
			seq.Add(day, ' ', month, ' ', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.ZA:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(year, '/', month, '/', day)
		}
	case cldr.ZH:
		switch region {
		default:
			seq.Add(year, '/', month, '/', day)
		case cldr.RegionMO, cldr.RegionSG:
			if script == cldr.Hans {
				seq.Add(year, symbols.Txt年, month, symbols.Txt月, day, symbols.Txt日)
			} else {
				seq.Add(day, '/', month, '/', year)
			}
		case cldr.RegionHK:
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.TG:
		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			seq.Add(day, '/', month, '/', year)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.GAA:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(month, '/', day, '/', year)
		}
	case cldr.BA:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		} else {
			seq.Add(day, '.', month, '.', year)
		}
	case cldr.BR, cldr.SC:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		} else {
			seq.Add(day, '/', month, '/', year)
		}
	case cldr.NSO:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		}
	}

	if opts.Month.numeric() && opts.Day.numeric() {
		if seq.Empty() {
			seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq
	}

	if seq.Empty() {
		seq.Add(year, '-', month, '-', day)
	}

	return seq
}
