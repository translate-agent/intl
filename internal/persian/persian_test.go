package persian

import (
	"testing"
	"time"
)

func TestFromGregorian(t *testing.T) {
	t.Parallel()

	tests := []struct {
		t                time.Time
		year, month, day int
	}{
		{time.Date(1096, 12, 31, 12, 0, 0, 0, time.UTC), 0, 0, 0},
		{time.Date(1097, 1, 1, 12, 0, 0, 0, time.UTC), 475, 10, 12},
		{time.Date(1097, 3, 20, 12, 0, 0, 0, time.UTC), 475, 12, 30},
		{time.Date(1582, 10, 14, 12, 0, 0, 0, time.UTC), 961, 7, 22},
		{time.Date(1582, 10, 15, 12, 0, 0, 0, time.UTC), 961, 7, 23},
		{time.Date(1600, 2, 28, 12, 0, 0, 0, time.UTC), 978, 12, 9},
		{time.Date(1600, 2, 29, 12, 0, 0, 0, time.UTC), 978, 12, 10},
		{time.Date(1600, 3, 1, 12, 0, 0, 0, time.UTC), 978, 12, 11},
		{time.Date(1700, 2, 28, 12, 0, 0, 0, time.UTC), 1078, 12, 10},
		{time.Date(1700, 3, 1, 12, 0, 0, 0, time.UTC), 1078, 12, 11},
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), 1348, 10, 11},
		{time.Date(2024, 2, 28, 12, 0, 0, 0, time.UTC), 1402, 12, 9},
		{time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC), 1402, 12, 10},
		{time.Date(2024, 3, 1, 12, 0, 0, 0, time.UTC), 1402, 12, 11},
		{time.Date(2024, 3, 18, 12, 0, 0, 0, time.UTC), 1402, 12, 28},
		{time.Date(2024, 3, 19, 12, 0, 0, 0, time.UTC), 1402, 12, 29},
		{time.Date(2024, 3, 20, 12, 0, 0, 0, time.UTC), 1403, 1, 1},
		{time.Date(2025, 3, 19, 12, 0, 0, 0, time.UTC), 1403, 12, 29},
		{time.Date(2025, 3, 20, 12, 0, 0, 0, time.UTC), 1403, 12, 30},
		{time.Date(2025, 3, 21, 12, 0, 0, 0, time.UTC), 1404, 1, 1},
		{time.Date(2025, 4, 20, 12, 0, 0, 0, time.UTC), 1404, 1, 31},
		{time.Date(2025, 4, 21, 12, 0, 0, 0, time.UTC), 1404, 2, 1},
		{time.Date(2025, 8, 23, 12, 0, 0, 0, time.UTC), 1404, 6, 1},
		{time.Date(2025, 9, 22, 12, 0, 0, 0, time.UTC), 1404, 6, 31},
		{time.Date(2025, 9, 23, 12, 0, 0, 0, time.UTC), 1404, 7, 1},
		{time.Date(2025, 10, 22, 12, 0, 0, 0, time.UTC), 1404, 7, 30},
		{time.Date(2025, 12, 22, 12, 0, 0, 0, time.UTC), 1404, 10, 1},
		{time.Date(2026, 1, 20, 12, 0, 0, 0, time.UTC), 1404, 10, 30},
		{time.Date(2026, 1, 21, 12, 0, 0, 0, time.UTC), 1404, 11, 1},
		{time.Date(2026, 2, 19, 12, 0, 0, 0, time.UTC), 1404, 11, 30},
		{time.Date(2026, 7, 11, 16, 0, 0, 0, time.UTC), 1405, 4, 20},
		{time.Date(2100, 1, 1, 12, 0, 0, 0, time.UTC), 1478, 10, 12},
		{time.Date(2100, 3, 21, 12, 0, 0, 0, time.UTC), 1479, 1, 1},
		{time.Date(2100, 12, 31, 12, 0, 0, 0, time.UTC), 1479, 10, 10},
		{time.Date(2200, 1, 1, 12, 0, 0, 0, time.UTC), 1578, 10, 11},
		{time.Date(2200, 3, 21, 12, 0, 0, 0, time.UTC), 1579, 1, 1},
		{time.Date(2200, 12, 31, 12, 0, 0, 0, time.UTC), 1579, 10, 10},
	}

	for _, tc := range tests {
		got := FromGregorian(tc.t)
		if got.Year != tc.year || got.Month != tc.month || got.Day != tc.day {
			t.Errorf("got %d-%d-%d, want %d-%d-%d", got.Year, got.Month, got.Day, tc.year, tc.month, tc.day)
		}

		if tc.t.Year() >= 1097 {
			if got.Hour != tc.t.Hour() ||
				got.Minute != tc.t.Minute() ||
				got.Second != tc.t.Second() ||
				got.Nanosecond != tc.t.Nanosecond() ||
				got.Weekday != tc.t.Weekday() {
				t.Errorf("Time units not preserved for %s: got %+v", tc.t, got)
			}
		} else {
			if got.Hour != 0 ||
				got.Minute != 0 ||
				got.Second != 0 ||
				got.Nanosecond != 0 ||
				got.Weekday != time.Weekday(0) {
				t.Errorf("Zero time expected for %s: got %+v", tc.t, got)
			}
		}
	}
}

func BenchmarkFromGregorian(b *testing.B) {
	t := time.Date(2024, 3, 20, 12, 34, 56, 789, time.UTC)

	for b.Loop() {
		_ = FromGregorian(t)
	}
}
