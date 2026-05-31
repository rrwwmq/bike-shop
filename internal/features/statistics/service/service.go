package statistics_service

import (
	"context"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: statisticsRepository,
	}
}

type StatisticsRepository interface {
	GetStatistics(ctx context.Context) (domain.Statistics, error)
}