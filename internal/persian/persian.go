// Package persian provides a minimal internal implementation of the Persian (Jalaali/Solar Hijri) calendar
// for converting Gregorian time values.
//
// Limitations:
//   - Only Gregorian years >= 1097 are supported. FromGregorian returns a zero Time struct for earlier dates.
//   - The implementation uses an arithmetic (tabular) 33-year cycle algorithm to match the behavior of
//     JavaScript's Intl.DateTimeFormat (V8/ICU). As a result, it may occasionally differ by one day from
//     the official astronomical Iranian civil calendar.
package persian

import "time"

// Time represents a moment in time in Persian (Jalali/Solar Hijri) Calendar.
type Time struct {
	Year       int
	Month      int // 1-12
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	Weekday    time.Weekday
}

const (
	minGregorianYear = 1097

	// toJDN constants.
	gregorianYearOffset        = 4800
	gregorianBaseDayAdjustment = 32075
	gregorianMonthCutoff       = 2
	gregorianMonthShift        = 10
	gregorianFourYearCycle     = 1461
	gregorianMonthFactor       = 367
	gregorianMonthsInYear      = 12
	gregorianCenturyFactor     = 3
	gregorianCenturyScale      = 100
	gregorianLeapAdjustment    = 4

	// jdnToPersian constants.
	jdnShamsiEpochOffset = 1365393
	daysIn33YearCycle    = 12053
	cycleOf33Years       = 33
	shamsiYearOffset     = -1595
	daysIn4YearCycle     = 1461
	fourYears            = 4
	daysInNormalYear     = 365
	daysInThreeYears     = 1095
	daysInTwoYears       = 730
	daysInLeapYearFirst  = 1096
	daysInLeapYearSecond = 731
	daysInLeapYearThird  = 366
	daysInFirstSixMonths = 186
	daysInMonth31        = 31
	daysInMonth30        = 30
	lastMonthsOffset     = 7
)

// toJDN converts a Gregorian date to Julian Day Number (JDN).
// Uses a standard Gregorian-to-JDN algorithm.
func toJDN(year, month, day int) int {
	var y, m int
	if month <= gregorianMonthCutoff {
		y = year + gregorianYearOffset - 1
		m = month + gregorianMonthShift
	} else {
		y = year + gregorianYearOffset
		m = month - gregorianMonthCutoff
	}

	return (gregorianFourYearCycle*y)/gregorianLeapAdjustment +
		(gregorianMonthFactor*m)/gregorianMonthsInYear -
		(gregorianCenturyFactor*((y+gregorianCenturyScale)/gregorianCenturyScale))/gregorianLeapAdjustment +
		day - gregorianBaseDayAdjustment
}

func jdnToPersianYear(jdn int) (y int, rem int) {
	daysSinceEpoch := jdn - jdnShamsiEpochOffset

	y = (daysSinceEpoch/daysIn33YearCycle)*cycleOf33Years + shamsiYearOffset
	rem = daysSinceEpoch % daysIn33YearCycle

	y += (rem / daysIn4YearCycle) * fourYears
	rem %= daysIn4YearCycle

	if rem > daysInNormalYear {
		switch {
		case rem > daysInThreeYears:
			y += 3
			rem -= daysInLeapYearFirst
		case rem > daysInTwoYears:
			y += 2
			rem -= daysInLeapYearSecond
		default:
			y++
			rem -= daysInLeapYearThird
		}
	}

	return y, rem
}

func remToPersianMonthDay(rem int) (m int, d int) {
	if rem < daysInFirstSixMonths {
		mOffset := rem / daysInMonth31

		return 1 + mOffset, 1 + rem - mOffset*daysInMonth31
	}

	offset := rem - daysInFirstSixMonths
	mOffset := offset / daysInMonth30

	return lastMonthsOffset + mOffset, 1 + offset - mOffset*daysInMonth30
}

// FromGregorian converts a Gregorian time.Time value to the Persian (Solar Hijri) calendar.
// It returns a Time struct containing all the date/time fields.
// If the Gregorian year is less than 1097, it returns a zero Time.
func FromGregorian(t time.Time) Time {
	gy, gmm, gd := t.Date()
	if gy < minGregorianYear {
		return Time{}
	}

	gm := int(gmm)

	jdn := toJDN(gy, gm, gd)
	year, rem := jdnToPersianYear(jdn)
	month, day := remToPersianMonthDay(rem)

	hour, minute, sec := t.Clock()

	return Time{
		Year:       year,
		Month:      month,
		Day:        day,
		Hour:       hour,
		Minute:     minute,
		Second:     sec,
		Nanosecond: t.Nanosecond(),
		Weekday:    t.Weekday(),
	}
}
