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

	if t.Options.Year != YearUnd {
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

// skipTest returns a reason to skip locale testing if formatting is not implemented yet.
// Returns empty string if all is fine.
func skipTest(locale language.Tag) string {
	return map[string]string{
		"ar-SA":       "islamic-umalqura calendar",
		"en-001":      "regional formatting",
		"en-150":      "regional formatting",
		"en-AE":       "regional formatting",
		"en-AG":       "regional formatting",
		"en-AI":       "regional formatting",
		"en-AT":       "regional formatting",
		"en-AU":       "regional formatting",
		"en-BB":       "regional formatting",
		"en-BE":       "regional formatting",
		"en-BI":       "regional formatting",
		"en-BM":       "regional formatting",
		"en-BS":       "regional formatting",
		"en-BW":       "regional formatting",
		"en-BZ":       "regional formatting",
		"en-CA":       "regional formatting",
		"en-CC":       "regional formatting",
		"en-CH":       "regional formatting",
		"en-CK":       "regional formatting",
		"en-CM":       "regional formatting",
		"en-CX":       "regional formatting",
		"en-CY":       "regional formatting",
		"en-DE":       "regional formatting",
		"en-DG":       "regional formatting",
		"en-DK":       "regional formatting",
		"en-DM":       "regional formatting",
		"en-ER":       "regional formatting",
		"en-FI":       "regional formatting",
		"en-FJ":       "regional formatting",
		"en-FK":       "regional formatting",
		"en-FM":       "regional formatting",
		"en-GB":       "regional formatting",
		"en-GD":       "regional formatting",
		"en-GG":       "regional formatting",
		"en-GH":       "regional formatting",
		"en-GI":       "regional formatting",
		"en-GM":       "regional formatting",
		"en-GU":       "regional formatting",
		"en-GY":       "regional formatting",
		"en-HK":       "regional formatting",
		"en-ID":       "regional formatting",
		"en-IE":       "regional formatting",
		"en-IL":       "regional formatting",
		"en-IM":       "regional formatting",
		"en-IN":       "regional formatting",
		"en-IO":       "regional formatting",
		"en-JE":       "regional formatting",
		"en-JM":       "regional formatting",
		"en-KE":       "regional formatting",
		"en-KI":       "regional formatting",
		"en-KN":       "regional formatting",
		"en-KY":       "regional formatting",
		"en-LC":       "regional formatting",
		"en-LR":       "regional formatting",
		"en-LS":       "regional formatting",
		"en-MG":       "regional formatting",
		"en-MS":       "regional formatting",
		"en-MO":       "regional formatting",
		"en-MT":       "regional formatting",
		"en-MU":       "regional formatting",
		"en-MV":       "regional formatting",
		"en-MW":       "regional formatting",
		"en-MY":       "regional formatting",
		"en-NA":       "regional formatting",
		"en-NF":       "regional formatting",
		"en-NG":       "regional formatting",
		"en-NL":       "regional formatting",
		"en-NR":       "regional formatting",
		"en-NU":       "regional formatting",
		"en-NZ":       "regional formatting",
		"en-PG":       "regional formatting",
		"en-PH":       "regional formatting",
		"en-PK":       "regional formatting",
		"en-PN":       "regional formatting",
		"en-PR":       "regional formatting",
		"en-PW":       "regional formatting",
		"en-RW":       "regional formatting",
		"en-SB":       "regional formatting",
		"en-SC":       "regional formatting",
		"en-SD":       "regional formatting",
		"en-SE":       "regional formatting",
		"en-SG":       "regional formatting",
		"en-SH":       "regional formatting",
		"en-SI":       "regional formatting",
		"en-SL":       "regional formatting",
		"en-SS":       "regional formatting",
		"en-SX":       "regional formatting",
		"en-SZ":       "regional formatting",
		"en-TC":       "regional formatting",
		"en-TK":       "regional formatting",
		"en-TO":       "regional formatting",
		"en-TT":       "regional formatting",
		"en-TV":       "regional formatting",
		"en-TZ":       "regional formatting",
		"en-UG":       "regional formatting",
		"en-VC":       "regional formatting",
		"en-VG":       "regional formatting",
		"en-VI":       "regional formatting",
		"en-VU":       "regional formatting",
		"en-WS":       "regional formatting",
		"en-ZA":       "regional formatting",
		"en-ZM":       "regional formatting",
		"en-ZW":       "regional formatting",
		"es-AR":       "regional formatting",
		"es-CL":       "regional formatting",
		"es-MX":       "regional formatting",
		"es-PA":       "regional formatting",
		"es-PR":       "regional formatting",
		"es-US":       "regional formatting",
		"ff-Adlm":     "regional formatting",
		"ff-Adlm-BF":  "regional formatting",
		"ff-Adlm-CM":  "regional formatting",
		"ff-Adlm-GH":  "regional formatting",
		"ff-Adlm-GM":  "regional formatting",
		"ff-Adlm-GN":  "regional formatting",
		"ff-Adlm-GW":  "regional formatting",
		"ff-Adlm-LR":  "regional formatting",
		"ff-Adlm-MR":  "regional formatting",
		"ff-Adlm-NE":  "regional formatting",
		"ff-Adlm-NG":  "regional formatting",
		"ff-Adlm-SL":  "regional formatting",
		"ff-Adlm-SN":  "regional formatting",
		"hi-Latn":     "regional formatting",
		"hi-Latn-IN":  "regional formatting",
		"hr-BA":       "regional formatting",
		"it-CH":       "regional formatting",
		"ks-Arab-IN":  "regional formatting",
		"ks-Deva":     "regional formatting",
		"ks-Deva-IN":  "regional formatting",
		"kxv-Deva":    "regional formatting",
		"kxv-Deva-IN": "regional formatting",
		"kxv-Orya":    "regional formatting",
		"kxv-Orya-IN": "regional formatting",
		"kxv-Telu":    "regional formatting",
		"kxv-Telu-IN": "regional formatting",
		"mn-Mong":     "regional formatting",
		"mn-Mong-CN":  "regional formatting",
		"mni-Beng-IN": "regional formatting",
		"mni-Mtei-IN": "regional formatting",
		"nl-BE":       "regional formatting",
		"pa-Arab":     "regional formatting",
		"pa-Arab-PK":  "regional formatting",
		"ps-PK":       "regional formatting",
		"sat-Deva-IN": "regional formatting",
		"sat-Olck-IN": "regional formatting",
		"sd-Arab-PK":  "regional formatting",
		"sd-Deva":     "regional formatting",
		"sd-Deva-IN":  "regional formatting",
		"se-FI":       "regional formatting",
		"sv-FI":       "regional formatting",
		"uz-Arab-AF":  "regional formatting",
		"uz-Cyrl":     "regional formatting",
		"uz-Cyrl-UZ":  "regional formatting",
		"vai-Latn":    "regional formatting",
		"vai-Latn-LR": "regional formatting",
		"yue-Hans":    "regional formatting",
		"yue-Hans-CN": "regional formatting",
		"yue-Hant-HK": "regional formatting",
		"zh-Hans-CN":  "regional formatting",
		"zh-Hans-HK":  "regional formatting",
		"zh-Hans-MO":  "regional formatting",
		"zh-Hans-SG":  "regional formatting",
		"zh-Hant":     "regional formatting",
		"zh-Hant-HK":  "regional formatting",
		"zh-Hant-MO":  "regional formatting",
		"zh-Hant-TW":  "regional formatting",
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
