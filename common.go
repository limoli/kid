package kid

import "errors"

var (
	ErrorInvalidCodeLength = errors.New("invalid code length")
	ErrorInvalidNumber     = errors.New("kid contains an invalid number")
)

type Generator interface {
	Generate(code string) (rune, error)
}

// Verify code using a generator (MOD10 or MOD11)
func Verify(code string, generator Generator) bool {
	if len(code) < 2 {
		return false
	}
	i, err := generator.Generate(code[:len(code)-1])
	return err == nil && i == rune(code[len(code)-1])
}

// A KID is a code between 2 and 25 digits.
// Excluding the check digit, the range is [1, 24].
func validateKidLengthWithoutCheckDigit(code string) (int, bool) {
	codeLength := len(code)
	return codeLength, codeLength >= 1 && codeLength <= 24
}
