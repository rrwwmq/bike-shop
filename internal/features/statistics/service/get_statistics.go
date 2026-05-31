package statistics_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *StatisticsService) GetStatistics(ctx context.Context) (domain.Statistics, error) {
	stats, err := s.statisticsRepository.GetStatistics(ctx)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get statistics: %w", err)
	}

	return stats, nil
}
