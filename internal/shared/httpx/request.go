package httpx

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
