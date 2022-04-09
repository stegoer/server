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

		if input.EvenDistribution {
			return model.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't use even distribution",
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
	imageData, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	encodeData, metadata, err := buildData(input, imageData)
	if err != nil {
		return "", err
	}

	lsbPosChannel := make(chan byte)
	go LSBPositions(byte(input.LsbUsed), lsbPosChannel)

	SetNRGBAValues(
		imageData,
		encodeData,
		pixelDataOffset,
		func() byte {
			return <-lsbPosChannel
		},
		input.Channel,
		metadata.GetDistributionDivisor(imageData),
	)

	imgBuffer, err := util.EncodeNRGBA(imageData.NRGBA)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(imgBuffer.Bytes()), nil
}

func buildData(
	input generated.EncodeImageInput,
	imageData util.ImageData,
) ([]byte, *Metadata, error) {
	encryptedData, err := cryptography.Encrypt(
		[]byte(input.Data),
		input.EncryptionKey,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("encode: %w", err)
	}

	encryptedLen := len(encryptedData)

	if max := maxBytesToEncode(
		imageData,
		input,
	) + metadataLength; encryptedLen >= max {
		return nil, nil, fmt.Errorf(
			"encode: max data length is %d, got %d",
			max,
			encryptedLen,
		)
	}

	metadata := MetadataFromEncodeInput(input, encryptedLen)
	metadata.EncodeIntoImageData(imageData)

	return encryptedData, &metadata, nil
}

// maxBytesToEncode calculates a maximum amount of bytes which can be encoded
// based on the bounds of given image.NRGBA and generated.EncodeImageInput.
func maxBytesToEncode(
	data util.ImageData,
	input generated.EncodeImageInput,
) int {
	return (data.Width * data.Height * input.Channel.Count() * input.LsbUsed) / bitLength
}
