package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtYearGregorian(locale language.Tag, digits digits, opt Year) fmtFunc {
	var suffix string

	switch lang, _ := locale.Base(); lang {
	case bg, mk:
		suffix = " г."
	case bs, hr, hu, sr:
		suffix = "."
	case ja, yue, zh:
		suffix = "年"
	case ko:
		suffix = "년"
	case lv:
		suffix = ". g."
	}

	yearDigits := convertYearDigits(digits, opt)

	return func(t time.Time) string { return yearDigits(t) + suffix }
}

func fmtYearBuddhist(locale language.Tag, digits digits, opts Options) fmtFunc {
	prefix := fmtEra(locale, opts.Era) + " "
	yearDigits := convertYearDigits(digits, opts.Year)

	return func(t time.Time) string { return prefix + yearDigits(t) }
}

func fmtYearPersian(locale language.Tag) func(y string) string {
	lang, _, region := locale.Raw()
	prefix := ""

	if lang != fa && (lang != uz || region != regionAF) {
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(y string) string { return prefix + y }
}
