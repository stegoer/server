// inspired by https://github.com/manakuro/golang-clean-architecture-ent-gqlgen/blob/main/pkg/entity/model/error.go

package util

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorCode represents an error code.
type ErrorCode string

const (
	// authorization represents ErrorCode related to authorization Error.
	authorization ErrorCode = "AUTHORIZATION_ERROR"

	// db represents ErrorCode related to database Error.
	db ErrorCode = "DB_ERROR"

	// internalServer represents ErrorCode related to internal server Error.
	internalServer ErrorCode = "INTERNAL_SERVER_ERROR"

	// notFound represents ErrorCode related to not found Error.
	notFound ErrorCode = "NOT_FOUND_ERROR"

	// validation represents ErrorCode related validation Error.
	validation ErrorCode = "VALIDATION_ERROR"
)

// NewAuthorizationError returns a new gqlerror.Error with authorization ErrorCode.
func NewAuthorizationError(
	ctx context.Context,
	message string,
) *gqlerror.Error {
	return newError(ctx, message, authorization)
}

// NewDBError returns a new gqlerror.Error with db ErrorCode.
func NewDBError(
	ctx context.Context,
	message string,
) *gqlerror.Error {
	return newError(ctx, message, db)
}

// NewInternalServerError returns a new gqlerror.Error with internalServer ErrorCode.
func NewInternalServerError(
	ctx context.Context,
	message string,
) *gqlerror.Error {
	return newError(ctx, message, internalServer)
}

// NewNotFoundError returns a new gqlerror.Error with notFound ErrorCode.
func NewNotFoundError(
	ctx context.Context,
	message string,
) *gqlerror.Error {
	return newError(ctx, message, notFound)
}

// NewValidationError returns a new gqlerror.Error with validation ErrorCode.
func NewValidationError(
	ctx context.Context,
	message string,
) *gqlerror.Error {
	return newError(ctx, message, validation)
}

// newError creates and returns a new gqlerror.Error.
func newError(
	ctx context.Context,
	message string,
	code ErrorCode,
) *gqlerror.Error {
	return &gqlerror.Error{ //nolint:exhaustivestruct
		Message:    message,
		Path:       graphql.GetPath(ctx),
		Extensions: map[string]interface{}{"code": code},
	}
}
