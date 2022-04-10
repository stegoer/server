package steganography

const (
	lsbMin byte = 1
	lsbMax byte = 8
)

// LSBSlice returns a slice with the lsbs based on the used byte.
func LSBSlice(used byte) []byte {
	var result []byte

	for position := lsbMin - 1; position < used; position++ {
		result = append(result, position)
	}

	return result
}

// ValidateLSB validates that the number n is within the LSB range.
func ValidateLSB(n byte) bool {
	return !(n > lsbMax || n < lsbMin)
}
