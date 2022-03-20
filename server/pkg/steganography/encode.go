package steganography

import (
	"bytes"
	"fmt"

	"github.com/kucera-lukas/stegoer/server/graph/generated"
	"github.com/kucera-lukas/stegoer/server/pkg/util"
)

const bitSize = 8

// Encode encodes a message into the given graphql.Upload file based on input.
func Encode(input generated.EncodeImageInput) (*bytes.Buffer, error) {
	data, err := FileToImageData(input.File.File)
	if err != nil {
		return nil, err
	}

	messageLength := len(input.Message)

	if messageLength > maxEncodeSize(data, input) {
		return nil, fmt.Errorf(
			"image isn't big enough for a message of length %d",
			messageLength,
		)
	}

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(data, input.Channel, pixelDataChannel)

	bitChannel := make(chan byte)
	go util.ByteArrToBits(
		append(splitToBytes(messageLength), []byte(input.Message)...),
		bitChannel,
	)

	lsbPosChannel := make(chan byte)
	go util.LSBPositions(byte(input.LsbUsed), lsbPosChannel)

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

	imgBuffer, err := EncodeNRGBA(data.NRGBA)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	return imgBuffer, nil
}

// maxEncodeSize calculates a maximum encode size
// based on the bounds of given image.NRGBA and generated.EncodeImageInput.
func maxEncodeSize(data ImageData, input generated.EncodeImageInput) int {
	channelCount := util.ChannelCount(input.Channel)

	return (data.Width * data.Height * channelCount) / (bitSize / input.LsbUsed)
}

// splitToBytes given an unsigned integer,
// will split this integer into its four bytes.
func splitToBytes(num int) []byte {
	one := byte(num >> bitSize * 3) //nolint:gomnd

	mask := 255

	two := byte((num >> bitSize * 2) & mask) //nolint:gomnd
	three := byte((num >> bitSize) & mask)
	four := byte(num & mask)

	return []byte{one, two, three, four}
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
