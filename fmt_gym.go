package intl

import (
	"time"

	"golang.org/x/text/language"
)

type EraYearMonth int

const (
	// eraYearMonth includes "era year month" and "year month era".
	eraYearMonth EraYearMonth = iota
	// eraMonthYear includes "era month year" and "month year era".
	eraMonthYear
)

func fmtEraYearMonthGregorian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	lang, _ := locale.Base()
	era := fmtEra(locale, opts.Era)
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearGregorian(locale)
	fmtMonth := fmtMonth(digits)
	prefix, middle, suffix := era+" ", " ", ""
	layout := eraYearMonth

	switch lang {
	case af, am, ar, ast, blo, bn, br, bs, ca, ccp, ceb, chr, cs, cy:
		layout = eraMonthYear
		prefix = ""
		suffix = " " + era
	case az:
		layout = eraMonthYear
	case be:
		layout = eraMonthYear
		prefix = ""
		suffix = " г. " + era
	case bg:
		layout = eraMonthYear
		prefix = ""
		middle = "."
		suffix = " " + era

		if opts.Month == MonthNumeric {
			opts.Month = Month2Digit
		}
	case cv:
		layout = eraMonthYear
		prefix = ""
		suffix = " ҫ. " + era
	}

	switch layout {
	default: // eraYearMonth
		return func(y int, m time.Month) string {
			return prefix + layoutYear(fmtYear(y, opts.Year)) + middle + fmtMonth(m, opts.Month) + suffix
		}
	case eraMonthYear:
		return func(y int, m time.Month) string {
			return prefix + fmtMonth(m, opts.Month) + middle + layoutYear(fmtYear(y, opts.Year)) + suffix
		}
	}
}

func fmtEraYearMonthPersian(locale language.Tag, digits digits, opts Options) func(y int, m time.Month) string {
	fmtYear := fmtYear(digits)
	layoutYear := fmtYearPersian(locale)
	fmtMonth := fmtMonth(digits)

	return func(y int, m time.Month) string {
		return layoutYear(fmtYear(y, opts.Year)) + " " + fmtMonth(m, opts.Month)
	}
}
