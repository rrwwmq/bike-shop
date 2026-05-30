package bikes_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *BikesService) GetBike(ctx context.Context, id int) (domain.Bike, error) {
	bike, err := s.bikesRepository.GetBike(ctx, id)
	if err != nil {
		return domain.Bike{}, fmt.Errorf("get bike: %w", err)
	}

	return bike, nil
}
