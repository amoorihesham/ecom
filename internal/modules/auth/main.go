package auth

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           int64     `json:"-"`
	PublicId     uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Fullname     string    `json:"full_name"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IAuthHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

var _ IAuthHandler = (*Handler)(nil)

func Initialize(mux *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo, logger)
	handler := NewAuthHandler(service, logger)

	mux.HandleFunc("POST /auth/register", handler.Create)
	mux.HandleFunc("POST /auth/login", handler.Create)
	mux.HandleFunc("POST /auth/refresh", handler.Create)
	mux.HandleFunc("POST /auth/logout", handler.Create)

}
