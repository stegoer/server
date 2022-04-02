package model

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ErrorCode string

const (
	// DBError is error code of database.
	DBError ErrorCode = "DB_ERROR"
	// GraphQLError is error code of graphql.
	GraphQLError ErrorCode = "GRAPHQL_ERROR"
	// AuthorizationError is error code of authorization error.
	AuthorizationError ErrorCode = "AUTHORIZATION_ERROR"
	// NotFoundError is error code of not found.
	NotFoundError ErrorCode = "NOT_FOUND_ERROR"
	// ValidationError is error code of validation.
	ValidationError ErrorCode = "VALIDATION_ERROR"
	// BadRequestError is error code of request.
	BadRequestError ErrorCode = "BAD_REQUEST_ERROR"
	// InternalServerError is error code of server error.
	InternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
)

// Error represents global error type.
type Error = gqlerror.Error

// NewDBError returns error related to database.
func NewDBError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, DBError)
}

// NewGraphQLError returns error related to graphql.
func NewGraphQLError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, GraphQLError)
}

// NewAuthorizationError returns error related to authorization.
func NewAuthorizationError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, AuthorizationError)
}

// NewNotFoundError returns error related to not found.
func NewNotFoundError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, NotFoundError)
}

// NewBadRequestError returns error related to bad request.
func NewBadRequestError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, BadRequestError)
}

// NewValidationError returns error related to validation.
func NewValidationError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, ValidationError)
}

// NewInternalServerError returns error related to internal server error.
func NewInternalServerError(
	ctx context.Context,
	message string,
) *Error {
	return newError(ctx, message, InternalServerError)
}

// newError creates a new Error.
func newError(
	ctx context.Context,
	message string,
	code ErrorCode,
) *Error {
	return &Error{ //nolint:exhaustivestruct
		Message:    message,
		Path:       graphql.GetPath(ctx),
		Extensions: map[string]interface{}{"code": code},
	}
}
