package intl

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"
	"slices"
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
	}

	*t = test

	return nil
}

//go:embed datetime.json
var data []byte

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
		if slices.Contains([]string{
			"hnj-Hmnp",
			"lrc-IR",
			"mzn-IR",
			"ps-AF",
			"prg-PL",
			"sdh-IR",
			"th-TH",
		}, locale.String()) {
			continue
		}

		t.Run(locale.String(), func(t *testing.T) {
			t.Parallel()

			for _, test := range cases {
				t.Run(fmt.Sprintf("%v: %s", test.Options, test.Output), func(t *testing.T) {
					t.Parallel()

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
	f := NewDateTimeFormat(locale, Options{})
	now := time.Now()

	var v string

	for range b.N {
		v = f.Format(now)
	}

	runtime.KeepAlive(v)
}
