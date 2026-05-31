package orders_postgres_repository

import core_repository_postgres_pool "github.com/rrwwmq/bike-shop/internal/core/repository/postgres/pool"

type OrdersRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewOrdersRepository(pool core_repository_postgres_pool.Pool) *OrdersRepository {
	return &OrdersRepository{
		pool: pool,
	}
}