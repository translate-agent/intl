package cldr

import "golang.org/x/text/language"

func MonthNames(locale string, context, width string) CalendarMonths {
	indexes := MonthLookup[locale]

	var i int

	// "abbreviated" width index is 0
	switch width {
	case "wide":
		i += 2 // 1*2
	case "narrow":
		i += 4 // 2*2
	}

	// "format" context index is 0
	if context == "stand-alone" {
		i++
	}

	if i >= 0 && i < len(indexes) { // isInBounds()
		if v := int(indexes[i]); v > 0 && v < len(CalendarMonthNames) { // isInBounds()
			return CalendarMonthNames[v]
		}
	}

	return CalendarMonths{}
}

func EraName(locale language.Tag) Era {
	era, ok := EraLookup[locale.String()]
	if ok {
		return era
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		era, ok := EraLookup[lang.String()+"-"+script.String()]
		if ok {
			return era
		}
	}

	if era, ok := EraLookup[lang.String()]; ok {
		return era
	}

	return EraLookup["aa"]
}
