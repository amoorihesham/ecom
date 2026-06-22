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
		httpx.Error(w, httpx.StatusFromCode(httpx.ErrBadRequest), httpx.ErrBadRequest, "invalid request body")
		return
	}

	created, _ := handler.service.Register(r.Context(), &payload)

	httpx.Created(w, created)

}
