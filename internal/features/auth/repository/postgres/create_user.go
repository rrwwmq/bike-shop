package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO bikeshop.users (email, password_hash, role, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, password_hash, role, created_at;
	`

	var m UserModel
	err := r.pool.QueryRow(ctx, query, user.Email, user.PasswordHash, user.Role, user.CreatedAt).Scan(
		&m.ID,
		&m.Email,
		&m.PasswordHash,
		&m.Role,
		&m.CreatedAt,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return m.ToDomain(), nil
}