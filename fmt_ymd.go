package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtYearMonthDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		layoutYearMonthDay = iota
		layoutDayMonthYear
		layoutMonthDayYear
		layoutYearDayMonth
	)

	layout := layoutYearMonthDay
	prefix := ""
	separator := "-"
	suffix := ""

	switch lang {
	default:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.ES:
		switch region {
		default:
			layout = layoutDayMonthYear
			separator = "/"
		case cldr.RegionCL:
			// year=numeric,month=numeric,day=numeric,out=02-01-2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02-01-24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			layout = layoutDayMonthYear

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionPA, cldr.RegionPR:
			// year=numeric,month=numeric,day=numeric,out=01/02/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=01/02/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			layout = layoutMonthDayYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
		layout = layoutDayMonthYear
		separator = "/"
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
		}
	case cldr.AK:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=1/02/24
		// year=2-digit,month=2-digit,day=numeric,out=01/2/24
		// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
		separator = "/"

		if opts.Year.twoDigit() {
			layout = layoutMonthDayYear
		}
	case cldr.EU, cldr.JA, cldr.YUE:
		// year=numeric,month=numeric,day=numeric,out=2024/1/2
		// year=numeric,month=numeric,day=2-digit,out=2024/1/02
		// year=numeric,month=2-digit,day=numeric,out=2024/01/2
		// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
		// year=2-digit,month=numeric,day=numeric,out=24/1/2
		// year=2-digit,month=numeric,day=2-digit,out=24/1/02
		// year=2-digit,month=2-digit,day=numeric,out=24/01/2
		// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
		separator = "/"
	case cldr.AR:
		// year=numeric,month=numeric,day=numeric,out=٢‏/١‏/٢٠٢٤
		// year=numeric,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٠٢٤
		// year=numeric,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٠٢٤
		// year=2-digit,month=numeric,day=numeric,out=٢‏/١‏/٢٤
		// year=2-digit,month=numeric,day=2-digit,out=٠٢‏/١‏/٢٤
		// year=2-digit,month=2-digit,day=numeric,out=٢‏/٠١‏/٢٤
		// year=2-digit,month=2-digit,day=2-digit,out=٠٢‏/٠١‏/٢٤
		layout = layoutDayMonthYear
		separator = "\u200f/"
	case cldr.AZ, cldr.HY, cldr.KK, cldr.UK:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=02.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Year.numeric() ||
			opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
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
		layout = layoutDayMonthYear
		separator = "."
	case cldr.BG:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.01.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		opts.Month = Month2Digit
		fallthrough
	case cldr.MK:
		// year=numeric,month=numeric,day=numeric,out=2.1.2024 г.
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024 г.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024 г.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024 г.
		// year=2-digit,month=numeric,day=numeric,out=2.1.24 г.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24 г.
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24 г.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24 г.
		layout = layoutDayMonthYear
		separator = "."
		suffix = " г."
	case cldr.EN:
		switch region {
		default:
			layout = layoutMonthDayYear
			separator = "/"
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
			separator = "/"

			if script == cldr.Shaw {
				layout = layoutMonthDayYear
			} else {
				layout = layoutDayMonthYear

				if opts.Month.numeric() && opts.Day.numeric() {
					opts.Month = Month2Digit
					opts.Day = Day2Digit
				}
			}
		case cldr.RegionAU, cldr.RegionSG:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionBE, cldr.RegionHK, cldr.RegionIE, cldr.RegionIN, cldr.RegionZW:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
		case cldr.RegionBW, cldr.RegionBZ:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=02/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear
			separator = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionMV:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02-1-2024
			// year=numeric,month=2-digit,day=numeric,out=2-01-2024
			// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
			// year=2-digit,month=numeric,day=numeric,out=02/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02-1-24
			// year=2-digit,month=2-digit,day=numeric,out=2-01-24
			// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
			layout = layoutDayMonthYear

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				separator = "/"
			}
		case cldr.RegionNZ:
			// year=numeric,month=numeric,day=numeric,out=2/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/01/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
			}
		case cldr.RegionZA:
			// year=numeric,month=numeric,day=numeric,out=2024/01/02
			// year=numeric,month=numeric,day=2-digit,out=2024/1/02
			// year=numeric,month=2-digit,day=numeric,out=2024/01/2
			// year=numeric,month=2-digit,day=2-digit,out=2024/01/02
			// year=2-digit,month=numeric,day=numeric,out=24/01/02
			// year=2-digit,month=numeric,day=2-digit,out=24/1/02
			// year=2-digit,month=2-digit,day=numeric,out=24/01/2
			// year=2-digit,month=2-digit,day=2-digit,out=24/01/02
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
		layout = layoutMonthDayYear
		separator = "/"
	case cldr.KS:
		separator = "/"

		if script == cldr.Deva && opts.Year.twoDigit() {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
		} else {
			// year=numeric,month=numeric,day=numeric,out=1/2/2024
			// year=numeric,month=numeric,day=2-digit,out=1/02/2024
			// year=numeric,month=2-digit,day=numeric,out=01/2/2024
			// year=numeric,month=2-digit,day=2-digit,out=01/02/2024
			// year=2-digit,month=numeric,day=numeric,out=1/2/24
			// year=2-digit,month=numeric,day=2-digit,out=1/02/24
			// year=2-digit,month=2-digit,day=numeric,out=01/2/24
			// year=2-digit,month=2-digit,day=2-digit,out=01/02/24
			layout = layoutMonthDayYear
		}
	case cldr.BR, cldr.GA, cldr.KEA, cldr.KGP, cldr.PT, cldr.SC, cldr.YRL:
		// year=numeric,month=numeric,day=numeric,out=02/01/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=02/01/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear
		separator = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.BS:
		layout = layoutDayMonthYear
		suffix = "."

		if script == cldr.Cyrl {
			// year=numeric,month=numeric,day=numeric,out=02.01.2024.
			// year=numeric,month=numeric,day=2-digit,out=02.1.2024.
			// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
			// year=2-digit,month=numeric,day=numeric,out=02.01.24.
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24.
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24.
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
			day := opts.Day
			separator = "."

			if opts.Month.numeric() {
				opts.Day = Day2Digit
			}

			if day.numeric() {
				opts.Month = Month2Digit
			}
		} else {
			// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
			// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
			// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
			// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024.
			// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
			// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
			// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
			// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24.
			separator = ". "
		}
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
			layout = layoutDayMonthYear
			separator = "/"
		}
	case cldr.CS, cldr.SK, cldr.SL:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024
		// year=numeric,month=2-digit,day=2-digit,out=02. 01. 2024
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24
		// year=2-digit,month=2-digit,day=2-digit,out=02. 01. 24
		layout = layoutDayMonthYear
		separator = ". "
	case cldr.CV, cldr.FO, cldr.KU, cldr.RO, cldr.RU, cldr.TK, cldr.TT:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.1.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.DZ, cldr.SI: // noop
		// year=numeric,month=numeric,day=numeric,out=2024-1-2
		// year=numeric,month=numeric,day=2-digit,out=2024-1-02
		// year=numeric,month=2-digit,day=numeric,out=2024-01-2
		// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
		// year=2-digit,month=numeric,day=numeric,out=24-1-2
		// year=2-digit,month=numeric,day=2-digit,out=24-1-02
		// year=2-digit,month=2-digit,day=numeric,out=24-01-2
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
		case opts.Year.numeric():
			opts.Year = YearNumeric
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month.numeric() && opts.Day.numeric():
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
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
		opts.Month = Month2Digit
		opts.Day = Day2Digit
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
			layout = layoutDayMonthYear
		} else {
			// year=numeric,month=numeric,day=numeric,out=2024-01-02
			// year=numeric,month=numeric,day=2-digit,out=2024-01-02
			// year=numeric,month=2-digit,day=numeric,out=2024-01-02
			// year=numeric,month=2-digit,day=2-digit,out=2024-01-02
			// year=2-digit,month=numeric,day=numeric,out=24-01-02
			// year=2-digit,month=numeric,day=2-digit,out=24-01-02
			// year=2-digit,month=2-digit,day=numeric,out=24-01-02
			// year=2-digit,month=2-digit,day=2-digit,out=24-01-02
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
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
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionCH:
			// year=numeric,month=numeric,day=numeric,out=02.01.2024
			// year=numeric,month=numeric,day=2-digit,out=02.01.2024
			// year=numeric,month=2-digit,day=numeric,out=02.01.2024
			// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
			// year=2-digit,month=numeric,day=numeric,out=02.01.24
			// year=2-digit,month=numeric,day=2-digit,out=02.1.24
			// year=2-digit,month=2-digit,day=numeric,out=2.01.24
			// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
			layout = layoutDayMonthYear
			separator = "."

			if opts.Year.numeric() || opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionBE:
			// year=numeric,month=numeric,day=numeric,out=02/01/2024
			// year=numeric,month=numeric,day=2-digit,out=02/01/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"

			if opts.Year.numeric() {
				if opts.Month.numeric() && opts.Day.numeric() {
					opts.Day = Day2Digit
				}

				opts.Month = Month2Digit
			}
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
			separator = "/"

			if opts.Month.numeric() && opts.Day.numeric() {
				layout = layoutMonthDayYear
			} else {
				layout = layoutDayMonthYear
			}

			break
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case cldr.NL:
		layout = layoutDayMonthYear

		if region == cldr.RegionBE {
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			separator = "/"
		}
	case cldr.FY, cldr.KOK:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "."
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
	case cldr.HR:
		layout = layoutDayMonthYear
		separator = ". "
		suffix = "."

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
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
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
		separator = ". "
		suffix = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
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
			layout = layoutDayMonthYear
			separator = "."
		}
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
			layout = layoutDayMonthYear

			if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
				opts.Month.numeric() && opts.Day.numeric() {
				separator = "/"
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				separator = "."
			}

			break
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
		layout = layoutDayMonthYear
		separator = "/"

		switch {
		case opts.Year.numeric():
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		case opts.Month.numeric() && opts.Day.numeric():
			opts.Year = Year2Digit
			opts.Month = Month2Digit
			opts.Day = Day2Digit
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
			layout = layoutMonthDayYear
			separator = "."
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		month = convertMonthDigits(digits, opts.Month)
		dayDigits := convertDayDigits(digits, opts.Day)

		return func(t cldr.TimeReader) string {
			return dayDigits(t) + "/" + month(t) + " " + yearDigits(t)
		}
	case cldr.KO:
		// year=numeric,month=numeric,day=numeric,out=2024. 1. 2.
		// year=numeric,month=numeric,day=2-digit,out=2024. 1. 02.
		// year=numeric,month=2-digit,day=numeric,out=2024. 01. 2.
		// year=numeric,month=2-digit,day=2-digit,out=2024. 01. 02.
		// year=2-digit,month=numeric,day=numeric,out=24. 1. 2.
		// year=2-digit,month=numeric,day=2-digit,out=24. 1. 02.
		// year=2-digit,month=2-digit,day=numeric,out=24. 01. 2.
		// year=2-digit,month=2-digit,day=2-digit,out=24. 01. 02.
		separator = ". "
		suffix = "."
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutYearDayMonth
		} else {
			layout = layoutDayMonthYear
			separator = "/"
		}
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
			layout = layoutDayMonthYear
			separator = "/"
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutMonthDayYear
			separator = "/"
		}
	case cldr.LV:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024.
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024.
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24.
		// year=2-digit,month=numeric,day=2-digit,out=02.1.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Year.twoDigit() && opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			suffix = "."
		}
	case cldr.AS, cldr.BRX, cldr.IA, cldr.JV, cldr.MI, cldr.RM, cldr.WO:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=02-01-24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonthYear
		}
	case cldr.MN:
		// year=numeric,month=numeric,day=numeric,out=2024.01.02
		// year=numeric,month=numeric,day=2-digit,out=2024.1.02
		// year=numeric,month=2-digit,day=numeric,out=2024.01.2
		// year=numeric,month=2-digit,day=2-digit,out=2024.01.02
		// year=2-digit,month=numeric,day=numeric,out=24.01.02
		// year=2-digit,month=numeric,day=2-digit,out=24.1.02
		// year=2-digit,month=2-digit,day=numeric,out=24.01.2
		// year=2-digit,month=2-digit,day=2-digit,out=24.01.02
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.MT, cldr.SBP:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		separator = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
		} else {
			layout = layoutDayMonthYear
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutYearDayMonth
			separator = " / "
		}
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
			layout = layoutDayMonthYear
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			separator = "/"
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonthYear
			separator = "."
		}
	case cldr.PL:
		// year=numeric,month=numeric,day=numeric,out=2.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=2.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		opts.Month = Month2Digit
		layout = layoutDayMonthYear
		separator = "."
	case cldr.QU:
		// year=numeric,month=numeric,day=numeric,out=02-01-2024
		// year=numeric,month=numeric,day=2-digit,out=02-01-2024
		// year=numeric,month=2-digit,day=numeric,out=02-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear

		if opts.Year.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			separator = "/"
		}
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
			layout = layoutMonthDayYear
			separator = "/"
		} else if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
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
			layout = layoutDayMonthYear
			separator = "."
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.SO:
		// year=numeric,month=numeric,day=numeric,out=1/2/2024
		// year=numeric,month=numeric,day=2-digit,out=1/02/2024
		// year=numeric,month=2-digit,day=numeric,out=01/2/2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=1/2/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		separator = "/"

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			layout = layoutMonthDayYear
		} else {
			layout = layoutDayMonthYear
		}
	case cldr.SR:
		// year=numeric,month=numeric,day=numeric,out=2. 1. 2024.
		// year=numeric,month=numeric,day=2-digit,out=02. 1. 2024.
		// year=numeric,month=2-digit,day=numeric,out=2. 01. 2024.
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024.
		// year=2-digit,month=numeric,day=numeric,out=2. 1. 24.
		// year=2-digit,month=numeric,day=2-digit,out=02. 1. 24.
		// year=2-digit,month=2-digit,day=numeric,out=2. 01. 24.
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24.
		layout = layoutDayMonthYear
		suffix = "."

		if opts.Month.twoDigit() && opts.Day.twoDigit() {
			separator = "."
		} else {
			separator = ". "
		}
	case cldr.SYR:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Month.numeric() {
			separator = "/"
		}
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
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonthYear
			separator = "."
		}
	case cldr.TE:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2/01/2024
		// year=numeric,month=2-digit,day=2-digit,out=02-01-2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02-1-24
		// year=2-digit,month=2-digit,day=numeric,out=2-01-24
		// year=2-digit,month=2-digit,day=2-digit,out=02-01-24
		layout = layoutDayMonthYear

		if opts.Year.numeric() && (!opts.Month.twoDigit() || !opts.Day.twoDigit()) ||
			opts.Month.numeric() && opts.Day.numeric() {
			separator = "/"
		}
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
			prefix = "#"
			separator = ")#"
		}
	case cldr.TR:
		// year=numeric,month=numeric,day=numeric,out=02.01.2024
		// year=numeric,month=numeric,day=2-digit,out=02.01.2024
		// year=numeric,month=2-digit,day=numeric,out=2.01.2024
		// year=numeric,month=2-digit,day=2-digit,out=02.01.2024
		// year=2-digit,month=numeric,day=numeric,out=02.01.24
		// year=2-digit,month=numeric,day=2-digit,out=02.01.24
		// year=2-digit,month=2-digit,day=numeric,out=2.01.24
		// year=2-digit,month=2-digit,day=2-digit,out=02.01.24
		layout = layoutDayMonthYear
		separator = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Day = Day2Digit
		}

		opts.Month = Month2Digit
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
			layout = layoutYearDayMonth
		}
	case cldr.YI:
		// year=numeric,month=numeric,day=numeric,out=2-1-2024
		// year=numeric,month=numeric,day=2-digit,out=02-1-2024
		// year=numeric,month=2-digit,day=numeric,out=2-01-2024
		// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
		// year=2-digit,month=numeric,day=numeric,out=2-1-24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2/01/24
		// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
		layout = layoutDayMonthYear

		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			separator = "/"
		}
	case cldr.YO:
		// year=numeric,month=numeric,day=numeric,out=2/1/2024
		// year=numeric,month=numeric,day=2-digit,out=02/1/2024
		// year=numeric,month=2-digit,day=numeric,out=2 01 2024
		// year=numeric,month=2-digit,day=2-digit,out=02 01 2024
		// year=2-digit,month=numeric,day=numeric,out=2/1/24
		// year=2-digit,month=numeric,day=2-digit,out=02/1/24
		// year=2-digit,month=2-digit,day=numeric,out=2 01 24
		// year=2-digit,month=2-digit,day=2-digit,out=02 01 24
		layout = layoutDayMonthYear

		if opts.Month.twoDigit() {
			separator = " "
		} else {
			separator = "/"
		}
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
			separator = "/"
		}
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
			separator = "/"
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
				month = convertMonthDigits(digits, opts.Month)
				dayDigits := convertDayDigits(digits, opts.Day)

				return func(t cldr.TimeReader) string {
					return yearDigits(t) + "年" + month(t) + "月" + dayDigits(t) + "日"
				}
			}

			fallthrough
		case cldr.RegionHK:
			// year=numeric,month=numeric,day=numeric,out=2/1/2024
			// year=numeric,month=numeric,day=2-digit,out=02/1/2024
			// year=numeric,month=2-digit,day=numeric,out=2/01/2024
			// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
			// year=2-digit,month=numeric,day=numeric,out=2/1/24
			// year=2-digit,month=numeric,day=2-digit,out=02/1/24
			// year=2-digit,month=2-digit,day=numeric,out=2/01/24
			// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
			layout = layoutDayMonthYear
			separator = "/"
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
		layout = layoutDayMonthYear

		if opts.Year.numeric() && opts.Month.twoDigit() && opts.Day.twoDigit() ||
			opts.Year.twoDigit() && (!opts.Month.numeric() || !opts.Day.numeric()) {
			separator = "/"
		} else {
			separator = "."
		}
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
			layout = layoutMonthDayYear
			separator = "/"
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // layoutYearMonthDay
		return func(t cldr.TimeReader) string {
			return prefix + yearDigits(t) + separator + month(t) + separator + dayDigits(t) + suffix
		}
	case layoutDayMonthYear:
		return func(t cldr.TimeReader) string {
			return dayDigits(t) + separator + month(t) + separator + yearDigits(t) + suffix
		}
	case layoutMonthDayYear:
		return func(t cldr.TimeReader) string {
			return month(t) + separator + dayDigits(t) + separator + yearDigits(t) + suffix
		}
	case layoutYearDayMonth:
		return func(t cldr.TimeReader) string {
			return yearDigits(t) + separator + dayDigits(t) + separator + month(t) + suffix
		}
	}
}

func fmtYearMonthDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()

	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		layoutYearMonthDay = iota
		layoutDayMonthYear
	)

	layout := layoutYearMonthDay

	// "lrc", "mzn", "ps", "uz"
	// year=numeric,month=numeric,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=numeric,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=2-digit,day=numeric,out=AP ۱۴۰۲-۱۰-۱۲
	// year=numeric,month=2-digit,day=2-digit,out=AP ۱۴۰۲-۱۰-۱۲
	// year=2-digit,month=numeric,day=numeric,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=numeric,day=2-digit,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=2-digit,day=numeric,out=AP ۰۲-۱۰-۱۲
	// year=2-digit,month=2-digit,day=2-digit,out=AP ۰۲-۱۰-۱۲
	opts.Month = Month2Digit
	opts.Day = Day2Digit
	prefix := ""
	separator := "-"

	switch lang {
	default:
		prefix = "AP "
	case cldr.CKB: // ckb-IR
		// year=numeric,month=numeric,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=numeric,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=numeric,out=١٢/١٠/١٤٠٢
		// year=numeric,month=2-digit,day=2-digit,out=١٢/١٠/١٤٠٢
		// year=2-digit,month=numeric,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=numeric,day=2-digit,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=numeric,out=١٢/١٠/٠٢
		// year=2-digit,month=2-digit,day=2-digit,out=١٢/١٠/٠٢
		layout = layoutDayMonthYear
		separator = "/"
	case cldr.FA: // fa-IR
		// year=numeric,month=numeric,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=numeric,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=numeric,out=۱۴۰۲/۱۰/۱۲
		// year=numeric,month=2-digit,day=2-digit,out=۱۴۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=numeric,day=2-digit,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=numeric,out=۰۲/۱۰/۱۲
		// year=2-digit,month=2-digit,day=2-digit,out=۰۲/۱۰/۱۲
		separator = "/"
	case cldr.UZ:
		if region != cldr.RegionAF {
			prefix = "AP "
		}
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonthYear {
		return func(v cldr.TimeReader) string {
			return dayDigits(v) + "/" + month(v) + "/" + yearDigits(v)
		}
	}

	return func(v cldr.TimeReader) string {
		return prefix + yearDigits(v) + separator + month(v) + separator + dayDigits(v)
	}
}

func fmtYearMonthDayBuddhist(_ language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	yearDigits := convertYearDigits(digits, opts.Year)
	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	// th-TH
	// year=numeric,month=numeric,day=numeric,out=2/1/2024
	// year=numeric,month=numeric,day=2-digit,out=02/1/2024
	// year=numeric,month=2-digit,day=numeric,out=2/01/2024
	// year=numeric,month=2-digit,day=2-digit,out=02/01/2024
	// year=2-digit,month=numeric,day=numeric,out=2/1/24
	// year=2-digit,month=numeric,day=2-digit,out=02/1/24
	// year=2-digit,month=2-digit,day=numeric,out=2/01/24
	// year=2-digit,month=2-digit,day=2-digit,out=02/01/24
	return func(t cldr.TimeReader) string {
		return dayDigits(t) + "/" + monthDigits(t) + "/" + yearDigits(t)
	}
}
