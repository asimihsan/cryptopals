package challenge3

import (
	hex "encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedXOR1(t *testing.T) {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	bestDecrypted := GetBestEnglishForFixedXor([]byte(input))
	assert.Equal(t, bestDecrypted, "Cooking MC's like a pound of bacon", "Should be equal")
}
