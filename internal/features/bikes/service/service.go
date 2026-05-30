package bikes_service

import (
	"context"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type BikesService struct {
	bikesRepository BikesRepository
}

func NewBikesService(bikesRepository BikesRepository) *BikesService {
	return &BikesService{
		bikesRepository: bikesRepository,
	}
}

type BikesRepository interface {
	CreateBike(ctx context.Context, bike domain.Bike) (domain.Bike, error)
	GetBikes(ctx context.Context) ([]domain.Bike, error)
	GetBike(ctx context.Context, id int) (domain.Bike, error)
	UpdateBike(ctx context.Context, id int, bike domain.Bike) (domain.Bike, error)
	DeleteBike(ctx context.Context, id int) error
}
