package intl

import "golang.org/x/text/language"

func fmtEraDayGregorian(locale language.Tag, digits digits, opts Options) func(d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtDay := fmtDay(digits)
	name := dayName(locale)

	switch lang {
	default:
		return func(d int) string { return "?" }
	case lv:
		// era=long,day=numeric,out=mūsu ērā 2
		// era=long,day=2-digit,out=mūsu ērā (diena: 02)
		// era=short,day=numeric,out=m.ē. (diena: 2)
		// era=short,day=2-digit,out=m.ē. (diena: 02)
		// era=narrow,day=numeric,out=m.ē. 2
		// era=narrow,day=2-digit,out=m.ē. 02
		if opts.Era == EraShort ||
			opts.Era == EraLong && opts.Day == Day2Digit {
			return func(d int) string { return era + " (" + name + ": " + fmtDay(d, opts.Day) + ")" }
		}

		return func(d int) string { return era + " " + fmtDay(d, opts.Day) }
	}
}
