package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func seqMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	month := opts.Month.symbolFormat()
	day := opts.Day.symbol()

	switch lang {
	default:
		return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.EN:
		switch region {
		default:
			return seq.Add(month, '/', day)
		case cldr.Region001, cldr.Region150, cldr.RegionAE, cldr.RegionAG, cldr.RegionAI, cldr.RegionAT, cldr.RegionBB,
			cldr.RegionBM, cldr.RegionBS, cldr.RegionBW, cldr.RegionBZ, cldr.RegionCC, cldr.RegionCK, cldr.RegionCM,
			cldr.RegionCX, cldr.RegionCY, cldr.RegionDE, cldr.RegionDG, cldr.RegionDK, cldr.RegionDM, cldr.RegionER,
			cldr.RegionFI, cldr.RegionFJ, cldr.RegionFK, cldr.RegionFM, cldr.RegionGB, cldr.RegionGD, cldr.RegionGG,
			cldr.RegionGH, cldr.RegionGI, cldr.RegionGM, cldr.RegionGY, cldr.RegionHK, cldr.RegionID, cldr.RegionIL,
			cldr.RegionIM, cldr.RegionIN, cldr.RegionIO, cldr.RegionJE, cldr.RegionJM, cldr.RegionKE, cldr.RegionKI,
			cldr.RegionKN, cldr.RegionKY, cldr.RegionLC, cldr.RegionLR, cldr.RegionLS, cldr.RegionMG, cldr.RegionMO,
			cldr.RegionMS, cldr.RegionMT, cldr.RegionMU, cldr.RegionMV, cldr.RegionMW, cldr.RegionMY, cldr.RegionNA,
			cldr.RegionNF, cldr.RegionNG, cldr.RegionNL, cldr.RegionNR, cldr.RegionNU, cldr.RegionPG, cldr.RegionPK,
			cldr.RegionPN, cldr.RegionPW, cldr.RegionRW, cldr.RegionSB, cldr.RegionSC, cldr.RegionSD, cldr.RegionSE,
			cldr.RegionSG, cldr.RegionSH, cldr.RegionSI, cldr.RegionSL, cldr.RegionSS, cldr.RegionSX, cldr.RegionSZ,
			cldr.RegionTC, cldr.RegionTK, cldr.RegionTO, cldr.RegionTT, cldr.RegionTV, cldr.RegionTZ, cldr.RegionUG,
			cldr.RegionVC, cldr.RegionVG, cldr.RegionVU, cldr.RegionWS, cldr.RegionZM:
			if script == cldr.Shaw {
				return seq.Add(month, '/', day)
			}

			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
			}

			return seq.Add(day, '/', month)
		case cldr.RegionAU, cldr.RegionBE, cldr.RegionIE, cldr.RegionNZ, cldr.RegionZW:
			return seq.Add(day, '/', month)
		case cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}

			return seq.Add(month, '-', day)
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
			}

			return seq.Add(day, '.', month)
		case cldr.RegionZA:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=01/02
			// month=2-digit,day=2-digit,out=02/01
			if opts.Month.twoDigit() && opts.Day.twoDigit() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
			}

			return seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
		}
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM)
	case cldr.HI:
		if script == cldr.Latn && opts.Month.numeric() && opts.Day.numeric() {
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		}

		fallthrough
	case cldr.AM, cldr.AGQ, cldr.AST, cldr.BAS, cldr.BM, cldr.CA, cldr.CY, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO,
		cldr.EL, cldr.EWO, cldr.FUR, cldr.GD, cldr.GL, cldr.HAW, cldr.ID, cldr.IG, cldr.KAB, cldr.KGP, cldr.KHQ, cldr.KM,
		cldr.KSF, cldr.KXV, cldr.LN, cldr.LO, cldr.LU, cldr.MAI, cldr.MFE, cldr.MG, cldr.MGH, cldr.ML, cldr.MNI, cldr.MUA,
		cldr.MY, cldr.NMG, cldr.NUS, cldr.PA, cldr.RN, cldr.SA, cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SU, cldr.SW,
		cldr.TO, cldr.TR, cldr.TWQ, cldr.UR, cldr.XNR, cldr.YAV, cldr.YO, cldr.YRL, cldr.ZGH:
		return seq.Add(day, '/', month)
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.UZ, cldr.VEC:
		return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
	case cldr.TI:
		return seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
	case cldr.KEA, cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		}

		return seq.Add(day, '/', month)
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BEZ, cldr.BLO, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.JA, cldr.JMC, cldr.KAA, cldr.KAM, cldr.KDE, cldr.KI, cldr.KLN,
		cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.MHN, cldr.NAQ, cldr.ND, cldr.NYN,
		cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TEO, cldr.TZM, cldr.VAI, cldr.VUN, cldr.XH, cldr.XOG,
		cldr.YUE:
		return seq.Add(month, '/', day)
	case cldr.KS:
		if script == cldr.Deva {
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=01-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(month, '/', day)
	case cldr.AR:
		return seq.Add(day, symbols.Txt02, month)
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		return seq.Add(day, '.', month)
	case cldr.MK:
		return seq.Add(symbols.Symbol_d, '.', month)
	case cldr.BG, cldr.PL:
		return seq.Add(day, '.', symbols.Symbol_MM)
	case cldr.LV:
		return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
	case cldr.DE, cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.IS, cldr.LB, cldr.SMN:
		return seq.Add(day, '.', month, '.')
	case cldr.NB, cldr.NN, cldr.NO: // d.M.
		return seq.Add(symbols.Symbol_d, '.', symbols.Symbol_M, '.')
	case cldr.SQ:
		return seq.Add(symbols.Symbol_d, '.', symbols.Symbol_M)
	case cldr.RO, cldr.RU:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM)
		}

		return seq.Add(day, '.', month)
	case cldr.SR:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '.', ' ', month, '.')
		}

		return seq.Add(day, '.', month, '.')
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.TA, cldr.TE, cldr.VI:
		var sep symbols.Symbol = '-'
		if opts.Month.numeric() && opts.Day.numeric() {
			sep = '/'
		}

		return seq.Add(day, sep, month)
	case cldr.BS:
		if script == cldr.Cyrl {
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.01.
			// month=2-digit,day=numeric,out=02.01.
			// month=2-digit,day=2-digit,out=02.01.
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
		}

		// month=numeric,day=numeric,out=2. 1.
		// month=numeric,day=2-digit,out=2. 1.
		// month=2-digit,day=numeric,out=2. 1.
		// month=2-digit,day=2-digit,out=2. 1.
		return seq.Add(symbols.Symbol_d, '.', ' ', symbols.Symbol_M, '.')
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			month = symbols.Symbol_MM
			day = symbols.Symbol_dd
		}

		return seq.Add(day, '.', ' ', month, '.')
	case cldr.CS, cldr.SK, cldr.SL:
		return seq.Add(day, '.', ' ', month, '.')
	case cldr.HU, cldr.KO:
		return seq.Add(month, '.', ' ', day, '.')
	case cldr.WAE:
		return seq.Add(day, '.', ' ', symbols.Symbol_LLL)
	case cldr.DZ, cldr.SI: // noop
		return seq.Add(month, '-', day)
	case cldr.ES:
		switch region {
		default:
			return seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
		case cldr.RegionCL:
			// month=numeric,day=numeric,out=02-01
			// month=numeric,day=2-digit,out=02-01
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit

				return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM)
			}

			return seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
		case cldr.RegionMX, cldr.RegionUS:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return seq.Add(day, '/', month)
		case cldr.RegionPA, cldr.RegionPR:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			if opts.Month.numeric() {
				return seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
			}

			return seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
		}
	case cldr.FF:
		if script == cldr.Adlm {
			return seq.Add(day, '-', month)
		}

		return seq.Add(day, '/', month)
	case cldr.FR:
		switch region {
		default:
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM)
		case cldr.RegionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month.numeric() && opts.Day.twoDigit() {
				return seq.Add(symbols.Symbol_M, '-', day)
			}

			return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		case cldr.RegionCH:
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
			}

			return seq.Add(day, '.', month)
		}
	case cldr.NL:
		if region == cldr.RegionBE {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return seq.Add(day, '/', month)
		}

		return seq.Add(day, '-', month)
	case cldr.FY, cldr.UG:
		return seq.Add(day, '-', month)
	case cldr.IU:
		return seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_MM, '-', day)
		}

		return seq.Add(month, '-', day)
	case cldr.MN:
		return seq.Add(symbols.Symbol_LLLLL, '/', symbols.Symbol_dd)
	case cldr.MS:
		if !opts.Month.numeric() || !opts.Day.numeric() {
			return seq.Add(day, '/', month)
		}

		return seq.Add(day, '-', month)
	case cldr.OM:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(day, '/', month)
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(month, '/', day)
		}

		return seq.Add(day, '-', month)
	case cldr.PCM:
		return seq.Add(day, ' ', '/', month)
	case cldr.SD:
		if script == cldr.Deva {
			// month=numeric,day=numeric,out=1/2
			// month=numeric,day=2-digit,out=1/02
			// month=2-digit,day=numeric,out=01/2
			// month=2-digit,day=2-digit,out=01/02
			return seq.Add(month, '/', day)
		}

		return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.SE:
		if region == cldr.RegionFI {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return seq.Add(day, '/', month)
		}

		return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.SV:
		if region == cldr.RegionFI {
			// month=numeric,day=numeric,out=2.1
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.1
			// month=2-digit,day=2-digit,out=02.01
			if opts.Day.numeric() {
				month = symbols.Symbol_M
			}

			return seq.Add(day, '.', month)
		}

		if opts.Month.twoDigit() && opts.Day.numeric() {
			month = symbols.Symbol_M
			day = symbols.Symbol_d
		}

		return seq.Add(day, '/', month)
	case cldr.ZH:
		switch region {
		default:
			return seq.Add(month, '/', day)
		case cldr.RegionHK, cldr.RegionMO:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			return seq.Add(day, '/', month)
		case cldr.RegionSG:
			return seq.Add(month, '-', day)
		}
	case cldr.II:
		// month=numeric,day=numeric,out=01ꆪ-02ꑍ
		// month=numeric,day=2-digit,out=01ꆪ-02ꑍ
		// month=2-digit,day=numeric,out=01ꆪ-02ꑍ
		// month=2-digit,day=2-digit,out=01ꆪ-02ꑍ
		return seq.Add(symbols.Symbol_MM, symbols.Txt03, symbols.Symbol_dd, symbols.Txtꑍ)
	case cldr.KOK:
		if script == cldr.Latn {
			return seq.Add(day, '/', month)
		}

		return seq.Add(day, '-', month)
	}
}

func seqMonthDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale)
	month := opts.Month.symbolFormat()

	if lang == cldr.TH {
		return seq.Add(opts.Day.symbol(), '/', month)
	}

	return seq.Add(month, '-', symbols.Symbol_dd)
}

func seqMonthDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale)

	switch lang {
	default:
		return seq.Add(symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.FA, cldr.PS:
		return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol())
	}
}
