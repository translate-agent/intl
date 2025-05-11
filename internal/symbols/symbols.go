package symbols

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

type Symbol byte

//nolint:asciicheck,revive
const (
	TxtSpace            Symbol = ' '  // " "
	TxtNumberSign       Symbol = '#'  // "#"
	TxtApostrophe       Symbol = '\'' // "'"
	TxtLeftParenthesis  Symbol = '('  // "("
	TxtRightParenthesis Symbol = ')'  // ")"
	TxtComma            Symbol = ','  // ","
	TxtHyphenMinus      Symbol = '-'  // "-"
	TxtFullStop         Symbol = '.'  // "."
	TxtSolidus          Symbol = '/'  // "/"
	TxtColon            Symbol = ':'  // ":"
	Txta                Symbol = 'a'  // "a"
	Txtm                Symbol = 'm'  // "m"

	Txt日     Symbol = iota + 128 // "日"
	Txt일                         // "일"
	Txtꑍ                         // "ꑍ"
	Txt年                         // "年"
	Txt月                         // "月"
	Txt년                         // "년"
	Txtс                         // "с"
	Txtҫ                         // "ҫ"
	Txtж                         // "ж"
	Txtթ                         // "թ"
	TxtNNBSP                     // " "
	Txtр                         // "р"
	Txt00                        // "г."
	Txt01                        // ". g."
	Txt02                        // "\u200f/"
	Txt03                        // "ꆪ-"
	Txt04                        // "tháng "
	Txt05                        // "ел"
	Txt06                        // "སྤྱི་ཟླ་"
	Txt07                        // "de"
	Txt08                        // "urteko"
	Txt09                        // "an"
	Txt10                        // "оны"

	symbolStart  // the start of CLDR symbols
	MonthUnit    // "month" in the local language
	DayUnit      // "day" in the local language
	Symbol_G     // G, era, abbreviated
	Symbol_GGGG  // GGGG, era, long
	Symbol_GGGGG // GGGGG, era, narrow
	Symbol_y     // y, year
	Symbol_yy    // yy, two-digit year
	Symbol_M     // M, month
	Symbol_MM    // MM, two-digit month
	Symbol_d     // d, day
	Symbol_dd    // dd, two-digit day
	Symbol_LLL   // LLL, stand-alone abbreviated
	Symbol_LLLLL // LLLLL, stand-alone narrow
	Symbol_MMM   // MMM, format abbreviated
	Symbol_MMMMM // MMMMM, format narrow
)

func (s Symbol) String() string {
	switch s {
	default:
		return string(s)
	case TxtNNBSP:
		return " "
	case Txt日:
		return "日"
	case Txt일:
		return "일"
	case Txtꑍ:
		return "ꑍ"
	case Txt年:
		return "年"
	case Txt月:
		return "月"
	case Txt년:
		return "년"
	case Txtҫ:
		return "ҫ"
	case Txtс:
		return "с"
	case Txtж:
		return "ж"
	case Txtթ:
		return "թ"
	case Txtр:
		return "р"
	case Txt00:
		return "г."
	case Txt01:
		return ". g."
	case Txt02:
		return "\u200f/"
	case Txt03:
		return "ꆪ-"
	case Txt04:
		return "tháng "
	case Txt05:
		return "ел"
	case Txt06:
		return "སྤྱི་ཟླ་"
	case Txt07:
		return "de"
	case Txt08:
		return "urteko"
	case Txt09:
		return "an"
	case Txt10:
		return "оны"
	}
}

type Seq struct {
	locale  language.Tag
	symbols []Symbol
}

func NewSeq(locale language.Tag) *Seq {
	return &Seq{locale: locale}
}

// Add appends one or more [Symbol] to the [Seq].
func (s *Seq) Add(symbol ...Symbol) *Seq {
	s.symbols = append(s.symbols, symbol...)

	return s
}

// AddSeq appends another [Seq].
func (s *Seq) AddSeq(seq *Seq) *Seq {
	s.symbols = append(s.symbols, seq.symbols...)

	return s
}

// Func returns [time.Time] formatting function.
func (s *Seq) Func() func(cldr.TimeReader) string {
	// concatenate all sequential text values
	var text cldr.Text

	digits := cldr.LocaleDigits(s.locale)
	fmt := make(cldr.Fmt, 0, len(s.symbols))

	for _, symbol := range s.symbols {
		if symbol < symbolStart {
			text += cldr.Text(symbol.String())
			continue
		}

		var symFmt cldr.FmtFunc

		//nolint:exhaustive
		switch symbol {
		case Symbol_G:
			symFmt = cldr.Text(cldr.EraName(s.locale)[1])
		case Symbol_GGGG:
			symFmt = cldr.Text(cldr.EraName(s.locale)[2])
		case Symbol_GGGGG:
			symFmt = cldr.Text(cldr.EraName(s.locale)[0])
		case Symbol_y:
			symFmt = cldr.YearNumeric(digits)
		case Symbol_yy:
			symFmt = cldr.YearTwoDigit(digits)
		case Symbol_M:
			symFmt = cldr.MonthNumeric(digits)
		case Symbol_MM:
			symFmt = cldr.MonthTwoDigit(digits)
		case Symbol_MMM:
			names := cldr.MonthNames(s.locale.String(), "format", "abbreviated")
			symFmt = cldr.Month(names)
		case Symbol_LLLLL:
			names := cldr.MonthNames(s.locale.String(), "stand-alone", "narrow")
			symFmt = cldr.Month(names)
		case Symbol_LLL:
			names := cldr.MonthNames(s.locale.String(), "stand-alone", "abbreviated")
			symFmt = cldr.Month(names)
		case Symbol_d:
			symFmt = cldr.DayNumeric(digits)
		case Symbol_dd:
			symFmt = cldr.DayTwoDigit(digits)
		case MonthUnit:
			symFmt = cldr.Text(cldr.UnitName(s.locale).Month)
		case DayUnit:
			symFmt = cldr.Text(cldr.UnitName(s.locale).Day)
		}

		if f, ok := symFmt.(cldr.Text); ok {
			text += f
			continue
		}

		if text != "" {
			fmt = append(fmt, text)
			text = ""
		}

		fmt = append(fmt, symFmt)
	}

	if text != "" {
		fmt = append(fmt, text)
	}

	return fmt.Format
}
