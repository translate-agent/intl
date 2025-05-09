package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	day := opts.Day.symbol()
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

func seqEraDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	day := opts.Day.symbol()
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
	}

	return seq.Add(era, ' ', day)
}

func seqEraDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	day := opts.Day.symbol()
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		return seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
	}

	return seq.Add(era, ' ', day)
}
