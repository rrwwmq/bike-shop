package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
	"github.com/rrwwmq/bike-shop/internal/core/jwtutil"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials: %w", core_errors.ErrInvalidArgument)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials: %w", core_errors.ErrInvalidArgument)
	}

	claims := jwtutil.TokenClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signed, nil
}
