package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func fmtMonthDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()

	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutMonthDay
	middle := "-"
	suffix := ""

	switch lang {
	default:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case cldr.EN:
		switch region {
		default:
			middle = "/"
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
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			middle = "/"

			if script == cldr.Shaw {
				break
			}

			layout = layoutDayMonth

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionAU, cldr.RegionBE, cldr.RegionIE, cldr.RegionNZ, cldr.RegionZW:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case cldr.RegionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-2
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionCH:
			// month=numeric,day=numeric,out=02.01
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			layout = layoutDayMonth
			middle = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		case cldr.RegionZA:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=01/02
			// month=2-digit,day=2-digit,out=02/01
			middle = "/"

			if opts.Month.twoDigit() && opts.Day.twoDigit() {
				layout = layoutDayMonth
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}
		}
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		layout = layoutDayMonth
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case cldr.HI:
		if script == cldr.Latn && opts.Month.numeric() && opts.Day.numeric() {
			// month=numeric,day=numeric,out=02/01
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cldr.AM, cldr.AGQ, cldr.AST, cldr.BAS, cldr.BM, cldr.CA, cldr.CY, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO,
		cldr.EL, cldr.EWO, cldr.FUR, cldr.GD, cldr.GL, cldr.HAW, cldr.ID, cldr.IG, cldr.KAB, cldr.KGP, cldr.KHQ, cldr.KM,
		cldr.KSF, cldr.KXV, cldr.LN, cldr.LO, cldr.LU, cldr.MAI, cldr.MFE, cldr.MG, cldr.MGH, cldr.ML, cldr.MNI, cldr.MUA,
		cldr.MY, cldr.NMG, cldr.NUS, cldr.PA, cldr.RN, cldr.SA, cldr.SEH, cldr.SES, cldr.SG, cldr.SHI, cldr.SU, cldr.SW,
		cldr.TO, cldr.TR, cldr.TWQ, cldr.UR, cldr.XNR, cldr.YAV, cldr.YO, cldr.YRL, cldr.ZGH:
		layout = layoutDayMonth
		middle = "/"
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.UZ, cldr.VEC:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutDayMonth
		middle = "/"
	case cldr.TI:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "/"
	case cldr.KEA, cldr.PT:
		layout = layoutDayMonth
		middle = "/"

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BEZ, cldr.BLO, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.JA, cldr.JMC, cldr.KAA, cldr.KAM, cldr.KDE, cldr.KI, cldr.KLN,
		cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.MHN, cldr.NAQ, cldr.ND, cldr.NYN,
		cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TEO, cldr.TZM, cldr.VAI, cldr.VUN, cldr.XH, cldr.XOG,
		cldr.YUE:
		middle = "/"
	case cldr.KS:
		if script == cldr.Deva {
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=01-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			middle = "/"
		}
	case cldr.AR:
		layout = layoutDayMonth
		middle = "\u200f/"
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		layout = layoutDayMonth
		middle = "."
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		layout = layoutDayMonth
		middle = "."
	case cldr.MK:
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "."
	case cldr.BG, cldr.PL:
		opts.Month = Month2Digit
		layout = layoutDayMonth
		middle = "."
	case cldr.LV:
		opts.Month = Month2Digit
		opts.Day = Day2Digit

		fallthrough
	case cldr.DE, cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.IS, cldr.LB, cldr.SMN:
		layout = layoutDayMonth
		middle = "."
		suffix = "."
	case cldr.NB, cldr.NN, cldr.NO:
		suffix = "."
		fallthrough
	case cldr.SQ:
		opts.Month = MonthNumeric
		opts.Day = DayNumeric
		layout = layoutDayMonth
		middle = "."
	case cldr.RO, cldr.RU:
		layout = layoutDayMonth
		middle = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.SR:
		layout = layoutDayMonth
		suffix = "."

		if opts.Month.numeric() && opts.Day.numeric() {
			middle = ". "
		} else {
			middle = "."
		}
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.TA, cldr.TE, cldr.VI:
		layout = layoutDayMonth

		if opts.Month.numeric() && opts.Day.numeric() {
			middle = "/"
		}
	case cldr.BS:
		layout = layoutDayMonth
		suffix = "."

		if script == cldr.Cyrl {
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.01.
			// month=2-digit,day=numeric,out=02.01.
			// month=2-digit,day=2-digit,out=02.01.
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			middle = "."
		} else {
			// month=numeric,day=numeric,out=2. 1.
			// month=numeric,day=2-digit,out=2. 1.
			// month=2-digit,day=numeric,out=2. 1.
			// month=2-digit,day=2-digit,out=2. 1.
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
			middle = ". "
		}
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		fallthrough
	case cldr.CS, cldr.SK, cldr.SL:
		layout = layoutDayMonth

		fallthrough
	case cldr.HU, cldr.KO:
		middle = ". "
		suffix = "."
	case cldr.WAE:
		month = fmtMonthName(locale.String(), "stand-alone", "abbreviated")
		layout = layoutDayMonth
		middle = ". "
	case cldr.DZ, cldr.SI: // noop
	case cldr.ES:
		switch region {
		default:
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
			layout = layoutDayMonth
			middle = "/"
		case cldr.RegionCL:
			// month=numeric,day=numeric,out=02-01
			// month=numeric,day=2-digit,out=02-01
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			layout = layoutDayMonth

			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				middle = "/"
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			}
		case cldr.RegionMX, cldr.RegionUS:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case cldr.RegionPA, cldr.RegionPR:
			// month=numeric,day=numeric,out=01/02
			// month=numeric,day=2-digit,out=01/02
			// month=2-digit,day=numeric,out=2/1
			// month=2-digit,day=2-digit,out=2/1
			middle = "/"

			if opts.Month.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			} else {
				layout = layoutDayMonth
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			}
		}
	case cldr.FF:
		layout = layoutDayMonth

		if script != cldr.Adlm {
			middle = "/"
		}
	case cldr.FR:
		switch region {
		default:
			opts.Month = Month2Digit
			opts.Day = Day2Digit
			layout = layoutDayMonth
			middle = "/"
		case cldr.RegionCA:
			// month=numeric,day=numeric,out=01-02
			// month=numeric,day=2-digit,out=1-02
			// month=2-digit,day=numeric,out=01-02
			// month=2-digit,day=2-digit,out=01-02
			if opts.Month.numeric() && opts.Day.twoDigit() {
				opts.Month = MonthNumeric
			} else {
				opts.Month = Month2Digit
			}

			opts.Day = Day2Digit
		case cldr.RegionCH:
			// month=numeric,day=numeric,out=02.01.
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.01
			// month=2-digit,day=2-digit,out=02.01
			layout = layoutDayMonth
			middle = "."

			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
				suffix = "."
			}
		}
	case cldr.NL:
		layout = layoutDayMonth

		if region == cldr.RegionBE {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			middle = "/"
		}
	case cldr.FY, cldr.UG:
		layout = layoutDayMonth
	case cldr.IU:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		middle = "/"
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
		}
	case cldr.MN:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
		opts.Day = Day2Digit
		middle = "/"
	case cldr.MS:
		layout = layoutDayMonth

		if !opts.Month.numeric() || !opts.Day.numeric() {
			middle = "/"
		}
	case cldr.OM:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		} else {
			layout = layoutDayMonth
			middle = "/"
		}
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			middle = "/"
		} else {
			layout = layoutDayMonth
		}
	case cldr.PCM:
		layout = layoutDayMonth
		middle = " /"
	case cldr.SD:
		if script == cldr.Deva {
			// month=numeric,day=numeric,out=1/2
			// month=numeric,day=2-digit,out=1/02
			// month=2-digit,day=numeric,out=01/2
			// month=2-digit,day=2-digit,out=01/02
			middle = "/"
		} else {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		} else {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}
	case cldr.SV:
		layout = layoutDayMonth

		if region == cldr.RegionFI {
			// month=numeric,day=numeric,out=2.1
			// month=numeric,day=2-digit,out=02.1
			// month=2-digit,day=numeric,out=2.1
			// month=2-digit,day=2-digit,out=02.01
			middle = "."

			if opts.Day.numeric() {
				opts.Month = MonthNumeric
			}

			break
		}

		middle = "/"

		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
			opts.Day = DayNumeric
		}
	case cldr.ZH:
		switch region {
		default:
			middle = "/"
		case cldr.RegionHK, cldr.RegionMO:
			// month=numeric,day=numeric,out=2/1
			// month=numeric,day=2-digit,out=02/1
			// month=2-digit,day=numeric,out=2/01
			// month=2-digit,day=2-digit,out=02/01
			layout = layoutDayMonth
			middle = "/"
		case cldr.RegionSG: // noop
		}
	case cldr.II:
		// month=numeric,day=numeric,out=01ꆪ-02ꑍ
		// month=numeric,day=2-digit,out=01ꆪ-02ꑍ
		// month=2-digit,day=numeric,out=01ꆪ-02ꑍ
		// month=2-digit,day=2-digit,out=01ꆪ-02ꑍ
		opts.Month = Month2Digit
		opts.Day = Day2Digit
		middle = "ꆪ-"
		suffix = "ꑍ"
	case cldr.KOK:
		layout = layoutDayMonth

		if script == cldr.Latn {
			middle = "/"
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t cldr.TimeReader) string {
			return dayDigits(t) + middle + month(t) + suffix
		}
	}

	return func(t cldr.TimeReader) string {
		return month(t) + middle + dayDigits(t) + suffix
	}
}

func fmtMonthDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	const (
		layoutMonthDay = iota
		layoutDayMonth
	)

	layout := layoutMonthDay

	if lang, _ := locale.Base(); lang == cldr.TH {
		layout = layoutDayMonth
	} else {
		opts.Day = Day2Digit
	}

	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	if layout == layoutDayMonth {
		return func(t cldr.TimeReader) string {
			return dayDigits(t) + "/" + monthDigits(t)
		}
	}

	return func(t cldr.TimeReader) string { return monthDigits(t) + "-" + dayDigits(t) }
}

func fmtMonthDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	middle := "-"

	switch lang, _ := locale.Base(); lang {
	default:
		opts.Month = Month2Digit
		opts.Day = Day2Digit
	case cldr.FA, cldr.PS:
		middle = "/"
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v cldr.TimeReader) string {
		return month(v) + middle + dayDigits(v)
	}
}
