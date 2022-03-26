package steganography

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/util"
)

// Decode decodes a message from the given graphql.Upload file.
func Decode(config *env.Config, input generated.DecodeImageInput) (string, error) {
	var (
		msgLength    uint64
		binaryBuffer bytes.Buffer
	)

	if !ValidateLSB(input.LsbUsed) {
		return "", fmt.Errorf(
			"%d is not a valid number of least significant bits used",
			input.LsbUsed,
		)
	}

	data, err := util.FileToImageData(input.File.File)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(data, input.Channel, pixelDataChannel)

	lsbPosChannel := make(chan byte)
	go LSBPositions(byte(input.LsbUsed), lsbPosChannel)

	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			var value byte

			switch {
			case pixelChannel.IsRed():
				value = pixelData.GetRed()
			case pixelChannel.IsGreen():
				value = pixelData.GetGreen()
			case pixelChannel.IsBlue():
				value = pixelData.GetBlue()
			}

			lsbPos := <-lsbPosChannel
			hasBit := util.HasBit(value, lsbPos)

			binaryBuffer.WriteRune(util.BoolToRune(hasBit))

			// get encoded message length
			if msgLength == 0 && binaryBuffer.Len() == bitLength*bitLength {
				msgLength, err = util.BinaryBufferToUint64(&binaryBuffer)
				if err != nil {
					return "", fmt.Errorf("decode: %w", err)
				}

				binaryBuffer.Reset()
			} else if msgLength != 0 && uint64(binaryBuffer.Len()) == bitLength*msgLength {
				byteSlice, err := util.BinaryBufferToBytes(&binaryBuffer)
				if err != nil {
					return "", fmt.Errorf("decode: %w", err)
				}

				msg, err := cryptography.Decrypt(
					byteSlice,
					[]byte(config.EncryptionKey),
				)
				if err != nil {
					return "", fmt.Errorf("decode: %w", err)
				}

				return string(msg), nil
			}
		}
	}

	return "", errors.New("decode: no message found")
}
