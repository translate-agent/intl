package intl

import "golang.org/x/text/language"

func fmtEraDayGregorian(locale language.Tag, digits digits, opts Options) func(d int) string {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtDay := fmtDay(digits)
	dayName := unitName(locale).Day
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Day == Day2Digit

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + dayName + ": "
		suffix = ")"
	}

	switch lang {
	case hi:
		if script != latn {
			break
		}

		fallthrough
	case en:
		if !withName {
			prefix = ""
			suffix = " " + era
		}
	case bg, cy:
		if withName {
			prefix = era + " (" + dayName + ": "
		} else {
			prefix = era + " "
		}
	case bs:
		if script == cyrl {
			break
		}

		fallthrough
	case cs, da, dsb, fo, hr, hsb, nb, nn, no, sk, sl:
		if withName {
			suffix = ".)"
		} else {
			suffix = "."
		}
	case ja, ko, yue, zh:
		if withName {
			suffix = dayName + ")"
		} else {
			suffix = dayName
		}
	case lt:
		opts.Day = Day2Digit
	}

	return func(d int) string { return prefix + fmtDay(d, opts.Day) + suffix }
}

func fmtEraDayPersian(locale language.Tag, digits digits, opts Options) func(d int) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtDay := fmtDay(digits)
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Day == Day2Digit

	prefix := era + " "
	suffix := ""

	if withName {
		prefix = era + " (" + unitName(locale).Day + ": "
		suffix = ")"
	}

	if lang == fa {
		if withName {
			prefix = era + " (" + unitName(locale).Day + ": "
		} else {
			prefix = era + " "
		}
	}

	return func(d int) string { return prefix + fmtDay(d, opts.Day) + suffix }
}

func fmtEraDayBuddhist(locale language.Tag, digits digits, opts Options) func(d int) string {
	era := fmtEra(locale, opts.Era)
	fmtDay := fmtDay(digits)
	prefix, suffix := era+" ", ""
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Day == Day2Digit

	if withName {
		prefix, suffix = era+" ("+unitName(locale).Day+": ", ")"
	}

	return func(d int) string { return prefix + fmtDay(d, opts.Day) + suffix }
}
