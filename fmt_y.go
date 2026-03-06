package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYear(locale language.Tag, opt Year) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale)
	year := opt.symbol()

	switch lang {
	default:
		seq.Add(year)
	case cldr.BG, cldr.MK:
		seq.Add(year, ' ', symbols.Txt00)
	case cldr.BS, cldr.HR, cldr.HU, cldr.SR:
		seq.Add(year, '.')
	case cldr.JA, cldr.YUE, cldr.ZH:
		seq.Add(year, symbols.Txt年)
	case cldr.KO:
		seq.Add(year, symbols.Txt년)
	case cldr.LV:
		seq.Add(year, symbols.Txt01)
	case cldr.TOK:
		seq.Add('#', year)
	}

	return seq
}

func seqYearBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Era.symbol(), ' ', opts.Year.symbol())
}

func seqYearPersian(locale language.Tag, opt Year) *symbols.Seq {
	lang, _, region := locale.Raw()
	seq := symbols.NewSeq(locale)

	if lang != cldr.FA && (lang != cldr.UZ || region != cldr.RegionAF) {
		seq.Add(symbols.Symbol_GGGGG, ' ')
	}

	return seq.Add(opt.symbol())
}
