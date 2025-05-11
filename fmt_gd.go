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

	// f applies the most frequent formatting
	f := func(withoutName ...symbols.Symbol) {
		if withName {
			seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
		} else {
			seq.Add(withoutName...)
		}
	}

	switch lang {
	default:
		f(era, ' ', day)
	case cldr.HI:
		if script != cldr.Latn {
			f(era, ' ', day)
		} else {
			f(day, ' ', era)
		}
	case cldr.KAA, cldr.EN, cldr.MHN:
		f(day, ' ', era)
	case cldr.BS:
		if script == cldr.Cyrl {
			f(era, ' ', day)
			break
		}

		fallthrough
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		if withName {
			seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, '.', ')')
		} else {
			seq.Add(era, ' ', day, '.')
		}
	case cldr.JA, cldr.KO, cldr.YUE, cldr.ZH:
		if withName {
			seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, symbols.DayUnit, ')')
		} else {
			seq.Add(era, ' ', day, symbols.DayUnit)
		}
	case cldr.LT:
		if withName {
			seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', symbols.Symbol_dd, ')')
		} else {
			seq.Add(era, ' ', symbols.Symbol_dd)
		}
	case cldr.II:
		if withName {
			seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, symbols.Txtꑍ, ')')
		} else {
			seq.Add(era, ' ', day, symbols.Txtꑍ)
		}
	}

	return seq
}

func seqEraDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	day := opts.Day.symbol()
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
	} else {
		seq.Add(era, ' ', day)
	}

	return seq
}

func seqEraDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	day := opts.Day.symbol()
	withName := opts.Era.short() || opts.Era.long() && opts.Day.twoDigit()

	if withName {
		seq.Add(era, ' ', '(', symbols.DayUnit, ':', ' ', day, ')')
	} else {
		seq.Add(era, ' ', day)
	}

	return seq
}
