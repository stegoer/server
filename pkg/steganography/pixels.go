package steganography

import (
	"bytes"
	"errors"
	"image/color"

	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

// ChannelType represents a pixel color channel.
type ChannelType byte

// PixelData represents data of one particular pixel of an image.
type PixelData struct {
	Width    int
	Height   int
	Channels []ChannelType
	Color    *color.NRGBA
}

const (
	// RedChannel represents the red ChannelType.
	RedChannel ChannelType = iota
	// GreenChannel represents the red ChannelType.
	GreenChannel
	// BlueChannel represents the red ChannelType.
	BlueChannel

	channelTypeSliceCapacity = 3
)

// IsRed returns whether the ChannelType represents a RedChannel.
func (ct ChannelType) IsRed() bool {
	return ct == RedChannel
}

// IsGreen returns whether the ChannelType represents a GreenChannel.
func (ct ChannelType) IsGreen() bool {
	return ct == GreenChannel
}

// IsBlue returns whether the ChannelType represents a BlueChannel.
func (ct ChannelType) IsBlue() bool {
	return ct == BlueChannel
}

// GetRed returns the underlying value of the RedChannel of the PixelData.
func (pd PixelData) GetRed() byte {
	return pd.Color.R
}

// GetGreen returns the underlying value of the GreenChannel of the PixelData.
func (pd PixelData) GetGreen() byte {
	return pd.Color.G
}

// GetBlue returns the underlying value of the BlueChannel of the PixelData.
func (pd PixelData) GetBlue() byte {
	return pd.Color.B
}

// GetChannelValue returns the value of the ChannelType of the PixelData.
func (pd PixelData) GetChannelValue(channel ChannelType) byte {
	switch {
	case channel.IsRed():
		return pd.GetRed()
	case channel.IsGreen():
		return pd.GetGreen()
	case channel.IsBlue():
		return pd.GetBlue()
	default:
		// should be unreachable
		return 0
	}
}

// SetRed sets the RedChannel of the PixelData.
func (pd *PixelData) SetRed(value byte) {
	pd.Color.R = value
}

// SetGreen sets the GreenChannel of the PixelData.
func (pd *PixelData) SetGreen(value byte) {
	pd.Color.G = value
}

// SetBlue sets the BlueChannel of the PixelData.
func (pd *PixelData) SetBlue(value byte) {
	pd.Color.B = value
}

// SetChannelValue sets the value of ChannelType of the PixelData on lsbPos.
func (pd PixelData) SetChannelValue(
	channel ChannelType,
	value byte,
	lsbPos byte,
) {
	switch {
	case channel.IsRed():
		pd.SetRed(util.GetUpdatedByte(value, pd.GetRed(), lsbPos))
	case channel.IsGreen():
		pd.SetGreen(util.GetUpdatedByte(value, pd.GetGreen(), lsbPos))
	case channel.IsBlue():
		pd.SetBlue(util.GetUpdatedByte(value, pd.GetBlue(), lsbPos))
	}
}

// NRGBAPixels sends PixelData of util.ImageData based on given parameters.
func NRGBAPixels(
	data util.ImageData,
	pixelOffset int,
	channel model.Channel,
	distributionDivisor int,
	resultChan chan PixelData,
) {
	var pixelCount int

	red := channel.IncludesRed()
	green := channel.IncludesGreen()
	blue := channel.IncludesBlue()

	for width := 0; width < data.Width; width++ {
		for height := 0; height < data.Height; height++ {
			pixelCount++

			if pixelCount <= pixelOffset || pixelCount%distributionDivisor != 0 {
				continue
			}

			channels := make([]ChannelType, 0, channelTypeSliceCapacity)
			nrgbaColor := data.NRGBA.NRGBAAt(width, height)

			if red {
				channels = append(channels, RedChannel)
			}

			if green {
				channels = append(channels, GreenChannel)
			}

			if blue {
				channels = append(channels, BlueChannel)
			}

			resultChan <- PixelData{
				Width:    width,
				Height:   height,
				Channels: channels,
				Color:    &nrgbaColor,
			}
		}
	}

	close(resultChan)
}

// SetNRGBAValues sets ChannelType values into util.ImageData based on params.
func SetNRGBAValues(
	imageData util.ImageData,
	encodeData []byte,
	pixelOffset int,
	lsbUsed byte,
	channel model.Channel,
	distributionDivisor int,
) {
	bitChannel := make(chan byte)
	go util.ByteArrToBits(encodeData, bitChannel)

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(
		imageData,
		pixelOffset,
		channel,
		distributionDivisor,
		pixelDataChannel,
	)

	lsbSlice := LSBSlice(lsbUsed)

	hasBits := true

	for pixelData := range pixelDataChannel {
	channelIterator:
		for _, pixelChannel := range pixelData.Channels {
			for _, lsbPos := range lsbSlice {
				dataBit, ok := <-bitChannel
				if !ok {
					hasBits = false

					break channelIterator
				}

				pixelData.SetChannelValue(pixelChannel, dataBit, lsbPos)
			}
		}

		imageData.NRGBA.SetNRGBA(pixelData.Width, pixelData.Height, *pixelData.Color)

		if !hasBits {
			return
		}
	}
}

// GetNRGBAValues returns bytes.Buffer with ChannelType values.
func GetNRGBAValues(
	imageData util.ImageData,
	pixelOffset int,
	lsbUsed byte,
	channel model.Channel,
	distributionDivisor int,
	bufferLength int,
) (*bytes.Buffer, error) {
	var binaryBuffer bytes.Buffer

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(
		imageData,
		pixelOffset,
		channel,
		distributionDivisor,
		pixelDataChannel,
	)

	lsbSlice := LSBSlice(lsbUsed)

	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			for _, lsbPos := range lsbSlice {
				hasBit := util.HasBit(
					pixelData.GetChannelValue(pixelChannel),
					lsbPos,
				)

				binaryBuffer.WriteRune(util.BoolToRune(hasBit))

				if binaryBuffer.Len() == bufferLength {
					return &binaryBuffer, nil
				}
			}
		}
	}

	return nil, errors.New("malformed image data")
}
