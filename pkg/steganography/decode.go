package steganography

import (
	"bytes"
	"context"
	"fmt"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

func ValidateDecodeInput(
	ctx context.Context,
	user *ent.User,
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

// Decode decodes a message from the given generated.DecodeImageInput input.
func Decode(input generated.DecodeImageInput) (string, error) {
	imageData, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	metadata, err := MetadataFromImageData(imageData)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	lsbPosChannel := make(chan byte)
	go LSBPositions(metadata.lsbUsed, lsbPosChannel)

	binaryBuffer, err := GetNRGBAValues(
		imageData,
		pixelDataOffset,
		func() byte {
			return <-lsbPosChannel
		},
		metadata.GetChannel(),
		metadata.GetDistributionDivisor(imageData),
		metadata.GetBinaryLength(),
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

	data, err := cryptography.Decrypt(byteSlice, encryptionKey)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	return string(data), nil
}
