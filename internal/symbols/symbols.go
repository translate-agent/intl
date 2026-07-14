package symbols

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

type Symbol byte

//nolint:asciicheck,revive
const (
	Txt日                    Symbol = iota + 128 // "日"
	Txt일                                        // "일"
	Txtꑍ                                        // "ꑍ"
	Txt年                                        // "年"
	Txt月                                        // "月"
	Txt년                                        // "년"
	Txtс                                        // "с"
	Txtҫ                                        // "ҫ"
	Txtж                                        // "ж"
	Txtթ                                        // "թ"
	TxtNNBSP                                    // " "
	Txtр                                        // "р"
	Txt00                                       // "г."
	Txt01                                       // ". g."
	Txt02                                       // "\u200f/"
	Txt03                                       // "ꆪ-"
	Txt04                                       // "tháng "
	Txt05                                       // "ел"
	Txt06                                       // "སྤྱི་ཟླ་"
	Txt08                                       // "urteko"
	Txt09                                       // "an"
	Txt10                                       // "оны"
	TxtArabicComma                              // "،"
	TxtCyrillicShortI                           // "й"
	TxtKurdishHam                               // "ھەم"
	TxtPersianOrdinalSuffix                     // "م"
	TxtHebrewHeDash                             // " ה-"
	TxtColognianDa                              // " dä "
	TxtBurmeseDay                               // " ရက် "
	TxtVietnameseNgay                           // ", ngày "
	TxtYiddishDem                               // " דעם "
	TxtYiddishTn                                // "טן"
	TxtDanishDen                                // " den "
	symbolStart                                 // the start of CLDR symbols
	MonthUnit                                   // "month" in the local language
	DayUnit                                     // "day" in the local language
	Symbol_G                                    // G, era, abbreviated
	Symbol_GGGG                                 // GGGG, era, long
	Symbol_GGGGG                                // GGGGG, era, narrow
	Symbol_y                                    // y, year
	Symbol_yy                                   // yy, two-digit year
	Symbol_M                                    // M, month
	Symbol_MM                                   // MM, two-digit month
	Symbol_d                                    // d, day
	Symbol_dd                                   // dd, two-digit day
	Symbol_LLL                                  // LLL, stand-alone abbreviated
	Symbol_LLLLL                                // LLLLL, stand-alone narrow
	Symbol_MMM                                  // MMM, format abbreviated
	Symbol_MMMMM                                // MMMMM, format narrow
	Symbol_E                                    // E, format abbreviated
	Symbol_EEEE                                 // EEEE, format wide
	Symbol_EEEEE                                // EEEEE, format narrow
	Symbol_EEEEEE                               // EEEEEE, format short
	Symbol_ccc                                  // ccc, stand-alone abbreviated
	Symbol_cccc                                 // cccc, stand-alone wide
	Symbol_ccccc                                // ccccc, stand-alone narrow
	Symbol_cccccc                               // cccccc, stand-alone short
)

//nolint:cyclop
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
	case Txt08:
		return "urteko"
	case Txt09:
		return "an"
	case Txt10:
		return "оны"
	case TxtArabicComma:
		return "،"
	case TxtCyrillicShortI:
		return "й"
	case TxtKurdishHam:
		return "ھەم"
	case TxtPersianOrdinalSuffix:
		return "م"
	case TxtHebrewHeDash:
		return " ה-"
	case TxtColognianDa:
		return " dä "
	case TxtBurmeseDay:
		return " ရက် "
	case TxtVietnameseNgay:
		return ", ngày "
	case TxtYiddishDem:
		return " דעם "
	case TxtYiddishTn:
		return "טן"
	case TxtDanishDen:
		return " den "
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

// Fmt returns the [cldr.Fmt] formatting sequence.
func (s *Seq) Fmt() cldr.Fmt {
	var text string

	digits := cldr.LocaleDigitsPtr(s.locale)
	fmt := make(cldr.Fmt, 0, len(s.symbols))

	for _, symbol := range s.symbols {
		if symbol < symbolStart {
			text += symbol.String()
			continue
		}

		var (
			item   cldr.FmtItem
			isText bool
		)

		//nolint:exhaustive
		switch symbol {
		case Symbol_G:
			text += cldr.EraName(s.locale)[1]
			isText = true
		case Symbol_GGGG:
			text += cldr.EraName(s.locale)[2]
			isText = true
		case Symbol_GGGGG:
			text += cldr.EraName(s.locale)[0]
			isText = true
		case Symbol_y:
			item = cldr.FmtItem{Kind: cldr.FmtKindYearNumeric, Digits: digits}
		case Symbol_yy:
			item = cldr.FmtItem{Kind: cldr.FmtKindYearTwoDigit, Digits: digits}
		case Symbol_M:
			item = cldr.FmtItem{Kind: cldr.FmtKindMonthNumeric, Digits: digits}
		case Symbol_MM:
			item = cldr.FmtItem{Kind: cldr.FmtKindMonthTwoDigit, Digits: digits}
		case Symbol_MMM:
			names := cldr.MonthNamesPtr(s.locale.String(), "format", "abbreviated")
			item = cldr.FmtItem{Kind: cldr.FmtKindMonth, Months: names}
		case Symbol_LLLLL:
			names := cldr.MonthNamesPtr(s.locale.String(), "stand-alone", "narrow")
			item = cldr.FmtItem{Kind: cldr.FmtKindMonth, Months: names}
		case Symbol_LLL:
			names := cldr.MonthNamesPtr(s.locale.String(), "stand-alone", "abbreviated")
			item = cldr.FmtItem{Kind: cldr.FmtKindMonth, Months: names}
		case Symbol_d:
			item = cldr.FmtItem{Kind: cldr.FmtKindDayNumeric, Digits: digits}
		case Symbol_dd:
			item = cldr.FmtItem{Kind: cldr.FmtKindDayTwoDigit, Digits: digits}
		case Symbol_E:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "format", "abbreviated")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_EEEE:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "format", "wide")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_EEEEE:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "format", "narrow")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_EEEEEE:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "format", "short")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_ccc:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "stand-alone", "abbreviated")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_cccc:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "stand-alone", "wide")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_ccccc:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "stand-alone", "narrow")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case Symbol_cccccc:
			names := cldr.WeekdayNamesPtr(s.locale.String(), "stand-alone", "short")
			item = cldr.FmtItem{Kind: cldr.FmtKindWeekday, Weekdays: names}
		case MonthUnit:
			text += cldr.UnitName(s.locale).Month
			isText = true
		case DayUnit:
			text += cldr.UnitName(s.locale).Day
			isText = true
		}

		if isText {
			continue
		}

		if text != "" {
			fmt = append(fmt, cldr.FmtItem{Kind: cldr.FmtKindText, Text: text})
			text = ""
		}

		fmt = append(fmt, item)
	}

	if text != "" {
		fmt = append(fmt, cldr.FmtItem{Kind: cldr.FmtKindText, Text: text})
	}

	return fmt
}
