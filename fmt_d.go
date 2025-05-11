package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqDay(locale language.Tag, opt Day) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale)

	day := opt.symbol()

	switch lang {
	default:
		seq.Add(day)
	case cldr.BS:
		seq.Add(day)

		if script, _ := locale.Script(); script != cldr.Cyrl {
			seq.Add('.')
		}
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		seq.Add(day, '.')
	case cldr.JA, cldr.YUE, cldr.ZH:
		seq.Add(day, symbols.Txt日)
	case cldr.KO:
		seq.Add(day, symbols.Txt일)
	case cldr.LT:
		seq.Add(symbols.Symbol_dd)
	case cldr.II:
		seq.Add(day, symbols.Txtꑍ)
	}

	return seq
}

func seqDayBuddhist(locale language.Tag, opt Day) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqDayPersian(locale language.Tag, opt Day) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}
