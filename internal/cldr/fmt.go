package cldr

type CalendarDate struct {
	Year    int
	Month   int // 1-12
	Day     int
	Weekday int // 0-6 (Sunday-Saturday)
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
	FmtKindWeekday
)

type FmtItem struct {
	Digits   *Digits
	Months   *CalendarMonths
	Weekdays *CalendarWeekdays
	Text     string
	Kind     FmtKind
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
		case FmtKindWeekday:
			b = append(b, item.Weekdays[t.Weekday]...)
		}
	}

	return string(b)
}
