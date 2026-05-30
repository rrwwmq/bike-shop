package bikes_postgres_repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (r *BikesRepository) GetBikes(ctx context.Context) ([]domain.Bike, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, brand, model, type, price, stock, description
		FROM bikeshop.bikes;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	bikes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Bike, error) {
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
		return m.ToDomain(), err
	})
	if err != nil {
		return nil, fmt.Errorf("collect rows: %w", err)
	}

	return bikes, nil
}
