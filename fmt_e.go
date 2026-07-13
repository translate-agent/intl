package intl

import (
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

func seqWeekday(locale language.Tag, opt Weekday) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqWeekdayBuddhist(locale language.Tag, opt Weekday) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}

func seqWeekdayPersian(locale language.Tag, opt Weekday) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opt.symbol())
}
