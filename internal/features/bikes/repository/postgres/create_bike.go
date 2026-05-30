package bikes_postgres_repository

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (r *BikesRepository) CreateBike(ctx context.Context, bike domain.Bike) (domain.Bike, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO bikeshop.bikes (brand, model, type, price, stock, description)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (brand, model) DO UPDATE SET stock = bikeshop.bikes.stock + EXCLUDED.stock
		RETURNING id, version, brand, model, type, price, stock, description;
	`

	row := r.pool.QueryRow(ctx, query, bike.Brand, bike.Model, bike.Type, bike.Price, bike.Stock, bike.Description)

	var m BikeModel
	err := row.Scan(
		&m.ID,
		&m.Version,
		&m.Brand,
		&m.Model,
		&m.Type,
		&m.Price,
		&m.Stock,
		&m.Description,
	)
	if err != nil {
		return domain.Bike{}, fmt.Errorf("scan error: %w", err)
	}

	return m.ToDomain(), nil
}
