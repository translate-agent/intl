package intl

import "golang.org/x/text/language"

func fmtYearGregorian(locale language.Tag) func(y string) string {
	switch lang, _ := locale.Base(); lang {
	default:
		return func(y string) string { return y }
	case bg:
		return func(y string) string { return y + " г." }
	case bs, hr, hu, sr:
		return func(y string) string { return y + "." }
	case ja, yue, zh:
		return func(y string) string { return y + "年" }
	case ko:
		return func(y string) string { return y + "년" }
	case lv:
		return func(y string) string { return y + ". g." }
	}
}

func fmtYearBuddhist(locale language.Tag) func(y string) string {
	return func(y string) string { return fmtEra(locale, EraNarrow) + " " + y }
}

func fmtYearPersian(locale language.Tag) func(y string) string {
	lang, _, region := locale.Raw()

	switch lang {
	case fa:
		return func(y string) string { return y }
	case uz:
		if region == regionAF {
			return func(y string) string { return y }
		}

		fallthrough
	default:
		return func(y string) string { return fmtEra(locale, EraNarrow) + " " + y }
	}
}
