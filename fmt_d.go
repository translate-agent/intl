package intl

import "golang.org/x/text/language"

func fmtDayGregorian(locale language.Tag, digits digits) func(v int, opt Day) string {
	fmt := fmtDay(digits)

	switch lang, _ := locale.Base(); lang.String() {
	default:
		return fmt
	case "bs", "cs", "da", "dsb", "fo", "hr", "hsb", "nb", "nn", "no", "sk", "sl":
		return func(v int, opt Day) string { return fmt(v, opt) + "." }
	case "ja", "yue", "zh":
		return func(v int, opt Day) string { return fmt(v, opt) + "日" }
	case "ko":
		return func(v int, opt Day) string { return fmt(v, opt) + "일" }
	case "lt":
		return func(v int, _ Day) string { return fmt(v, Day2Digit) }
	}
}

func fmtDayBuddhist(_ language.Tag, digits digits) func(v int, opt Day) string {
	return fmtDay(digits)
}

func fmtDayPersian(_ language.Tag, digits digits) func(v int, opt Day) string {
	return fmtDay(digits)
}
