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
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

type Era byte

const (
	EraUnd Era = iota
	EraNarrow
	EraShort
	EraLong
)

// MustParseEra converts a string representation of an era format to the [Era] type.
// It panics if the input string is not a valid era format.
func MustParseEra(s string) Era {
	v, err := ParseEra(s)
	if err != nil {
		panic(err)
	}

	return v
}

// String returns the string representation of the [Era].
// It converts the [Era] constant to its corresponding string value.
//
// Returns:
//   - "narrow" for [EraNarrow]
//   - "short" for [EraShort]
//   - "long" for [EraLong]
//   - "" for any other value (including [EraUnd])
func (e Era) String() string {
	switch e {
	default:
		return ""
	case EraNarrow:
		return "narrow"
	case EraShort:
		return "short"
	case EraLong:
		return "long"
	}
}

func (e Era) und() bool    { return e == EraUnd }
func (e Era) narrow() bool { return e == EraNarrow }
func (e Era) short() bool  { return e == EraShort }
func (e Era) long() bool   { return e == EraLong }

func (e Era) symbol() symbols.Symbol {
	switch e {
	default:
		return symbols.Symbol_GGGGG
	case EraShort:
		return symbols.Symbol_G
	case EraLong:
		return symbols.Symbol_GGGG
	}
}

// ParseEra converts a string representation of a year format to the [Era] type.
//
// Parameters:
//   - s: A string representing the era format. Valid values are "narrow", "short", "long" or an empty string.
//
// Returns:
//   - Era: The corresponding [Era] constant ([EraNarrow], [EraShort], or [EraLong]).
//   - error: An error if the input string is not a valid era format.
func ParseEra(s string) (Era, error) {
	switch s {
	default:
		return EraUnd, fmt.Errorf(`bad era value "%s", want "narrow", "short", "long" or ""`, s)
	case "":
		return EraUnd, nil
	case "narrow":
		return EraNarrow, nil
	case "short":
		return EraShort, nil
	case "long":
		return EraLong, nil
	}
}

// Year is year option for [Options].
type Year byte

const (
	YearUnd Year = iota
	YearNumeric
	Year2Digit
)

// MustParseYear converts a string representation of a year format to the [Year] type.
// It panics if the input string is not a valid year format.
func MustParseYear(s string) Year {
	v, err := ParseYear(s)
	if err != nil {
		panic(err)
	}

	return v
}

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

func (y Year) und() bool      { return y == YearUnd }
func (y Year) numeric() bool  { return y == YearNumeric }
func (y Year) twoDigit() bool { return y == Year2Digit }

func (y Year) symbol() symbols.Symbol {
	if y.twoDigit() {
		return symbols.Symbol_yy
	}

	return symbols.Symbol_y
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

// MustParseMonth converts a string representation of a month format to the [Month] type.
// It panics if the input string is not a valid month format.
func MustParseMonth(s string) Month {
	v, err := ParseMonth(s)
	if err != nil {
		panic(err)
	}

	return v
}

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

func (m Month) und() bool      { return m == MonthUnd }
func (m Month) numeric() bool  { return m == MonthNumeric }
func (m Month) twoDigit() bool { return m == Month2Digit }

func (m Month) symbolFormat() symbols.Symbol { return m.symbol("format") }

// func (m Month) symbolStandAlone() symbols.Symbol { return m.symbol("stand-alone") }

// TODO(jhorsts): define iota for context values.
func (m Month) symbol(context string) symbols.Symbol {
	switch m {
	default: // M
		return symbols.Symbol_M
	case Month2Digit:
		return symbols.Symbol_MM
	case MonthNarrow:
		if context == "format" {
			return symbols.Symbol_MMMMM
		}

		return symbols.Symbol_LLLLL
	case MonthLong:
		if context == "format" {
			return symbols.Symbol_MMM
		}

		return symbols.Symbol_LLL
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

// Day represents the format for displaying days.
type Day byte

const (
	DayUnd Day = iota
	DayNumeric
	Day2Digit
)

// MustParseDay converts a string representation of a year format to the [Day] type.
// It panics if the input string is not a valid day format.
func MustParseDay(s string) Day {
	v, err := ParseDay(s)
	if err != nil {
		panic(err)
	}

	return v
}

// String returns the string representation of the Day format.
// It converts the Day constant to its corresponding string value.
//
// Returns:
//   - "numeric" for [DayNumeric]
//   - "2-digit" for [Day2Digit]
//   - "" for any other value (including [DayUnd])
func (d Day) String() string {
	switch d {
	default:
		return ""
	case DayNumeric:
		return "numeric"
	case Day2Digit:
		return "2-digit"
	}
}

func (d Day) und() bool      { return d == DayUnd }
func (d Day) numeric() bool  { return d == DayNumeric }
func (d Day) twoDigit() bool { return d == Day2Digit }

func (d Day) symbol() symbols.Symbol {
	if d.twoDigit() {
		return symbols.Symbol_dd
	}

	return symbols.Symbol_d
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

// Options defines configuration parameters for [NewDateTimeFormat].
// It allows customization of the date and time representations in formatted output.
type Options struct {
	Era   Era
	Year  Year
	Month Month
	Day   Day
}

// DateTimeFormat encapsulates the configuration and functionality for
// formatting dates and times according to specific locales and options.
type DateTimeFormat struct {
	fmt fmtFunc
}

// NewDateTimeFormat creates a new [DateTimeFormat] instance for the specified locale and options.
//
// This function initializes a [DateTimeFormat] with the default calendar based on the
// given locale. It supports different calendar systems including Gregorian, Persian, and Buddhist calendars.
func NewDateTimeFormat(locale language.Tag, options Options) DateTimeFormat {
	switch cldr.DefaultCalendar(locale) {
	default:
		return DateTimeFormat{fmt: gregorianDateTimeFormat(locale, options)}
	case cldr.CalendarTypePersian:
		return DateTimeFormat{fmt: persianDateTimeFormat(locale, options)}
	case cldr.CalendarTypeBuddhist:
		return DateTimeFormat{fmt: buddhistDateTimeFormat(locale, options)}
	}
}

// Format formats the given [time.Time] value according to the [DateTimeFormat]'s configuration.
//
// This method applies the formatting options specified in the [DateTimeFormat] instance
// to the provided time value.
func (f DateTimeFormat) Format(t time.Time) string {
	return f.fmt(cldr.TimeReader(t))
}

// fmtFunc is date time formatter for a particular calendar.
type fmtFunc func(cldr.TimeReader) string

//nolint:cyclop
func gregorianDateTimeFormat(locale language.Tag, opts Options) fmtFunc {
	var seq *symbols.Seq

	switch {
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraYearMonthDay(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und():
		seq = seqEraYearMonth(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Day.und():
		seq = seqEraYearDay(locale, opts)
	case !opts.Era.und() && !opts.Year.und():
		seq = seqEraYear(locale, opts)
	case !opts.Era.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraMonthDay(locale, opts)
	case !opts.Era.und() && !opts.Month.und():
		seq = seqEraMonth(locale, opts)
	case !opts.Era.und() && !opts.Day.und():
		seq = seqEraDay(locale, opts)
	case !opts.Era.und():
		seq = seqEraYearMonthDay(locale, opts)
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqYearMonthDay(locale, opts)
	case !opts.Year.und() && !opts.Month.und():
		seq = seqYearMonth(locale, opts)
	case !opts.Year.und() && !opts.Day.und():
		seq = seqYearDay(locale, opts)
	case !opts.Year.und():
		seq = seqYear(locale, opts.Year)
	case !opts.Month.und() && !opts.Day.und():
		seq = seqMonthDay(locale, opts)
	case !opts.Month.und():
		seq = seqMonth(locale, opts.Month)
	case !opts.Day.und():
		seq = seqDay(locale, opts.Day)
	}

	return seq.Func()
}

//nolint:cyclop
func persianDateTimeFormat(locale language.Tag, opts Options) fmtFunc {
	var seq *symbols.Seq

	switch {
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraYearMonthDayPersian(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und():
		seq = seqEraYearMonthPersian(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Day.und():
		seq = seqEraYearDayPersian(locale, opts)
	case !opts.Era.und() && !opts.Year.und():
		seq = seqEraYearPersian(locale, opts)
	case !opts.Era.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraMonthDayPersian(locale, opts)
	case !opts.Era.und() && !opts.Month.und():
		seq = seqEraMonthPersian(locale, opts)
	case !opts.Era.und() && !opts.Day.und():
		seq = seqEraDayPersian(locale, opts)
	case !opts.Era.und():
		seq = seqEraYearMonthDayPersian(locale, opts)
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqYearMonthDayPersian(locale, opts)
	case !opts.Year.und() && !opts.Month.und():
		seq = seqYearMonthPersian(locale, opts)
	case !opts.Year.und() && !opts.Day.und():
		seq = seqYearDayPersian(locale, opts)
	case !opts.Year.und():
		seq = seqYearPersian(locale, opts.Year)
	case !opts.Month.und() && !opts.Day.und():
		seq = seqMonthDayPersian(locale, opts)
	case !opts.Month.und():
		seq = seqMonthPersian(locale, opts.Month)
	case !opts.Day.und():
		seq = seqDayPersian(locale, opts.Day)
	}

	f := seq.Func()

	return func(t cldr.TimeReader) string {
		v, _ := t.(time.Time) // t is always [time.Time]
		return f(persionTime(ptime.New(v)))
	}
}

//nolint:cyclop
func buddhistDateTimeFormat(locale language.Tag, opts Options) fmtFunc {
	var seq *symbols.Seq

	switch {
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraYearMonthDayBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und():
		seq = seqEraYearMonthBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Year.und() && !opts.Day.und():
		seq = seqEraYearDayBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Year.und():
		seq = seqEraYearBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqEraMonthDayBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Month.und():
		seq = seqEraMonthBuddhist(locale, opts)
	case !opts.Era.und() && !opts.Day.und():
		seq = seqEraDayBuddhist(locale, opts)
	case !opts.Era.und():
		seq = seqEraYearMonthDayBuddhist(locale, opts)
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		seq = seqYearMonthDayBuddhist(locale, opts)
	case !opts.Year.und() && !opts.Month.und():
		seq = seqYearMonthBuddhist(locale, opts)
	case !opts.Year.und() && !opts.Day.und():
		seq = seqYearDayBuddhist(locale, opts)
	case !opts.Year.und():
		seq = seqYearBuddhist(locale, opts)
	case !opts.Month.und() && !opts.Day.und():
		seq = seqMonthDayBuddhist(locale, opts)
	case !opts.Month.und():
		seq = seqMonthBuddhist(locale, opts.Month)
	case !opts.Day.und():
		seq = seqDayBuddhist(locale, opts.Day)
	}

	f := seq.Func()

	return func(t cldr.TimeReader) string {
		v, _ := t.(time.Time)          // t is always [time.Time]
		return f(v.AddDate(543, 0, 0)) //nolint:mnd
	}
}

type persionTime ptime.Time

func (p persionTime) Year() int {
	return ptime.Time(p).Year()
}

func (p persionTime) Month() time.Month {
	return time.Month(ptime.Time(p).Month())
}

func (p persionTime) Day() int {
	return ptime.Time(p).Day()
}
