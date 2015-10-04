package main

// cd challenge6
// go run $GOPATH/src/github.com/asimihsan/cryptopals/challenge6/challenge6.go

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"sort"
	"unicode"
)

func Hamming(input1, input2 []byte) (int, error) {
	if len(input1) != len(input2) {
		return 0, errors.New("Input lengths must match")
	}
	result := 0
	for i := range input1 {
		c1 := input1[i]
		c2 := input2[i]
		for j := uint(0); j < 8; j++ {
			mask := byte(1 << j)
			if c1&mask != c2&mask {
				result++
			}
		}
	}
	return result, nil
}

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

func getKeySizeScore(keySize int, ciphertext []byte, iterations int) (result float64) {
	distances := make([]int, iterations)
	for i := 0; i <= iterations; i++ {
		first := ciphertext[keySize*i : keySize*(i+1)]
		second := ciphertext[keySize*(i+1) : keySize*(i+2)]
		distance, _ := Hamming(first, second)
		distances = append(distances, distance)
	}
	var sum int
	for _, distance := range distances {
		sum += distance
	}
	return float64(sum) / float64(len(distances)*keySize)
}

type bestKeySizesResult struct {
	KeySize int
	Score   float64
}
type byScore []bestKeySizesResult

func (s byScore) Len() int {
	return len(s)
}

func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func getKeySizesAndScores(ciphertext []byte, keySizeIterations int) (results []bestKeySizesResult) {
	for i := 2; i <= 40; i++ {
		score := getKeySizeScore(i, ciphertext, keySizeIterations)
		result := bestKeySizesResult{KeySize: i, Score: score}
		results = append(results, result)
	}
	sort.Sort(byScore(results))
	return results
}

func getBestKeySizes(ciphertext []byte, numKeySizes int, keySizeIterations int) (bestKeySizes []int) {
	keySizesAndScores := getKeySizesAndScores(ciphertext, keySizeIterations)
	for i := 0; i < numKeySizes; i++ {
		bestKeySizes = append(bestKeySizes, keySizesAndScores[i].KeySize)
	}
	return bestKeySizes
}

func splitIntoKeySizeChunks(ciphertext []byte, keySize int) (output [][]byte) {
	for i := 0; i < keySize; i++ {
		output = append(output, make([]byte, 0))
	}
	for i, c := range ciphertext {
		output[i%keySize] = append(output[i%keySize], c)
	}
	return output
}

func guessKey(keySize int, chunks [][]byte) []byte {
	key := make([]byte, keySize)
	for i, chunk := range chunks {
		bestScore, keyByte, _, _ := GetSingleByteXORBestKeyAndPlaintext(chunk)
		if bestScore == 0 {
			return nil
		}
		key[i] = keyByte
	}
	return key
}

func Score(input []byte) int {
	score := 0
	for _, b := range input {
		if b == 9 || b == 10 || b == 13 {
			score++
			continue
		}
		if b < 32 || b > unicode.MaxASCII {
			// if invalid ASCII, can't be plaintext
			return 0
		}
		if b == 32 {
			score += 3
		} else if b == 69 || b == 101 {
			score += 2
		} else if b == 84 || b == 116 {
			score++
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

func GetSingleByteXORBestKeyAndPlaintext(input []byte) (int, byte, []byte, error) {
	bestScore := 0
	var bestKey byte
	var bestPlaintext []byte
	for b := 0; b <= 255; b++ {
		result, _ := SingleByteXOR(input, byte(b))
		currentScore := Score(result)
		if currentScore > bestScore {
			bestScore = currentScore
			bestKey = byte(b)
			bestPlaintext = result
		}
	}
	return bestScore, bestKey, bestPlaintext, nil
}

func RepeatingKeyXOR(input []byte, key []byte) ([]byte, error) {
	result := make([]byte, len(input))
	for i, c := range input {
		result[i] = c ^ key[i%len(key)]
	}
	return result, nil
}

func main() {
	ciphertext, _ := base64DecodeFile("6.txt")
	keySizes := getBestKeySizes(ciphertext, 5, 5)
	for _, keySize := range keySizes {
		chunks := splitIntoKeySizeChunks(ciphertext, keySize)
		key := guessKey(keySize, chunks)
		if key != nil {
			log.Printf("key: %s", string(key))
			plaintext, _ := RepeatingKeyXOR(ciphertext, key)
			log.Printf("plaintext: %s", string(plaintext))
		}
	}
}
