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
	"golang.org/x/text/language"
)

type calendarType int

const (
	calendarTypeGregorian calendarType = iota
	calendarTypeBuddhist
	calendarTypePersian
	calendarTypeIslamicUmalqura
)

// String implements [fmt.Stringer] interface.
func (t calendarType) String() string {
	switch t {
	default:
		return "gregorian"
	case calendarTypeBuddhist:
		return "buddhist"
	case calendarTypePersian:
		return "persian"
	case calendarTypeIslamicUmalqura:
		return "islamic-umalqura"
	}
}

type Era byte

const (
	EraUnd Era = iota
	EraNarrow
	EraShort
	EraLong
)

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

// MustParseEra converts a string representation of an era format to the [Era] type.
// It panics if the input string is not a valid era format.
func MustParseEra(s string) Era {
	v, err := ParseEra(s)
	if err != nil {
		panic(err)
	}

	return v
}

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

func (y Year) und() bool      { return y == YearUnd }
func (y Year) numeric() bool  { return y == YearNumeric }
func (y Year) twoDigit() bool { return y == Year2Digit }

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

func (m Month) und() bool      { return m == MonthUnd }
func (m Month) numeric() bool  { return m == MonthNumeric }
func (m Month) twoDigit() bool { return m == Month2Digit }

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
	d := localeDigits(locale)

	switch defaultCalendar(locale) {
	default:
		return DateTimeFormat{fmt: gregorianDateTimeFormat(locale, d, options)}
	case calendarTypePersian:
		return DateTimeFormat{fmt: persianDateTimeFormat(locale, d, options)}
	case calendarTypeBuddhist:
		return DateTimeFormat{fmt: buddhistDateTimeFormat(locale, d, options)}
	}
}

// Format formats the given [time.Time] value according to the [DateTimeFormat]'s configuration.
//
// This method applies the formatting options specified in the [DateTimeFormat] instance
// to the provided time value.
func (f DateTimeFormat) Format(v time.Time) string {
	return f.fmt(v)
}

// fmtFunc is date time formatter for a particular calendar.
type fmtFunc func(time.Time) string

// fmtYear formats year.
func fmtYear(digits digits, opt Year) func(v int) string {
	if opt.twoDigit() {
		return digits.twoDigit
	}

	return digits.numeric
}

// fmtMonth returns month formatting func.
func fmtMonth(digits digits, opt Month) func(v int) string {
	if opt.twoDigit() {
		return digits.twoDigit
	}

	return digits.numeric
}

// fmtMonthName formats month as name.
//
// TODO(jhorsts): ensure this is rectified before release v0.1.0 - when formatting of date is complete.
// The "context" is always "stand-alone".
func fmtMonthName(locale string, context, width string) func(v int) string {
	indexes := monthLookup[locale]

	var i int

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

	var names calendarMonths

	if i >= 0 && i < len(indexes) { // isInBounds()
		if v := int(indexes[i]); v > 0 && v < len(calendarMonthNames) { // isInBounds()
			names = calendarMonthNames[v]
		}
	}

	return func(v int) string {
		v--

		if v >= 0 && v < len(names) { // isInBounds()
			return names[v]
		}

		return ""
	}
}

// fmtDay formats day as numeric.
func fmtDay(digits digits, opt Day) func(v int) string {
	if opt.twoDigit() {
		return digits.twoDigit
	}

	return digits.numeric
}

//nolint:cyclop
func gregorianDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		return func(_ time.Time) string {
			return ""
		}
	case !opts.Era.und() && (!opts.Year.und() && !opts.Month.und() && !opts.Day.und() ||
		opts.Year.und() && opts.Month.und() && opts.Day.und()):
		layout := fmtEraYearMonthDayGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Year(), t.Month(), t.Day())
		}
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraYearMonthGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Year(), t.Month())
		}
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && opts.Day.und():
		layout := fmtEraYearGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Year())
		}
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraMonthGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Month())
		}
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraYearDayGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Year(), t.Day())
		}
	case !opts.Era.und() && opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraDayGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Day())
		}
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtEraMonthDayGregorian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(t.Month(), t.Day())
		}
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtYearMonthDayGregorian(locale, digits, opts)

		return func(v time.Time) string {
			return layout(v.Year(), v.Month(), v.Day())
		}
	case !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtYearDayGregorian(locale, digits, opts)

		return func(v time.Time) string {
			return layout(v.Year(), v.Day())
		}
	case !opts.Year.und() && !opts.Month.und():
		layout := fmtYearMonthGregorian(locale, digits, opts)

		return func(v time.Time) string {
			return layout(v.Year(), v.Month())
		}
	case !opts.Month.und() && !opts.Day.und():
		layout := fmtMonthDayGregorian(locale, digits, opts)

		return func(v time.Time) string {
			return layout(v.Month(), v.Day())
		}
	case !opts.Year.und():
		layout := fmtYearGregorian(locale)
		fmt := fmtYear(digits, opts.Year)

		return func(v time.Time) string {
			return layout(fmt(v.Year()))
		}
	case !opts.Month.und():
		layout := fmtMonthGregorian(locale, digits, opts.Month)

		return func(v time.Time) string {
			return layout(int(v.Month()))
		}
	case !opts.Day.und():
		layout := fmtDayGregorian(locale, digits, opts.Day)

		return func(v time.Time) string {
			return layout(v.Day())
		}
	}
}

//nolint:cyclop
func persianDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		return func(_ time.Time) string {
			return ""
		}
	case !opts.Era.und() && (!opts.Year.und() && !opts.Month.und() && !opts.Day.und() ||
		opts.Year.und() && opts.Month.und() && opts.Day.und()):
		layout := fmtEraYearMonthDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)
			return layout(t.Year(), time.Month(t.Month()), t.Day())
		}
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraYearMonthPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)
			return layout(t.Year(), time.Month(t.Month()))
		}
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraMonthPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)
			return layout(time.Month(t.Month()))
		}
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtEraMonthDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)
			return layout(time.Month(t.Month()), t.Day())
		}
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraYearDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)
			return layout(t.Year(), t.Day())
		}
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && opts.Day.und():
		layout := fmtEraYearPersian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(ptime.New(t).Year())
		}
	case !opts.Era.und() && opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraDayPersian(locale, digits, opts)

		return func(t time.Time) string {
			return layout(ptime.New(t).Day())
		}
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtYearMonthDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)

			return layout(t.Year(), time.Month(t.Month()), t.Day())
		}
	case !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtYearDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)

			return layout(t.Year(), t.Day())
		}
	case !opts.Year.und() && !opts.Month.und():
		layout := fmtYearMonthPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)

			return layout(t.Year(), time.Month(t.Month()))
		}
	case !opts.Month.und() && !opts.Day.und():
		layout := fmtMonthDayPersian(locale, digits, opts)

		return func(v time.Time) string {
			t := ptime.New(v)

			return layout(time.Month(t.Month()), t.Day())
		}
	case !opts.Year.und():
		layout := fmtYearPersian(locale)
		fmt := fmtYear(digits, opts.Year)

		return func(v time.Time) string {
			return layout(fmt(ptime.New(v).Year()))
		}
	case !opts.Month.und():
		layout := fmtMonthPersian(locale, digits, opts.Month)

		return func(v time.Time) string {
			return layout(int(ptime.New(v).Month()))
		}
	case !opts.Day.und():
		layout := fmtDayPersian(locale, digits, opts.Day)

		return func(v time.Time) string {
			return layout(ptime.New(v).Day())
		}
	}
}

//nolint:cyclop
func buddhistDateTimeFormat(locale language.Tag, digits digits, opts Options) fmtFunc {
	switch {
	default:
		return func(_ time.Time) string {
			return ""
		}
	case !opts.Era.und() && (!opts.Year.und() && !opts.Month.und() && !opts.Day.und() ||
		opts.Year.und() && opts.Month.und() && opts.Day.und()):
		layout := fmtEraYearMonthDayBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Month(), v.Day())
		})
	case !opts.Era.und() && !opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraYearMonthBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Month())
		})
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraYearDayBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Day())
		})
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && opts.Day.und():
		layout := fmtEraMonthBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Month())
		})
	case !opts.Era.und() && opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtEraDayBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Day())
		})
	case !opts.Era.und() && opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtEraMonthDayBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Month(), v.Day())
		})
	case !opts.Era.und() && !opts.Year.und() && opts.Month.und() && opts.Day.und():
		layout := fmtEraYearBuddhist(locale, digits, opts)

		return (func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year())
		})
	case !opts.Year.und() && !opts.Month.und() && !opts.Day.und():
		layout := fmtYearMonthDayBuddhist(locale, digits, opts)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Month(), v.Day())
		}
	case !opts.Year.und() && opts.Month.und() && !opts.Day.und():
		layout := fmtYearDayBuddhist(locale, digits, opts)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Day())
		}
	case !opts.Year.und() && !opts.Month.und():
		layout := fmtYearMonthBuddhist(locale, digits, opts)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Year(), v.Month())
		}
	case !opts.Month.und() && !opts.Day.und():
		layout := fmtMonthDayBuddhist(locale, digits, opts)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Month(), v.Day())
		}
	case !opts.Year.und():
		layout := fmtYearBuddhist(locale, EraNarrow)
		fmt := fmtYear(digits, opts.Year)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(fmt(v.Year()))
		}
	case !opts.Month.und():
		layout := fmtMonthBuddhist(locale, digits, opts.Month)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(int(v.Month()))
		}
	case !opts.Day.und():
		layout := fmtDayBuddhist(locale, digits, opts.Day)

		return func(v time.Time) string {
			v = v.AddDate(543, 0, 0) //nolint:mnd

			return layout(v.Day())
		}
	}
}
