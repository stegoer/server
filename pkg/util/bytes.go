package util

import "encoding/binary"

// Uint64ToBytes turns int to []byte representing uint64.
func Uint64ToBytes(num uint64) []byte {
	b := make([]byte, bitLen)
	binary.LittleEndian.PutUint64(b, num)

	return b
}

// BytesToUint64 turns []byte to uint64.
func BytesToUint64(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

// BytesToBool turns []byte to bool.
func BytesToBool(bytes []byte) bool {
	return BytesToUint64(bytes) != 0
}

func GetUpdatedByte(newBit byte, value byte, lsbPos byte) byte {
	hasBit := HasBit(value, lsbPos)

	switch {
	case newBit == 0 && hasBit:
		return ClearBit(value, lsbPos)
	case newBit == 1 && !hasBit:
		return SetBit(value, lsbPos)
	default:
		return value
	}
}
