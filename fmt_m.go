package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqMonth(locale language.Tag, opt Month) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale)

	switch lang {
	default:
		seq.Add(opt.symbolFormat())

		switch lang {
		case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
			seq.Add('.')
		case cldr.JA, cldr.YUE, cldr.ZH, cldr.KO:
			seq.Add(symbols.MonthUnit)
		}
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		seq.Add(symbols.Symbol_MM)
	case cldr.MN:
		seq.Add(symbols.Symbol_LLLLL)
	case cldr.WAE:
		seq.Add(symbols.Symbol_LLL)
	}

	return seq
}

func seqMonthBuddhist(locale language.Tag, opt Month) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbolFormat())
}

func seqMonthPersian(locale language.Tag, opt Month) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbolFormat())
}
