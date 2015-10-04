package challenge3

import (
	"bytes"
	hex "encoding/hex"
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

var (
	englishFrequencies   = []byte("etaoinsrhldcumfpgwybvkxjqz")
	numberOfLettersToUse = 20
	absentPenalty        = 100
)

func SingleByteXOR(input []byte, key byte) ([]byte, error) {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return result, nil
}

func IsPrintable(input string) bool {
	for _, r := range input {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func EnglishScore(input string) int {
	if !IsPrintable(input) {
		return math.MaxInt32
	}
	inputLower := strings.ToLower(input)
	count := GetCharactersInCountOrder(inputLower)
	score := 0
	i := 0
	for expectedIndex, c := range englishFrequencies {
		actualIndex := bytes.IndexByte(count, c)
		if actualIndex == -1 {
			score += absentPenalty
		} else {
			score += (expectedIndex - actualIndex) * (expectedIndex - actualIndex)
		}
		i++
		if i > numberOfLettersToUse {
			break
		}
	}
	return score
}

type sortedMap struct {
	m map[byte]int
	b []byte
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i int, j int) bool {
	return (sm.m[sm.b[i]] == sm.m[sm.b[j]] && sm.b[i] > sm.b[j]) ||
		(sm.m[sm.b[i]] > sm.m[sm.b[j]])
}

func (sm *sortedMap) Swap(i int, j int) {
	sm.b[i], sm.b[j] = sm.b[j], sm.b[i]
}

func sortedKeys(m map[byte]int) []byte {
	sm := new(sortedMap)
	sm.m = m
	sm.b = make([]byte, len(m))
	i := 0
	for key, _ := range m {
		sm.b[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.b
}

func GetCharactersInCountOrder(input string) []byte {
	count := make(map[byte]int)
	for _, c := range input {
		count[byte(c)]++
	}
	sortedCount := sortedKeys(count)
	return sortedCount
}

func GetBestEnglishForFixedXor(input []byte) string {
	smallestScore := math.MaxInt32
	bestDecrypted := ""
	for key := 0; key < 256; key++ {
		result, _ := SingleByteXOR(input, byte(key))
		score := EnglishScore(string(result))
		if score < smallestScore {
			smallestScore = score
			bestDecrypted = string(result)
		}
	}
	return bestDecrypted
}

func main() {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	bestDecrypted := GetBestEnglishForFixedXor([]byte(input))
	fmt.Printf(bestDecrypted)
}
