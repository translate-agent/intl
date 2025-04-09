package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtYearGregorian(locale language.Tag, digits cldr.Digits, opt Year) fmtFunc {
	var suffix string

	switch lang, _ := locale.Base(); lang {
	case cldr.BG, cldr.MK:
		suffix = " г."
	case cldr.BS, cldr.HR, cldr.HU, cldr.SR:
		suffix = "."
	case cldr.JA, cldr.YUE, cldr.ZH:
		suffix = "年"
	case cldr.KO:
		suffix = "년"
	case cldr.LV:
		suffix = ". g."
	}

	yearDigits := convertYearDigits(digits, opt)

	return func(t timeReader) string { return yearDigits(t) + suffix }
}

func fmtYearBuddhist(locale language.Tag, digits cldr.Digits, opts Options) fmtFunc {
	prefix := fmtEra(locale, opts.Era) + " "
	yearDigits := convertYearDigits(digits, opts.Year)

	return func(t timeReader) string { return prefix + yearDigits(t) }
}

func fmtYearPersian(locale language.Tag) func(y string) string {
	lang, _, region := locale.Raw()
	prefix := ""

	if lang != cldr.FA && (lang != cldr.UZ || region != cldr.RegionAF) {
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(y string) string { return prefix + y }
}
