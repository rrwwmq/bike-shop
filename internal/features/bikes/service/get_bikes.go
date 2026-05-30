package bikes_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *BikesService) GetBikes(ctx context.Context) ([]domain.Bike, error) {
	bikes, err := s.bikesRepository.GetBikes(ctx)
	if err != nil {
		return nil, fmt.Errorf("get bikes: %w", err)
	}

	return bikes, nil
}
