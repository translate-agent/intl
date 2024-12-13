package intl

import "golang.org/x/text/language"

// TODO(jhorsts): temporary era formatting until [era] option is added.
// This allows to format era in year formats where it is required.
//
// [era]: https://github.com/translate-agent/intl/issues/25
func fmtEra(locale language.Tag, opt Era) string {
	era, ok := eraLookup[locale.String()]
	if ok && opt > 0 && int(opt) <= len(era) { // isInBounds()
		return era[opt-1]
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		era, ok := eraLookup[lang.String()+"-"+script.String()]
		if ok && opt > 0 && int(opt) <= len(era) {
			return era[opt-1]
		}
	}

	if era, ok := eraLookup[lang.String()]; ok && opt > 0 && int(opt) <= len(era) {
		return era[opt-1]
	}

	return "CE"
}
