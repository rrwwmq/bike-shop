package bikes_service

import (
	"context"
	"fmt"
)

func (s *BikesService) DeleteBike(ctx context.Context, id int) error {
	if err := s.bikesRepository.DeleteBike(ctx, id); err != nil {
		return fmt.Errorf("delete bike: %w", err)
	}

	return nil
}
