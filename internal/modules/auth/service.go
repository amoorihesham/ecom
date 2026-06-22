package auth

import (
	"context"
	"log/slog"
)

type IAuthRespository interface {
	// Get(ctx context.Context, userId uuid.UUID) (*User, httpx.AppError)
	Create(ctx context.Context, user *RegisterRequest) (*User, error)
}

type AuthService struct {
	repo   IAuthRespository
	logger *slog.Logger
}

func NewAuthService(repo IAuthRespository, logger *slog.Logger) *AuthService {
	return &AuthService{repo: repo, logger: logger}
}

func (service *AuthService) Register(ctx context.Context, user *RegisterRequest) (*User, error) {

	cUser, err := service.repo.Create(ctx, user)

	return cUser, err
}

// func (service *AuthService) Login(ctx context.Context, userId uuid.UUID) (*User, error) {
// 	user, err := service.repo.Get(ctx, userId)

// 	return user, err
// }

// func (service *AuthService) Refresh(ctx context.Context, token string) (*User, error) {
// 	user, err := service.repo.Get(ctx, userId)

// 	return user, err
// }

// func (service *AuthService) Logout(ctx context.Context, token string) error {
// 	user, err := service.repo.Get(ctx, userId)

// 	return user, err
// }
