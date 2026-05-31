package auth_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, email, password_hash, role, created_at
		FROM bikeshop.users
		WHERE email = $1;
	`

	var m UserModel
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&m.ID,
		&m.Email,
		&m.PasswordHash,
		&m.Role,
		&m.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user email=%s: %w", email, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return m.ToDomain(), nil
}