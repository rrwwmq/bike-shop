package bikes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *BikesRepository) GetBike(ctx context.Context, id int) (domain.Bike, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, brand, model, type, price, stock, description
		FROM bikeshop.bikes
		WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)

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
