package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqDay(locale language.Tag, opt Day) *symbols.Seq {
	seq := symbols.NewSeq(locale)

	day := symbols.Symbol_d
	if opt.twoDigit() {
		day = symbols.Symbol_dd
	}

	switch lang, _ := locale.Base(); lang {
	default:
		return seq.Add(day)
	case cldr.BS:
		if script, _ := locale.Script(); script == cldr.Cyrl {
			return seq.Add(day)
		}

		return seq.Add(day, '.')
	case cldr.CS, cldr.DA, cldr.DSB, cldr.FO, cldr.HR, cldr.HSB, cldr.IE, cldr.NB, cldr.NN, cldr.NO, cldr.SK, cldr.SL:
		return seq.Add(day, '.')
	case cldr.JA, cldr.YUE, cldr.ZH:
		return seq.Add(day, symbols.Txt日)
	case cldr.KO:
		return seq.Add(day, symbols.Txt일)
	case cldr.LT:
		return seq.Add(symbols.Symbol_dd)
	case cldr.II:
		return seq.Add(day, symbols.Txtꑍ)
	}
}

func fmtDayBuddhist(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	return cldr.Fmt{convertDayDigitsFmt(digits, opt)}.Format
}

func fmtDayPersian(_ language.Tag, digits cldr.Digits, opt Day) fmtFunc {
	return cldr.Fmt{convertDayDigitsFmt(digits, opt)}.Format
}
