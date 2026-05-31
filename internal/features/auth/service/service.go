package auth_service

import (
	"context"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type AuthService struct {
	authRepository AuthRepository
	jwtSecret      string
}

func NewAuthService(authRepository AuthRepository, jwtSecret string) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		jwtSecret:      jwtSecret,
	}
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}