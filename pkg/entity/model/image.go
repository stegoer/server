package model

import (
	"StegoLSB/ent"
	"StegoLSB/ent/image"
	"StegoLSB/graph/generated"
)

// Image is the model entity for the Image schema.
type Image = ent.Image

// ImageConnection is the connection containing edges to Image.
type ImageConnection = ent.ImageConnection

// ImageWhereInput represents input for filtering Image queries.
type ImageWhereInput = ent.ImageWhereInput

// ImageOrderInput represents order of Image queries.
type ImageOrderInput = ent.ImageOrder

// Channel represents the image channel
type Channel = image.Channel

// NewImageInput represents a mutation input for creating users.
type NewImageInput = generated.NewImage
