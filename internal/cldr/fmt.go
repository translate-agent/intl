package cldr

import (
	"strings"
	"time"
)

type timeReader interface {
	Year() int
	Month() time.Month
	Day() int
}

type Fmt []FmtFunc

func (f Fmt) Format(t timeReader) string {
	var b strings.Builder

	for _, fn := range f {
		fn(b, t)
	}

	return b.String()
}

type FmtFunc func(strings.Builder, timeReader)

func MonthNumeric(digits Digits) FmtFunc {
	return func(b strings.Builder, t timeReader) {
		digits.appendNumeric(b, int(t.Month()))
	}
}

func MonthTwoDigit(digits Digits) FmtFunc {
	return func(b strings.Builder, t timeReader) {
		digits.appendTwoDigit(b, int(t.Month()))
	}
}

func Month(months []string) FmtFunc {
	return func(b strings.Builder, t timeReader) {
		b.WriteString(months[t.Month()-1])
	}
}

func DayNumeric(digits Digits) FmtFunc {
	return func(b strings.Builder, t timeReader) {
		digits.appendNumeric(b, t.Day())
	}
}

func DayTwoDigit(digits Digits) FmtFunc {
	return func(b strings.Builder, t timeReader) {
		digits.appendTwoDigit(b, t.Day())
	}
}
