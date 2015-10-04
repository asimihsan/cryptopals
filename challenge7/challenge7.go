package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
)

func base64DecodeFile(filepath string) ([]byte, error) {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	output := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	outputLength, err := base64.StdEncoding.Decode(output, input)
	if err != nil {
		return nil, err
	}
	return output[:outputLength], nil
}

func AESECBDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Couldn't initialize AES cipher")
	}

	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		return nil, errors.New("Need a multiple of blocksize")
	}

	plaintext := make([]byte, len(ciphertext))
	plaintextOriginal := plaintext
	for len(plaintext) > 0 {
		block.Decrypt(plaintext, ciphertext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	return plaintextOriginal, nil
}

func main() {
	ciphertext, err := base64DecodeFile("7.txt")
	if err != nil {
		log.Panic("failed to load 7.txt input file")
	}
	key := []byte("YELLOW SUBMARINE")
	plaintext, err := AESECBDecrypt(ciphertext, key)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(string(plaintext))
}
