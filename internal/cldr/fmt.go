package cldr

type CalendarDate struct {
	Year  int
	Month int // 1-12
	Day   int
}

type FmtKind int

const (
	FmtKindText FmtKind = iota
	FmtKindYearNumeric
	FmtKindYearTwoDigit
	FmtKindMonthNumeric
	FmtKindMonthTwoDigit
	FmtKindMonth
	FmtKindDayNumeric
	FmtKindDayTwoDigit
)

type FmtItem struct {
	Digits *Digits
	Months *CalendarMonths
	Text   string
	Kind   FmtKind
}

type Fmt []FmtItem

func (f Fmt) Format(t CalendarDate) string {
	var buf [64]byte

	b := buf[:0]

	for i := range f {
		item := &f[i]
		switch item.Kind {
		case FmtKindText:
			b = append(b, item.Text...)
		case FmtKindYearNumeric:
			b = item.Digits.appendNumeric(b, t.Year)
		case FmtKindYearTwoDigit:
			b = item.Digits.appendTwoDigit(b, t.Year)
		case FmtKindMonthNumeric:
			b = item.Digits.appendNumeric(b, t.Month)
		case FmtKindMonthTwoDigit:
			b = item.Digits.appendTwoDigit(b, t.Month)
		case FmtKindMonth:
			b = append(b, item.Months[t.Month-1]...)
		case FmtKindDayNumeric:
			b = item.Digits.appendNumeric(b, t.Day)
		case FmtKindDayTwoDigit:
			b = item.Digits.appendTwoDigit(b, t.Day)
		}
	}

	return string(b)
}
