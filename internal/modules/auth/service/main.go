package service

import (
	"context"
	"ecom/internal/modules/auth/models"
	repository "ecom/internal/modules/auth/repo"
	bcryptx "ecom/internal/shared/bcrypt"
	"ecom/internal/shared/httpx"
	"ecom/internal/shared/jwt"
	"log/slog"
)

func NewAuthService(repo repository.AuthRespository, logger *slog.Logger, hasher bcryptx.HashService, jwt jwt.JWTService) *Service {
	return &Service{repo: repo, logger: logger, hasher: hasher, jwt: jwt}
}

func (service *Service) Register(ctx context.Context, user *models.RegisterRequest) (*models.User, error) {
	exists, err := service.repo.IsExistByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, httpx.NewError(httpx.ErrConflict, "Email in use")
	}

	hasedPassword, err := service.hasher.Hash(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	cUser, err := service.repo.Create(ctx, &models.RegisterRequest{Email: user.Email, Fullname: user.Fullname, PasswordHash: hasedPassword})

	return cUser, err
}

func (service *Service) Login(ctx context.Context, user *models.LoginRequest) (*string, error) {
	existing, err := service.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, httpx.NewError(httpx.ErrNotFound, "User not found")
	}

	match := service.hasher.Compare(user.Password, existing.PasswordHash)
	if !match {

		return nil, httpx.NewError(httpx.ErrBadRequest, "Invalid credentials")
	}

	accessToken, err := service.jwt.Generate(httpx.AuthUser{PublicID: existing.PublicId.String(), Role: existing.Role})
	if err != nil {
		return nil, httpx.NewError(httpx.ErrConflict, "Can not generate your access token")
	}

	return &accessToken, nil
}
