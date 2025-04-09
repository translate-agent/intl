package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func fmtYearMonthGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	var month fmtFunc

	lang, script, region := locale.Raw()
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		layoutYearMonth = iota
		layoutMonthYear
	)

	layout := layoutYearMonth
	prefix := ""
	middle := "-"
	suffix := ""

	switch lang {
	default:
		opts.Month = Month2Digit
	case cldr.AF, cldr.AS, cldr.IA, cldr.JV, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		opts.Month = Month2Digit
		layout = layoutMonthYear
	case cldr.EN:
		switch region {
		default:
			layout = layoutMonthYear
			middle = "/"
		case cldr.Region001, cldr.Region150, cldr.RegionAE, cldr.RegionAG, cldr.RegionAI, cldr.RegionAT, cldr.RegionAU,
			cldr.RegionBB, cldr.RegionBE, cldr.RegionBM, cldr.RegionBS, cldr.RegionBW, cldr.RegionBZ, cldr.RegionCC,
			cldr.RegionCK, cldr.RegionCM, cldr.RegionCX, cldr.RegionCY, cldr.RegionDE, cldr.RegionDG, cldr.RegionDK,
			cldr.RegionDM, cldr.RegionER, cldr.RegionFI, cldr.RegionFJ, cldr.RegionFK, cldr.RegionFM, cldr.RegionGB,
			cldr.RegionGD, cldr.RegionGG, cldr.RegionGH, cldr.RegionGI, cldr.RegionGM, cldr.RegionGY, cldr.RegionHK,
			cldr.RegionID, cldr.RegionIE, cldr.RegionIL, cldr.RegionIM, cldr.RegionIN, cldr.RegionIO, cldr.RegionJE,
			cldr.RegionJM, cldr.RegionKE, cldr.RegionKI, cldr.RegionKN, cldr.RegionKY, cldr.RegionLC, cldr.RegionLR,
			cldr.RegionLS, cldr.RegionMG, cldr.RegionMO, cldr.RegionMS, cldr.RegionMT, cldr.RegionMU, cldr.RegionMV,
			cldr.RegionMW, cldr.RegionMY, cldr.RegionNA, cldr.RegionNF, cldr.RegionNG, cldr.RegionNL, cldr.RegionNR,
			cldr.RegionNU, cldr.RegionNZ, cldr.RegionPG, cldr.RegionPK, cldr.RegionPN, cldr.RegionPW, cldr.RegionRW,
			cldr.RegionSB, cldr.RegionSC, cldr.RegionSD, cldr.RegionSG, cldr.RegionSH, cldr.RegionSI, cldr.RegionSL,
			cldr.RegionSS, cldr.RegionSX, cldr.RegionSZ, cldr.RegionTC, cldr.RegionTK, cldr.RegionTO, cldr.RegionTT,
			cldr.RegionTV, cldr.RegionTZ, cldr.RegionUG, cldr.RegionVC, cldr.RegionVG, cldr.RegionVU, cldr.RegionWS,
			cldr.RegionZA, cldr.RegionZM, cldr.RegionZW:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"

			if script != cldr.Shaw {
				opts.Month = Month2Digit
			}
		case cldr.RegionCA, cldr.RegionSE:
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			opts.Month = Month2Digit
		case cldr.RegionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			opts.Month = Month2Digit
			layout = layoutMonthYear
			middle = "."
		}
	case cldr.AGQ, cldr.AK, cldr.AM, cldr.ASA, cldr.AST, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BLO, cldr.BM, cldr.BRX,
		cldr.CA, cldr.CEB, cldr.CGG, cldr.CHR, cldr.CKB, cldr.CS, cldr.CY, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA,
		cldr.DYO, cldr.EBU, cldr.EE, cldr.EL, cldr.EWO, cldr.FIL, cldr.FUR, cldr.GD, cldr.GL, cldr.GUZ, cldr.HA, cldr.HAW,
		cldr.ID, cldr.IG, cldr.JMC, cldr.KAA, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KLN, cldr.KM, cldr.KSB,
		cldr.KSF, cldr.KXV, cldr.LAG, cldr.LG, cldr.LN, cldr.LO, cldr.LU, cldr.LUO, cldr.LUY, cldr.MAI, cldr.MAS, cldr.MER,
		cldr.MFE, cldr.MG, cldr.MGH, cldr.MHN, cldr.MNI, cldr.MUA, cldr.NAQ, cldr.ND, cldr.NMG, cldr.NUS, cldr.NYN,
		cldr.OM, cldr.PCM, cldr.RN, cldr.ROF, cldr.RWK, cldr.SA, cldr.SAQ, cldr.SBP, cldr.SES, cldr.SG, cldr.SHI, cldr.SK,
		cldr.SL, cldr.SO, cldr.SU, cldr.SW, cldr.TEO, cldr.TWQ, cldr.TZM, cldr.UR, cldr.VAI, cldr.VUN, cldr.XH, cldr.XNR,
		cldr.XOG, cldr.YAV, cldr.YO, cldr.ZGH:
		layout = layoutMonthYear
		middle = "/"
	case cldr.PA:
		if script == cldr.Arab {
			// year=numeric,month=numeric,out=۲۰۲۴-۰۱
			// year=numeric,month=2-digit,out=۲۰۲۴-۰۱
			// year=2-digit,month=numeric,out=۲۴-۰۱
			// year=2-digit,month=2-digit,out=۲۴-۰۱
			opts.Month = Month2Digit
		} else {
			layout = layoutMonthYear
			middle = "/"
		}
	case cldr.KS:
		if script == cldr.Deva {
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			opts.Month = Month2Digit
		} else {
			layout = layoutMonthYear
			middle = "/"
		}
	case cldr.HI:
		layout = layoutMonthYear
		middle = "/"

		if script == cldr.Latn {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			opts.Month = Month2Digit
		}
	case cldr.AR:
		layout = layoutMonthYear
		middle = "\u200f/"
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.PL, cldr.RO, cldr.RU, cldr.TK, cldr.TT,
		cldr.UK:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
	case cldr.UZ:
		opts.Month = Month2Digit
		layout = layoutMonthYear

		if script == cldr.Cyrl {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			middle = "/"
		} else {
			middle = "."
		}
	case cldr.BE, cldr.DA, cldr.DSB, cldr.ET, cldr.HSB, cldr.IE, cldr.KA, cldr.LB, cldr.NB, cldr.NN, cldr.NO, cldr.SMN,
		cldr.SQ:
		layout = layoutMonthYear
		middle = "."
	case cldr.BG:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
		suffix = " г."
	case cldr.MK:
		layout = layoutMonthYear
		middle = "."
		suffix = " г."
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.OR, cldr.TA, cldr.TE, cldr.TO:
		layout = layoutMonthYear

		if opts.Month.numeric() {
			middle = "/"
		}
	case cldr.BR, cldr.GA, cldr.IT, cldr.IU, cldr.KEA, cldr.KGP, cldr.PT, cldr.SC, cldr.SEH, cldr.SYR, cldr.VEC, cldr.YRL:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "/"
	case cldr.BS:
		layout = layoutMonthYear

		if script == cldr.Cyrl {
			opts.Month = Month2Digit
			middle = "."
			suffix = "."

			break
		}

		middle = "/"

		if !opts.Month.numeric() {
			middle = ". "
			suffix = "."
		}

		if opts.Month.numeric() {
			opts.Month = Month2Digit
		} else {
			opts.Month = MonthNumeric
		}
	case cldr.DE:
		layout = layoutMonthYear
		middle = "."

		if opts.Month.numeric() {
			middle = "/"
		}
	case cldr.DZ, cldr.SI: // noop
	case cldr.ES:
		switch region {
		default:
			opts.Month = MonthNumeric
			layout = layoutMonthYear
			middle = "/"
		case cldr.RegionAR:
			// year=numeric,month=numeric,out=1-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=1-24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear

			if !opts.Month.numeric() {
				middle = "/"
			}

			opts.Month = MonthNumeric
		case cldr.RegionCL:
			// year=numeric,month=numeric,out=01-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01-24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear

			if opts.Month.numeric() {
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
				middle = "/"
			}
		case cldr.RegionMX, cldr.RegionUS:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"
		case cldr.RegionPA, cldr.RegionPR:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=1/24
			layout = layoutMonthYear
			middle = "/"

			if opts.Month.numeric() {
				opts.Month = Month2Digit
			} else {
				opts.Month = MonthNumeric
			}
		}
	case cldr.TI:
		opts.Month = MonthNumeric
		layout = layoutMonthYear
		middle = "/"
	case cldr.YUE:
		if script == cldr.Hans {
			// year=numeric,month=numeric,out=2024年1月
			// year=numeric,month=2-digit,out=2024年1月
			// year=2-digit,month=numeric,out=24年1月
			// year=2-digit,month=2-digit,out=24年1月
			opts.Month = MonthNumeric
			middle = "年"
			suffix = "月"
		} else {
			middle = "/"
		}
	case cldr.EU, cldr.JA:
		middle = "/"
	case cldr.FI, cldr.HE:
		opts.Month = MonthNumeric
		layout = layoutMonthYear
		middle = "."
	case cldr.FF:
		layout = layoutMonthYear

		if script != cldr.Adlm {
			middle = "/"
		}
	case cldr.FR:
		opts.Month = Month2Digit

		switch region {
		default:
			layout = layoutMonthYear
			middle = "/"
		case cldr.RegionCA: // noop
		// year=numeric,month=numeric,out=2024-01
		// year=numeric,month=2-digit,out=2024-01
		// year=2-digit,month=numeric,out=24-01
		// year=2-digit,month=2-digit,out=24-01
		case cldr.RegionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			layout = layoutMonthYear
			middle = "."
		}
	case cldr.NL:
		layout = layoutMonthYear

		if region == cldr.RegionBE {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			middle = "/"
		}
	case cldr.FY, cldr.KOK, cldr.MS, cldr.UG:
		layout = layoutMonthYear
	case cldr.GSW:
		if !opts.Month.numeric() {
			layout = layoutMonthYear
			middle = "."
		}
	case cldr.HR:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = ". "
		suffix = "."
	case cldr.HU:
		middle = ". "
		suffix = "."
	case cldr.IS:
		layout = layoutMonthYear
		middle = ". "
	case cldr.KKJ:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = " "
	case cldr.KO:
		opts.Month = MonthNumeric
		middle = ". "
		suffix = "."
	case cldr.LV:
		opts.Month = Month2Digit
		layout = layoutMonthYear
		middle = "."
		suffix = "."
	case cldr.MN:
		month = fmtMonthName(locale.String(), "stand-alone", "narrow")
		middle = " "
	case cldr.YI:
		if !opts.Month.numeric() {
			layout = layoutMonthYear
			middle = "/"
		}

		opts.Month = Month2Digit
	case cldr.SD:
		if script == cldr.Deva {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			layout = layoutMonthYear
			middle = "/"
		} else {
			opts.Month = Month2Digit
		}
	case cldr.SE:
		opts.Month = Month2Digit

		if region == cldr.RegionFI {
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			layout = layoutMonthYear
			middle = "."
		}
	case cldr.SR:
		layout = layoutMonthYear
		suffix = "."

		if opts.Month.numeric() {
			middle = ". "
		} else {
			middle = "."
		}
	case cldr.TR:
		layout = layoutMonthYear
		middle = "."

		if opts.Month.numeric() {
			middle = "/"
		}

		opts.Month = Month2Digit
	case cldr.VI:
		layout = layoutMonthYear

		if opts.Month.numeric() {
			middle = "/"
		} else {
			prefix = "tháng "
			middle = ", "
		}
	case cldr.ZH:
		middle = "/"

		switch script {
		case cldr.Hant:
			switch region {
			default:
			// year=numeric,month=numeric,out=2024/1
			// year=numeric,month=2-digit,out=2024/01
			// year=2-digit,month=numeric,out=24/1
			// year=2-digit,month=2-digit,out=24/01
			case cldr.RegionHK, cldr.RegionMO:
				// year=numeric,month=numeric,out=1/2024
				// year=numeric,month=2-digit,out=01/2024
				// year=2-digit,month=numeric,out=1/24
				// year=2-digit,month=2-digit,out=01/24
				layout = layoutMonthYear
			}
		case cldr.Hans:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			if region == cldr.RegionHK {
				layout = layoutMonthYear
				break
			}

			fallthrough
		default:
			if !opts.Month.numeric() {
				opts.Month = MonthNumeric
				middle = "年"
				suffix = "月"
			}
		}
	}

	if month == nil {
		month = convertMonthDigits(digits, opts.Month)
	}

	if layout == layoutMonthYear {
		return func(t timeReader) string {
			return prefix + month(t) + middle + yearDigits(t) + suffix
		}
	}

	return func(t timeReader) string {
		return prefix + yearDigits(t) + middle + month(t) + suffix
	}
}

func fmtYearMonthBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	yearDigits := convertYearDigits(digits, opts.Year)

	if lang, _ := locale.Base(); lang == cldr.TH {
		monthDigits := convertMonthDigits(digits, opts.Month)

		return func(t timeReader) string {
			return monthDigits(t) + "/" + yearDigits(t)
		}
	}

	monthDigits := convertMonthDigits(digits, Month2Digit)

	return func(t timeReader) string {
		return yearDigits(t) + "-" + monthDigits(t)
	}
}

func fmtYearMonthPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()
	yearDigits := convertYearDigits(digits, opts.Year)
	month := convertMonthDigits(digits, Month2Digit)

	prefix := ""
	separator := "-"

	switch lang {
	case cldr.CKB: // ckb-IR
		// year=numeric,month=numeric,out=١٠/١٤٠٢
		// year=numeric,month=2-digit,out=١٠/١٤٠٢
		// year=2-digit,month=numeric,out=١٠/٠٢
		// year=2-digit,month=2-digit,out=١٠/٠٢
		return func(v timeReader) string {
			return month(v) + "/" + yearDigits(v)
		}
	case cldr.FA:
		separator = "/"
	case cldr.PS:
		prefix = fmtEra(locale, EraNarrow) + " "
		separator = "/"
	case cldr.UZ:
		if region == cldr.RegionAF {
			// year=numeric,month=numeric,out=۱۴۰۲-۱۰
			// year=numeric,month=2-digit,out=۱۴۰۲-۱۰
			// year=2-digit,month=numeric,out=۰۲-۱۰
			// year=2-digit,month=2-digit,out=۰۲-۱۰
			break
		}

		fallthrough
	default:
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(v timeReader) string {
		return prefix + yearDigits(v) + separator + month(v)
	}
}
