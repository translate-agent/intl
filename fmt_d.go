package intl

import "golang.org/x/text/language"

func fmtDayGregorian(locale language.Tag, digits digits, opt Day) func(d int) string {
	suffix := ""

	switch lang, _ := locale.Base(); lang {
	case bs:
		if script, _ := locale.Script(); script == cyrl {
			break
		}

		suffix = "."
	case cs, da, dsb, fo, hr, hsb, ie, nb, nn, no, sk, sl:
		suffix = "."
	case ja, yue, zh:
		suffix = "日"
	case ko:
		suffix = "일"
	case lt:
		opt = Day2Digit
	case ii:
		suffix = "ꑍ"
	}

	day := fmtDay(digits, opt)

	return func(d int) string { return day(d) + suffix }
}

func fmtDayBuddhist(_ language.Tag, digits digits, opt Day) func(d int) string {
	return fmtDay(digits, opt)
}

func fmtDayPersian(_ language.Tag, digits digits, opt Day) func(d int) string {
	return fmtDay(digits, opt)
}
