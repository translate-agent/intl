package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	seq := symbols.NewSeq(locale)

	var era symbols.Symbol

	switch opts.Era {
	default:
		era = symbols.Symbol_GGGGG
	case EraShort:
		era = symbols.Symbol_G
	case EraLong:
		era = symbols.Symbol_GGGG
	}

	day := symbols.Symbol_d
	if opts.Day.twoDigit() {
		day = symbols.Symbol_dd
	}

	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	switch lang {
	default:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
		}

		return seq.Add(era, ' ', day)
	case cldr.HI:
		if script != cldr.Latn {
			if withName {
				return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
			}

			return seq.Add(era, ' ', day)
		}

		fallthrough
	case cldr.KAA, cldr.EN, cldr.MHN:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
		}

		return seq.Add(day, ' ', era)
	case cldr.BS:
		if script == cldr.Cyrl {
			if withName {
				return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
			}

			return seq.Add(era, ' ', day)
		}

		fallthrough
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, '.', ')')
		}

		return seq.Add(era, ' ', day, '.')
	case cldr.JA, cldr.KO, cldr.YUE, cldr.ZH:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, symbols.DayUnit, ')')
		}

		return seq.Add(era, ' ', day, symbols.DayUnit)
	case cldr.LT:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', symbols.Symbol_dd, ')')
		}

		return seq.Add(era, ' ', symbols.Symbol_dd)
	case cldr.II:
		if withName {
			return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, symbols.Txtꑍ, ')')
		}

		return seq.Add(era, ' ', day, symbols.Txtꑍ)
	}
}

func fmtEraDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + cldr.UnitName(locale).Day + ": "
		suffix = ")"
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(v cldr.TimeReader) string { return prefix + dayDigits(v) + suffix }
}

func fmtEraDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	prefix, suffix := era+" ", ""
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		prefix, suffix = era+" ("+cldr.UnitName(locale).Day+": ", ")"
	}

	dayDigits := convertDayDigits(digits, opts.Day)

	return func(t cldr.TimeReader) string { return prefix + dayDigits(t) + suffix }
}
