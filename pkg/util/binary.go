package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strconv"
)

const (
	binaryBase = 2
	bitSize    = 32
)

// BinaryBufferToString turns the data from bytes.Buffer into a string.
//func BinaryBufferToString(binBuffer *bytes.Buffer) (string, error) {
//	var textBuilder strings.Builder
//
//	bufferLen := binBuffer.Len()
//
//	err := validateBufferLength(bufferLen)
//	if err != nil {
//		return "", err
//	}
//
//	for i := 0; i < bufferLen; i += bitLen {
//		parsedInt, err := parseInt64(binBuffer)
//		if err != nil {
//			return "", err
//		}
//
//		// returns a nil error
//		textBuilder.WriteRune(rune(parsedInt))
//	}
//
//	return textBuilder.String(), nil
//}

// BinaryBufferToUint64 turns the data from bytes.Buffer into an uint64.
func BinaryBufferToUint64(binBuffer *bytes.Buffer) (uint64, error) {
	byteSlice, err := BinaryBufferToBytes(binBuffer)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(byteSlice), nil
}

// BinaryBufferToBytes turns the data from bytes.Buffer into a slice of bytes.
func BinaryBufferToBytes(binBuffer *bytes.Buffer) ([]byte, error) {
	var byteBuffer bytes.Buffer

	bufferLen := binBuffer.Len()

	err := validateBufferLength(bufferLen)
	if err != nil {
		return nil, err
	}

	for i := 0; i < bufferLen; i += bitLen {
		parsedInt, err := parseInt64(binBuffer)
		if err != nil {
			return nil, err
		}

		// returns a nil error
		byteBuffer.WriteByte(byte(parsedInt))
	}

	return byteBuffer.Bytes(), nil
}

func parseInt64(binBuffer *bytes.Buffer) (int64, error) {
	strChunk, err := io.ReadAll(io.LimitReader(binBuffer, bitLen))
	if err != nil {
		return 0, fmt.Errorf("failed reading from buffer: %w", err)
	}

	parsedInt, err := strconv.ParseInt(string(strChunk), binaryBase, bitSize)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse %s as a string: %w",
			strChunk,
			err,
		)
	}

	return parsedInt, nil
}

func validateBufferLength(bufferLen int) error {
	if (bufferLen % bitLen) != 0 {
		return errors.New("invalid buffer length")
	}

	return nil
}
