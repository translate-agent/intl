package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqMonth(locale language.Tag, opt Month) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	lang, _ := locale.Base()

	switch lang {
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		return seq.Add(symbols.Symbol_MM)
	case cldr.MN:
		return seq.Add(symbols.Symbol_LLLLL)
	case cldr.WAE:
		return seq.Add(symbols.Symbol_LLL)
	}

	month := symbols.Symbol_M
	if opt == Month2Digit {
		month = symbols.Symbol_MM
	}

	seq.Add(month)

	switch lang {
	default:
		return seq
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		return seq.Add('.')
	case cldr.JA, cldr.YUE, cldr.ZH, cldr.KO:
		return seq.Add(symbols.MonthUnit)
	}
}

func seqMonthBuddhist(locale language.Tag, opt Month) *symbols.Seq {
	seq := symbols.NewSeq(locale)

	month := symbols.Symbol_M
	if opt == Month2Digit {
		month = symbols.Symbol_MM
	}

	seq.Add(month)

	return seq
}

func seqMonthPersian(locale language.Tag, opt Month) *symbols.Seq {
	seq := symbols.NewSeq(locale)

	month := symbols.Symbol_M
	if opt == Month2Digit {
		month = symbols.Symbol_MM
	}

	seq.Add(month)

	return seq
}
