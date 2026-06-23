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
	service.logger.Debug("IsExistByEmail", "err", err)

	if exists {
		return nil, httpx.NewError(httpx.ErrConflict, "Email in use")
	}
	service.logger.Debug("exists", "err", err)

	hasedPassword, err := service.hasher.Hash(user.PasswordHash)
	if err != nil {
		return nil, err
	}
	service.logger.Debug("hasedPassword", "err", err)

	cUser, err := service.repo.Create(ctx, &models.RegisterRequest{Email: user.Email, Fullname: user.Fullname, PasswordHash: hasedPassword})
	service.logger.Debug("CREATE", "err", err)
	return cUser, err
}
