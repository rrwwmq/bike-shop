package bikes_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *BikesService) CreateBike(ctx context.Context, bike domain.Bike) (domain.Bike, error) {
	bike, err := s.bikesRepository.CreateBike(ctx, bike)
	if err != nil {
		return domain.Bike{}, fmt.Errorf("create bike: %w", err)
	}

	return bike, nil
}
