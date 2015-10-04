package challenge2

import (
	hex "encoding/hex"
	"errors"
	"fmt"
)

func FixedXOR(input []byte, key []byte) ([]byte, error) {
	if len(input) != len(key) {
		return nil, errors.New("Key size must be equal to input size")
	}
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i]
	}
	return output, nil
}

func main() {
	input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	key, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	output, _ := FixedXOR([]byte(input), []byte(key))
	fmt.Printf(hex.EncodeToString(output))
}
