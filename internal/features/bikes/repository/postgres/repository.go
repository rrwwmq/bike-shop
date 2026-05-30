package bikes_postgres_repository

import core_repository_postgres_pool "github.com/rrwwmq/bike-shop/internal/core/repository/postgres/pool"

type BikesRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewBikesRepository(pool core_repository_postgres_pool.Pool) *BikesRepository {
	return &BikesRepository{
		pool: pool,
	}
}
