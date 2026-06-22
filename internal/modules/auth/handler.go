package auth

import (
	"context"
	"ecom/internal/shared/httpx"
	"encoding/json"
	"log/slog"
	"net/http"
)

type RegisterRequest struct {
	Email        string `json:"email"`
	Fullname     string `json:"fullname"`
	PasswordHash string `json:"password"`
}
type IAuthService interface {
	Register(ctx context.Context, user *RegisterRequest) (*User, error)
	// Login(ctx context.Context, userId uuid.UUID) (*User, error)
	// Refresh(ctx context.Context, token string) (*User, error)
	// Logout(ctx context.Context, token string) error
}

type Handler struct {
	service IAuthService
	logger  *slog.Logger
}

var _ IAuthService = (*AuthService)(nil)

func NewAuthHandler(service IAuthService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		httpx.Error(w, httpx.StatusFromCode(httpx.ErrBadRequest), httpx.ErrBadRequest, "error")
		return
	}
	created, err := handler.service.Register(r.Context(), &payload)
	if err != nil {
		if err, ok := err.(*httpx.AppError); ok {
			httpx.Error(w, httpx.StatusFromCode(err.Code), err.Code, err.Message)
			return
		}
	}

	httpx.Created(w, created)
}

// func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
// 	token, err := handler.service.Login(r.Context(), uuid.UUID{})

// }

// func (handler *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
// 	token, err := handler.service.Login(r.Context(), uuid.UUID{})

// }

// func (handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
// 	token, err := handler.service.Login(r.Context(), uuid.UUID{})

// }
