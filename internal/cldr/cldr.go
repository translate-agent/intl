package cldr

import "golang.org/x/text/language"

func MonthNames(locale string, context, width string) CalendarMonths {
	indexes := findMonthIndexes(locale)

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
	eraVal, okVal := findEra(locale.String())
	if okVal {
		return eraVal
	}

	lang, _ := locale.Base()

	if script, confidence := locale.Script(); confidence == language.Exact {
		eraVal, okVal = findEra(lang.String() + "-" + script.String())
		if okVal {
			return eraVal
		}
	}

	eraVal, okVal = findEra(lang.String())
	if okVal {
		return eraVal
	}

	eraVal, _ = findEra("aa")

	return eraVal
}

func MonthNamesPtr(locale string, context, width string) *CalendarMonths {
	indexes := findMonthIndexes(locale)

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
		if v := int(indexes[i]); v >= 0 && v < len(CalendarMonthNames) { // isInBounds()
			return &CalendarMonthNames[v]
		}
	}

	return &CalendarMonthNames[0]
}

type Fields struct {
	Month, Day string
}

type fieldsEntry struct {
	locale string
	fields Fields
}

func findFields(locale string) (Fields, bool) {
	low, high := 0, len(fieldsLookup)-1

	for low <= high {
		mid := int(uint(low+high) >> 1)
		midVal := &fieldsLookup[mid]

		switch {
		case midVal.locale < locale:
			low = mid + 1
		case midVal.locale > locale:
			high = mid - 1
		default:
			return midVal.fields, true
		}
	}

	return Fields{}, false
}

// Era contains era names: 0 - narrow, 1 - short, 2 - long.
// The index is the value of iota [Era].
type Era [3]string

type eraEntry struct {
	locale string
	era    Era
}

func findEra(locale string) (Era, bool) {
	low, high := 0, len(eraLookup)-1

	for low <= high {
		mid := int(uint(low+high) >> 1)
		midVal := &eraLookup[mid]

		switch {
		case midVal.locale < locale:
			low = mid + 1
		case midVal.locale > locale:
			high = mid - 1
		default:
			return midVal.era, true
		}
	}

	return Era{}, false
}

type CalendarMonths [12]string

// monthIndexes contains indexes of months names in Gregorian calendar, it has 6 indexes
// for all variations of "width" and "context".
//
//	0 - abbreviated, format
//	1 - abbreviated, stand-alone
//	2 - wide, format
//	3 - wide, stand-alone
//	4 - narrow, format
//	5 - narrow, stand-alone
type monthIndexes [6]int16

type monthEntry struct {
	locale  string
	indexes monthIndexes
}

func findMonthIndexes(locale string) monthIndexes {
	low, high := 0, len(monthLookup)-1

	for low <= high {
		mid := int(uint(low+high) >> 1)
		midVal := &monthLookup[mid]

		switch {
		case midVal.locale < locale:
			low = mid + 1
		case midVal.locale > locale:
			high = mid - 1
		default:
			return midVal.indexes
		}
	}

	return monthIndexes{}
}

type NumberingSystem int
