package httpx

import "net/http"

func StatusFromCode(code ErrorCode) int {
	switch code {

	case ErrBadRequest:
		return http.StatusBadRequest

	case ErrUnauthorized:
		return http.StatusUnauthorized

	case ErrForbidden:
		return http.StatusForbidden

	case ErrNotFound:
		return http.StatusNotFound

	case ErrConflict:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
