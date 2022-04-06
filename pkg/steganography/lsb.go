package steganography

const (
	lsbMin = 1
	lsbMax = 8
)

// LSBPositions infinitely sends the least significant bit positions.
func LSBPositions(used byte, resultChan chan byte) {
	var position byte

	for position = 0; position <= used; position++ {
		if position == used {
			position = 0
		}

		resultChan <- position
	}

	close(resultChan)
}

// ValidateLSB validates that the number n is within the LSB range.
func ValidateLSB(n int) bool {
	return !(n > lsbMax || n < lsbMin)
}
