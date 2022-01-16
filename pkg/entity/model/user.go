package model

import (
	"StegoLSB/ent"
	"StegoLSB/graph/generated"
)

// User is the model entity for the User schema.
type User = ent.User

// NewUserInput represents a mutation input for creating users.
type NewUserInput = generated.NewUser

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput = generated.UpdateUser
