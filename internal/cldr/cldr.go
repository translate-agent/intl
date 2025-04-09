package cldr

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
