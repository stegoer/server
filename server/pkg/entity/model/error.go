package model

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	// DBError is error code of database.
	DBError = "DB_ERROR"
	// GraphQLError is error code of graphql.
	GraphQLError = "GRAPHQL_ERROR"
	// AuthorizationError is error code of authorization error.
	AuthorizationError = "AUTHORIZATION_ERROR"
	// NotFoundError is error code of not found.
	NotFoundError = "NOT_FOUND_ERROR"
	// ValidationError is error code of validation.
	ValidationError = "VALIDATION_ERROR"
	// BadRequestError is error code of request.
	BadRequestError = "BAD_REQUEST_ERROR"
	// InternalServerError is error code of server error.
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

// Error represents gqlerror.Error type.
type Error = gqlerror.Error

// Extensions represents Error extensions.
type Extensions map[string]interface{}

// NewDBError returns error related to database.
func NewDBError(
	ctx context.Context,
	message string,
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": DBError},
	)
}

// NewGraphQLError returns error related to graphql.
func NewGraphQLError(
	ctx context.Context,
	message string,
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": GraphQLError},
	)
}

// NewAuthorizationError returns error related to authorization.
func NewAuthorizationError(
	ctx context.Context,
	message string,
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": AuthorizationError},
	)
}

// NewNotFoundError returns error related to not found.
func NewNotFoundError(
	ctx context.Context,
	message string,
	value interface{},
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": NotFoundError, "value": value},
	)
}

// NewBadRequestError returns error related to bad request.
func NewBadRequestError(
	ctx context.Context,
	message string,
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": BadRequestError},
	)
}

// NewValidationError returns error related to validation.
func NewValidationError(
	ctx context.Context,
	message string,
	value interface{},
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": ValidationError, "value": value},
	)
}

// NewInternalServerError returns error related to internal server error.
func NewInternalServerError(
	ctx context.Context,
	message string,
) *Error {
	return newError(
		ctx,
		message,
		Extensions{"code": InternalServerError},
	)
}

// newError creates a new Error.
func newError(
	ctx context.Context,
	message string,
	extensions Extensions,
) *Error {
	return &gqlerror.Error{ //nolint:exhaustivestruct
		Path:       graphql.GetPath(ctx),
		Message:    message,
		Extensions: extensions,
	}
}
