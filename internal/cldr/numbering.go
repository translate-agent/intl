package cldr

import (
	"unicode/utf8"

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

func (d *Digits) appendTwoDigit(b []byte, number int) []byte {
	if d[0] == '0' {
		if number < 10 { //nolint:mnd
			//nolint:gosec // number is < 10, won't overflow
			digit := byte('0' + number)

			return append(b, '0', digit)
		}

		ones := number % 10      //nolint:mnd
		tens := number / 10 % 10 //nolint:mnd

		return append(b, byte('0'+tens), byte('0'+ones))
	}

	if number < 10 { //nolint:mnd
		b = utf8.AppendRune(b, d[0])
		b = utf8.AppendRune(b, d[number])

		return b
	}

	ones := number % 10      //nolint:mnd
	tens := number / 10 % 10 //nolint:mnd

	b = utf8.AppendRune(b, d[tens])
	b = utf8.AppendRune(b, d[ones])

	return b
}

func (d *Digits) appendNumeric(b []byte, number int) []byte {
	if d[0] == '0' {
		if number < 10 { //nolint:mnd
			//nolint:gosec // number is < 10, won't overflow
			digit := byte('0' + number)

			return append(b, digit)
		}

		var buf [19]byte

		i := len(buf)

		for number > 0 {
			i--
			buf[i] = byte('0' + number%10)
			number /= 10
		}

		return append(b, buf[i:]...)
	}

	// single digit
	if number < 10 { //nolint:mnd
		return utf8.AppendRune(b, d[number])
	}

	// Use fixed-size stack array instead of make([]int, 0, maxDigits)
	var chars [19]int

	i := 0

	for number > 0 {
		chars[i] = number % 10 //nolint:mnd
		i++
		number /= 10
	}

	for j := i - 1; j >= 0; j-- {
		b = utf8.AppendRune(b, d[chars[j]])
	}

	return b
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

func LocaleDigitsPtr(locale language.Tag) *Digits {
	if i := defaultNumberingSystem(locale); i >= 0 && int(i) < len(numberingSystems) { // isInBounds()
		return &numberingSystems[i]
	}

	return &numberingSystems[NumberingSystemLatn]
}
