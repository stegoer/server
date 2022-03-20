package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/stegoer/server/ent/schema"
)

const (
	binaryBase = 2
	bitLen     = 8
	bitSize    = 32
)

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

// LSBPositions infinitely sends the least significant bits positions.
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

// BinaryBufferToString turns the data from bytes.Buffer into a string.
func BinaryBufferToString(binBuffer *bytes.Buffer) (string, error) {
	var textBuilder strings.Builder

	bufferLen := binBuffer.Len()

	if bufferLen%bitLen != 0 {
		return "", errors.New("invalid buffer length")
	}

	for i := 0; i < bufferLen; i += bitLen {
		strChunk, err := io.ReadAll(io.LimitReader(binBuffer, bitLen))
		if err != nil {
			return "", fmt.Errorf("failed reading from buffer: %w", err)
		}

		parsedInt, err := strconv.ParseInt(string(strChunk), binaryBase, bitSize)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s as a string: %w", strChunk, err)
		}

		textBuilder.WriteRune(rune(parsedInt))
	}

	return textBuilder.String(), nil
}

// BinaryBufferToInt turns the data from bytes.Buffer into an int.
func BinaryBufferToInt(binBuffer *bytes.Buffer) (int, error) {
	var byteBuffer bytes.Buffer

	bufferLen := binBuffer.Len()

	if bufferLen%bitLen != 0 {
		return 0, errors.New("invalid buffer length")
	}

	for i := 0; i < bufferLen; i += bitLen {
		strChunk, err := io.ReadAll(io.LimitReader(binBuffer, bitLen))
		if err != nil {
			return 0, fmt.Errorf("failed reading from buffer: %w", err)
		}

		parsedInt, err := strconv.ParseInt(string(strChunk), binaryBase, bitSize)
		if err != nil {
			return 0, fmt.Errorf("failed to parse %s as a string: %w", strChunk, err)
		}

		byteBuffer.WriteByte(byte(parsedInt))
	}

	return int(ByteArrayToInt(byteBuffer.Bytes())), nil
}

// ByteArrayToInt given four bytes,
// will return the 32 bit unsigned integer
// which is the composition of those four bytes (one is MSB).
func ByteArrayToInt(byteArr []byte) (ret uint32) {
	ret = uint32(byteArr[0])
	ret <<= 8
	ret |= uint32(byteArr[1])
	ret <<= 8
	ret |= uint32(byteArr[2])
	ret <<= 8
	ret |= uint32(byteArr[3])

	return
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

func ValidLSBUsed(n int) bool {
	return !(n > schema.LsbMax || n < schema.LsbMin)
}
