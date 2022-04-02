package steganography

import (
	"bytes"
	"context"
	"errors"
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
	var binaryBuffer bytes.Buffer

	data, err := util.FileToImageData(input.Upload.File)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	metadata, err := MetadataFromImageData(data)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	pixelDataChannel := make(chan PixelData)
	go NRGBAPixels(
		data,
		pixelDataOffset,
		metadata.GetChannel(),
		pixelDataChannel,
	)

	lsbPosChannel := make(chan byte)
	go LSBPositions(metadata.lsbUsed, lsbPosChannel)

	expectedBinaryLength := metadata.GetBinaryLength()

	for pixelData := range pixelDataChannel {
		for _, pixelChannel := range pixelData.Channels {
			value := pixelData.GetChannelValue(pixelChannel)
			lsbPos := <-lsbPosChannel
			hasBit := util.HasBit(value, lsbPos)

			binaryBuffer.WriteRune(util.BoolToRune(hasBit))

			if binaryBuffer.Len() == expectedBinaryLength {
				return decodeData(&binaryBuffer, input.EncryptionKey)
			}
		}
	}

	return "", errors.New("decode: no message found")
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
