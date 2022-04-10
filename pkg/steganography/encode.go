package steganography

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/gqlgen"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/util"
)

// ValidateEncodeInput validates the generated.EncodeImageInput.
func ValidateEncodeInput(
	ctx context.Context,
	user *ent.User,
	input gqlgen.EncodeImageInput,
) error {
	if user == nil {
		if input.EncryptionKey != nil {
			return util.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify encryption key",
			)
		}

		if input.LsbUsed != 1 {
			return util.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify least significant bits",
			)
		}

		if input.Channel != util.ChannelRedGreenBlue {
			return util.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't specify channel",
			)
		}

		if input.EvenDistribution {
			return util.NewAuthorizationError(
				ctx,
				"encode: unauthorized users can't use even distribution",
			)
		}
	}

	if !ValidateLSB(byte(input.LsbUsed)) {
		return util.NewValidationError(
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
func Encode(input gqlgen.EncodeImageInput) (string, error) {
	imageData, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	encodeData, metadata, err := buildData(input, imageData)
	if err != nil {
		return "", err
	}

	SetNRGBAValues(
		imageData,
		encodeData,
		pixelDataOffset,
		byte(input.LsbUsed),
		input.Channel,
		metadata.GetDistributionDivisor(imageData),
	)

	imgBuffer, err := util.EncodeNRGBA(imageData.NRGBA)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	return base64.StdEncoding.EncodeToString(imgBuffer.Bytes()), nil
}

func buildData(
	input gqlgen.EncodeImageInput,
	imageData util.ImageData,
) ([]byte, *Metadata, error) {
	encryptedData, err := cryptography.Encrypt(
		[]byte(input.Data),
		input.EncryptionKey,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("encode: %w", err)
	}

	encodedLen := base64.RawURLEncoding.EncodedLen(len(encryptedData))
	encodeSlice := make([]byte, encodedLen)
	metadata := MetadataFromEncodeInput(input, encodedLen)
	available := imageData.PixelCount() - pixelDataOffset
	needed := metadata.PixelsNeeded()

	if needed > available {
		return nil, nil, fmt.Errorf(
			"encode: need %d pixels, but only %d is available with current config",
			needed,
			available,
		)
	}

	// encode everything only after we validate that we'll be able to proceed
	metadata.EncodeIntoImageData(imageData)
	base64.RawURLEncoding.Encode(encodeSlice, encryptedData)

	return encodeSlice, &metadata, nil
}
