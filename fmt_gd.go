package intl

import "golang.org/x/text/language"

func fmtEraDayGregorian(locale language.Tag, digits digits, opts Options) func(d int) string {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Day == Day2Digit

	prefix, suffix := era+" ", ""
	if withName {
		prefix, suffix = era+" ("+name+": ", ")"
	}

	switch lang {
	case hi:
		if script != latn {
			break
		}

		fallthrough
	case en:
		if !withName {
			prefix, suffix = "", " "+era
		}
	case bg, cy:
		prefix = era + " "
		if withName {
			prefix = era + " (" + name + ": "
		}
	case bs:
		if script == cyrl {
			break
		}

		fallthrough
	case cs, da, dsb, fo, hr, hsb, nb, nn, no, sk, sl:
		suffix = "."
		if withName {
			suffix = ".)"
		}
	case ja, ko, yue, zh:
		suffix = name
		if withName {
			suffix = name + ")"
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
	prefix, suffix := era+" ", ""
	withName := opts.Era == EraShort || opts.Era == EraLong && opts.Day == Day2Digit

	if withName {
		prefix, suffix = era+" ("+dayName(locale)+": ", ")"
	}

	if lang == fa {
		if withName {
			prefix = era + " (" + dayName(locale) + ": "
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
		name := dayName(locale)
		prefix, suffix = era+" ("+name+": ", ")"
	}

	return func(d int) string { return prefix + fmtDay(d, opts.Day) + suffix }
}
