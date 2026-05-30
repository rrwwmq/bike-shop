package bikes_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *BikesRepository) DeleteBike(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM bikeshop.bikes
		WHERE id = $1
		RETURNING id;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var deletedID int
	err := row.Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("bike id=%d: %w", id, core_errors.ErrNotFound)
		}
		return fmt.Errorf("scan error: %w", err)
	}

	return nil
}
