package handler

import (
	"ecom/internal/modules/auth/service"
	"log/slog"
	"net/http"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.AuthService
	logger  *slog.Logger
}
