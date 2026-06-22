package service

import (
	"context"
	"ecom/internal/modules/auth/models"
	repository "ecom/internal/modules/auth/repo"
	bcryptx "ecom/internal/shared/bcrypt"
	"ecom/internal/shared/jwt"
	"log/slog"
)

type AuthService interface {
	Register(ctx context.Context, user *models.RegisterRequest) (*models.User, error)
}

type Service struct {
	repo   repository.AuthRespository
	logger *slog.Logger
	hasher bcryptx.HashService
	jwt    jwt.JWTService
}
