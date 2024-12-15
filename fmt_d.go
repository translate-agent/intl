package intl

import "golang.org/x/text/language"

func fmtDayGregorian(locale language.Tag, digits digits) func(d int, opt Day) string {
	fmt := fmtDay(digits)

	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case bs:
		if script, _ := locale.Script(); script == cyrl {
			return fmt
		}

		suffix = "."
	case cs, da, dsb, fo, hr, hsb, ie, nb, nn, no, sk, sl:
		suffix = "."
	case ja, yue, zh:
		suffix = "日"
	case ko:
		suffix = "일"
	case lt:
		return func(d int, _ Day) string { return fmt(d, Day2Digit) }
	case ii:
		suffix = "ꑍ"
	}

	return func(d int, opt Day) string { return fmt(d, opt) + suffix }
}

func fmtDayBuddhist(_ language.Tag, digits digits) func(d int, opt Day) string {
	return fmtDay(digits)
}

func fmtDayPersian(_ language.Tag, digits digits) func(d int, opt Day) string {
	return fmtDay(digits)
}
