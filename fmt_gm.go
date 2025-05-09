package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	month := opts.Month.symbolFormat()
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	switch lang {
	case cldr.EN, cldr.KAA, cldr.MHN:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(month, ' ', era)
	case cldr.BG, cldr.CY, cldr.MK:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(era, ' ', month)
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		month = symbols.Symbol_MM

		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(era, ' ', month)
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, '.', ')')
		}

		return seq.Add(era, ' ', month, '.')
	case cldr.HI:
		if script != cldr.Latn {
			break
		}

		if opts.Era.long() && opts.Month.numeric() || opts.Era.narrow() {
			return seq.Add(month, ' ', era)
		}
	case cldr.MN:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', symbols.Symbol_LLLLL, ')')
		}

		return seq.Add(era, ' ', symbols.Symbol_LLLLL)
	case cldr.WAE:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', symbols.Symbol_MMM, ')')
		}

		return seq.Add(era, ' ', symbols.Symbol_MMM)
	case cldr.JA, cldr.KO, cldr.ZH, cldr.YUE:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, symbols.MonthUnit, ')')
		}

		return seq.Add(era, ' ', month, symbols.MonthUnit)
	}

	if withName {
		return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
	}

	return seq.Add(era, ' ', month)
}

func seqEraMonthPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	month := opts.Month.symbolFormat()
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	if withName {
		return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
	}

	return seq.Add(era, ' ', month)
}

func seqEraMonthBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	month := opts.Month.symbolFormat()
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	if withName {
		return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
	}

	return seq.Add(era, ' ', month)
}
