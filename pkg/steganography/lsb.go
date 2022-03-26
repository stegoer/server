package steganography

import "github.com/stegoer/server/ent/schema"

// LSBPositions infinitely sends the least significant bit positions.
func LSBPositions(used byte, resultChan chan byte) {
	var position byte

	for position = 0; position <= used; position++ {
		resultChan <- position

		if position == used {
			position = 0
		}
	}

	close(resultChan)
}

// ValidateLSB validates that the number n is within the LSB range.
func ValidateLSB(n int) bool {
	return !(n > schema.LsbMax || n < schema.LsbMin)
}
