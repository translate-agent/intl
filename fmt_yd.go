package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqYearDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	seq := symbols.NewSeq(locale)
	year := seqYear(locale, opts.Year)
	day := seqDay(locale, opts.Day)
	withName := !opts.Year.twoDigit() || !opts.Day.numeric()

	switch lang {
	default:
		seq.AddSeq(year).Add(' ')

		if withName {
			seq.Add('(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		} else {
			seq.AddSeq(day)
		}
	case cldr.EN, cldr.KAA, cldr.MHN:
		if withName {
			seq.AddSeq(year).Add(' ').Add('(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		} else {
			seq.AddSeq(day).Add(' ').AddSeq(year)
		}
	case cldr.HI:
		switch {
		default:
			seq.AddSeq(year).Add(' ').AddSeq(day)
		case withName:
			seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		case script == cldr.Latn:
			seq.AddSeq(day).Add(' ').AddSeq(year)
		}
	}

	return seq
}

func seqYearDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	year := seqYearPersian(locale, opts.Year)
	day := opts.Day.symbol()

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		return seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ', day).Add(')')
	}

	return seq.AddSeq(year).Add(' ', day)
}

func seqYearDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	year := seqYearBuddhist(locale, opts)
	day := opts.Day.symbol()

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		return seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ', day).Add(')')
	}

	return seq.AddSeq(year).Add(' ', day)
}
