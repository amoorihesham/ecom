package repository

import (
	"context"
	"database/sql"
	"ecom/internal/modules/auth/models"
	"ecom/internal/shared/httpx"

	"github.com/google/uuid"
)

type AuthRepository struct {
	db *sql.DB
}

var _ AuthRespository = (*AuthRepository)(nil)

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(ctx context.Context, user *models.RegisterRequest) (*models.User, error) {
	q := `INSERT INTO USERS (email, password_hash, full_name)
	VALUES ($1, $2, $3) RETURNING id, public_id, email, full_name, role, created_at`

	var created models.User
	row := r.db.QueryRowContext(ctx, q, user.Email, user.PasswordHash, user.Fullname)

	if err := row.Scan(&created.ID, &created.PublicId, &created.Email, &created.Fullname, &created.Role, &created.CreatedAt); err != nil {
		return nil, httpx.HandleError(err)
	}

	return &created, nil
}

func (r *AuthRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	q := `SELECT id, public_id, email, password_hash, full_name, role, created_at FROM users WHERE email = $1`

	var found models.User
	row := r.db.QueryRowContext(ctx, q, email)
	if err := row.Scan(&found.ID, &found.PublicId, &found.Email, &found.PasswordHash, &found.Fullname, &found.Role, &found.CreatedAt); err != nil {

		return nil, httpx.HandleError(err)
	}
	return &found, nil

}

func (r *AuthRepository) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	q := `SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
		)`
	var exists bool

	if err := r.db.QueryRowContext(ctx, q, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
func (r *AuthRepository) IsExistById(ctx context.Context, userId uuid.UUID) (bool, error) {
	q := `SELECT id FROM users WHERE public_id = $1`
	var id int64

	row := r.db.QueryRowContext(ctx, q, userId)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
