package httpx

import (
	"encoding/json"
	"net/http"
)

type ErrorEnvlop struct {
	Code    int
	Message string
	Details map[string]string
}

func SendErrorResponse(w http.ResponseWriter, envlop *ErrorEnvlop) {
	json.NewEncoder(w).Encode(envlop)
}
