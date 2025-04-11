package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqYearMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)

	year := symbols.Symbol_y
	if opts.Year == Year2Digit {
		year = symbols.Symbol_yy
	}

	month := symbols.Symbol_M
	if opts.Month == Month2Digit {
		month = symbols.Symbol_MM
	}

	switch lang {
	default:
		return seq.Add(year, '-', symbols.Symbol_MM)
	case cldr.AF, cldr.AS, cldr.IA, cldr.JV, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		return seq.Add(symbols.Symbol_MM, '-', year)
	case cldr.EN:
		switch region {
		default:
			return seq.Add(month, '/', year)
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
			if script != cldr.Shaw {
				return seq.Add(symbols.Symbol_MM, '/', year)
			}

			return seq.Add(month, '/', year)
		case cldr.RegionCA, cldr.RegionSE:
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return seq.Add(year, '-', symbols.Symbol_MM)
		case cldr.RegionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return seq.Add(symbols.Symbol_MM, '.', year)
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
		return seq.Add(month, '/', year)
	case cldr.PA:
		if script == cldr.Arab {
			// year=numeric,month=numeric,out=۲۰۲۴-۰۱
			// year=numeric,month=2-digit,out=۲۰۲۴-۰۱
			// year=2-digit,month=numeric,out=۲۴-۰۱
			// year=2-digit,month=2-digit,out=۲۴-۰۱
			return seq.Add(year, '-', symbols.Symbol_MM)
		}

		return seq.Add(month, '/', year)
	case cldr.KS:
		if script == cldr.Deva {
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return seq.Add(year, '-', symbols.Symbol_MM)
		}

		return seq.Add(month, '/', year)
	case cldr.HI:
		if script == cldr.Latn {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			return seq.Add(symbols.Symbol_MM, '/', year)
		}

		return seq.Add(month, '/', year)
	case cldr.AR:
		return seq.Add(month, symbols.Txt02, year)
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.PL, cldr.RO, cldr.RU, cldr.TK, cldr.TT,
		cldr.UK:
		return seq.Add(symbols.Symbol_MM, '.', year)
	case cldr.UZ:
		if script == cldr.Cyrl {
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=01/24
			return seq.Add(symbols.Symbol_MM, '/', year)
		}

		return seq.Add(symbols.Symbol_MM, '.', year)
	case cldr.BE, cldr.DA, cldr.DSB, cldr.ET, cldr.HSB, cldr.IE, cldr.KA, cldr.LB, cldr.NB, cldr.NN, cldr.NO, cldr.SMN,
		cldr.SQ:
		return seq.Add(month, '.', year)
	case cldr.BG:
		return seq.Add(symbols.Symbol_MM, '.', year, symbols.Txt00)
	case cldr.MK:
		return seq.Add(month, '.', year, symbols.Txt00)
	case cldr.BN, cldr.CCP, cldr.GU, cldr.KN, cldr.MR, cldr.OR, cldr.TA, cldr.TE, cldr.TO:
		if opts.Month.numeric() {
			return seq.Add(month, '/', year)
		}

		return seq.Add(month, '-', year)
	case cldr.BR, cldr.GA, cldr.IT, cldr.IU, cldr.KEA, cldr.KGP, cldr.PT, cldr.SC, cldr.SEH, cldr.SYR, cldr.VEC, cldr.YRL:
		return seq.Add(symbols.Symbol_MM, '/', year)
	case cldr.BS:
		if script == cldr.Cyrl {
			return seq.Add(symbols.Symbol_MM, '.', year, '.')
		}

		if !opts.Month.numeric() {
			return seq.Add(symbols.Symbol_M, '.', ' ', year, '.')
		}

		return seq.Add(symbols.Symbol_MM, '/', year)
	case cldr.DE:
		if opts.Month.numeric() {
			return seq.Add(month, '/', year)
		}

		return seq.Add(month, '.', year)
	case cldr.DZ, cldr.SI:
		return seq.Add(year, '-', month)
	case cldr.ES:
		switch region {
		default:
			return seq.Add(symbols.Symbol_M, '/', year)
		case cldr.RegionAR:
			// year=numeric,month=numeric,out=1-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=1-24
			// year=2-digit,month=2-digit,out=1/24
			if !opts.Month.numeric() {
				return seq.Add(symbols.Symbol_M, '/', year)
			}

			return seq.Add(symbols.Symbol_M, '-', year)
		case cldr.RegionCL:
			// year=numeric,month=numeric,out=01-2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01-24
			// year=2-digit,month=2-digit,out=1/24
			if opts.Month.numeric() {
				return seq.Add(symbols.Symbol_MM, '-', year)
			}

			return seq.Add(symbols.Symbol_M, '/', year)
		case cldr.RegionMX, cldr.RegionUS:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return seq.Add(month, '/', year)
		case cldr.RegionPA, cldr.RegionPR:
			// year=numeric,month=numeric,out=01/2024
			// year=numeric,month=2-digit,out=1/2024
			// year=2-digit,month=numeric,out=01/24
			// year=2-digit,month=2-digit,out=1/24
			if opts.Month.numeric() {
				return seq.Add(symbols.Symbol_MM, '/', year)
			}

			return seq.Add(symbols.Symbol_M, '/', year)
		}
	case cldr.TI:
		return seq.Add(symbols.Symbol_M, '/', year)
	case cldr.YUE:
		if script == cldr.Hans {
			// year=numeric,month=numeric,out=2024年1月
			// year=numeric,month=2-digit,out=2024年1月
			// year=2-digit,month=numeric,out=24年1月
			// year=2-digit,month=2-digit,out=24年1月
			return seq.Add(year, symbols.Txt年, symbols.Symbol_M, symbols.Txt月)
		}

		return seq.Add(year, '/', month)
	case cldr.EU, cldr.JA:
		return seq.Add(year, '/', month)
	case cldr.FI, cldr.HE:
		return seq.Add(symbols.Symbol_M, '.', year)
	case cldr.FF:
		if script != cldr.Adlm {
			return seq.Add(month, '/', year)
		}

		return seq.Add(month, '-', year)
	case cldr.FR:
		switch region {
		default:
			return seq.Add(symbols.Symbol_MM, '/', year)
		case cldr.RegionCA: // noop
			// year=numeric,month=numeric,out=2024-01
			// year=numeric,month=2-digit,out=2024-01
			// year=2-digit,month=numeric,out=24-01
			// year=2-digit,month=2-digit,out=24-01
			return seq.Add(year, '-', symbols.Symbol_MM)
		case cldr.RegionCH:
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return seq.Add(symbols.Symbol_MM, '.', year)
		}
	case cldr.NL:
		if region == cldr.RegionBE {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return seq.Add(month, '/', year)
		}

		return seq.Add(month, '-', year)
	case cldr.FY, cldr.KOK, cldr.MS, cldr.UG:
		return seq.Add(month, '-', year)
	case cldr.GSW:
		if !opts.Month.numeric() {
			return seq.Add(month, '.', year)
		}

		return seq.Add(year, '-', month)
	case cldr.HR:
		return seq.Add(symbols.Symbol_MM, '.', ' ', year, '.')
	case cldr.HU:
		return seq.Add(year, '.', ' ', month, '.')
	case cldr.IS:
		return seq.Add(month, '.', ' ', year)
	case cldr.KKJ:
		return seq.Add(symbols.Symbol_MM, ' ', year)
	case cldr.KO:
		return seq.Add(year, '.', ' ', symbols.Symbol_M, '.')
	case cldr.LV:
		return seq.Add(symbols.Symbol_MM, '.', year, '.')
	case cldr.MN:
		return seq.Add(year, ' ', symbols.Symbol_LLLLL)
	case cldr.YI:
		if !opts.Month.numeric() {
			return seq.Add(month, '/', year)
		}

		return seq.Add(year, '-', symbols.Symbol_MM)
	case cldr.SD:
		if script == cldr.Deva {
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			return seq.Add(month, '/', year)
		}

		return seq.Add(year, '-', symbols.Symbol_MM)
	case cldr.SE:
		if region == cldr.RegionFI {
			// year=numeric,month=numeric,out=01.2024
			// year=numeric,month=2-digit,out=01.2024
			// year=2-digit,month=numeric,out=01.24
			// year=2-digit,month=2-digit,out=01.24
			return seq.Add(symbols.Symbol_MM, '.', year)
		}

		return seq.Add(year, '-', symbols.Symbol_MM)
	case cldr.SR:
		if opts.Month.numeric() {
			return seq.Add(month, '.', ' ', year, '.')
		}

		return seq.Add(month, '.', year, '.')
	case cldr.TR:
		if opts.Month.numeric() {
			return seq.Add(symbols.Symbol_MM, '/', year)
		}

		return seq.Add(symbols.Symbol_MM, '.', year)
	case cldr.VI:
		if opts.Month.numeric() {
			return seq.Add(month, '/', year)
		}

		return seq.Add(symbols.Txt04, month, ',', ' ', year)
	case cldr.ZH:
		switch script {
		case cldr.Hant:
			switch region {
			default:
				// year=numeric,month=numeric,out=2024/1
				// year=numeric,month=2-digit,out=2024/01
				// year=2-digit,month=numeric,out=24/1
				// year=2-digit,month=2-digit,out=24/01
				return seq.Add(year, '/', month)
			case cldr.RegionHK, cldr.RegionMO:
				// year=numeric,month=numeric,out=1/2024
				// year=numeric,month=2-digit,out=01/2024
				// year=2-digit,month=numeric,out=1/24
				// year=2-digit,month=2-digit,out=01/24
				return seq.Add(month, '/', year)
			}
		case cldr.Hans:
			// year=numeric,month=numeric,out=1/2024
			// year=numeric,month=2-digit,out=01/2024
			// year=2-digit,month=numeric,out=1/24
			// year=2-digit,month=2-digit,out=01/24
			if region == cldr.RegionHK {
				return seq.Add(month, '/', year)
			}

			fallthrough
		default:
			if !opts.Month.numeric() {
				return seq.Add(year, symbols.Txt年, symbols.Symbol_M, symbols.Txt月)
			}

			return seq.Add(year, '/', month)
		}
	}
}

func fmtYearMonthBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	yearDigits := convertYearDigits(digits, opts.Year)

	if lang, _ := locale.Base(); lang == cldr.TH {
		monthDigits := convertMonthDigits(digits, opts.Month)

		return func(t cldr.TimeReader) string {
			return monthDigits(t) + "/" + yearDigits(t)
		}
	}

	monthDigits := convertMonthDigits(digits, Month2Digit)

	return func(t cldr.TimeReader) string {
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
		return func(v cldr.TimeReader) string {
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

	return func(v cldr.TimeReader) string {
		return prefix + yearDigits(v) + separator + month(v)
	}
}
