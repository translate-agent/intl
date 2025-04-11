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
		return seq.Add(symbols.Txt00)
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

func fmtYearGregorian(locale language.Tag, digits cldr.Digits, opt Year) fmtFunc {
	fmt := cldr.Fmt{convertYearDigitsFmt(digits, opt)}

	switch lang, _ := locale.Base(); lang {
	default:
		return fmt.Format
	case cldr.BG, cldr.MK:
		return append(fmt, cldr.Text(" г.")).Format
	case cldr.BS, cldr.HR, cldr.HU, cldr.SR:
		return append(fmt, cldr.Text(".")).Format
	case cldr.JA, cldr.YUE, cldr.ZH:
		return append(fmt, cldr.Text("年")).Format
	case cldr.KO:
		return append(fmt, cldr.Text("년")).Format
	case cldr.LV:
		return append(fmt, cldr.Text(". g.")).Format
	}
}

func fmtYearBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	return cldr.Fmt{
		cldr.Text(fmtEra(locale, opts.Era) + " "),
		convertYearDigitsFmt(digits, opts.Year),
	}.Format
}

func fmtYearPersian(locale language.Tag) func(y string) string {
	lang, _, region := locale.Raw()
	prefix := ""

	if lang != cldr.FA && (lang != cldr.UZ || region != cldr.RegionAF) {
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(y string) string { return prefix + y }
}
