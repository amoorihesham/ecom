package database

import (
	"ecom/internal/shared/httpx"

	"github.com/lib/pq"
)

func DBToAppError(err error) error {

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {

		switch pqErr.Code {

		case "23505": // unique_violation
			switch pqErr.Constraint {
			case "users_email_key":
				return httpx.NewError(httpx.ErrConflict, "email already exist", err)
			default:
				return httpx.NewError(httpx.ErrConflict, "resource already exist", err)
			}

		case "23503": // foreign_key_violation
			return httpx.NewError(
				httpx.ErrBadRequest,
				"invalid reference",
				err,
			)

		default:
			return httpx.NewError(
				httpx.ErrInternal,
				"database error",
				err,
			)
		}
	}

	return httpx.NewError(
		httpx.ErrInternal,
		"unknown error",
		err,
	)
}
