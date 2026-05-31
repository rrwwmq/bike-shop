package auth_postgres_repository

import core_repository_postgres_pool "github.com/rrwwmq/bike-shop/internal/core/repository/postgres/pool"

type AuthRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewAuthRepository(pool core_repository_postgres_pool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}