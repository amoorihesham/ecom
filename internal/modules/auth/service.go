package auth

import (
	"context"
	"log/slog"
)

type IAuthRespository interface {
	Create(ctx context.Context, user *User) (*User, error)
}

type AuthService struct {
	repo   IAuthRespository
	logger *slog.Logger
}

func NewAuthService(repo IAuthRespository, logger *slog.Logger) *AuthService {
	return &AuthService{repo: repo, logger: logger}
}

func (service *AuthService) Create(ctx context.Context, user *User) (*User, error) {
	cUser, err := service.repo.Create(ctx, user)

	return cUser, err
}
