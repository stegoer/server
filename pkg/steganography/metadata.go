package steganography

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

const (
	metadataLength       = 12
	metadataBinaryLength = metadataLength * bitLength
	metadataPixelOffset  = 0
	metadataLsbPos       = 1
)

// Metadata represents information which was used to encode data into an image.
type Metadata struct {
	length  uint64
	lsbUsed uint8
	red     bool
	green   bool
	blue    bool
}

func (md Metadata) GetBinaryLength() int {
	return int(md.length) * bitLength
}

func (md Metadata) GetChannel() model.Channel {
	switch {
	case md.red && md.green && md.blue:
		return model.ChannelRedGreenBlue
	case md.red && md.green && !md.blue:
		return model.ChannelRedGreen
	case md.red && !md.green && md.blue:
		return model.ChannelRedBlue
	case md.red && !md.green && !md.blue:
		return model.ChannelRed
	case !md.red && md.green && md.blue:
		return model.ChannelGreenBlue
	case !md.red && md.green && !md.blue:
		return model.ChannelGreen
	case !md.red && !md.green && md.blue:
		return model.ChannelBlue
	default:
		// should be unreachable
		return model.ChannelRedGreenBlue
	}
}

func (md Metadata) ToByteArr() []byte {
	result := util.Uint64ToBytes(md.length)
	result = append(result, md.lsbUsed)
	result = append(
		result,
		[]byte{
			util.BoolToBit(md.red),
			util.BoolToBit(md.green),
			util.BoolToBit(md.blue),
		}...,
	)

	return result
}

func (md Metadata) EncodeIntoImageData(data util.ImageData) {
	bitChannel := make(chan byte)
	go util.ByteArrToBits(md.ToByteArr(), bitChannel)

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(
		data,
		metadataPixelOffset,
		model.ChannelRedGreenBlue,
		pixelDataChannel,
	)

pixelIterator:
	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			dataBit, ok := <-bitChannel
			// there are no more bits in the bit channel
			if !ok {
				break pixelIterator
			}

			pixelData.SetChannelValue(pixelChannel, dataBit, metadataLsbPos)
		}

		data.NRGBA.SetNRGBA(pixelData.Width, pixelData.Height, *pixelData.Color)
	}
}

func MetadataFromEncodeInput(
	input generated.EncodeImageInput,
	messageLength int,
) Metadata {
	return Metadata{
		length:  uint64(messageLength),
		lsbUsed: uint8(input.LsbUsed),
		red:     input.Channel.IncludesRed(),
		green:   input.Channel.IncludesGreen(),
		blue:    input.Channel.IncludesBlue(),
	}
}

func MetadataFromBinaryBuffer(binaryBuffer *bytes.Buffer) (*Metadata, error) {
	byteSlice, err := util.BinaryBufferToBytes(binaryBuffer)
	if err != nil {
		return nil, fmt.Errorf("metadata: %w", err)
	}

	if len(byteSlice) != metadataLength {
		return nil, errors.New(
			"buffer length does not match expected metadata length",
		)
	}

	return &Metadata{
		length:  util.BytesToUint64(byteSlice[0:8]),
		lsbUsed: byteSlice[8],
		red:     util.BitToBool(byteSlice[9]),
		green:   util.BitToBool(byteSlice[10]),
		blue:    util.BitToBool(byteSlice[11]),
	}, nil
}

func MetadataFromImageData(data util.ImageData) (*Metadata, error) {
	var binaryBuffer bytes.Buffer

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(
		data,
		metadataPixelOffset,
		model.ChannelRedGreenBlue,
		pixelDataChannel,
	)

	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			value := pixelData.GetChannelValue(pixelChannel)
			hasBit := util.HasBit(value, metadataLsbPos)

			binaryBuffer.WriteRune(util.BoolToRune(hasBit))

			if binaryBuffer.Len() == metadataBinaryLength {
				return MetadataFromBinaryBuffer(&binaryBuffer)
			}
		}
	}

	return nil, errors.New("metadata: not found")
}
