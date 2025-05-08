package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqEraMonth(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := symbolEra(opts.Era)
	month := opts.Month.symbolFormat()
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	switch lang {
	case cldr.EN, cldr.KAA, cldr.MHN:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(month, ' ', era)
	case cldr.BG, cldr.CY, cldr.MK:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(era, ' ', month)
	case cldr.BR, cldr.FO, cldr.GA, cldr.LT, cldr.UK, cldr.UZ:
		month = Month2Digit.symbolFormat()

		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
		}

		return seq.Add(era, ' ', month)
	case cldr.HR, cldr.NB, cldr.NN, cldr.NO, cldr.SK:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, '.', ')')
		}

		return seq.Add(era, ' ', month, '.')
	case cldr.HI:
		if script != cldr.Latn {
			break
		}

		if opts.Era.long() && opts.Month.numeric() || opts.Era.narrow() {
			return seq.Add(month, ' ', era)
		}
	case cldr.MN:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', symbols.Symbol_LLLLL, ')')
		}

		return seq.Add(era, ' ', symbols.Symbol_LLLLL)
	case cldr.WAE:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', symbols.Symbol_MMM, ')')
		}

		return seq.Add(era, ' ', symbols.Symbol_MMM)
	case cldr.JA, cldr.KO, cldr.ZH, cldr.YUE:
		if withName {
			return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, symbols.MonthUnit, ')')
		}

		return seq.Add(era, ' ', month, symbols.MonthUnit)
	}

	if withName {
		return seq.Add(era, ' ', '(', symbols.MonthUnit, ':', ' ', month, ')')
	}

	return seq.Add(era, ' ', month)
}

func fmtEraMonthPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	monthName := cldr.UnitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	month := convertMonthDigits(digits, opts.Month)

	return func(v cldr.TimeReader) string { return prefix + month(v) + suffix }
}

func fmtEraMonthBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	era := fmtEra(locale, opts.Era)
	monthDigits := convertMonthDigits(digits, opts.Month)
	monthName := cldr.UnitName(locale).Month
	withName := opts.Era.short() || opts.Era.long() && opts.Month.twoDigit()

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + monthName + ": "
		suffix = ")"
	}

	return func(t cldr.TimeReader) string { return prefix + monthDigits(t) + suffix }
}
