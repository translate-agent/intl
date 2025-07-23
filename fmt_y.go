package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYear(locale language.Tag, opt Year) *symbols.Seq {
	lang, _ := locale.Base()
	seq := symbols.NewSeq(locale).Add(opt.symbol())

	switch lang {
	case cldr.BG, cldr.MK:
		seq.Add(' ', symbols.Txt00)
	case cldr.BS, cldr.HR, cldr.HU, cldr.SR:
		seq.Add('.')
	case cldr.JA, cldr.YUE, cldr.ZH:
		seq.Add(symbols.Txt年)
	case cldr.KO:
		seq.Add(symbols.Txt년)
	case cldr.LV:
		seq.Add(symbols.Txt01)
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
