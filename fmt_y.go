package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYear(locale language.Tag, opts Options) *symbols.Seq {
	lang, _ := locale.Base()
	region, regionConfidence := locale.Region()
	seq := symbols.NewSeq(locale)
	year := opts.Year.symbol()

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
	case cldr.TH, cldr.SHN:
		if region == cldr.RegionTH {
			seq.Add(opts.Era.symbol(), ' ', year)
		} else {
			seq.Add(year)
		}
	case cldr.LRC, cldr.MZN, cldr.PS, cldr.CKB:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(symbols.Symbol_GGGGG, ' ', year)
		} else {
			seq.Add(year)
		}
	case cldr.UZ:
		if region == cldr.RegionAF && regionConfidence != language.Exact {
			seq.Add(symbols.Symbol_GGGGG, ' ', year)
		} else {
			seq.Add(year)
		}
	}

	return seq
}
