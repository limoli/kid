package kid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLuhn10_Generate(t *testing.T) {
	m := map[string]rune{
		"390362452109003080489288": '4',
		"992021383":                '3',
		"54817431":                 '7',
		"31004006799114005781207":  '4',
		"31000713454736768884029":  '3',
		"31000712436818000100010":  '8',
		"994168985":                '4',
	}
	luhn10 := Luhn10{}
	for k, v := range m {
		r, err := luhn10.Generate(k)
		assert.Nil(t, err)
		assert.Equal(t, v, r)
	}
}

func TestLuhn10_Verify(t *testing.T) {
	m := map[string]bool{
		"":                          false,
		"1234567890":                false,
		"3903624521090030804892884": true,
		"9920213833":                true,
		"548174317":                 true,
		"21090001177060030":         true,
		"310040067991140057812074":  true,
		"310007134547367688840293":  true,
		"310007124368180001000108":  true,
		"9941689854":                true,
	}
	for kid, expectedIsValid := range m {
		assert.Equal(t, expectedIsValid, Verify(kid, new(Luhn10)), kid)
	}
}
