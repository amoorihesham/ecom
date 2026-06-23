package handler

import (
	"ecom/internal/modules/auth/models"
	"ecom/internal/modules/auth/service"
	"ecom/internal/shared/httpx"
	"log/slog"
	"net/http"
)

var _ service.AuthService = (*service.Service)(nil)

func NewAuthHandler(service *service.Service, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterRequest
	if err := httpx.Decode(r, &payload); err != nil {
		appErr := httpx.HandleError(err)
		httpx.Error(w, httpx.StatusFromCode(appErr.Code), appErr.Code, "invalid request body")
		return
	}

	created, err := handler.service.Register(r.Context(), &payload)

	if err != nil {
		appErr := httpx.HandleError(err)
		httpx.Error(w, httpx.StatusFromCode(appErr.Code), appErr.Code, appErr.Message)
		return
	}

	httpx.Created(w, created)

}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.LoginRequest
	if err := httpx.Decode(r, &payload); err != nil {
		appErr := httpx.HandleError(err)
		httpx.Error(w, httpx.StatusFromCode(appErr.Code), appErr.Code, "invalid request body")
		return
	}

	accessToken, err := handler.service.Login(r.Context(), &payload)

	if err != nil {
		appErr := httpx.HandleError(err)
		httpx.Error(w, httpx.StatusFromCode(appErr.Code), appErr.Code, appErr.Message)
		return
	}
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    *accessToken,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	httpx.Success(w, accessToken)

}
