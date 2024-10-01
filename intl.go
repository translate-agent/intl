package intl

import (
	"strconv"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
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

type DateTimeFormat struct {
	locale   language.Tag
	calendar string
	digits   digits
	options  Options
}

func NewDateTimeFormat(locale language.Tag, options Options) *DateTimeFormat {
	var d digits

	if i := defaultNumberingSystem(locale); i > 0 && int(i) < len(numberingSystems) { // isInBounds()
		d = numberingSystems[i]
	}

	return &DateTimeFormat{
		locale:   locale,
		options:  options,
		calendar: defaultCalendar(locale),
		digits:   d,
	}
}

func (f *DateTimeFormat) Format(v time.Time) string {
	switch {
	default:
		return ""
	case f.options.Year != YearUnd:
		switch f.calendar {
		default: // gregorian
			return fmtYear(f.fmtYear(v), f.locale)
		case "persian":
			return fmtYear(f.fmtPersianYear(v), f.locale)
		}
	case f.options.Day != DayUnd:
		switch f.calendar {
		default: // gregorian
			return f.fmtDay(v.Day())
		case "persian":
			return f.fmtDay(ptime.New(v).Day())
		}
	}
}

func (f *DateTimeFormat) fmtYear(v time.Time) string {
	s := v.Format("06")
	if f.options.Year == YearNumeric {
		s = v.Format("2006")
	}

	return f.fmtNumeral(s)
}

func (f *DateTimeFormat) fmtPersianYear(v time.Time) string {
	year := strconv.Itoa(ptime.New(v).Year())

	switch f.options.Year {
	default:
		panic("invalid year option")
	case YearNumeric:
		return f.fmtNumeral(year)
	case Year2Digit:
		const last2digits = 2

		if len(year) > last2digits {
			return f.fmtNumeral(year[len(year)-last2digits:])
		}

		return f.fmtNumeral(year)
	}
}

func (f *DateTimeFormat) fmtNumeral(s string) string {
	if f.digits[0] == 0 { // latn
		return s
	}

	var r string

	// s contains only digits
	for _, digit := range []byte(s) {
		if i := int(digit - '0'); i >= 0 && i < len(f.digits) { // isInBounds()
			r += string(f.digits[i])
		}
	}

	return r
}
