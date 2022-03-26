package steganography

import (
	"image/color"

	entImage "github.com/stegoer/server/ent/image"
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

func (pd *PixelData) SetRed(value byte) {
	pd.Color.R = value
}

func (pd *PixelData) SetGreen(value byte) {
	pd.Color.G = value
}

func (pd *PixelData) SetBlue(value byte) {
	pd.Color.B = value
}

func NRGBAPixels(
	data util.ImageData,
	channel entImage.Channel,
	resultChan chan PixelData,
) {
	red := IncludesRedChannel(channel)
	green := IncludesGreenChannel(channel)
	blue := IncludesBlueChannel(channel)

	for width := 0; width < data.Width; width++ {
		for height := 0; height < data.Height; height++ {
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
