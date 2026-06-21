package httpx

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
