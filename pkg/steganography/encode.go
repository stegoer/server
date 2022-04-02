package steganography

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

func ValidateEncodeInput(
	ctx context.Context,
	user *ent.User,
	input generated.EncodeImageInput,
) *model.Error {
	if user == nil {
		if input.EncryptionKey != nil {
			return model.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify encryption key",
			)
		}

		if input.LsbUsed != 1 {
			return model.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify least significant bits",
			)
		}

		if input.Channel != model.ChannelRedGreenBlue {
			return model.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify channel",
			)
		}
	}

	if !ValidateLSB(input.LsbUsed) {
		return model.NewValidationError(
			ctx,
			fmt.Sprintf(
				"encode: %d is not a valid number of least significant bits used",
				input.LsbUsed,
			),
		)
	}

	return nil
}

// Encode encodes a message into the given graphql.Upload file based on input.
// Returns the image data base64 encoded.
func Encode(input generated.EncodeImageInput) (string, error) {
	data, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	encodeData, err := buildData(input, data)
	if err != nil {
		return "", err
	}

	bitChannel := make(chan byte)
	go util.ByteArrToBits(encodeData, bitChannel)

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(data, pixelDataOffset, input.Channel, pixelDataChannel)

	lsbPosChannel := make(chan byte)
	go LSBPositions(byte(input.LsbUsed), lsbPosChannel)

pixelIterator:
	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			dataBit, ok := <-bitChannel
			// there are no more bits in the bit channel
			if !ok {
				break pixelIterator
			}

			lsbPos := <-lsbPosChannel

			pixelData.SetChannelValue(pixelChannel, dataBit, lsbPos)
		}

		data.NRGBA.SetNRGBA(pixelData.Width, pixelData.Height, *pixelData.Color)
	}

	imgBuffer, err := util.EncodeNRGBA(data.NRGBA)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(imgBuffer.Bytes()), nil
}

func buildData(
	input generated.EncodeImageInput,
	imageData util.ImageData,
) ([]byte, error) {
	encryptedData, err := cryptography.Encrypt(
		[]byte(input.Data),
		input.EncryptionKey,
	)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	encryptedLen := len(encryptedData)

	if max := maxBytesToEncode(
		imageData,
		input,
	); metadataLength+encryptedLen >= max {
		return nil, fmt.Errorf(
			"encode: max data length is %d, got %d",
			max,
			encryptedLen,
		)
	}

	MetadataFromEncodeInput(input, encryptedLen).EncodeIntoImageData(imageData)

	return encryptedData, nil
}

// maxBytesToEncode calculates a maximum amount of bytes which can be encoded
// based on the bounds of given image.NRGBA and generated.EncodeImageInput.
func maxBytesToEncode(
	data util.ImageData,
	input generated.EncodeImageInput,
) int {
	return (data.Width * data.Height * input.Channel.Count() * input.LsbUsed) / bitLength
}
