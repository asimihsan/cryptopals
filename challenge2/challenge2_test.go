package challenge2

import (
	hex "encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedXOR1(t *testing.T) {
	input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	key, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	output, err := FixedXOR([]byte(input), []byte(key))
	expected_output, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")
	assert.Equal(t, output, expected_output, "Should be equal")
	assert.Nil(t, err)
}
