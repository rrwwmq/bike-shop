package bikes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *BikesRepository) UpdateBike(ctx context.Context, id int, bike domain.Bike) (domain.Bike, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE bikeshop.bikes
		SET brand = $1, model = $2, type = $3, price = $4, stock = $5, description = $6, version = version + 1
		WHERE id = $7
		RETURNING id, version, brand, model, type, price, stock, description;
	`

	row := r.pool.QueryRow(ctx, query, bike.Brand, bike.Model, bike.Type, bike.Price, bike.Stock, bike.Description, id)

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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Bike{}, fmt.Errorf("bike id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.Bike{}, fmt.Errorf("scan error: %w", err)
	}

	return m.ToDomain(), nil
}
