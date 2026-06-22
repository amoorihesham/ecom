package auth

import (
	"database/sql"
	"ecom/internal/modules/auth/handler"
	repository "ecom/internal/modules/auth/repo"
	"ecom/internal/modules/auth/service"
	bcryptx "ecom/internal/shared/bcrypt"
	"log/slog"
	"net/http"
)

var _ handler.AuthHandler = (*handler.Handler)(nil)

func Initialize(mux *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	hasher := bcryptx.New(&bcryptx.Config{Cost: 12})
	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo, logger, hasher)
	handler := handler.NewAuthHandler(service, logger)

	mux.HandleFunc("POST /auth/register", handler.Register)

}
