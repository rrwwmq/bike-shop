package auth_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, email string, password string) (domain.User, error) {
	_, err := s.authRepository.GetUserByEmail(ctx, email)
	if err == nil {
		return domain.User{}, fmt.Errorf("email already taken: %w", core_errors.ErrConflict)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user := domain.NewUserUninitialized(email, string(hash))

	user, err = s.authRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
