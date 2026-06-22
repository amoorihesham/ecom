package httpx

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return NewError(ErrBadRequest, "Bad request can not parse request payload", err)
	}
	return nil
}
