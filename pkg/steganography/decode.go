package steganography

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/util"
)

const (
	bitLength = 8
)

// Decode decodes a message from the given graphql.Upload file.
//nolint:cyclop
func Decode(input generated.DecodeImageInput) (string, error) {
	var (
		msgLength    int
		binaryBuffer bytes.Buffer
	)

	data, err := FileToImageData(input.File.File)
	if err != nil {
		return "", err
	}

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(data, input.Channel, pixelDataChannel)

	lsbPosChannel := make(chan byte)
	go util.LSBPositions(byte(input.LsbUsed), lsbPosChannel)

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
			if msgLength == 0 && binaryBuffer.Len() == 32 {
				msgLength, err = util.BinaryBufferToInt(&binaryBuffer)
				if err != nil {
					return "", fmt.Errorf("decode: %w", err)
				}

				binaryBuffer.Reset()
			} else if msgLength != 0 && binaryBuffer.Len() == bitLength*msgLength {
				msg, err := util.BinaryBufferToString(&binaryBuffer)
				if err != nil {
					return "", fmt.Errorf("decode: %w", err)
				}

				return msg, nil
			}
		}
	}

	return "", errors.New("decode: no message found")
}
