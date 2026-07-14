package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:gocognit,cyclop
func seqDayWeekday(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	weekday := opts.Weekday.symbolFormat()
	day := opts.Day.symbol()

	switch lang {
	default:
		return seq.Add(day, ',', ' ', weekday)
	case cldr.AF, cldr.AM, cldr.AS, cldr.AST, cldr.BAS, cldr.BLO, cldr.BR, cldr.CA, cldr.DJE,
		cldr.DOI, cldr.DUA, cldr.DYO, cldr.EE, cldr.EL, cldr.ES, cldr.FR, cldr.GA, cldr.GL, cldr.GU,
		cldr.HAW, cldr.HI, cldr.IA, cldr.IE, cldr.IT, cldr.JGO, cldr.KKJ, cldr.KSF, cldr.KXV,
		cldr.LN, cldr.LU, cldr.MAI, cldr.MGH, cldr.MUA, cldr.NL, cldr.NMG, cldr.NUS, cldr.RM,
		cldr.RN, cldr.RO, cldr.SA, cldr.SBP, cldr.SC, cldr.SCN, cldr.SU, cldr.SV, cldr.SW, cldr.TH,
		cldr.TO, cldr.TWQ, cldr.VEC, cldr.XNR, cldr.YAV:
		return seq.Add(weekday, ' ', day)
	case cldr.AGQ, cldr.BN, cldr.CCP, cldr.CEB, cldr.CHR, cldr.DZ, cldr.EWO, cldr.FIL, cldr.FUR,
		cldr.KA, cldr.KM, cldr.KN, cldr.MR, cldr.MS, cldr.NE, cldr.OR, cldr.PCM, cldr.SI, cldr.TA,
		cldr.TI, cldr.TK, cldr.TR, cldr.UG, cldr.UR, cldr.YUE, cldr.ZU:
		return seq.Add(day, ' ', weekday)
	case cldr.AR, cldr.SYR:
		return seq.Add(weekday, symbols.TxtArabicComma, ' ', day)
	case cldr.AZ:
		if opts.Weekday == WeekdayNarrow && opts.Day.numeric() {
			return seq.Add(weekday, ' ', day)
		}

		return seq.Add(day, ' ', weekday)
	case cldr.BG, cldr.ET, cldr.GD, cldr.HA, cldr.ID, cldr.JV, cldr.KEA, cldr.KGP, cldr.LO, cldr.MI,
		cldr.MK, cldr.PL, cldr.PT, cldr.RU, cldr.RW, cldr.SQ, cldr.UK, cldr.WO, cldr.YO, cldr.YRL:
		return seq.Add(weekday, ',', ' ', day)
	case cldr.BS, cldr.DE, cldr.DSB, cldr.HR, cldr.HSB, cldr.LB, cldr.LV, cldr.SL:
		return seq.Add(weekday, ',', ' ', day, '.')
	case cldr.CKB:
		return seq.Add(weekday, ' ', day, symbols.TxtKurdishHam)
	case cldr.CS, cldr.FI, cldr.FO, cldr.GSW, cldr.IS, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SMN,
		cldr.SR, cldr.WAE:
		return seq.Add(weekday, ' ', day, '.')
	case cldr.DA:
		return seq.Add(weekday, symbols.TxtDanishDen, day, '.')
	case cldr.EN:
		if script == cldr.Dsrt || script == cldr.Shaw {
			return seq.Add(day, ' ', weekday)
		}

		switch region {
		default:
			return seq.Add(weekday, ' ', day)
		case cldr.RegionAS, cldr.RegionBI, cldr.RegionGU, cldr.RegionJP, cldr.RegionMH,
			cldr.RegionMP, cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionUS, cldr.RegionVI,
			cldr.RegionZZ:
			return seq.Add(day, ' ', weekday)
		}
	case cldr.FA:
		return seq.Add(weekday, ' ', day, symbols.TxtPersianOrdinalSuffix)
	case cldr.FF:
		if script == cldr.Adlm {
			return seq.Add(weekday, ' ', day)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.HE:
		return seq.Add(weekday, symbols.TxtHebrewHeDash, day)
	case cldr.HU:
		return seq.Add(day, '.', ',', ' ', weekday)
	case cldr.II:
		if opts.Weekday == WeekdayNarrow {
			return seq.Add(day, weekday, ',', ' ', symbols.Txtꑍ)
		}

		return seq.Add(day, symbols.Txtꑍ, ',', ' ', weekday)
	case cldr.JA:
		if opts.Weekday == WeekdayLong {
			return seq.Add(day, symbols.Txt日, weekday)
		}

		return seq.Add(day, symbols.Txt日, '(', weekday, ')')
	case cldr.KK:
		if script == cldr.Arab {
			return seq.Add(day, ' ', weekday)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.KO:
		if opts.Weekday == WeekdayLong {
			return seq.Add(day, symbols.Txt일, ' ', weekday)
		}

		return seq.Add(day, symbols.Txt일, ' ', '(', weekday, ')')
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(day, ',', ' ', weekday)
		}

		return seq.Add(day, ' ', weekday)
	case cldr.KSH:
		return seq.Add(weekday, symbols.TxtColognianDa, day, '.')
	case cldr.MN:
		return seq.Add(symbols.Symbol_dd, '.', ' ', weekday)
	case cldr.MY:
		return seq.Add(day, symbols.TxtBurmeseDay, weekday)
	case cldr.SD:
		if script == cldr.Deva {
			return seq.Add(day, ' ', weekday)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.SE:
		if region == cldr.RegionFI {
			return seq.Add(day, ' ', weekday)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.SHN:
		return seq.Add(day, ' ', '-', ' ', weekday)
	case cldr.TOK:
		if opts.Weekday == WeekdayNarrow && opts.Day.numeric() {
			return seq.Add(weekday, ',', ' ', day)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.VAI:
		if script == cldr.Latn {
			return seq.Add(weekday, ' ', day)
		}

		return seq.Add(day, ',', ' ', weekday)
	case cldr.VI:
		return seq.Add(weekday, symbols.TxtVietnameseNgay, day)
	case cldr.YI:
		return seq.Add(weekday, symbols.TxtYiddishDem, day, symbols.TxtYiddishTn)
	case cldr.ZH:
		if script == cldr.Hant {
			return seq.Add(day, ' ', weekday)
		}

		return seq.Add(day, symbols.Txt日, weekday)
	}
}

func seqDayWeekdayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return seqDayWeekday(locale, opts)
}

func seqDayWeekdayPersian(locale language.Tag, opts Options) *symbols.Seq {
	return seqDayWeekday(locale, opts)
}
