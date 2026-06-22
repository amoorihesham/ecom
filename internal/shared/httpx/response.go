package httpx

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   *AppError   `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func Success(w http.ResponseWriter, data any) {
	JSON(w, http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Created(w http.ResponseWriter, data any) {
	JSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, code ErrorCode, message string) {
	JSON(w, status, Response{
		Success: false,
		Error: &AppError{
			Code:    code,
			Message: message,
		},
	})
}
