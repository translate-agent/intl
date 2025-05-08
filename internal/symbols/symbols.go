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
	Txt00                        // "г."
	Txt01                        // ". g."
	Txt02                        // "\u200f/"
	Txt03                        // "ꆪ-"
	Txt04                        // "tháng "
	Txt05                        // "ел"

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
		return ""
	case TxtFullStop:
		return "."
	case TxtComma:
		return ","
	case TxtSpace:
		return " "
	case TxtNNBSP:
		return " "
	case TxtSolidus:
		return "/"
	case TxtHyphenMinus:
		return "-"
	case TxtColon:
		return ":"
	case TxtLeftParenthesis:
		return "("
	case TxtRightParenthesis:
		return ")"
	case TxtNumberSign:
		return "#"
	case Txtm:
		return "m"
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
	}
}

type Seq struct {
	locale  language.Tag
	symbols []Symbol
}

func NewSeq(locale language.Tag) *Seq {
	return &Seq{locale: locale}
}

func (s *Seq) Add(symbol ...Symbol) *Seq {
	s.symbols = append(s.symbols, symbol...)

	return s
}

func (s *Seq) AddSeq(seq *Seq) *Seq {
	s.symbols = append(s.symbols, seq.symbols...)

	return s
}

// Func returns [time.Time] formatting function.
func (s *Seq) Func() func(cldr.TimeReader) string {
	digits := cldr.LocaleDigits(s.locale)
	fmt := make(cldr.Fmt, 0, len(s.symbols))

	for _, symbol := range s.symbols {
		if symbol < symbolStart {
			fmt = append(fmt, cldr.Text(symbol.String()))
			continue
		}

		//nolint:exhaustive
		switch symbol {
		case Symbol_G:
			fmt = append(fmt, cldr.Text(cldr.EraName(s.locale)[1]))
		case Symbol_GGGG:
			fmt = append(fmt, cldr.Text(cldr.EraName(s.locale)[2]))
		case Symbol_GGGGG:
			fmt = append(fmt, cldr.Text(cldr.EraName(s.locale)[0]))
		case Symbol_y:
			fmt = append(fmt, cldr.YearNumeric(digits))
		case Symbol_yy:
			fmt = append(fmt, cldr.YearTwoDigit(digits))
		case Symbol_M:
			fmt = append(fmt, cldr.MonthNumeric(digits))
		case Symbol_MM:
			fmt = append(fmt, cldr.MonthTwoDigit(digits))
		case Symbol_MMM:
			names := cldr.MonthNames(s.locale.String(), "format", "abbreviated")
			fmt = append(fmt, cldr.Month(names))
		case Symbol_LLLLL:
			names := cldr.MonthNames(s.locale.String(), "stand-alone", "narrow")
			fmt = append(fmt, cldr.Month(names))
		case Symbol_LLL:
			names := cldr.MonthNames(s.locale.String(), "stand-alone", "abbreviated")
			fmt = append(fmt, cldr.Month(names))
		case Symbol_d:
			fmt = append(fmt, cldr.DayNumeric(digits))
		case Symbol_dd:
			fmt = append(fmt, cldr.DayTwoDigit(digits))
		case MonthUnit:
			fmt = append(fmt, cldr.Text(cldr.UnitName(s.locale).Month))
		case DayUnit:
			fmt = append(fmt, cldr.Text(cldr.UnitName(s.locale).Day))
		}
	}

	return fmt.Format
}
