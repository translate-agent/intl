package intl

import (
	"golang.org/x/text/language"
)

func fmtYearDayGregorian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, script, _ := locale.Raw()
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)

	const (
		layoutYearDay = iota
		layoutDayYear
	)

	withName := opts.Year != Year2Digit || opts.Day != DayNumeric
	dayName := unitName(locale).Day
	layout := layoutYearDay
	middle := " "
	suffix := ""

	if withName {
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	switch lang {
	case bg, mk:
		if withName {
			middle = " (" + dayName + ": "
		} else {
			middle = " "
		}
	case kaa, en, mhn:
		if !withName {
			layout = layoutDayYear
		}
	case hi:
		if !withName && script == latn {
			layout = layoutDayYear
		}
	}

	fmtDay := fmtDayGregorian(locale, digits, opts.Day)

	if layout == layoutDayYear {
		return func(y, d int) string {
			return fmtDay(d) + middle + layoutYear(fmtYear(y, opts.Year)) + suffix
		}
	}

	// layoutYearDay
	return func(y, d int) string {
		return layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d) + suffix
	}
}

func fmtYearDayPersian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, _, region := locale.Raw()
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)

	prefix := ""
	middle := " "
	suffix := ""

	if opts.Year != Year2Digit || opts.Day != DayNumeric {
		dayName := unitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	if !(lang == fa || lang == uz && region == regionAF) {
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	fmtDay := fmtDayGregorian(locale, digits, opts.Day)

	return func(y, d int) string {
		return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d) + suffix
	}
}

func fmtYearDayBuddhist(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	layoutYear := fmtYearBuddhist(locale, EraNarrow)
	fmtYear := fmtYear(digits)

	middle := " "
	suffix := ""

	if opts.Year != Year2Digit || opts.Day != DayNumeric {
		dayName := unitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	fmtDay := fmtDayBuddhist(locale, digits, opts.Day)

	return func(y, d int) string {
		return layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d) + suffix
	}
}
