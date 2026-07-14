package cldr_test

import (
	"testing"

	"go.expect.digital/intl/internal/cldr"
)

func TestWeekdayNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		locale   string
		context  string
		width    string
		expected cldr.CalendarWeekdays
	}{
		{
			locale:  "en",
			context: "format",
			width:   "wide",
			expected: cldr.CalendarWeekdays{
				"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
			},
		},
		{
			locale:  "en",
			context: "format",
			width:   "abbreviated",
			expected: cldr.CalendarWeekdays{
				"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
			},
		},
		{
			locale:  "lv",
			context: "format",
			width:   "wide",
			expected: cldr.CalendarWeekdays{
				"svētdiena", "pirmdiena", "otrdiena", "trešdiena", "ceturtdiena", "piektdiena", "sestdiena",
			},
		},
		{
			locale:  "lv",
			context: "format",
			width:   "abbreviated",
			expected: cldr.CalendarWeekdays{
				"svētd.", "pirmd.", "otrd.", "trešd.", "ceturtd.", "piektd.", "sestd.",
			},
		},
		{
			locale:  "nonexistent-locale",
			context: "format",
			width:   "wide",
			expected: cldr.CalendarWeekdays{
				"", "", "", "", "", "", "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.locale+"-"+tc.context+"-"+tc.width, func(t *testing.T) {
			t.Parallel()

			got := cldr.WeekdayNames(tc.locale, tc.context, tc.width)
			if got != tc.expected {
				t.Errorf("WeekdayNames(%q, %q, %q) = %v; expected %v", tc.locale, tc.context, tc.width, got, tc.expected)
			}
		})
	}
}

func TestWeekdayNamesPtr(t *testing.T) {
	t.Parallel()

	got := cldr.WeekdayNamesPtr("en", "format", "wide")
	expected := cldr.CalendarWeekdays{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}

	if *got != expected {
		t.Errorf("WeekdayNamesPtr(%q, %q, %q) = %v; expected %v", "en", "format", "wide", *got, expected)
	}

	// Non-existent fallback test
	fallbackGot := cldr.WeekdayNamesPtr("nonexistent-locale", "format", "wide")
	if fallbackGot == nil {
		t.Errorf("expected non-nil pointer for nonexistent-locale")
	}
}
