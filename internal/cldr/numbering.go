package cldr

import (
	"strings"

	"golang.org/x/text/language"
)

// Digits represents a set of numeral glyphs for a specific numeral system.
// It is an array of 10 runes, where each index corresponds to a digit (0-9)
// in the decimal system, and the value at that index is the corresponding
// glyph in the represented numeral system.
//
// For example:
//
//	Digits{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} // represents Latin numerals
//	Digits{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'} // represents Arabic-Indic numerals
//
// A special case is when Digits[0] is 0, which is used to represent Latin numerals
// and triggers special handling in some methods.
type Digits [10]rune

func (d Digits) TwoDigit(number int) string {
	if number < 10 { //nolint:mnd
		return string(d[0]) + string(d[number])
	}

	last := number % 10 //nolint:mnd
	number /= 10
	beforeLast := number % 10 //nolint:mnd

	return string(d[beforeLast]) + string(d[last])
}

func (d Digits) Numeric(number int) string {
	// single digit
	if number < 10 { //nolint:mnd
		return string(d[number])
	}

	const maxDigits = 4 // based on digits in the current Gregorian calendar year

	// more than one digit
	chars := make([]int, 0, maxDigits)

	for number > 0 {
		chars = append(chars, number%10) //nolint:mnd
		number /= 10
	}

	var sb strings.Builder

	const bytesPerRune = 4

	sb.Grow(len(chars) * bytesPerRune)

	for i := len(chars) - 1; i >= 0; i-- {
		sb.WriteRune(d[chars[i]])
	}

	return sb.String()
}

func defaultNumberingSystem(locale language.Tag) NumberingSystem {
	lang, script, region := locale.Raw()

	switch script {
	case Latn:
		return NumberingSystemLatn
	case Adlm:
		return NumberingSystemAdlm
	}

	switch lang {
	default:
		return NumberingSystemLatn
	case AR:
		switch region {
		default:
			return NumberingSystemArab
		case Region001, RegionAE, RegionDZ, RegionEH, RegionLY, RegionMA, RegionTN, RegionZZ:
			return NumberingSystemLatn
		}
	case CKB, SDH:
		return NumberingSystemArab
	case SD:
		if script == Deva {
			return NumberingSystemLatn
		}

		return NumberingSystemArab
	case AS, BN, MNI:
		return NumberingSystemBeng
	case BGN, FA, LRC, MZN, PS:
		return NumberingSystemArabext
	case BGC, BHO, MR, NE, RAJ, SA:
		return NumberingSystemDeva
	case CCP:
		return NumberingSystemCakm
	case DZ:
		return NumberingSystemTibt
	case KS:
		if script == Deva {
			return NumberingSystemLatn
		}

		return NumberingSystemArabext
	case HNJ:
		return NumberingSystemHmnp
	case MY:
		return NumberingSystemMymr
	case NQO:
		return NumberingSystemNkoo
	case PA:
		if script == Arab {
			return NumberingSystemArabext
		}

		return NumberingSystemLatn
	case SAT:
		return NumberingSystemOlck
	case UR:
		if region == RegionIN {
			return NumberingSystemArabext
		}

		return NumberingSystemLatn
	case UZ:
		if script == Arab {
			return NumberingSystemArabext
		}

		return NumberingSystemLatn
	}
}

func LocaleDigits(locale language.Tag) Digits {
	if i := defaultNumberingSystem(locale); i >= 0 && int(i) < len(numberingSystems) { // isInBounds()
		return numberingSystems[i]
	}

	return numberingSystems[NumberingSystemLatn]
}
