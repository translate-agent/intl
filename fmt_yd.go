package intl

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/text/language"
)

func fmtYearDayGregorian(locale language.Tag, digits digits, opts Options) fmtFunc {
	lang, script, _ := locale.Raw()
	year := fmtYearGregorian(locale, digits, opts.Year)

	const (
		layoutYearDay = iota
		layoutDayYear
	)

	withName := !opts.Year.twoDigit() || !opts.Day.numeric()
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

	day := fmtDayGregorian(locale, digits, opts.Day)

	if layout == layoutDayYear {
		return func(t time.Time) string {
			return day(t) + middle + year(t) + suffix
		}
	}

	// layoutYearDay
	return func(t time.Time) string {
		return year(t) + middle + day(t) + suffix
	}
}

func fmtYearDayPersian(locale language.Tag, digits digits, opts Options) fmtPersianFunc {
	year := fmtYearPersian(locale)
	yearDigits := convertYearDigitsPersian(digits, opts.Year)

	prefix := ""
	middle := " "
	suffix := ""

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		dayName := unitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	day := fmtDayPersian(locale, digits, opts.Day)

	return func(v ptime.Time) string {
		return prefix + year(yearDigits(v)) + middle + day(v) + suffix
	}
}

func fmtYearDayBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	year := fmtYearBuddhist(locale, digits, opts)

	middle := " "
	suffix := ""

	if !opts.Year.twoDigit() || !opts.Day.numeric() {
		dayName := unitName(locale).Day
		middle = " (" + dayName + ": "
		suffix = ")"
	}

	day := fmtDayBuddhist(locale, digits, opts.Day)

	return func(t time.Time) string {
		return year(t) + middle + day(t) + suffix
	}
}
