package intl

import (
	"time"

	"golang.org/x/text/language"
)

type Year byte

func (y Year) String() string {
	switch y {
	default:
		return ""
	case YearNumeric:
		return "numeric"
	case Year2Digit:
		return "2-digit"
	}
}

const (
	YearUnd Year = iota
	YearNumeric
	Year2Digit
)

type Day byte

func (y Day) String() string {
	switch y {
	default:
		return ""
	case DayNumeric:
		return "numeric"
	case Day2Digit:
		return "2-digit"
	}
}

const (
	DayUnd Day = iota
	DayNumeric
	Day2Digit
)

type Options struct {
	Year Year
	Day  Day
}

type digits [10]rune

func (d digits) Sprint(s string) string {
	if d[0] == 0 { // latn
		return s
	}

	var r string

	// s contains only digits
	for _, digit := range []byte(s) {
		if i := int(digit - '0'); i >= 0 && i < len(d) { // isInBounds()
			r += string(d[i])
		}
	}

	return r
}

type DateTimeFormat struct {
	fmt      dateTimeFormatter
	locale   language.Tag
	calendar string
	options  Options
}

func NewDateTimeFormat(locale language.Tag, options Options) *DateTimeFormat {
	var d digits

	if i := defaultNumberingSystem(locale); i > 0 && int(i) < len(numberingSystems) { // isInBounds()
		d = numberingSystems[i]
	}

	var fmt dateTimeFormatter

	switch defaultCalendar(locale) {
	default:
		fmt = &gregorianDateTimeFormat{
			fmtYear: fmtYear(locale),
			fmtDay:  fmtDay(locale, d),
			digits:  d,
		}
	case "persian":
		fmt = &persianDateTimeFormat{
			fmtYear: fmtYear(locale),
			fmtDay:  fmtDay(locale, d),
			digits:  d,
		}
	}

	return &DateTimeFormat{
		locale:   locale,
		options:  options,
		calendar: defaultCalendar(locale),
		fmt:      fmt,
	}
}

func (f *DateTimeFormat) Format(v time.Time) string {
	f.fmt.SetTime(v)

	switch {
	default:
		return ""
	case f.options.Year != YearUnd:
		s := "06"
		if f.options.Year == YearNumeric {
			s = "2006"
		}

		return f.fmt.Year(s)
	case f.options.Day != DayUnd:
		s := "2"
		if f.options.Day == Day2Digit {
			s = "02"
		}

		return f.fmt.Day(s)
	}
}

func (f *DateTimeFormat) fmtYear(v time.Time) string {
	s := "06"
	if f.options.Year == YearNumeric {
		s = "2006"
	}

	return f.fmt.Year(s)
}

// func (f *DateTimeFormat) fmtPersianYear(v time.Time) string {
// 	year := strconv.Itoa(ptime.New(v).Year())

// 	switch f.options.Year {
// 	default:
// 		panic("invalid year option")
// 	case YearNumeric:
// 		return f.fmtNumeral(year)
// 	case Year2Digit:
// 		const last2digits = 2

// 		if len(year) > last2digits {
// 			return f.fmtNumeral(year[len(year)-last2digits:])
// 		}

// 		return f.fmtNumeral(year)
// 	}
// }

// dateTimeFormatter is date time formatter for a specific calendar.
type dateTimeFormatter interface {
	SetTime(time.Time)
	Year(format string) string
	Day(format string) string
}
