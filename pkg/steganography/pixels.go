package steganography

import (
	"bytes"
	"errors"
	"image/color"

	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

type ChannelType byte

type PixelData struct {
	Width    int
	Height   int
	Channels []ChannelType
	Color    *color.NRGBA
}

const (
	RedChannel ChannelType = iota
	GreenChannel
	BlueChannel

	channelTypeSliceCapacity = 3
)

func (ct ChannelType) IsRed() bool {
	return ct == RedChannel
}

func (ct ChannelType) IsGreen() bool {
	return ct == GreenChannel
}

func (ct ChannelType) IsBlue() bool {
	return ct == BlueChannel
}

func (pd PixelData) GetRed() byte {
	return pd.Color.R
}

func (pd PixelData) GetGreen() byte {
	return pd.Color.G
}

func (pd PixelData) GetBlue() byte {
	return pd.Color.B
}

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

func (pd *PixelData) SetRed(value byte) {
	pd.Color.R = value
}

func (pd *PixelData) SetGreen(value byte) {
	pd.Color.G = value
}

func (pd *PixelData) SetBlue(value byte) {
	pd.Color.B = value
}

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
		for _, pixelChannel := range pixelData.Channels {
			for _, lsbPos := range lsbSlice {
				dataBit, ok := <-bitChannel
				if !ok {
					hasBits = false

					break
				}

				pixelData.SetChannelValue(pixelChannel, dataBit, lsbPos)
			}

			if !hasBits {
				break
			}
		}

		imageData.NRGBA.SetNRGBA(pixelData.Width, pixelData.Height, *pixelData.Color)

		if !hasBits {
			return
		}
	}
}

func GetNRGBAValues(
	imageData util.ImageData,
	pixelOffset int,
	lsbUsed byte,
	channel model.Channel,
	distributionDivisor int,
	bufferLength uint64,
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

				if uint64(binaryBuffer.Len()) == bufferLength {
					return &binaryBuffer, nil
				}
			}
		}
	}

	return nil, errors.New("requested NRGBA values not found")
}
