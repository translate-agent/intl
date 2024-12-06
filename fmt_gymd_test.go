package intl

import (
	"runtime"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func BenchmarkFmtEraYearMonthDayGregorian(b *testing.B) {
	now := time.Now()

	fmt := NewDateTimeFormat(language.Latvian, Options{
		Era:   EraLong,
		Year:  YearNumeric,
		Month: MonthNumeric,
		Day:   DayNumeric,
	})

	var s string

	for range b.N {
		s = fmt.Format(now)
	}

	runtime.KeepAlive(s)
}
