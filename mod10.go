package kid

import "unicode"

// Luhn MOD10 implements Generator interface
type Luhn10 struct{}

// Generate check digit using MOD10
func (l *Luhn10) Generate(code string) (rune, error) {
	codeLength, isValid := validateKidLengthWithoutCheckDigit(code)
	if !isValid {
		return 0, ErrorInvalidCodeLength
	}

	sum := 0
	parity := (codeLength + 1) % 2
	for i, n := range code {
		if !unicode.IsDigit(n) {
			return 0, ErrorInvalidNumber
		}
		d := int(n - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	return rune(sum*9%10) + '0', nil
}