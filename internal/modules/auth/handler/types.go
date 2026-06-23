package handler

import (
	"ecom/internal/modules/auth/service"
	"log/slog"
	"net/http"
)

type AuthHandler interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type Handler struct {
	service service.AuthService
	logger  *slog.Logger
}
