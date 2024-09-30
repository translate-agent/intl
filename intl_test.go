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
		{"hnj-Hmnp", Options{Year: Year2Digit}}:  {},
		{"hnj-Hmnp", Options{Year: YearNumeric}}: {},
		{"lrc-IR", Options{Year: Year2Digit}}:    {},
		{"lrc-IR", Options{Year: YearNumeric}}:   {},
		{"mzn-IR", Options{Year: Year2Digit}}:    {},
		{"mzn-IR", Options{Year: YearNumeric}}:   {},
		{"nb", Options{Day: Day2Digit}}:          {},
		{"nb-NO", Options{Day: Day2Digit}}:       {},
		{"nb-NO", Options{Day: DayNumeric}}:      {},
		{"nn-NO", Options{Day: Day2Digit}}:       {},
		{"nn-NO", Options{Day: DayNumeric}}:      {},
		{"ps-AF", Options{Year: Year2Digit}}:     {},
		{"ps-AF", Options{Year: YearNumeric}}:    {},
		{"prg-PL", Options{Year: Year2Digit}}:    {},
		{"prg-PL", Options{Year: YearNumeric}}:   {},
		{"sdh-IR", Options{Year: Year2Digit}}:    {},
		{"sdh-IR", Options{Year: YearNumeric}}:   {},
		{"th-TH", Options{Year: Year2Digit}}:     {},
		{"th-TH", Options{Year: YearNumeric}}:    {},
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

func BenchmarkNewDateTime(b *testing.B) {
	locale := language.MustParse("fa-IR")

	var v *DateTimeFormat

	for range b.N {
		v = NewDateTimeFormat(locale, Options{})
	}

	runtime.KeepAlive(v)
}

func BenchmarkDateTime_Format(b *testing.B) {
	locale := language.MustParse("fa-IR")
	f1 := NewDateTimeFormat(locale, Options{}).Format
	f2 := NewDateTimeFormat(locale, Options{Year: YearNumeric}).Format
	f3 := NewDateTimeFormat(locale, Options{Year: Year2Digit}).Format
	f4 := NewDateTimeFormat(locale, Options{Day: DayNumeric}).Format
	f5 := NewDateTimeFormat(locale, Options{Day: Day2Digit}).Format
	now := time.Now()

	var v1, v2, v3, v4, v5 string

	for range b.N {
		v1, v2, v3, v4, v5 = f1(now), f2(now), f3(now), f4(now), f5(now)
	}

	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(v3)
	runtime.KeepAlive(v4)
	runtime.KeepAlive(v5)
}
