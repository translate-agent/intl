package intl

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
)

// Year is year option for [Options].
type Year byte

const (
	YearUnd Year = iota
	YearNumeric
	Year2Digit
)

// String returns the string representation of the [Year].
// It converts the [Year] constant to its corresponding string value.
//
// Returns:
//   - "numeric" for [YearNumeric]
//   - "2-digit" for [Year2Digit]
//   - "" for any other value (including [YearUnd])
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

// ParseYear converts a string representation of a year format to the [Year] type.
//
// Parameters:
//   - s: A string representing the year format. Valid values are "numeric", "2-digit", or an empty string.
//
// Returns:
//   - Year: The corresponding [Year] constant ([YearNumeric], [Year2Digit], or [YearUnd]).
//   - error: An error if the input string is not a valid year format.
func ParseYear(s string) (Year, error) {
	switch s {
	default:
		return YearUnd, fmt.Errorf(`bad year value "%s", want "numeric", "2-digit" or ""`, s)
	case "":
		return YearUnd, nil
	case "numeric":
		return YearNumeric, nil
	case "2-digit":
		return Year2Digit, nil
	}
}

// MustParseYear converts a string representation of a year format to the [Year] type.
// It panics if the input string is not a valid year format.
func MustParseYear(s string) Year {
	v, err := ParseYear(s)
	if err != nil {
		panic(err)
	}

	return v
}

// Day represents the format for displaying days.
type Day byte

const (
	DayUnd Day = iota
	DayNumeric
	Day2Digit
)

// String returns the string representation of the Day format.
// It converts the Day constant to its corresponding string value.
//
// Returns:
//   - "numeric" for [DayNumeric]
//   - "2-digit" for [Day2Digit]
//   - "" for any other value (including [DayUnd])
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

// ParseDay converts a string representation of a year format to the [Day] type.
//
// Parameters:
//   - s: A string representing the day format. Valid values are "numeric", "2-digit", or an empty string.
//
// Returns:
//   - Year: The corresponding [Day] constant ([DayNumeric], [Day2Digit], or [DayUnd]).
//   - error: An error if the input string is not a valid day format.
func ParseDay(s string) (Day, error) {
	switch s {
	default:
		return DayUnd, fmt.Errorf(`bad day value "%s", want "numeric", "2-digit" or ""`, s)
	case "":
		return DayUnd, nil
	case "numeric":
		return DayNumeric, nil
	case "2-digit":
		return Day2Digit, nil
	}
}

// MustParseDay converts a string representation of a year format to the [Day] type.
// It panics if the input string is not a valid day format.
func MustParseDay(s string) Day {
	v, err := ParseDay(s)
	if err != nil {
		panic(err)
	}

	return v
}

// Options defines configuration parameters for [NewDateTimeFormat].
// It allows customization of the date and time representations in formatted output.
type Options struct {
	Year Year
	Day  Day
}

// digits represents a set of numeral glyphs for a specific numeral system.
// It is an array of 10 runes, where each index corresponds to a digit (0-9)
// in the decimal system, and the value at that index is the corresponding
// glyph in the represented numeral system.
//
// For example:
//
//	digits{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} // represents Latin numerals
//	digits{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'} // represents Arabic-Indic numerals
//
// A special case is when digits[0] is 0, which is used to represent Latin numerals
// and triggers special handling in some methods.
type digits [10]rune

// Sprint converts a string of digits to the corresponding digits in the
// numeral system represented by d.
//
// If d[0] is 0 (representing Latin numerals), the function returns the
// input string unchanged.
//
// For other numeral systems, it replaces each digit in the input string
// with the corresponding digit from d.
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

// DateTimeFormat encapsulates the configuration and functionality for
// formatting dates and times according to specific locales and options.
type DateTimeFormat struct {
	fmt      dateTimeFormatter
	locale   language.Tag
	calendar string
	options  Options
}

// NewDateTimeFormat creates a new [DateTimeFormat] instance for the specified locale and options.
//
// This function initializes a [DateTimeFormat] with the default calendar based on the
// given locale. It supports different calendar systems including Gregorian, Persian, and Buddhist calendars.
func NewDateTimeFormat(locale language.Tag, options Options) *DateTimeFormat {
	var d digits

	if i := defaultNumberingSystem(locale); i > 0 && int(i) < len(numberingSystems) { // isInBounds()
		d = numberingSystems[i]
	}

	var fmt dateTimeFormatter

	switch defaultCalendar(locale) {
	default:
		fmt = &gregorianDateTimeFormat{
			fmtYear: fmtYearGregorian(locale),
			fmtDay:  fmtDayGregorian(locale, d),
			digits:  d,
		}
	case "persian":
		fmt = &persianDateTimeFormat{
			fmtYear: fmtYearPersian(locale),
			fmtDay:  fmtDayPersian(locale, d),
			digits:  d,
		}
	case "buddhist":
		fmt = &buddhistDateTimeFormat{
			fmtYear: fmtYearBuddhist(locale),
			fmtDay:  fmtDayBuddhist(locale, d),
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

// Format formats the given [time.Time] value according to the [DateTimeFormat]'s configuration.
//
// This method applies the formatting options specified in the [DateTimeFormat] instance
// to the provided time value.
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

// dateTimeFormatter is date time formatter for a calendar.
type dateTimeFormatter interface {
	SetTime(time.Time)
	Year(format string) string
	Day(format string) string
}
