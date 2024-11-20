package intl

import "golang.org/x/text/language"

func fmtDayGregorian(locale language.Tag, digits digits) func(d int, opt Day) string {
	fmt := fmtDay(digits)

	switch lang, _ := locale.Base(); lang {
	default:
		return fmt
	case bs:
		if script, _ := locale.Script(); script == cyrl {
			return fmt
		}

		return func(d int, opt Day) string { return fmt(d, opt) + "." }
	case cs, da, dsb, fo, hr, hsb, nb, nn, no, sk, sl:
		return func(d int, opt Day) string { return fmt(d, opt) + "." }
	case ja, yue, zh:
		return func(d int, opt Day) string { return fmt(d, opt) + "日" }
	case ko:
		return func(d int, opt Day) string { return fmt(d, opt) + "일" }
	case lt:
		return func(d int, _ Day) string { return fmt(d, Day2Digit) }
	}
}

func fmtDayBuddhist(_ language.Tag, digits digits) func(d int, opt Day) string {
	return fmtDay(digits)
}

func fmtDayPersian(_ language.Tag, digits digits) func(d int, opt Day) string {
	return fmtDay(digits)
}
