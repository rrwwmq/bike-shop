package auth_postgres_repository

import (
	"time"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type UserModel struct {
	ID           int
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

func (m *UserModel) ToDomain() domain.User {
	return domain.NewUser(
		m.ID,
		m.Email,
		m.PasswordHash,
		m.Role,
		m.CreatedAt,
	)
}