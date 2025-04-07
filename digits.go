package intl

import (
	"strings"

	"golang.org/x/text/language"
)

// digits represents a set of numeral glyphs for a specific numeral system.
// It is an array of 10 runes, where each index corresponds to a digit (0-9)
// in the decimal system, and the value at that index is the corresponding
// glyph in the represented numeral system.
//
// For example:
//
//	digits{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} // represents Latin numerals
//	digits{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'} // represents Arabic-Indic numerals
//
// A special case is when digits[0] is 0, which is used to represent Latin numerals
// and triggers special handling in some methods.
type digits [10]rune

func (d digits) twoDigit(number int) string {
	if number < 10 {
		return string(d[0]) + string(d[number])
	}

	last := number % 10
	number = number / 10
	beforeLast := number % 10

	return string(d[beforeLast]) + string(d[last])
}

func (d digits) numeric(number int) string {
	// single digit
	if number < 10 {
		return string(d[number])
	}

	// more than one digit
	chars := make([]int, 0, 4)

	for number > 0 {
		chars = append(chars, number%10)
		number /= 10
	}

	var sb strings.Builder

	sb.Grow(len(chars) * 4)

	for i := len(chars) - 1; i >= 0; i-- {
		sb.WriteRune(d[chars[i]])
	}

	return sb.String()
}

func localeDigits(locale language.Tag) digits {
	if i := defaultNumberingSystem(locale); i >= 0 && int(i) < len(numberingSystems) { // isInBounds()
		return numberingSystems[i]
	}

	return numberingSystems[numberingSystemLatn]
}
