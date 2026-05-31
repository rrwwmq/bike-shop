package statistics_postgres_repository

import core_repository_postgres_pool "github.com/rrwwmq/bike-shop/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewStatisticsRepository(pool core_repository_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{pool: pool}
}