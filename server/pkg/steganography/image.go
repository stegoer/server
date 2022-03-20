package steganography

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"

	entImage "github.com/kucera-lukas/stegoer/server/ent/image"
	"github.com/kucera-lukas/stegoer/server/pkg/util"
)

type ImageData struct {
	NRGBA  *image.NRGBA
	Width  int
	Height int
}

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
	data ImageData,
	channel entImage.Channel,
	resultChan chan PixelData,
) {
	red := util.IncludesRedChannel(channel)
	green := util.IncludesGreenChannel(channel)
	blue := util.IncludesBlueChannel(channel)

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

// FileToImageData reads file and returns ImageData.
func FileToImageData(file io.Reader) (ImageData, error) {
	img, err := ReadImageFile(file)
	if err != nil {
		return ImageData{NRGBA: nil, Width: 0, Height: 0}, err
	}

	nrgba, width, height := ImageToNRGBA(img)

	return ImageData{
		NRGBA:  nrgba,
		Width:  width,
		Height: height,
	}, nil
}

// ReadImageFile reads given file and returns image.Image.
func ReadImageFile(file io.Reader) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image file: %w", err)
	}

	return img, nil
}

// ImageToNRGBA converts image.Image to image.NRGBA.
func ImageToNRGBA(img image.Image) (*image.NRGBA, int, int) {
	bounds := img.Bounds()

	width, height := bounds.Dx(), bounds.Dy()
	ret := image.NewNRGBA(image.Rect(0, 0, width, height))

	draw.Draw(ret, ret.Bounds(), img, bounds.Min, draw.Src)

	return ret, width, height
}

// EncodeNRGBA encodes given nrgba image into a bytes.Buffer.
func EncodeNRGBA(nrgba *image.NRGBA) (*bytes.Buffer, error) {
	imgBuffer := new(bytes.Buffer)

	if err := png.Encode(imgBuffer, nrgba); err != nil {
		return nil, fmt.Errorf("error encoding NRGBA image: %w", err)
	}

	return imgBuffer, nil
}
