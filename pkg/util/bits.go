package util

const bitLen = 8

// ByteArrToBits turns given byteArr into bits and sends it over a channel.
func ByteArrToBits(byteArr []byte, resultChan chan byte) {
	var position byte

	for _, b := range byteArr {
		for position = bitLen; position > 0; position-- {
			resultChan <- BoolToBit(HasBit(b, position-1))
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

// BoolToBit turns a bool into a bit.
func BoolToBit(b bool) byte {
	if b {
		return 1
	}

	return 0
}

// BitToBool turns a bit inti bool
func BitToBool(b byte) bool {
	return b != 0
}
