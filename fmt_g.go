package intl

import "golang.org/x/text/language"

// TODO(jhorsts): temporary era formatting until [era] option is added.
// This allows to format era in year formats where it is required.
//
// [era]: https://github.com/translate-agent/intl/issues/25
func fmtEra(locale language.Tag, opt Era) string {
	f := func(v era) string {
		switch opt {
		default:
			return v.narrow
		case EraLong:
			return v.long
		case EraShort:
			return v.short
		}
	}

	era, ok := eraLookup[locale.String()]
	if ok {
		return f(era)
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		if era, ok := eraLookup[lang.String()+"-"+script.String()]; ok {
			return f(era)
		}
	}

	if era, ok := eraLookup[lang.String()]; ok {
		return f(era)
	}

	return "CE"
}
