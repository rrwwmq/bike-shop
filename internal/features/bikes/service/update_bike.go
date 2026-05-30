package bikes_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *BikesService) UpdateBike(ctx context.Context, id int, bike domain.Bike) (domain.Bike, error) {
	bike, err := s.bikesRepository.UpdateBike(ctx, id, bike)
	if err != nil {
		return domain.Bike{}, fmt.Errorf("update bike: %w", err)
	}

	return bike, nil
}
