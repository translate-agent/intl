package intl

import (
	"strconv"
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
	ones := number % 10      //nolint:mnd
	tens := number / 10 % 10 //nolint:mnd

	// latn
	if d[0] == '0' {
		return string([]byte{byte('0' + tens), byte('0' + ones)})
	}

	return string([]rune{d[tens], d[ones]})
}

func (d digits) numeric(number int) string {
	// latn
	if d[0] == '0' {
		return strconv.Itoa(number)
	}

	const maxDigits = 4 // based on digits in the current Gregorian calendar year

	chars := make([]rune, 0, 4)

	for number > 0 {
		if v := number % 10; v >= 0 && v < 10 { // isInBounds()
			chars = append(chars, d[v]) //nolint:mnd
		}

		number /= 10
	}

	var sb strings.Builder

	const bytesPerRune = 4 // at least one digit

	sb.Grow(bytesPerRune)

	for i := len(chars) - 1; i >= 0; i-- {
		sb.WriteRune(chars[i])
	}

	return sb.String()
}

func localeDigits(locale language.Tag) digits {
	if i := defaultNumberingSystem(locale); i >= 0 && int(i) < len(numberingSystems) { // isInBounds()
		return numberingSystems[i]
	}

	return numberingSystems[numberingSystemLatn]
}
