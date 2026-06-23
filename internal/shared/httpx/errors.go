package httpx

import (
	"github.com/lib/pq"
)

type ErrorCode string

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

var (
	ErrBadRequest   ErrorCode = "BAD_REQUEST"
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrForbidden    ErrorCode = "FORBIDDEN"
	ErrNotFound     ErrorCode = "NOT_FOUND"
	ErrConflict     ErrorCode = "CONFLICT"
	ErrInternal     ErrorCode = "INTERNAL_ERROR"
)

func (ap *AppError) Error() string {
	return ap.Err.Error()
}

func NewError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func HandleError(err error) *AppError {
	if dbErr, ok := err.(*pq.Error); ok {
		return handleDatabaseErr(dbErr)
	}

	if appErr, ok := err.(*AppError); ok {
		return handleAppErr(appErr)
	}

	return &AppError{Code: ErrInternal, Message: "Internal server error"}

}

func handleAppErr(err *AppError) *AppError {
	return NewError(err.Code, err.Message)
}

func handleDatabaseErr(err *pq.Error) *AppError {
	switch err.Code {
	case "23505": // unique_violation
		switch err.Constraint {
		case "users_email_key":
			return NewError(ErrConflict, "email already exist")
		default:
			return NewError(ErrConflict, "resource already exist")
		}

	case "23503": // foreign_key_violation
		return NewError(
			ErrBadRequest,
			"invalid reference",
		)

	default:
		return NewError(
			ErrInternal,
			"database error",
		)
	}

}
