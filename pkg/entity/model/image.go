package model

import (
	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/image"
)

// Image is the model entity for the Image schema.
type Image = ent.Image

// ImageConnection is the connection containing edges to Image.
type ImageConnection = ent.ImageConnection

// ImageWhereInput represents input for filtering Image queries.
type ImageWhereInput = ent.ImageWhereInput

// ImageOrderInput represents order of Image queries.
type ImageOrderInput = ent.ImageOrder

// ImageOrderField represents the ordering field of Image queries.
type ImageOrderField = ent.ImageOrderField

// Channel represents the image channel.
type Channel = image.Channel
