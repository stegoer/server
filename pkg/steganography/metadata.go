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
	metadataLength                   = 13
	metadataBinaryLength             = metadataLength * util.BitLength
	metadataPixelOffset              = 0
	metadataLsbPos              byte = 1
	metadataDistributionDivisor      = 1
)

// Metadata represents information which was used to encode data into an image.
type Metadata struct {
	length           uint64
	lsbUsed          byte
	red              bool
	green            bool
	blue             bool
	evenDistribution bool
}

func (md Metadata) GetBinaryLength() uint64 {
	return md.length * util.BitLength
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

// PixelsNeeded returns needed pixels for encoding data based on its Metadata.
func (md Metadata) PixelsNeeded() uint64 {
	return md.GetBinaryLength() / uint64(
		md.lsbUsed) / uint64(md.GetChannel().Count())
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
			util.BoolToBit(md.evenDistribution),
		}...,
	)

	return result
}

func (md Metadata) GetDistributionDivisor(imageData util.ImageData) int {
	switch md.evenDistribution {
	case true:
		pixelsAvailable := imageData.PixelCount() - pixelDataOffset

		if divisor := int(pixelsAvailable / md.PixelsNeeded()); divisor > 0 {
			return divisor
		}

		fallthrough
	default:
		return 1
	}
}

func (md Metadata) EncodeIntoImageData(imageData util.ImageData) {
	SetNRGBAValues(
		imageData,
		md.ToByteArr(),
		metadataPixelOffset,
		metadataLsbPos,
		model.ChannelRedGreenBlue,
		metadataDistributionDivisor,
	)
}

func MetadataFromEncodeInput(
	input generated.EncodeImageInput,
	messageLength int,
) Metadata {
	return Metadata{
		length:           uint64(messageLength),
		lsbUsed:          byte(input.LsbUsed),
		red:              input.Channel.IncludesRed(),
		green:            input.Channel.IncludesGreen(),
		blue:             input.Channel.IncludesBlue(),
		evenDistribution: input.EvenDistribution,
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
		length:           util.BytesToUint64(byteSlice[0:8]),
		lsbUsed:          byteSlice[8],
		red:              util.BitToBool(byteSlice[9]),
		green:            util.BitToBool(byteSlice[10]),
		blue:             util.BitToBool(byteSlice[11]),
		evenDistribution: util.BitToBool(byteSlice[12]),
	}, nil
}

func MetadataFromImageData(imageData util.ImageData) (*Metadata, error) {
	binaryBuffer, err := GetNRGBAValues(
		imageData,
		metadataPixelOffset,
		metadataLsbPos,
		model.ChannelRedGreenBlue,
		metadataDistributionDivisor,
		metadataBinaryLength,
	)
	if err != nil {
		return nil, err
	}

	return MetadataFromBinaryBuffer(binaryBuffer)
}
