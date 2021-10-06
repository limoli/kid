package kid

import "unicode"

type Luhn11 struct{}

func (l *Luhn11) Generate(code string) (rune, error) {
	codeLength, isValid := validateKidLengthWithoutCheckDigit(code)
	if !isValid {
		return 0, ErrorInvalidCodeLength
	}

	sum := 0
	for i, n := range code {
		if !unicode.IsDigit(n) {
			return 0, ErrorInvalidNumber
		}
		number := int(n - '0')
		factor := (codeLength-i-1)%6 + 2
		sum += number * factor
	}

	var checkDigit rune
	remainder := sum % 11
	switch remainder {
	case 0:
		checkDigit = '0'
		break
	case 1:
		checkDigit = '-'
		break
	default:
		checkDigit = rune(11-remainder) + '0'
	}

	return checkDigit, nil
}