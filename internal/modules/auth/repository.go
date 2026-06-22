package auth

import (
	"context"
	"database/sql"
	"ecom/internal/shared/database"
)

type AuthRepository struct {
	db *sql.DB
}

var _ IAuthRespository = (*AuthRepository)(nil)

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(ctx context.Context, user *RegisterRequest) (*User, error) {
	q := `INSERT INTO USERS (email, password_hash, full_name)
	VALUES ($1, $2, $3) RETURNING id, public_id, email, full_name, role, created_at`

	var created User
	row := r.db.QueryRowContext(ctx, q, user.Email, user.PasswordHash, user.Fullname)

	if err := row.Scan(&created.ID, &created.PublicId, &created.Email, &created.Fullname, &created.Role, &created.CreatedAt); err != nil {
		return nil, database.DBToAppError(err)
	}

	return &created, nil
}

// func (r *AuthRepository) Get(ctx context.Context, userId uuid.UUID) (*User, error) {
// 	q := `SELECT * FROM users WHERE public_id = $1`

// 	var found User
// 	row := r.db.QueryRowContext(ctx, q, userId)

// 	if err := row.Scan(&found.ID, &found.PublicId, &found.Email, &found.Fullname, &found.Role, &found.CreatedAt); err != nil {
// 		return &User{}, nil
// 	}

// 	return &found, nil
// }
