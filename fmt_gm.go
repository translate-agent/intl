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
			seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		} else {
			seq.Add(month, ' ', era)
		}
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		month = symbols.Symbol_MM
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		if withName {
			seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, '.', ')')
		} else {
			seq.Add(era, ' ', month, '.')
		}
	case cldr.HI:
		if script == cldr.Latn && (opts.Era.long() && opts.Month.numeric() || opts.Era.narrow()) {
			seq.Add(month, ' ', era)
		}
	case cldr.MN:
		month = symbols.Symbol_LLLLL
	case cldr.WAE:
		month = symbols.Symbol_MMM
	case cldr.JA, cldr.KO, cldr.ZH, cldr.YUE:
		if withName {
			seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, symbols.MonthUnit, ')')
		} else {
			seq.Add(era, ' ', month, symbols.MonthUnit)
		}
	}

	if seq.Empty() {
		if withName {
			seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		} else {
			seq.Add(era, ' ', month)
		}
	}

	return seq
}
