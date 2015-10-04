package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHamming_1(t *testing.T) {
	result, err := Hamming([]byte("this is a test"), []byte("wokka wokka!!!"))
	if assert.NoError(t, err) {
		assert.Equal(t, result, 37, "Should be equal")
	}
}

func TestHamming_2(t *testing.T) {
	result, err := Hamming([]byte("foobar"), []byte("foobar"))
	if assert.NoError(t, err) {
		assert.Equal(t, result, 0, "Should be equal")
	}
}

func TestHamming_3(t *testing.T) {
	_, err := Hamming([]byte("foobar"), []byte(""))
	assert.EqualError(t, err, "Input lengths must match")
}
