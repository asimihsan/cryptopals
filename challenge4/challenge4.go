package main

// cat 4.txt | go run challenge4.go

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"log"
	"os"
	"unicode"
)

func Score(input []byte) int {
	score := 0
	for _, b := range input {
		if b < 32 || b > unicode.MaxASCII {
			// if invalid ASCII, can't be plaintext
			return score
		}
		if b == 32 { // space
			score += 3
		} else if b == 69 || b == 101 { // E, e
			score += 2
		} else if b == 84 || b == 116 { // T, t
			score += 1
		}
	}
	return score
}

func SingleByteXOR(input []byte, key byte) ([]byte, error) {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return result, nil
}

func GetSingleByteXORBestKeyAndPlaintext(scanner *bufio.Scanner) (byte, []byte, error) {
	bestScore := 0
	var bestKey byte
	var bestPlaintext []byte
	for scanner.Scan() {
		input, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Printf("failed to hex decode string: %s", err)
			return 0, nil, err
		}
		for b := 0; b <= 255; b++ {
			result, _ := SingleByteXOR(input, byte(b))
			result = bytes.TrimSuffix(result, []byte("\n"))
			currentScore := Score(result)
			if currentScore > bestScore {
				bestScore = currentScore
				bestKey = byte(b)
				bestPlaintext = result
			}
		}
	}
	return bestKey, bestPlaintext, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bestKey, bestPlaintext, _ := GetSingleByteXORBestKeyAndPlaintext(scanner)
	log.Printf("key %x, plaintext: %s", bestKey, bestPlaintext)
}
