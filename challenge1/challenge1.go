package challenge1

import (
	"bytes"
	base64 "encoding/base64"
	hex "encoding/hex"
	"errors"
	"fmt"
)

func ConvertHexToBase64(input []byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(data, input)
	if err != nil {
		return nil, errors.New("Can't convert hex input to raw data")
	}
	b := &bytes.Buffer{}
	e := base64.NewEncoder(base64.StdEncoding, b)
	e.Write(data)
	e.Close()
	return b.Bytes(), nil
}

func main() {
	encoded, _ := ConvertHexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println(string(encoded))
}
