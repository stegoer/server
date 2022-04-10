package steganography

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

// ValidateDecodeInput validates the generated.DecodeImageInput.
func ValidateDecodeInput(
	ctx context.Context,
	user *model.User,
	input generated.DecodeImageInput,
) *model.Error {
	if user == nil && input.EncryptionKey != nil {
		return model.NewAuthorizationError(
			ctx,
			"decode: unauthorized users can't specify encryption key",
		)
	}

	return nil
}

// Decode decodes the data from the given generated.DecodeImageInput input.
func Decode(input generated.DecodeImageInput) (string, error) {
	imageData, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	metadata, err := MetadataFromImageData(imageData)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	binaryBuffer, err := GetNRGBAValues(
		imageData,
		pixelDataOffset,
		metadata.lsbUsed,
		metadata.GetChannel(),
		metadata.GetDistributionDivisor(imageData),
		int(metadata.GetBinaryLength()),
	)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	return decodeData(binaryBuffer, input.EncryptionKey)
}

func decodeData(
	binaryBuffer *bytes.Buffer,
	encryptionKey *string,
) (string, error) {
	byteSlice, err := util.BinaryBufferToBytes(binaryBuffer)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	decodeSlice := make([]byte, base64.RawURLEncoding.DecodedLen(len(byteSlice)))

	bytesWritten, err := base64.RawURLEncoding.Decode(decodeSlice, byteSlice)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	data, err := cryptography.Decrypt(decodeSlice[:bytesWritten], encryptionKey)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	return string(data), nil
}
