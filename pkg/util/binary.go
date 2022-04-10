package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

const (
	binaryBase = 2
	bitSize    = 32
)

// BinaryBufferToBytes turns the data from bytes.Buffer into a slice of bytes.
func BinaryBufferToBytes(binBuffer *bytes.Buffer) ([]byte, error) {
	var byteBuffer bytes.Buffer

	bufferLen := binBuffer.Len()

	if err := validateBufferLength(bufferLen); err != nil {
		return nil, err
	}

	for i := 0; i < bufferLen; i += BitLength {
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
	byteChunk, err := io.ReadAll(io.LimitReader(binBuffer, BitLength))
	if err != nil {
		return 0, fmt.Errorf("failed reading from buffer: %w", err)
	}

	parsedInt, err := strconv.ParseInt(string(byteChunk), binaryBase, bitSize)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse %s as a string: %w",
			byteChunk,
			err,
		)
	}

	return parsedInt, nil
}

func validateBufferLength(bufferLen int) error {
	if (bufferLen % BitLength) != 0 {
		return errors.New("invalid buffer length")
	}

	return nil
}
