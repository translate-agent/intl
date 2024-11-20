package intl

import "golang.org/x/text/language"

// TODO(jhorsts): temporary era formatting until [era] option is added.
// This allows to format era in year formats where it is required.
//
// [era]: https://github.com/translate-agent/intl/issues/25
func fmtEra(locale language.Tag) string {
	lang, _ := locale.Base()

	switch lang.String() {
	default:
		return "AP"
	case "th":
		return "พ.ศ."
	}
}
