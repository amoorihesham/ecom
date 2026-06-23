package auth

import (
	"database/sql"
	"ecom/internal/modules/auth/handler"
	repository "ecom/internal/modules/auth/repo"
	"ecom/internal/modules/auth/service"
	bcryptx "ecom/internal/shared/bcrypt"
	"ecom/internal/shared/jwt"
	"log/slog"
	"net/http"
	"time"
)

var _ handler.AuthHandler = (*handler.Handler)(nil)

func Initialize(mux *http.ServeMux, db *sql.DB, logger *slog.Logger, cfg *AuthConfig) {
	hasher := bcryptx.New(&bcryptx.Config{Cost: cfg.BcryptCost})
	jwt := jwt.New(&jwt.Config{Secret: cfg.JWTSecret, Issuer: "ecom-api", AccessExpiry: 15 * time.Minute})
	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo, logger, hasher, jwt)
	handler := handler.NewAuthHandler(service, logger)

	mux.HandleFunc("POST /auth/register", handler.Register)

}
