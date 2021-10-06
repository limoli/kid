package kid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLuhn11_Generate(t *testing.T) {
	m := map[string]rune{
		"036532":                 '7',
		"1234567890":             '3',
		"2440997":                '-',
		"2597817":                '-',
		"1081925130007365887":    '-',
		"0004740765018948619306": '-',
	}
	luhn11 := Luhn11{}
	for k, v := range m {
		r, err := luhn11.Generate(k)
		assert.Nil(t, err)
		assert.Equal(t, v, r)
	}
}

func TestLuhn11_Verify(t *testing.T) {
	m := map[string]bool{
		"":                        false,
		"1234567890":              false,
		"0365327":                 true,
		"2440997-":                true,
		"2597817-":                true,
		"1081925130007365887-":    true,
		"0004740765018948619306-": true,
		"12345678903":             true,
		"19333485":                true,
		"04866952901086741000039": true,
	}
	for kid, expectedIsValid := range m {
		assert.Equal(t, expectedIsValid, Verify(kid, new(Luhn11)), kid)
	}
}
