package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func seqYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)

	year := symbols.Symbol_y
	if opts.Year.twoDigit() {
		year = symbols.Symbol_yy
	}

	month := symbols.Symbol_M
	if opts.Month.twoDigit() {
		month = symbols.Symbol_MM
	}

	day := symbols.Symbol_d
	if opts.Day.twoDigit() {
		day = symbols.Symbol_dd
	}

	switch lang {
	default:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.ES:
		switch region {
		default:
			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionCL:
			// year=numeric,month=numeric,day=numeric,out=02-01-2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02-01-24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
			}

			return seq.Add(day, '-', month, '-', year)
		case cldr.RegionPA, cldr.RegionPR:
			// year=numeric,month=numeric,day=numeric,out=01/02/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=01/02/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd, '/', year)
			}

			return seq.Add(month, '/', day, '/', year)
		}
	case cldr.AGQ, cldr.AM, cldr.ASA, cldr.AST, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BM, cldr.BN, cldr.CA, cldr.CCP,
		cldr.CGG, cldr.CY, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.EBU, cldr.EL, cldr.EWO, cldr.GD, cldr.GL,
		cldr.GU, cldr.HAW, cldr.HI, cldr.ID, cldr.IG, cldr.KM, cldr.KN, cldr.KSF, cldr.KXV, cldr.LN, cldr.LO, cldr.LU,
		cldr.MAI, cldr.MGH, cldr.ML, cldr.MNI, cldr.MR, cldr.MS, cldr.MUA, cldr.MY, cldr.NMG, cldr.NNH, cldr.NUS, cldr.PCM,
		cldr.RN, cldr.SA, cldr.SU, cldr.SW, cldr.TA, cldr.TO, cldr.TWQ, cldr.UR, cldr.VI, cldr.XNR, cldr.YAV:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		return seq.Add(day, '/', month, '/', year)
	case cldr.PA:
		if script == cldr.Arab && opts.Month.numeric() && opts.Day.numeric() {
			// year=numeric,month=numeric,day=numeric,out=۲۰۲۴-۰۱-۰۲
			// year=numeric,month=numeric,day=2-digit,out=۰۲/۱/۲۰۲۴
			// year=numeric,month=2-digit,day=numeric,out=۲/۰۱/۲۰۲۴
			// year=numeric,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۰۲۴
			// year=2-digit,month=numeric,day=numeric,out=۲۴-۰۱-۰۲
			// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱/۲۴
			// year=2-digit,month=2-digit,day=numeric,out=۲/۰۱/۲۴
			// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۰۱/۲۴
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		return seq.Add(day, '/', month, '/', year)
	case cldr.AK:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		if opts.Year.twoDigit() {
			return seq.Add(month, '/', day, '/', year)
		}

		return seq.Add(year, '/', month, '/', day)
	case cldr.EU, cldr.JA, cldr.YUE:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		return seq.Add(year, '/', month, '/', day)
	case cldr.AR:
		// year=numeric,month=numeric,day=numeric,out=٢‏/١‏/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٠٢٤
		// year=2-digit,month=numeric,day=numeric,out=٢‏/١‏/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٤
		// year=2-digit,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٤
		// year=2-digit,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٤
		return seq.Add(day, symbols.Txt02, month, symbols.Txt02, year)
	case cldr.AZ, cldr.HY, cldr.KK, cldr.UK:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=02.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Year.numeric() ||
			opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.BE, cldr.DA, cldr.DE, cldr.DSB, cldr.ET, cldr.FI, cldr.HE, cldr.HSB, cldr.IE, cldr.IS, cldr.KA, cldr.LB,
		cldr.NB, cldr.NN, cldr.NO, cldr.SMN, cldr.SQ:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		return seq.Add(day, '.', month, '.', year)
	case cldr.BG:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.01.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		return seq.Add(day, '.', symbols.Symbol_MM, '.', year, ' ', symbols.Txt00)
	case cldr.MK:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.1.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		return seq.Add(day, '.', month, '.', year, ' ', symbols.Txt00)
	case cldr.EN:
		switch region {
		default:
			return seq.Add(month, '/', day, '/', year)
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
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if script == cldr.Shaw {
				return seq.Add(month, '/', day, '/', year)
			}

			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionAU, cldr.RegionSG:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Year.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionBE, cldr.RegionHK, cldr.RegionIE, cldr.RegionIN, cldr.RegionZW:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionBW, cldr.RegionBZ:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionCA, cldr.RegionSE:
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-1-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-2
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-1-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-2
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}

			return seq.Add(year, '-', month, '-', day)
		case cldr.RegionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
			}

			return seq.Add(day, '.', month, '.', year)
		case cldr.RegionMV:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '-', month, '-', year)
		case cldr.RegionNZ:
			// year=numeric,month=numeric,day=numeric,out=2/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(day, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionZA:
			// year=numeric,month=numeric,day=numeric,out=2024/01/02
			// year=numeric,month=numeric,day=2-digit,out=2024/1/02
			// year=numeric,month=2-digit,day=numeric,out=2024/01/2
			// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
			// year=2-digit,month=numeric,day=numeric,out=24/01/02
			// year=2-digit,month=numeric,day=2-digit,out=24/1/02
			// year=2-digit,month=2-digit,day=numeric,out=24/01/2
			// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(year, '/', symbols.Symbol_MM, '/', symbols.Symbol_dd)
			}

			return seq.Add(year, '/', month, '/', day)
		}
	case cldr.BLO, cldr.CEB, cldr.CHR, cldr.EE, cldr.FIL, cldr.KAA, cldr.MHN, cldr.OM, cldr.OR, cldr.TI, cldr.XH:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		return seq.Add(month, '/', day, '/', year)
	case cldr.KS:
		if script == cldr.Deva && opts.Year.twoDigit() {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return seq.Add(day, '/', month, '/', year)
		}

		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		return seq.Add(month, '/', day, '/', year)
	case cldr.BR, cldr.GA, cldr.KEA, cldr.KGP, cldr.PT, cldr.SC, cldr.YRL:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.BS:
		if script == cldr.Cyrl {
			// year=numeric,month=numeric,day=numeric,out=02.01.2024.
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024.
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
			// year=2-digit,month=numeric,day=numeric,out=02.01.24.
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24.
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24.
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year, '.')
			}

			return seq.Add(day, '.', month, '.', year, '.')
		}

		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
		// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
		// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
		return seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
	case cldr.CKB:
		// year=numeric,month=numeric,day=numeric,out=٢/١/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٢٠٢٤-١-٠٢
		// year=numeric,month=2-digit,day=numeric,out=٢٠٢٤-٠١-٢
		// year=numeric,month=2-digit,day=2-digit,out=٢٠٢٤-٠١-٠٢
		// year=2-digit,month=numeric,day=numeric,out=٢/١/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٢٤-١-٠٢
		// year=2-digit,month=2-digit,day=numeric,out=٢٤-٠١-٢
		// year=2-digit,month=2-digit,day=2-digit,out=٢٤-٠١-٠٢
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.CS, cldr.SK, cldr.SL:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024
		// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24
		// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24
		return seq.Add(day, '.', ' ', month, '.', ' ', year)
	case cldr.CV, cldr.FO, cldr.KU, cldr.RO, cldr.RU, cldr.TK, cldr.TT:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.DZ, cldr.SI: // noop
		// year=numeric,month=numeric,day=numeric,out=2024-1-2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-1-2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		return seq.Add(year, '-', month, '-', day)
	case cldr.EO:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		switch {
		default:
			return seq.Add(year, '-', month, '-', day)
		case opts.Year.numeric():
			return seq.Add(symbols.Symbol_y, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		case opts.Month.numeric() && opts.Day.numeric():
			return seq.Add(symbols.Symbol_yy, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.KAB, cldr.KHQ, cldr.KSH, cldr.MFE, cldr.ZGH, cldr.PS, cldr.SEH, cldr.SES, cldr.SG, cldr.SHI:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-01-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-02
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.FF:
		if script == cldr.Adlm {
			// year=numeric,month=numeric,day=numeric,out=𞥒-𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=numeric,day=2-digit,out=𞥐𞥒-𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=2-digit,day=numeric,out=𞥒-𞥐𞥑-𞥒𞥐𞥒𞥔
			// year=numeric,month=2-digit,day=2-digit,out=𞥐𞥒-𞥐𞥑-𞥒𞥐𞥒𞥔
			// year=2-digit,month=numeric,day=numeric,out=𞥒-𞥑-𞥒𞥔
			// year=2-digit,month=numeric,day=2-digit,out=𞥐𞥒-𞥑-𞥒𞥔
			// year=2-digit,month=2-digit,day=numeric,out=𞥒-𞥐𞥑-𞥒𞥔
			// year=2-digit,month=2-digit,day=2-digit,out=𞥐𞥒-𞥐𞥑-𞥒𞥔
			return seq.Add(day, '-', month, '-', year)
		}

		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=24-01-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-02
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.FR:
		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionCA:
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-1-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-2
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-1-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-2
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}

			return seq.Add(year, '-', month, '-', day)
		case cldr.RegionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.01.2024
			// year=numeric,month=2-digit,day=numeric,out=02.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
			}

			return seq.Add(day, '.', month, '.', year)
		case cldr.RegionBE:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Year.numeric() {
				if opts.Month.numeric() && opts.Day.numeric() {
					return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
				}

				return seq.Add(day, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		}
	case cldr.VAI:
		if script == cldr.Latn {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(month, '/', day, '/', year)
			}

			return seq.Add(day, '/', month, '/', year)
		}

		fallthrough
	case cldr.FUR, cldr.GUZ, cldr.JMC, cldr.KAM, cldr.KDE, cldr.KI, cldr.KLN, cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO,
		cldr.LUY, cldr.MAS, cldr.MER, cldr.NAQ, cldr.ND, cldr.NYN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.TEO, cldr.TZM,
		cldr.VUN, cldr.XOG:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.NL:
		if region == cldr.RegionBE {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(day, '-', month, '-', year)
	case cldr.FY, cldr.KOK:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		return seq.Add(day, '-', month, '-', year)
	case cldr.GSW:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.HA, cldr.SAT:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.HR:
		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=02. 01. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
			// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=02. 01. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', ' ', symbols.Symbol_MM, '.', ' ', year, '.')
			}

			return seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
		case cldr.RegionBA:
			// year=numeric,month=numeric,day=numeric,out=02. 01. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 01. 2024.
			// year=numeric,month=2-digit,day=numeric,out=02. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			if opts.Year.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', ' ', symbols.Symbol_MM, '.', ' ', year, '.')
			}

			return seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
		}
	case cldr.HU:
		// year=numeric,month=numeric,day=numeric,out=2024. 01. 02.
		// year=numeric,month=numeric,day=2-digit,out=2024. 1. 02.
		// year=numeric,month=2-digit,day=numeric,out=2024. 01. 2.
		// year=numeric,month=2-digit,day=2-digit,out=2024. 01. 02.
		// year=2-digit,month=numeric,day=numeric,out=24. 01. 02.
		// year=2-digit,month=numeric,day=2-digit,out=24. 1. 02.
		// year=2-digit,month=2-digit,day=numeric,out=24. 01. 2.
		// year=2-digit,month=2-digit,day=2-digit,out=24. 01. 02.
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '.', ' ', symbols.Symbol_MM, '.', ' ', symbols.Symbol_dd, '.')
		}

		return seq.Add(year, '.', ' ', month, '.', ' ', day, '.')
	case cldr.NDS, cldr.PRG:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '.', month, '.', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.IT:
		if region == cldr.RegionCH {
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
				opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
			}

			return seq.Add(day, '.', month, '.', year)
		}

		fallthrough
	case cldr.VEC, cldr.UZ:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=02/01/2024
		// year=numeric,month=2-digit,day=numeric,out=02/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		switch {
		default:
			return seq.Add(day, '/', month, '/', year)
		case opts.Year.numeric():
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		case opts.Month.numeric() && opts.Day.numeric():
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', symbols.Symbol_yy)
		}
	case cldr.JGO:
		// year=numeric,month=numeric,day=numeric,out=1.2.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1.2.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(month, '.', day, '.', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.KKJ:
		// year=numeric,month=numeric,day=numeric,out=02/01 2024
		// year=numeric,month=numeric,day=2-digit,out=02/1 2024
		// year=numeric,month=2-digit,day=numeric,out=2/01 2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01 2024
		// year=2-digit,month=numeric,day=numeric,out=02/01 24
		// year=2-digit,month=numeric,day=2-digit,out=02/1 24
		// year=2-digit,month=2-digit,day=numeric,out=2/01 24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01 24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, ' ', year)
		}

		return seq.Add(day, '/', month, ' ', year)
	case cldr.KO:
		// year=numeric,month=numeric,day=numeric,out=2024. 1. 2.
		// year=numeric,month=numeric,day=2-digit,out=2024. 1. 02.
		// year=numeric,month=2-digit,day=numeric,out=2024. 01. 2.
		// year=numeric,month=2-digit,day=2-digit,out=2024. 01. 02.
		// year=2-digit,month=numeric,day=numeric,out=24. 1. 2.
		// year=2-digit,month=numeric,day=2-digit,out=24. 1. 02.
		// year=2-digit,month=2-digit,day=numeric,out=24. 01. 2.
		// year=2-digit,month=2-digit,day=2-digit,out=24. 01. 02.
		return seq.Add(year, '.', ' ', month, '.', ' ', day, '.')
	case cldr.KY:
		// year=numeric,month=numeric,day=numeric,out=2024-02-01
		// year=numeric,month=numeric,day=2-digit,out=2024-02-01
		// year=numeric,month=2-digit,day=numeric,out=2024-02-01
		// year=numeric,month=2-digit,day=2-digit,out=2024-02-01
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() {
			return seq.Add(year, '-', symbols.Symbol_dd, '-', symbols.Symbol_MM)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.LIJ, cldr.VMW:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.LKT, cldr.ZU:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		if opts.Year.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(month, '/', day, '/', year)
	case cldr.LV:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Year.twoDigit() && opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '.', symbols.Symbol_MM, '.', year, '.')
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.AS, cldr.BRX, cldr.IA, cldr.JV, cldr.MI, cldr.RM, cldr.WO:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		}

		return seq.Add(day, '-', month, '-', year)
	case cldr.RW:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.MN:
		// year=numeric,month=numeric,day=numeric,out=2024.01.02
		// year=numeric,month=numeric,day=2-digit,out=2024.1.02
		// year=numeric,month=2-digit,day=numeric,out=2024.01.2
		// year=numeric,month=2-digit,day=2-digit,out=2024.01.02
		// year=2-digit,month=numeric,day=numeric,out=24.01.02
		// year=2-digit,month=numeric,day=2-digit,out=24.1.02
		// year=2-digit,month=2-digit,day=numeric,out=24.01.2
		// year=2-digit,month=2-digit,day=2-digit,out=24.01.02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '.', symbols.Symbol_MM, '.', symbols.Symbol_dd)
		}

		return seq.Add(year, '.', month, '.', day)
	case cldr.MT, cldr.SBP:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(month, '/', day, '/', year)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.NE:
		// year=numeric,month=numeric,day=numeric,out=२०२४-०१-०२
		// year=numeric,month=numeric,day=2-digit,out=२०२४-०१-०२
		// year=numeric,month=2-digit,day=numeric,out=२०२४-०१-०२
		// year=numeric,month=2-digit,day=2-digit,out=२०२४-०१-०२
		// year=2-digit,month=numeric,day=numeric,out=२४/१/२
		// year=2-digit,month=numeric,day=2-digit,out=२४/१/०२
		// year=2-digit,month=2-digit,day=numeric,out=२४/०१/२
		// year=2-digit,month=2-digit,day=2-digit,out=२४/०१/०२
		if opts.Year.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(year, '/', month, '/', day)
	case cldr.NQO:
		// year=numeric,month=numeric,day=numeric,out=߂߀߂߄ / ߀߂ / ߀߁
		// year=numeric,month=numeric,day=2-digit,out=߂߀߂߄-߁-߀߂
		// year=numeric,month=2-digit,day=numeric,out=߂߀߂߄-߀߁-߂
		// year=numeric,month=2-digit,day=2-digit,out=߂߀߂߄-߀߁-߀߂
		// year=2-digit,month=numeric,day=numeric,out=߂߄ / ߀߂ / ߀߁
		// year=2-digit,month=numeric,day=2-digit,out=߂߄-߁-߀߂
		// year=2-digit,month=2-digit,day=numeric,out=߂߄-߀߁-߂
		// year=2-digit,month=2-digit,day=2-digit,out=߂߄-߀߁-߀߂
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, ' ', '/', ' ', symbols.Symbol_dd, ' ', '/', ' ', symbols.Symbol_MM)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.OC:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '/', symbols.Symbol_MM, '/', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.OS:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=24-01-02
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.PL:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		return seq.Add(day, '.', symbols.Symbol_MM, '.', year)
	case cldr.QU:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-01-2024
		// year=numeric,month=2-digit,day=numeric,out=02-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() {
			return seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM, '-', year)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.SAH:
		// year=numeric,month=numeric,day=numeric,out=2024-01-02
		// year=numeric,month=numeric,day=2-digit,out=2024-01-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-02
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		if opts.Year.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(year, '/', month, '/', day)
	case cldr.SD:
		if script == cldr.Deva {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			return seq.Add(month, '/', day, '/', year)
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.SE:
		if region == cldr.RegionFI {
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
			}

			return seq.Add(day, '.', month, '.', year)
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.SO:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(month, '/', day, '/', year)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.SR:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
		if opts.Month.twoDigit() && opts.Day.twoDigit() {
			return seq.Add(day, '.', month, '.', year, '.')
		}

		return seq.Add(day, '.', ' ', month, '.', ' ', year, '.')
	case cldr.SYR:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		if opts.Month.numeric() {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(day, '-', month, '-', year)
	case cldr.SZL:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.TE:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(day, '-', month, '-', year)
	case cldr.TOK:
		// year=numeric,month=numeric,day=numeric,out=#2024)#1)#2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=#24)#1)#2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add('#', year, ')', '#', month, ')', '#', day)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.TR:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(symbols.Symbol_dd, '.', symbols.Symbol_MM, '.', year)
		}

		return seq.Add(day, '.', symbols.Symbol_MM, '.', year)
	case cldr.UG:
		// year=numeric,month=numeric,day=numeric,out=2024-2-1
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-2-1
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '-', day, '-', month)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.YI:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(day, '-', month, '-', year)
	case cldr.YO:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2 01 2024
		// year=numeric,month=2-digit,day=2-digit,out=02 01 2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2 01 24
		// year=2-digit,month=2-digit,day=2-digit,out=02 01 24
		if opts.Month.twoDigit() {
			return seq.Add(day, ' ', month, ' ', year)
		}

		return seq.Add(day, '/', month, '/', year)
	case cldr.ZA:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(year, '/', month, '/', day)
		}

		return seq.Add(year, '-', month, '-', day)
	case cldr.ZH:
		switch region {
		default:
			// year=numeric,month=numeric,day=numeric,out=2024/1/2
			// year=numeric,month=numeric,day=2-digit,out=2024/1/02
			// year=numeric,month=2-digit,day=numeric,out=2024/01/2
			// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
			// year=2-digit,month=numeric,day=numeric,out=24/1/2
			// year=2-digit,month=numeric,day=2-digit,out=24/1/02
			// year=2-digit,month=2-digit,day=numeric,out=24/01/2
			// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
			return seq.Add(year, '/', month, '/', day)
		case cldr.RegionMO, cldr.RegionSG:
			if script == cldr.Hans {
				// year=numeric,month=numeric,day=numeric,out=2024年1月2日
				// year=numeric,month=numeric,day=2-digit,out=2024年1月02日
				// year=numeric,month=2-digit,day=numeric,out=2024年01月2日
				// year=numeric,month=2-digit,day=2-digit,out=2024年01月02日
				// year=2-digit,month=numeric,day=numeric,out=24年1月2日
				// year=2-digit,month=numeric,day=2-digit,out=24年1月02日
				// year=2-digit,month=2-digit,day=numeric,out=24年01月2日
				// year=2-digit,month=2-digit,day=2-digit,out=24年01月02日
				return seq.Add(year, symbols.Txt年, month, symbols.Txt月, day, symbols.Txt日)
			}

			return seq.Add(day, '/', month, '/', year)
		case cldr.RegionHK:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			return seq.Add(day, '/', month, '/', year)
		}
	case cldr.TG:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2.1.24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			return seq.Add(day, '/', month, '/', year)
		}

		return seq.Add(day, '.', month, '.', year)
	case cldr.GAA:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
		// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(month, '/', day, '/', year)
		}

		return seq.Add(year, '-', month, '-', day)
	}
}

func seqYearMonthDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	lang, _, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	year := opts.Year.symbol()
	month := Month2Digit.symbol("format")
	day := Day2Digit.symbol()

	switch lang {
	case cldr.CKB: // ckb-IR
		// year=numeric,month=numeric,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=numeric,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=2-digit,month=numeric,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=numeric,day=2-digit,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=2-digit,out=١٢/١٠/٠٢
		return seq.Add(day, '/', month, '/', year)
	case cldr.FA: // fa-IR
		// year=numeric,month=numeric,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=numeric,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۱۰/۱۲
		return seq.Add(year, '/', month, '/', day)
	case cldr.UZ:
		if region == cldr.RegionAF {
			return seq.Add(year, '-', month, '-', day)
		}
	}

	return seq.Add(EraShort.symbol(), ' ', year, '-', month, '-', day)
}

func seqYearMonthDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	// th-TH
	// year=numeric,month=numeric,day=numeric,out=2/1/2024
	// year=numeric,month=numeric,day=2-digit,out=02/1/2024
	// year=numeric,month=2-digit,day=numeric,out=2/01/2024
	// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
	// year=2-digit,month=numeric,day=numeric,out=2/1/24
	// year=2-digit,month=numeric,day=2-digit,out=02/1/24
	// year=2-digit,month=2-digit,day=numeric,out=2/01/24
	// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
	return symbols.NewSeq(locale).Add(opts.Day.symbol(), '/', opts.Month.symbol("format"), '/', opts.Year.symbol())
}
