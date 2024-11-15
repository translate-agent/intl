package main

import (
	"fmt"
	"log/slog"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_buildFmtMD(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		in  [4]string
		out string
	}{
		{
			in: [4]string{"d/M", "", "", "dd-MM"},
			out: `if opts.Month == MonthNumeric && opts.Day == DayNumeric {
	return fmtDay(d, DayNumeric) + "/" + fmtMonth(m, MonthNumeric)
}
return fmtDay(d, cmp.Or(opts.Day, Day2Digit)) + "-" + fmtMonth(m, cmp.Or(opts.Month, Month2Digit))`,
		},
		{
			in: [4]string{"d.M.", "", "", "d. M."},
			out: `if opts.Month == MonthNumeric && opts.Day == DayNumeric {
	return fmtDay(d, DayNumeric) + "." + fmtMonth(m, MonthNumeric) + "."
}
return fmtDay(d, DayNumeric) + ". " + fmtMonth(m, MonthNumeric) + "."`,
		},
		{
			in:  [4]string{"d.MM", "d.MM", "dd.MM", "dd.MM"},
			out: `return fmtDay(d, cmp.Or(opts.Day, Day2Digit)) + "." + fmtMonth(m, Month2Digit)`,
		},
		{
			in: [4]string{"MMMMM/dd", "MMMMM/dd", "MMMMM/dd", "MMMMM/dd"},
			out: `fmtMonth = fmtMonthName(locale.String(), "stand-alone", "narrow")
return fmtMonth(m, opts.Month) + "/" + fmtDay(d, Day2Digit)`,
		},
		{
			in: [4]string{"MM-dd", "d/MM", "dd/M", "dd/MM"},
			out: `if opts.Month == MonthNumeric && opts.Day == DayNumeric {
	return fmtMonth(m, Month2Digit) + "-" + fmtDay(d, Day2Digit)
}
return fmtDay(d, cmp.Or(opts.Day, Day2Digit)) + "/" + fmtMonth(m, cmp.Or(opts.Month, Month2Digit))`,
		},
		{ // wae-CH
			in: [4]string{"d. MMM", "d. MMM", "dd. MMM", "dd. MMM"},
			out: `fmtMonth = fmtMonthName(locale.String(), "stand-alone", "abbreviated")
return fmtDay(d, cmp.Or(opts.Day, DayNumeric)) + ". " + fmtMonth(m, opts.Month)`,
		},
	} {
		t.Run(fmt.Sprintf("%+v", test.in), func(t *testing.T) {
			t.Parallel()

			got := buildFmtMD(test.in[0], test.in[1], test.in[2], test.in[3], slog.Default())

			if got != test.out {
				t.Errorf("\nwant %q\ngot  %q", test.out, got)
			}
		})
	}
}

func Test_yearMonthPatterns(t *testing.T) {
	t.Parallel()

	for _, test := range []struct{ in, out [3]string }{
		{ // th
			in:  [3]string{"", "", "M/y"},
			out: [3]string{"M/y", "MM/y", "M/y"},
		},
		{
			in:  [3]string{"LL-y", "", ""},
			out: [3]string{"MM-y", "MM-y", "MM-y"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", test.in), func(t *testing.T) {
			t.Parallel()

			ym, ymm, yyyym := yearMonthPatterns(test.in[0], test.in[1], test.in[2])

			if v := cmp.Diff(ym, parseDatePattern(test.out[0])); v != "" {
				t.Errorf("\nwant yM %v\ngot      %v", test.out[0], ym)
			}

			if v := cmp.Diff(ymm, parseDatePattern(test.out[1])); v != "" {
				t.Errorf("\nwant yMM %v\ngot        %v", test.out[1], ymm)
			}

			if v := cmp.Diff(yyyym, parseDatePattern(test.out[2])); v != "" {
				t.Errorf("\nwant yyyyM %v\ngot         %v", test.out[2], yyyym)
			}
		})
	}
}

func Test_monthDayPatterns(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		in, out [4]string
	}{
		{
			in:  [4]string{"d/M", "", "", ""},
			out: [4]string{"d/M", "d/MM", "dd/M", "dd/MM"},
		},
		{
			in:  [4]string{"d/M", "", "", "dd-MM"},
			out: [4]string{"d/M", "d-MM", "dd-M", "dd-MM"},
		},
		{ // af-ZA
			in:  [4]string{"dd-MM", "", "", ""},
			out: [4]string{"dd-MM", "dd-MM", "dd-MM", "dd-MM"},
		},
		{ // om-ET
			in:  [4]string{"MM-dd", "", "", "dd/MM"},
			out: [4]string{"MM-dd", "d/MM", "dd/M", "dd/MM"},
		},
		{
			in:  [4]string{"d.M.", "", "", "d. M."},
			out: [4]string{"d.M.", "d. M.", "d. M.", "d. M."},
		},
		{ // bg-BG
			in:  [4]string{"d.MM", "", "", ""},
			out: [4]string{"d.MM", "d.MM", "dd.MM", "dd.MM"},
		},
		{ // sv-SE
			in:  [4]string{"d/M", "d/M", "", "dd/MM"},
			out: [4]string{"d/M", "d/M", "dd/M", "dd/MM"},
		},
		{ // sv
			in:  [4]string{"", "", "", "d/M"},
			out: [4]string{"d/M", "d/M", "d/M", "d/M"},
		},
		{ // wae-CH
			in:  [4]string{"d. MMM", "", "", ""},
			out: [4]string{"d. MMM", "d. MMM", "dd. MMM", "dd. MMM"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", test.in), func(t *testing.T) {
			t.Parallel()

			md, mmd, mdd, mmdd := monthDayPatterns(test.in[0], test.in[1], test.in[2], test.in[3])

			if v := cmp.Diff(md, parseDatePattern(test.out[0])); v != "" {
				t.Errorf("\nwant md %v\ngot      %v", test.out[0], md)
			}

			if v := cmp.Diff(mmd, parseDatePattern(test.out[1])); v != "" {
				t.Errorf("\nwant mmd %v\ngot      %v", test.out[1], mmd)
			}

			if v := cmp.Diff(mdd, parseDatePattern(test.out[2])); v != "" {
				t.Errorf("\nwant mdd %v\ngot      %v", test.out[2], mdd)
			}

			if v := cmp.Diff(mmdd, parseDatePattern(test.out[3])); v != "" {
				t.Errorf("\nwant mmdd %v\ngot      %v", test.out[3], mmdd)
			}
		})
	}
}

func Test_GroupLayouts(t *testing.T) {
	t.Parallel()

	layout := func(id, fmt string) Layout {
		return Layout{ID: parseDatePattern(id), Pattern: parseDatePattern(fmt)}
	}

	for _, test := range []struct {
		name string
		in   []Layout
		out  []LayoutGroup
	}{{
		name: "same pattern, different formatting",
		in:   []Layout{layout("Md", "d/M"), layout("MMd", "d/MM"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
		out: []LayoutGroup{
			{
				Layouts: []Layout{layout("Md", "d/M"), layout("MMd", "d/MM"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
			},
		},
	}, {
		name: "3 same patterns and other",
		in:   []Layout{layout("Md", "d/M"), layout("MMd", "d-MM"), layout("Mdd", "dd-M"), layout("MMdd", "dd-MM")},
		out: []LayoutGroup{
			{
				Expr:         "opts.Month == MonthNumeric && opts.Day == DayNumeric",
				FmtTypeMonth: FmtTypeNumericOnly,
				FmtTypeDay:   FmtTypeNumericOnly,
				Layouts:      []Layout{layout("Md", "d/M")},
			},
			{
				Layouts: []Layout{layout("MMd", "d-MM"), layout("Mdd", "dd-M"), layout("MMdd", "dd-MM")},
			},
		},
	}, {
		name: "3 same patterns and other with different formatting",
		in:   []Layout{layout("Md", "MM-dd"), layout("MMd", "d/MM"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
		out: []LayoutGroup{
			{
				Expr:         "opts.Month == MonthNumeric && opts.Day == DayNumeric",
				FmtTypeMonth: FmtType2DigitOnly,
				FmtTypeDay:   FmtType2DigitOnly,
				Layouts:      []Layout{layout("Md", "MM-dd")},
			},
			{
				Layouts: []Layout{layout("MMd", "d/MM"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
			},
		},
	}, {
		name: "3 same patterns with different formatting and 1 other",
		in:   []Layout{layout("Md", "d.M."), layout("MMd", "d. M."), layout("Mdd", "d. M."), layout("MMdd", "d. M.")},
		out: []LayoutGroup{
			{
				Expr:         "opts.Month == MonthNumeric && opts.Day == DayNumeric",
				FmtTypeMonth: FmtTypeNumericOnly,
				FmtTypeDay:   FmtTypeNumericOnly,
				Layouts:      []Layout{layout("Md", "d.M.")},
			},
			{
				FmtTypeMonth: FmtTypeNumericOnly,
				FmtTypeDay:   FmtTypeNumericOnly,
				Layouts:      []Layout{layout("MMd", "d. M."), layout("Mdd", "d. M."), layout("MMdd", "d. M.")},
			},
		},
	}, { // sv-SE
		in: []Layout{layout("Md", "d/M"), layout("MMd", "d/M"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
		out: []LayoutGroup{
			{
				Expr:         "opts.Month == Month2Digit && opts.Day == DayNumeric",
				FmtTypeMonth: FmtTypeNumericOnly,
				FmtTypeDay:   FmtTypeNumericOnly,
				Layouts:      []Layout{layout("MMd", "d/M")},
			},
			{
				Layouts: []Layout{layout("Md", "d/M"), layout("Mdd", "dd/M"), layout("MMdd", "dd/MM")},
			},
		},
	}} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			t.Log(test.in)

			got := groupLayouts(test.in...)

			if diff := cmp.Diff(test.out, got); diff != "" {
				t.Errorf("GroupLayouts() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_ParseDatePattern(t *testing.T) {
	t.Parallel()

	cases := []struct {
		pattern  string
		elements []DatePatternItem
	}{
		{
			pattern:  "",
			elements: []DatePatternItem{},
		},
		{
			pattern: "d",
			elements: []DatePatternItem{
				{Value: "d", Literal: false},
			},
		},
		{
			pattern: "'", // NOTE(mvilks): should be invalid pattern but we don't care
			elements: []DatePatternItem{
				{Value: "'", Literal: true},
			},
		},
		{
			pattern: "yMd",
			elements: []DatePatternItem{
				{Value: "y", Literal: false},
				{Value: "M", Literal: false},
				{Value: "d", Literal: false},
			},
		},
		{
			pattern: "yyyy.MM.dd.",
			elements: []DatePatternItem{
				{Value: "yyyy", Literal: false},
				{Value: ".", Literal: true},
				{Value: "MM", Literal: false},
				{Value: ".", Literal: true},
				{Value: "dd", Literal: false},
				{Value: ".", Literal: true},
			},
		},
		{
			pattern: "y. 'g.' M",
			elements: []DatePatternItem{
				{Value: "y", Literal: false},
				{Value: ". g. ", Literal: true},
				{Value: "M", Literal: false},
			},
		},
		{
			pattern: "d''M''y. 'a''b' E",
			elements: []DatePatternItem{
				{Value: "d", Literal: false},
				{Value: "'", Literal: true},
				{Value: "M", Literal: false},
				{Value: "'", Literal: true},
				{Value: "y", Literal: false},
				{Value: ". a'b ", Literal: true},
				{Value: "E", Literal: false},
			},
		},
		{
			pattern: "G y. 'gada' d. MMM, E – G y. 'gada' d. MMM, E",
			elements: []DatePatternItem{
				{Value: "G", Literal: false},
				{Value: " ", Literal: true},
				{Value: "y", Literal: false},
				{Value: ". gada ", Literal: true},
				{Value: "d", Literal: false},
				{Value: ". ", Literal: true},
				{Value: "MMM", Literal: false},
				{Value: ", ", Literal: true},
				{Value: "E", Literal: false},
				{Value: " – ", Literal: true},
				{Value: "G", Literal: false},
				{Value: " ", Literal: true},
				{Value: "y", Literal: false},
				{Value: ". gada ", Literal: true},
				{Value: "d", Literal: false},
				{Value: ". ", Literal: true},
				{Value: "MMM", Literal: false},
				{Value: ", ", Literal: true},
				{Value: "E", Literal: false},
			},
		},
		{
			pattern: "EEEE, 'ngày' dd MMMM 'năm' y G",
			elements: []DatePatternItem{
				{Value: "EEEE", Literal: false},
				{Value: ", ngày ", Literal: true},
				{Value: "dd", Literal: false},
				{Value: " ", Literal: true},
				{Value: "MMMM", Literal: false},
				{Value: " năm ", Literal: true},
				{Value: "y", Literal: false},
				{Value: " ", Literal: true},
				{Value: "G", Literal: false},
			},
		},
		{
			pattern: "'Ngày' dd",
			elements: []DatePatternItem{
				{Value: "Ngày ", Literal: true},
				{Value: "dd", Literal: false},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.pattern, func(t *testing.T) {
			t.Parallel()

			elems := parseDatePattern(test.pattern)
			if !slices.Equal(elems, test.elements) {
				t.Errorf("want %v, got %v", test.elements, elems)
			}
		})
	}
}
