package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"golang.org/x/text/language"
)

func fmtEra(locale language.Tag, opt Era) string {
	if opt.und() {
		opt = EraNarrow
	}

	era, ok := cldr.EraLookup[locale.String()]
	if ok && opt > 0 && int(opt) <= len(era) { // isInBounds()
		return era[opt-1]
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		era, ok := cldr.EraLookup[lang.String()+"-"+script.String()]
		if ok && opt > 0 && int(opt) <= len(era) {
			return era[opt-1]
		}
	}

	if era, ok := cldr.EraLookup[lang.String()]; ok && opt > 0 && int(opt) <= len(era) {
		return era[opt-1]
	}

	return "CE"
}
