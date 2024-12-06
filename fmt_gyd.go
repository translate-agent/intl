package intl

import "golang.org/x/text/language"

type EraYearDay int

const (
	eraYearDay EraYearDay = iota
	eraDayYear
)

func fmtEraYearDayGregorian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, script, _ := locale.Raw()
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	prefix, middle, suffix := era+" ", " ("+name+": ", ")"
	layout := eraYearDay

	switch lang {
	case af, am, ar, ast, blo, bn, br, ca, ccp, ceb, chr:
		prefix = ""
		middle = " " + era + " (" + name + ": "
	case be:
		prefix = ""
		middle = " г. " + era + " (" + name + ": "
		suffix = ")"
	case cv:
		prefix = ""
		middle = " ҫ. " + era + " (" + name + ": "
		suffix = ")"
	case bg, cy:
		prefix = ""
		middle = " " + era + " (" + name + ": "
	case brx:
		prefix = era
	case bs:
		prefix = ""
		middle = " " + era + " (" + name + ": "

		if script != cyrl {
			suffix = ".)"
		}
	case ckb:
		prefix = era + " "
		middle = " (" + name + ": "
		suffix = ")"
	case cs:
		prefix = ""
		middle = " " + era + " (" + name + ": "
		suffix = ".)"
	}

	switch layout {
	default: // eraYearDay
		return func(y, d int) string {
			return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d, opts.Day) + suffix
		}
	case eraDayYear:
		return func(y, d int) string {
			return prefix + fmtDay(d, opts.Day) + middle + layoutYear(fmtYear(y, opts.Year)) + suffix
		}
	}
}

func fmtEraYearDayPersian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	era := fmtEra(locale, opts.Era)
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDay(digits)
	name := dayName(locale)
	prefix, middle, suffix := era+" ", " ("+name+": ", ")"

	return func(y, d int) string {
		return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtDay(d, opts.Day) + suffix
	}
}
