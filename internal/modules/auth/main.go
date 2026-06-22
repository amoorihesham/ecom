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
	Register(w http.ResponseWriter, r *http.Request)
	// Login(w http.ResponseWriter, r *http.Request)
	// Refresh(w http.ResponseWriter, r *http.Request)
	// Logout(w http.ResponseWriter, r *http.Request)
}

var _ IAuthHandler = (*Handler)(nil)

func Initialize(mux *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo, logger)
	handler := NewAuthHandler(service, logger)

	mux.HandleFunc("POST /auth/register", handler.Register)
	// mux.HandleFunc("POST /auth/login", handler.Login)
	// mux.HandleFunc("POST /auth/refresh", handler.Refresh)
	// mux.HandleFunc("POST /auth/logout", handler.Logout)

}
