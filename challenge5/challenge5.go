package challenge5

func RepeatingKeyXOR(input []byte, key []byte) ([]byte, error) {
	result := make([]byte, len(input))
	for i, c := range input {
		result[i] = c ^ key[i%len(key)]
	}
	return result, nil
}
