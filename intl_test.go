package intl

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"
)

type Test struct {
	Output  string
	Options Options
}

type Tests struct {
	Date  time.Time               `json:"date"`
	Tests map[language.Tag][]Test `json:"tests"`
}

func (t *Test) UnmarshalJSON(b []byte) error {
	var data [2]any // first options, second formatted output

	if err := json.Unmarshal(b, &data); err != nil {
		return fmt.Errorf("unmarshal test data: %w", err)
	}

	out, ok := data[1].(string)
	if !ok {
		panic("want formatted string value")
	}

	test := Test{Output: out}

	if o, ok := data[0].(map[string]any); ok {
		if v, ok := o["year"].(string); ok {
			switch v {
			case "numeric":
				test.Options.Year = YearNumeric
			case "2-digit":
				test.Options.Year = Year2Digit
			}
		}

		if v, ok := o["day"].(string); ok {
			switch v {
			case "numeric":
				test.Options.Day = DayNumeric
			case "2-digit":
				test.Options.Day = Day2Digit
			}
		}

		if v, ok := o["month"].(string); ok {
			switch v {
			case "numeric":
				test.Options.Month = MonthNumeric
			case "2-digit":
				test.Options.Month = Month2Digit
			}
		}
	}

	*t = test

	return nil
}

//go:embed tests.json
var data []byte

// skipTest returns true for locales where formatting is hard to be determined for given cases.
func skipTest(locale language.Tag, options Options) bool {
	type key struct {
		locale  string
		options Options
	}

	_, ok := map[key]struct{}{
		// CLDR stipulates arabext numbering. Why Node.js uses latn?
		{"bgn-PK", Options{Year: Year2Digit}}:    {},
		{"bgn-PK", Options{Year: YearNumeric}}:   {},
		{"bgn-PK", Options{Day: Day2Digit}}:      {},
		{"bgn-PK", Options{Day: DayNumeric}}:     {},
		{"bgn-PK", Options{Month: MonthNumeric}}: {},
		{"bgn-PK", Options{Month: Month2Digit}}:  {},

		// CLDR stipulates hmnr numbering. Why Node.js uses latn?
		{"hnj-Hmnp", Options{Year: Year2Digit}}:    {},
		{"hnj-Hmnp", Options{Year: YearNumeric}}:   {},
		{"hnj-Hmnp", Options{Day: Day2Digit}}:      {},
		{"hnj-Hmnp", Options{Day: DayNumeric}}:     {},
		{"hnj-Hmnp", Options{Month: MonthNumeric}}: {},
		{"hnj-Hmnp", Options{Month: Month2Digit}}:  {},

		{"sdh-IR", Options{Year: Year2Digit}}:    {},
		{"sdh-IR", Options{Year: YearNumeric}}:   {},
		{"sdh-IR", Options{Month: Month2Digit}}:  {},
		{"sdh-IR", Options{Month: MonthNumeric}}: {},
		{"sdh-IR", Options{Day: Day2Digit}}:      {},
		{"sdh-IR", Options{Day: DayNumeric}}:     {},

		// depends on localised era
		{"th-TH", Options{Year: Year2Digit}}:  {},
		{"th-TH", Options{Year: YearNumeric}}: {},
	}[key{locale.String(), options}]

	return ok
}

func TestDateTime_Format(t *testing.T) {
	t.Parallel()

	var tests Tests

	if err := json.Unmarshal(data, &tests); err != nil {
		panic(err)
	}

	if len(tests.Tests) == 0 {
		t.Error("no tests found")
	}

	for locale, cases := range tests.Tests {
		t.Run(locale.String(), func(t *testing.T) {
			t.Parallel()

			for _, test := range cases {
				t.Run(fmt.Sprintf("%+v: %s", test.Options, test.Output), func(t *testing.T) {
					t.Parallel()

					if skipTest(locale, test.Options) {
						t.Skip()
					}

					got := NewDateTimeFormat(locale, test.Options).Format(tests.Date)

					// replace space with non-breaking space. Latest CLDR uses non-breaking space.
					if strings.ContainsRune(got, ' ') {
						test.Output = strings.ReplaceAll(test.Output, " ", " ")
					}

					if test.Output != got {
						t.Errorf("want '%s', got '%s'", test.Output, got)
						t.Logf("\n%v\n%v", []rune(test.Output), []rune(got))
					}
				})
			}
		})
	}
}

var locales = []string{
	"fa-IR", // persian calendar, arabext numerals
	"lv-LV", // gregorian calendar, latn numerals
	"dz-BT", // gregorian calendar, tibt numerals
}

func BenchmarkNewDateTime(b *testing.B) {
	var v *DateTimeFormat

	for _, s := range locales {
		locale := language.MustParse(s)

		b.Run(s, func(b *testing.B) {
			for range b.N {
				v = NewDateTimeFormat(locale, Options{})
			}
		})
	}

	runtime.KeepAlive(v)
}

func BenchmarkDateTime_Format(b *testing.B) {
	var v1, v2, v3, v4, v5, v6, v7 string

	now := time.Now()

	for _, s := range locales {
		locale := language.MustParse(s)
		f1 := NewDateTimeFormat(locale, Options{}).Format
		f2 := NewDateTimeFormat(locale, Options{Year: YearNumeric}).Format
		f3 := NewDateTimeFormat(locale, Options{Year: Year2Digit}).Format
		f4 := NewDateTimeFormat(locale, Options{Month: MonthNumeric}).Format
		f5 := NewDateTimeFormat(locale, Options{Month: Month2Digit}).Format
		f6 := NewDateTimeFormat(locale, Options{Day: DayNumeric}).Format
		f7 := NewDateTimeFormat(locale, Options{Day: Day2Digit}).Format

		b.Run(s, func(b *testing.B) {
			for range b.N {
				v1, v2, v3, v4, v5, v6, v7 = f1(now), f2(now), f3(now), f4(now), f5(now), f6(now), f7(now)
			}
		})
	}

	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(v3)
	runtime.KeepAlive(v4)
	runtime.KeepAlive(v5)
	runtime.KeepAlive(v6)
	runtime.KeepAlive(v7)
}
