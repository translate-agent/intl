// Package intl provides internationalization and localization support for date and time formatting,
// inspired by and reflecting the ECMAScript Intl API.
//
// This package offers a flexible and extensible way to format dates and times according to
// various locales, calendar systems, and custom options, similar to the JavaScript Intl.DateTimeFormat
// object. It supports multiple numeral systems and calendar types, including Gregorian, Persian,
// and Buddhist calendars.
//
// Key features:
//   - Locale-aware date and time formatting
//   - Support for different calendar systems
//   - Customizable date and time formatting options
//   - Handling of various numeral systems
//   - API design similar to ECMAScript Intl.DateTimeFormat
//
// The main types in this package are:
//   - DateTimeFormat: Encapsulates formatting logic for dates and times, similar to Intl.DateTimeFormat
//   - Options: Configures formatting options for date and time representations
//
// Usage:
//
//	formatter := intl.NewDateTimeFormat(locale, Options{Day: Day2Digit})
//	now := formatter.Format(time.Now())
package intl

import (
	"fmt"
	"strconv"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/text/language"
)

type calendarType int

const (
	calendarTypeGregorian calendarType = iota
	calendarTypeBuddhist
	calendarTypePersian
	calendarTypeIslamicUmalqura
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

// Month represents the format for displaying months.
type Month byte

const (
	MonthUnd Month = iota
	MonthNumeric
	Month2Digit
	MonthLong
	MonthShort
	MonthNarrow
)

// String returns the string representation of the [Month].
// It converts the [Month] constant to its corresponding string value.
//
// Returns:
//   - "numeric" for [MonthNumeric]
//   - "2-digit" for [Month2Digit]
//   - "long" for [MonthLong]
//   - "short" for [MonthShort]
//   - "narrow" for [MonthNarrow]
//   - "" for any other value (including [MonthUnd])
func (m Month) String() string {
	switch m {
	default:
		return ""
	case MonthNumeric:
		return "numeric"
	case Month2Digit:
		return "2-digit"
	case MonthLong:
		return "long"
	case MonthShort:
		return "short"
	case MonthNarrow:
		return "narrow"
	}
}

// ParseMonth converts a string representation of a month format to the [Month] type.
//
// Parameters:
//   - s: A string representing the month format. Valid values are "numeric", "2-digit",
//     "long", "short", "narrow", or an empty string.
//
// Returns:
//   - Month: The corresponding [Month] constant ([MonthNumeric], [Month2Digit],
//     [MonthLong], [MonthShort], [MonthNarrow], or [MonthUnd]).
//   - error: An error if the input string is not a valid month format.
func ParseMonth(s string) (Month, error) {
	switch s {
	default:
		return MonthUnd, fmt.Errorf(`bad month value "%s", want "numeric", "2-digit", "long", "short", "narrow" or ""`, s)
	case "":
		return MonthUnd, nil
	case "numeric":
		return MonthNumeric, nil
	case "2-digit":
		return Month2Digit, nil
	case "long":
		return MonthLong, nil
	case "short":
		return MonthShort, nil
	case "narrow":
		return MonthNarrow, nil
	}
}

// MustParseMonth converts a string representation of a month format to the [Month] type.
// It panics if the input string is not a valid month format.
func MustParseMonth(s string) Month {
	v, err := ParseMonth(s)
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
	Year  Year
	Month Month
	Day   Day
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
	format fmtFunc
}

// NewDateTimeFormat creates a new [DateTimeFormat] instance for the specified locale and options.
//
// This function initializes a [DateTimeFormat] with the default calendar based on the
// given locale. It supports different calendar systems including Gregorian, Persian, and Buddhist calendars.
func NewDateTimeFormat(locale language.Tag, options Options) DateTimeFormat {
	var d digits

	if i := defaultNumberingSystem(locale); i > 0 && int(i) < len(numberingSystems) { // isInBounds()
		d = numberingSystems[i]
	}

	switch defaultCalendar(locale) {
	default:
		return DateTimeFormat{format: gregorianDateTimeFormat(locale, d, options)}
	case calendarTypePersian:
		return DateTimeFormat{format: persianDateTimeFormat(locale, d, options)}
	case calendarTypeBuddhist:
		return DateTimeFormat{format: buddhistDateTimeFormat(locale, d, options)}
	}
}

// Format formats the given [time.Time] value according to the [DateTimeFormat]'s configuration.
//
// This method applies the formatting options specified in the [DateTimeFormat] instance
// to the provided time value.
func (f DateTimeFormat) Format(v time.Time) string {
	return f.format(v)
}

// fmtFunc is date time formatter for a particular calendar.
type fmtFunc func(time.Time) string

// fmtYear formats year as numeric.
func fmtYear(digits digits) func(year int, f string) string {
	return func(year int, f string) string {
		s := strconv.Itoa(year)

		if f == "06" {
			switch n := len(s); n {
			default:
				s = s[n-2:]
			case 1:
				s = "0" + s
			case 0, 2: //nolint:mnd // noop, isSliceInBounds()
			}
		}

		return digits.Sprint(s)
	}
}

// fmtMonth formats month as numeric. The assumption is that f is "1" or "01".
func fmtMonth(digits digits) func(m int, f string) string {
	return func(m int, f string) string {
		if f == "01" && m <= 9 {
			return digits.Sprint("0" + strconv.Itoa(m))
		}

		return digits.Sprint(strconv.Itoa(m))
	}
}

// fmtMonthName formats month as name.
func fmtMonthName(locale string, calendar calendarType, context, width string) func(int, string) string {
	indexes := monthLookup[locale]

	// multiply by 6
	i := calendar<<2 + calendar<<1 //nolint:mnd

	// "abbreviated" width index is 0
	switch width {
	case "wide":
		i += 2 // 1*2
	case "narrow":
		i += 4 // 2*2
	}

	// "format" context index is 0
	if context == "stand-alone" {
		i++
	}

	names := calendarMonthNames[int(indexes[i])]

	return func(m int, _ string) string {
		return names[m-1]
	}
}

// fmtDay formats day as numeric. The assumption is that f is "2" or "02".
func fmtDay(digits digits) func(d int, f string) string {
	return func(d int, f string) string {
		if f == "02" && d <= 9 {
			return digits.Sprint("0" + strconv.Itoa(d))
		}

		return digits.Sprint(strconv.Itoa(d))
	}
}

func gregorianDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		panic(fmt.Errorf("gregorian calendar formatter: NOT IMPLEMENTED: %+v", opts))
	case opts.Year != YearUnd:
		s := "06"
		if opts.Year == YearNumeric {
			s = "2006"
		}

		layout := fmtYearGregorian(locale)
		fmt := fmtYear(digits)

		return func(v time.Time) string {
			return layout(fmt(v.Year(), s))
		}
	case opts.Month != MonthUnd:
		s := "1"
		if opts.Month == Month2Digit {
			s = "01"
		}

		layout := fmtMonthGregorian(locale, digits)

		return func(v time.Time) string {
			return layout(int(v.Month()), s)
		}
	case opts.Day != DayUnd:
		s := "2"
		if opts.Day == Day2Digit {
			s = "02"
		}

		layout := fmtDayGregorian(locale, digits)

		return func(v time.Time) string {
			return layout(v.Day(), s)
		}
	}
}

func persianDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		panic(fmt.Errorf("persian calendar formatter: NOT IMPLEMENTED: %+v", opts))
	case opts.Year != YearUnd:
		s := "06"
		if opts.Year == YearNumeric {
			s = "2006"
		}

		layout := fmtYearPersian(locale)
		fmt := fmtYear(digits)

		return func(v time.Time) string {
			return layout(fmt(ptime.New(v).Year(), s))
		}
	case opts.Month != MonthUnd:
		s := "1"
		if opts.Month == Month2Digit {
			s = "01"
		}

		layout := fmtMonthPersian(locale, digits)

		return func(v time.Time) string {
			return layout(int(ptime.New(v).Month()), s)
		}
	case opts.Day != DayUnd:
		s := "2"
		if opts.Day == Day2Digit {
			s = "02"
		}

		layout := fmtDayPersian(locale, digits)

		return func(v time.Time) string {
			return layout(ptime.New(v).Day(), s)
		}
	}
}

func buddhistDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		panic(fmt.Errorf("buddhist calendar formatter: NOT IMPLEMENTED: %+v", opts))
	case opts.Year != YearUnd:
		s := "06"
		if opts.Year == YearNumeric {
			s = "2006"
		}

		layout := fmtYearBuddhist(locale)
		fmt := fmtYear(digits)

		return func(v time.Time) string {
			return layout(fmt(v.Year(), s))
		}
	case opts.Month != MonthUnd:
		s := "1"
		if opts.Month == Month2Digit {
			s = "01"
		}

		layout := fmtMonthBuddhist(locale, digits)

		return func(v time.Time) string {
			return layout(int(v.Month()), s)
		}
	case opts.Day != DayUnd:
		s := "2"
		if opts.Day == Day2Digit {
			s = "02"
		}

		layout := fmtDayBuddhist(locale, digits)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Day(), s)
		}
	}
}
