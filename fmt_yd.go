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
	case cldr.BG, cldr.MK:
		seq.AddSeq(year).Add(symbols.TxtNNBSP)

		if withName {
			seq.Add('(', symbols.DayUnit, ':', symbols.TxtNNBSP).AddSeq(day).Add(')')
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

func fmtYearDayPersian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigits(digits, opts.Year)

	prefix := ""
	middle := " "
	suffix := ""

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		dayName := cldr.UnitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	day := fmtDayPersian(locale, digits, opts.Day)

	return func(v cldr.TimeReader) string {
		return prefix + year(yearDigits(v)) + middle + day(v) + suffix
	}
}

func fmtYearDayBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)

	middle := " "
	suffix := ""

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		dayName := cldr.UnitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	day := fmtDayBuddhist(locale, digits, opts.Day)

	return func(t cldr.TimeReader) string {
		return year(t) + middle + day(t) + suffix
	}
}
