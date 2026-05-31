package domain

import "time"

type User struct {
	ID           int
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

func NewUser(id int, email string, passwordHash string, role string, createdAt time.Time) User {
	return User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    createdAt,
	}
}

func NewUserUninitialized(email string, passwordHash string) User {
	return NewUser(
		UninitializedID,
		email,
		passwordHash,
		"customer",
		time.Now(),
	)
}
