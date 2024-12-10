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

func (t *Test) String() string {
	var sb strings.Builder

	if t.Options.Era != EraUnd {
		sb.WriteString("era=")
		sb.WriteString(t.Options.Era.String())
	}

	if t.Options.Year != YearUnd {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}

		sb.WriteString("year=")
		sb.WriteString(t.Options.Year.String())
	}

	if t.Options.Month != MonthUnd {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}

		sb.WriteString("month=")
		sb.WriteString(t.Options.Month.String())
	}

	if t.Options.Day != DayUnd {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}

		sb.WriteString("day=")
		sb.WriteString(t.Options.Day.String())
	}

	if sb.Len() > 0 {
		sb.WriteRune(',')
	}

	sb.WriteString("out=" + t.Output)

	return sb.String()
}

type Tests []Test

func (t Tests) String() string {
	var sb strings.Builder

	for _, test := range t {
		sb.WriteString(test.String() + "\n")
	}

	return sb.String()
}

type AllTests struct {
	Date  time.Time              `json:"date"`
	Tests map[language.Tag]Tests `json:"tests"`
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
		if v, ok := o["era"].(string); ok {
			test.Options.Era = MustParseEra(v)
		}

		if v, ok := o["year"].(string); ok {
			test.Options.Year = MustParseYear(v)
		}

		if v, ok := o["month"].(string); ok {
			test.Options.Month = MustParseMonth(v)
		}

		if v, ok := o["day"].(string); ok {
			test.Options.Day = MustParseDay(v)
		}
	}

	*t = test

	return nil
}

//go:embed tests.json
var data []byte

// skipTest returns a reason to skip locale testing if formatting is not implemented yet.
// Returns empty string if all is fine.
func skipTest(locale language.Tag) string {
	return map[string]string{
		"ar-SA": "islamic-umalqura calendar",
	}[locale.String()]
}

func TestDateTime_Format(t *testing.T) {
	t.Parallel()

	var tests AllTests

	if err := json.Unmarshal(data, &tests); err != nil {
		panic(err)
	}

	if len(tests.Tests) == 0 {
		t.Error("no tests found")
	}

	for locale, cases := range tests.Tests {
		t.Run(locale.String(), func(t *testing.T) {
			t.Parallel()

			t.Logf("calendar type: %s", defaultCalendar(locale))
			t.Logf("cases:\n%s", cases)

			for _, test := range cases {
				if reason := skipTest(locale); reason != "" {
					t.Skip(reason)
				}

				got := NewDateTimeFormat(locale, test.Options).Format(tests.Date)

				// replace space with non-breaking space. Latest CLDR uses non-breaking space.
				if strings.ContainsRune(got, ' ') {
					test.Output = strings.ReplaceAll(test.Output, " ", " ")
				}

				if test.Output != got {
					t.Errorf("%s\nwant '%s'\ngot  '%s'", test.String(), test.Output, got)
					t.Logf("\n%v\n%v\n", []rune(test.Output), []rune(got))
				}
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
	var v DateTimeFormat

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
	var v1, v2, v3, v4, v5 string

	now := time.Now()

	for _, s := range locales {
		locale := language.MustParse(s)
		f1 := NewDateTimeFormat(locale, Options{Era: EraLong}).Format
		f2 := NewDateTimeFormat(locale, Options{Year: Year2Digit}).Format
		f3 := NewDateTimeFormat(locale, Options{Month: Month2Digit}).Format
		f4 := NewDateTimeFormat(locale, Options{Day: Day2Digit}).Format
		f5 := NewDateTimeFormat(locale, Options{Era: EraShort, Year: YearNumeric, Month: MonthNumeric, Day: DayNumeric}).
			Format

		b.Run(s, func(b *testing.B) {
			for range b.N {
				v1, v2, v3, v4, v5 = f1(now), f2(now), f3(now), f4(now), f5(now)
			}
		})
	}

	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(v3)
	runtime.KeepAlive(v4)
	runtime.KeepAlive(v5)
}
