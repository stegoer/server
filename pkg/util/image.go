package util

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
)

// ImageData represents a data of an image.NRGBA and Width and Height.
type ImageData struct {
	NRGBA  *image.NRGBA
	Width  int
	Height int
}

// PixelCount returns the number of pixels of the image represented.
func (id ImageData) PixelCount() uint64 {
	return uint64(id.Width * id.Height)
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
	img, format, err := image.Decode(file)

	if err != nil {
		return nil, fmt.Errorf("failed to decode image file: %w", err)
	} else if format != "png" {
		return nil, fmt.Errorf("unsupported image format: %s", format)
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

// EncodeNRGBA encodes given image.NRGBA into a bytes.Buffer.
func EncodeNRGBA(nrgba *image.NRGBA) (*bytes.Buffer, error) {
	imgBuffer := new(bytes.Buffer)

	if err := png.Encode(imgBuffer, nrgba); err != nil {
		return nil, fmt.Errorf("error encoding NRGBA image: %w", err)
	}

	return imgBuffer, nil
}
