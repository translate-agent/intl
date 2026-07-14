package cldr

import (
	"strings"

	"golang.org/x/text/language"
)

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

func resolveWeekdayIndex(indexes weekdayIndexes, context, width string) int {
	var preferred []int

	switch width {
	case "abbreviated":
		if context == "format" {
			preferred = []int{0, 1, 4, 5}
		} else {
			preferred = []int{1, 0, 5, 4}
		}
	case "short":
		if context == "format" {
			preferred = []int{2, 3, 0, 1, 4, 5}
		} else {
			preferred = []int{3, 2, 1, 0, 5, 4}
		}
	case "wide":
		if context == "format" {
			preferred = []int{4, 5}
		} else {
			preferred = []int{5, 4}
		}
	case "narrow":
		if context == "format" {
			preferred = []int{6, 7, 0, 1, 4, 5}
		} else {
			preferred = []int{7, 6, 1, 0, 5, 4}
		}
	}

	for _, idx := range preferred {
		if idx >= 0 && idx < len(indexes) {
			if v := int(indexes[idx]); v > 0 && v < len(CalendarWeekdayNames) {
				return v
			}
		}
	}

	return 0
}

func WeekdayNames(locale string, context, width string) CalendarWeekdays {
	indexes := findWeekdayIndexes(locale)

	v := resolveWeekdayIndex(indexes, context, width)
	if v > 0 {
		return CalendarWeekdayNames[v]
	}

	return CalendarWeekdays{}
}

func WeekdayNamesPtr(locale string, context, width string) *CalendarWeekdays {
	indexes := findWeekdayIndexes(locale)

	v := resolveWeekdayIndex(indexes, context, width)
	if v > 0 {
		return &CalendarWeekdayNames[v]
	}

	return &CalendarWeekdayNames[0]
}

type Fields struct {
	Month, Day string
}

type fieldsEntry struct {
	locale string
	fields Fields
}

func findLookup[T any](locale string, list []T, getLocale func(*T) string) (*T, bool) {
	for {
		low, high := 0, len(list)-1

		for low <= high {
			mid := low + (high-low)/2 //nolint:mnd
			midVal := &list[mid]

			loc := getLocale(midVal)

			switch {
			case loc < locale:
				low = mid + 1
			case loc > locale:
				high = mid - 1
			default:
				return midVal, true
			}
		}

		idx := strings.LastIndexAny(locale, "-_")
		if idx == -1 {
			break
		}

		locale = locale[:idx]
	}

	return nil, false
}

func findFields(locale string) (Fields, bool) {
	if entry, ok := findLookup(locale, fieldsLookup[:], func(e *fieldsEntry) string { return e.locale }); ok {
		return entry.fields, true
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
	if entry, ok := findLookup(locale, eraLookup[:], func(e *eraEntry) string { return e.locale }); ok {
		return entry.era, true
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
	if entry, ok := findLookup(locale, monthLookup[:], func(e *monthEntry) string { return e.locale }); ok {
		return entry.indexes
	}

	return monthIndexes{}
}

type CalendarWeekdays [7]string

// weekdayIndexes contains indexes of weekdays names in Gregorian calendar, it has 8 indexes
// for all variations of "width" and "context".
//
//	0 - abbreviated, format
//	1 - abbreviated, stand-alone
//	2 - short, format
//	3 - short, stand-alone
//	4 - wide, format
//	5 - wide, stand-alone
//	6 - narrow, format
//	7 - narrow, stand-alone
type weekdayIndexes [8]int16

type weekdayEntry struct {
	locale  string
	indexes weekdayIndexes
}

func findWeekdayIndexes(locale string) weekdayIndexes {
	if entry, ok := findLookup(locale, weekdayLookup[:], func(e *weekdayEntry) string { return e.locale }); ok {
		return entry.indexes
	}

	return weekdayIndexes{}
}

type NumberingSystem int
