package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtYearDayGregorian(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	lang, script, _ := locale.Raw()
	year := fmtYearGregorian(locale, digits, opts.Year)

	const (
		layoutYearDay = iota
		layoutDayYear
	)

	withName := !opts.Year.twoDigit() || !opts.Day.numeric()
	dayName := cldr.UnitName(locale).Day
	layout := layoutYearDay
	middle := " "
	suffix := ""

	if withName {
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	switch lang {
	case cldr.BG, cldr.MK:
		if withName {
			middle = " (" + dayName + ": "
		} else {
			middle = " "
		}
	case cldr.KAA, cldr.EN, cldr.MHN:
		if !withName {
			layout = layoutDayYear
		}
	case cldr.HI:
		if !withName && script == cldr.Latn {
			layout = layoutDayYear
		}
	}

	day := fmtDayGregorian(locale, digits, opts.Day)

	if layout == layoutDayYear {
		return func(t cldr.TimeReader) string {
			return day(t) + middle + year(t) + suffix
		}
	}

	// layoutYearDay
	return func(t cldr.TimeReader) string {
		return year(t) + middle + day(t) + suffix
	}
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
