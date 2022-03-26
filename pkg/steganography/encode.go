package steganography

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/util"
)

const bitLength = 8

// Encode encodes a message into the given graphql.Upload file based on input.
func Encode(config *env.Config, input generated.EncodeImageInput) (*bytes.Buffer, error) {
	if !ValidateLSB(input.LsbUsed) {
		return nil, fmt.Errorf(
			"%d is not a valid number of least significant bits used",
			input.LsbUsed,
		)
	}

	data, err := util.FileToImageData(input.File.File)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	encryptedMsg, err := cryptography.Encrypt(
		[]byte(input.Message),
		[]byte(config.EncryptionKey),
	)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	messageLength := len(encryptedMsg)

	if bitLength+messageLength > maxBytesToEncode(data, input) {
		return nil, fmt.Errorf(
			"encode: image isn't big enough for a message of length %d",
			bitLength*messageLength,
		)
	}

	bitChannel := make(chan byte)
	go util.ByteArrToBits(
		append(
			intToUint64Bytes(messageLength),
			encryptedMsg...,
		),
		bitChannel,
	)

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(data, input.Channel, pixelDataChannel)

	lsbPosChannel := make(chan byte)
	go LSBPositions(byte(input.LsbUsed), lsbPosChannel)

pixelIterator:
	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			msgBit, ok := <-bitChannel
			// there are no more bits in the message
			if !ok {
				break pixelIterator
			}

			lsbPos := <-lsbPosChannel

			switch {
			case pixelChannel.IsRed():
				pixelData.SetRed(getUpdatedByte(msgBit, pixelData.GetRed(), lsbPos))
			case pixelChannel.IsGreen():
				pixelData.SetGreen(getUpdatedByte(msgBit, pixelData.GetGreen(), lsbPos))
			case pixelChannel.IsBlue():
				pixelData.SetBlue(getUpdatedByte(msgBit, pixelData.GetBlue(), lsbPos))
			}
		}

		data.NRGBA.SetNRGBA(pixelData.Width, pixelData.Height, *pixelData.Color)
	}

	imgBuffer, err := util.EncodeNRGBA(data.NRGBA)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	return imgBuffer, nil
}

// maxBytesToEncode calculates a maximum amount of bytes which can be encoded
// based on the bounds of given image.NRGBA and generated.EncodeImageInput.
func maxBytesToEncode(
	data util.ImageData,
	input generated.EncodeImageInput,
) int {
	channelCount := ChannelCount(input.Channel)

	return (data.Width * data.Height * channelCount * input.LsbUsed) / bitLength
}

// intToUint64Bytes turns int to []byte representing uint64.
func intToUint64Bytes(num int) []byte {
	b := make([]byte, bitLength)
	binary.LittleEndian.PutUint64(b, uint64(num))

	return b
}

func getUpdatedByte(msgBit byte, value byte, lsbPos byte) byte {
	hasBit := util.HasBit(value, lsbPos)

	switch {
	case msgBit == 0 && hasBit:
		return util.ClearBit(value, lsbPos)
	case msgBit == 1 && !hasBit:
		return util.SetBit(value, lsbPos)
	default:
		return value
	}
}
