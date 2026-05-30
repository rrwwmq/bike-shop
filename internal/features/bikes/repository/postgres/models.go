package bikes_postgres_repository

import "github.com/rrwwmq/bike-shop/internal/core/domain"

type BikeModel struct {
	ID          int
	Version     int
	Brand       string
	Model       string
	Type        string
	Price       float64
	Stock       int
	Description string
}

func (m *BikeModel) ToDomain() domain.Bike {
	return domain.NewBike(
		m.ID,
		m.Version,
		m.Brand,
		m.Model,
		m.Type,
		m.Price,
		m.Stock,
		m.Description,
	)
}
