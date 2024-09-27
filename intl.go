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
	case YearUnknown:
		return ""
	case YearNumeric:
		return "numeric"
	case Year2Digit:
		return "2-digit"
	}
}

const (
	YearUnknown Year = iota
	YearNumeric
	Year2Digit
)

type Options struct {
	Year Year
}

type DateTimeFormat struct {
	locale  language.Tag
	options Options
}

func NewDateTimeFormat(locale language.Tag, options Options) *DateTimeFormat {
	return &DateTimeFormat{locale: locale, options: options}
}

func (f *DateTimeFormat) Format(v time.Time) string {
	calendars := calendarPreferences(f.locale)

	if len(calendars) == 0 {
		return fmtYear(f.fmtYear(v), f.locale)
	}

	switch calendars[0] {
	default:
		return fmtYear(f.fmtYear(v), f.locale)
	case "persian":
		return fmtYear(f.fmtPersianYear(v), f.locale)
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
		return f.fmtNumeral(year[len(year)-2:])
	}
}

func (f *DateTimeFormat) fmtNumeral(s string) string {
	num := defaultNumberingSystem(f.locale)
	if num == numberingSystemLatn {
		return s
	}

	digits := numberingSystems[num]

	var r string

	for _, c := range s {
		r += string(digits[c-'0'])
	}

	return r
}
