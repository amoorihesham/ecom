package auth

import (
	"context"
	"log/slog"
	"net/http"
)

type IAuthService interface {
	Create(ctx context.Context, user *User) (*User, error)
}

type Handler struct {
	service IAuthService
	logger  *slog.Logger
}

var _ IAuthService = (*AuthService)(nil)

func NewAuthHandler(service IAuthService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	u, err := handler.service.Create(r.Context(), &User{})
	handler.logger.Info("USER", "user", u)
	handler.logger.Info("Error", "err", err)
}
