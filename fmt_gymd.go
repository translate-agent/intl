package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqEraYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)

	era := symbolEra(opts.Era)
	year := opts.Year.symbol()
	month := Month2Digit.symbolFormat()
	day := Day2Digit.symbol()

	switch lang {
	default:
		return seq.Add(era, ' ', year, '-', month, '-', Day2Digit.symbol())
	case cldr.EN:
		switch region {
		default:
			if script == cldr.Dsrt || script == cldr.Shaw || region == cldr.RegionZZ {
				return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
			}

			return seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year, ' ', era)
		case cldr.RegionAE, cldr.RegionAS, cldr.RegionBI, cldr.RegionCA, cldr.RegionGU, cldr.RegionMH, cldr.RegionMP,
			cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionUS,
			cldr.RegionVI:
			return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
		case cldr.RegionCH:
			return seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', era)
		case cldr.RegionGB:
			if script == cldr.Shaw {
				return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
			}

			return seq.Add(day, '/', month, '/', year, ' ', era)
		}
	case cldr.BRX, cldr.LV, cldr.MNI:
		return seq.Add(era, ' ', day, '-', month, '-', year)
	case cldr.DA, cldr.DSB, cldr.HSB, cldr.IE, cldr.KA, cldr.SQ:
		return seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', era)
	case cldr.MK:
		return seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', symbols.Txt00, ' ', era)
	case cldr.ET, cldr.PL:
		return seq.Add(opts.Day.symbol(), '.', month, '.', year, ' ', era)
	case cldr.BE, cldr.CV, cldr.DE, cldr.FO, cldr.HY, cldr.NB, cldr.NN, cldr.NO, cldr.RO, cldr.RU:
		return seq.Add(symbols.Symbol_dd, '.', month, '.', year, ' ', era)
	case cldr.SR:
		return seq.Add(opts.Day.symbol(), '.', month, '.', year, '.', ' ', era)
	case cldr.BG:
		return seq.Add(day, '.', month, '.', year, ' ', symbols.Txt00, ' ', era)
	case cldr.FI:
		return seq.Add(opts.Month.symbolFormat(), '.', opts.Day.symbol(), '.', year, ' ', era)
	case cldr.FR:
		if region == cldr.RegionCA {
			return seq.Add(year, '-', month, '-', day, ' ', era)
		}

		return seq.Add(day, '/', month, '/', year, ' ', era)
	case cldr.AM, cldr.AS, cldr.ES, cldr.GD, cldr.GL, cldr.HE, cldr.EL, cldr.ID, cldr.IS, cldr.JV, cldr.NL, cldr.SU,
		cldr.SW, cldr.TA, cldr.XNR, cldr.UR, cldr.VI, cldr.YO:
		return seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year, ' ', era)
	case cldr.GA, cldr.IT, cldr.KEA, cldr.PT, cldr.SC, cldr.SYR, cldr.VEC:
		return seq.Add(day, '/', month, '/', year, ' ', era)
	case cldr.CEB, cldr.CHR, cldr.CY, cldr.BLO, cldr.FIL, cldr.KAA, cldr.MHN, cldr.ML, cldr.NE, cldr.OR, cldr.PS, cldr.SD,
		cldr.SO, cldr.TI, cldr.XH, cldr.ZU:
		return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
	case cldr.AR, cldr.IA, cldr.BN, cldr.CA, cldr.MAI, cldr.RM, cldr.UK, cldr.WO:
		return seq.Add(day, '-', month, '-', year, ' ', era)
	case cldr.LT, cldr.SV:
		return seq.Add(year, '-', month, '-', day, ' ', era)
	case cldr.BS:
		if script == cldr.Cyrl {
			return seq.Add(era, ' ', year, '-', month, '-', day)
		}

		return seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, '.', ' ', era)
	case cldr.FF:
		if script == cldr.Adlm {
			return seq.Add(opts.Day.symbol(), '-', opts.Month.symbolFormat(), '-', year, ' ', era)
		}

		return seq.Add(era, ' ', year, '-', month, '-', day)
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(era, ' ', year, '-', month, '-', day)
		}

		return seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
	case cldr.UZ:
		if script != cldr.Cyrl {
			return seq.Add(day, '.', month, '.', year, ' ', era)
		}

		return seq.Add(era, ' ', year, '-', month, '-', day)
	case cldr.AZ:
		if script == cldr.Cyrl {
			return seq.Add(era, ' ', year, '-', month, '-', day)
		}

		return seq.Add(era, ' ', opts.Day.symbol(), ' ', symbols.Symbol_MMM, ' ', year)
	case cldr.KU, cldr.TK, cldr.TR:
		return seq.Add(era, ' ', day, '.', month, '.', year)
	case cldr.HU:
		return seq.Add(era, ' ', year, '.', ' ', month, '.', ' ', day, '.')
	case cldr.CS, cldr.SK, cldr.SL:
		return seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, ' ', era)
	case cldr.HR:
		return seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, '.', ' ', era)
	case cldr.HI:
		if script == cldr.Latn {
			return seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year, ' ', era)
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	case cldr.ZH:
		if script == cldr.Hant {
			return seq.Add(era, ' ', year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
		}

		return seq.Add(era, ' ', year, '-', month, '-', day)
	case cldr.KXV:
		if script == cldr.Deva || script == cldr.Orya || script == cldr.Telu {
			return seq.Add(era, ' ', year, '-', month, '-', day)
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	case cldr.JA:
		return seq.Add(era, year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
	case cldr.KO, cldr.MY:
		return seq.Add(era, ' ', year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
	case cldr.MR, cldr.QU:
		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	case cldr.TO:
		return seq.Add(day, ' ', month, ' ', year, ' ', era)
	case cldr.KK:
		return seq.Add(Day2Digit.symbol(), '-', month, '-', era, ' ', year)
	case cldr.LO:
		return seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', era, ' ', year)
	case cldr.PA:
		if script == cldr.Arab {
			return seq.Add(era, ' ', year, '-', month, '-', day)
		}

		return seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', era, ' ', year)
	case cldr.KOK:
		if script == cldr.Latn {
			return seq.Add(opts.Day.symbol(), '-', opts.Month.symbolFormat(), '-', year, ' ', era)
		}

		return seq.Add(era, ' ', year, '-', month, '-', day)
	}
}

func fmtEraYearMonthDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, _, region := locale.Raw()

	era := fmtEra(locale, opts.Era)
	yearDigits := convertYearDigits(digits, opts.Year)

	const (
		eraYearMonthDay = iota
		eraMonthDayYear
	)

	layout := eraMonthDayYear
	separator := "/"
	suffix := " " + era

	switch lang {
	case cldr.CKB:
		if region == cldr.RegionIR {
			layout = eraYearMonthDay
			separator = "-"
		}
	case cldr.LRC, cldr.MZN, cldr.UZ:
		layout = eraYearMonthDay
		separator = "-"
	case cldr.PS:
		layout = eraYearMonthDay

		if !opts.Era.narrow() {
			separator = "-"
		}
	}

	month := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	switch layout {
	default: // eraMonthDayYear
		return func(v cldr.TimeReader) string {
			return month(v) + separator + dayDigits(v) + separator + yearDigits(v) + suffix
		}
	case eraYearMonthDay:
		prefix := era + " "

		return func(v cldr.TimeReader) string {
			return prefix + yearDigits(v) + separator + month(v) + separator + dayDigits(v)
		}
	}
}

func fmtEraYearMonthDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)
	monthDigits := convertMonthDigits(digits, opts.Month)
	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t cldr.TimeReader) string {
		return dayDigits(t) + "/" + monthDigits(t) + "/" + year(t)
	}
}
