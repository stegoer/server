package util

const bitLen = 8

// ByteArrToBits turns given string into bits and sends it over a channel.
func ByteArrToBits(byteArr []byte, resultChan chan byte) {
	var position byte

	for _, b := range byteArr {
		for position = bitLen; position > 0; position-- {
			// need to offset starting position by 1
			if HasBit(b, position-1) {
				resultChan <- 1
			} else {
				resultChan <- 0
			}
		}
	}

	close(resultChan)
}

// SetBit sets the bit at pos in the integer n.
func SetBit(n byte, pos byte) byte {
	n |= 1 << pos

	return n
}

// ClearBit clears the bit at pos in n.
func ClearBit(n byte, pos byte) byte {
	n &= ^(1 << pos)

	return n
}

// HasBit returns whether the byte n has a bit set on the pos.
func HasBit(n byte, pos byte) bool {
	val := n & (1 << pos)

	return val > 0
}

// BoolToRune turns a bool into a rune.
func BoolToRune(b bool) rune {
	if b {
		return '1'
	}

	return '0'
}
