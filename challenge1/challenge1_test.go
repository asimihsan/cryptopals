package challenge1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHexToBase64(t *testing.T) {
	encoded, err := ConvertHexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	assert.Equal(t, encoded, []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"), "they should be equal")
	assert.Nil(t, err)
}
