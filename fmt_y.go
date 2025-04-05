package intl

import "golang.org/x/text/language"

func fmtYearGregorian(locale language.Tag) func(y string) string {
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

	return func(y string) string { return y + suffix }
}

func fmtYearBuddhist(locale language.Tag, era Era) func(y string) string {
	prefix := fmtEra(locale, era) + " "
	return func(y string) string { return prefix + y }
}

func fmtYearPersian(locale language.Tag) func(y string) string {
	lang, _, region := locale.Raw()
	prefix := ""

	if lang != fa && (lang != uz || region != regionAF) {
		prefix = fmtEra(locale, EraNarrow) + " "
	}

	return func(y string) string { return prefix + y }
}
