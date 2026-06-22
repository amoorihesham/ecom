package auth

import (
	"context"
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

var _ IAuthRespository = (*AuthRepository)(nil)

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(ctx context.Context, user *User) (*User, error) {
	q := `INSERT INTO USERS (email, password_hash, fullname)
	VALUES ($1, $2, $3) RETURNING *`

	var created User
	row := r.db.QueryRowContext(ctx, q, user.Email, user.PasswordHash, user.Fullname)

	if err := row.Scan(&created.ID, &created.PublicId, &created.Email, &created.Fullname, &created.Role, &created.CreatedAt); err != nil {
		return &User{}, nil
	}

	return &created, nil
}
