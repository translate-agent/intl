package intl

import "golang.org/x/text/language"

func fmtYearGregorian(locale language.Tag) func(v string) string {
	switch lang, _ := locale.Base(); lang.String() {
	default:
		return func(v string) string { return v }
	case "bg":
		return func(v string) string { return v + " г." }
	case "bs", "hr", "hu", "sr":
		return func(v string) string { return v + "." }
	case "ja", "yue", "zh":
		return func(v string) string { return v + "年" }
	case "ko":
		return func(v string) string { return v + "년" }
	case "lv":
		return func(v string) string { return v + ". g." }
	}
}

func fmtYearBuddhist(locale language.Tag) func(v string) string {
	return func(v string) string { return fmtEra(locale) + " " + v }
}

func fmtYearPersian(locale language.Tag) func(v string) string {
	if lang, _ := locale.Base(); lang.String() == "fa" {
		return func(v string) string { return v }
	}

	return func(v string) string { return fmtEra(locale) + " " + v }
}
