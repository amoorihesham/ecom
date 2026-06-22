package repository

import (
	"context"
	"ecom/internal/modules/auth/models"

	"github.com/google/uuid"
)

type AuthRespository interface {
	// Get(ctx context.Context, userId uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.RegisterRequest) (*models.User, error)
	IsExistByEmail(ctx context.Context, email string) (bool, error)
	IsExistById(ctx context.Context, userId uuid.UUID) (bool, error)
}
