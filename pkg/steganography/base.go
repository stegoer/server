package steganography

const (
	// pixelDataOffset represents the amount of pixels to offset by when
	// encoding the actual data into an image. 1 is added so not even the
	// last pixel used by Metadata is altered in any way.
	pixelDataOffset = metadataBinaryLength/3 + 1
)
