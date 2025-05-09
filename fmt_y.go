package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYear(locale language.Tag, opt Year) *symbols.Seq {
	seq := symbols.NewSeq(locale)

	year := symbols.Symbol_y
	if opt == Year2Digit {
		year = symbols.Symbol_yy
	}

	seq.Add(year)

	switch lang, _ := locale.Base(); lang {
	default:
		return seq
	case cldr.BG, cldr.MK:
		return seq.Add(' ', symbols.Txt00)
	case cldr.BS, cldr.HR, cldr.HU, cldr.SR:
		return seq.Add('.')
	case cldr.JA, cldr.YUE, cldr.ZH:
		return seq.Add(symbols.Txt年)
	case cldr.KO:
		return seq.Add(symbols.Txt년)
	case cldr.LV:
		return seq.Add(symbols.Txt01)
	}
}

func seqYearBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Era.symbol(), ' ', opts.Year.symbol())
}

func seqYearPersian(locale language.Tag, opt Year) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	lang, _, region := locale.Raw()

	if lang != cldr.FA && (lang != cldr.UZ || region != cldr.RegionAF) {
		seq.Add(EraNarrow.symbol(), ' ')
	}

	return seq.Add(opt.symbol())
}
