package orders_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *OrdersRepository) UpdateOrderStatus(ctx context.Context, id int, status string, completedAt *time.Time) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE bikeshop.orders
		SET status = $1, completed_at = $2
		WHERE id = $3
		RETURNING id, full_name, email, address, status, total_price, created_at, completed_at;
	`

	var m OrderModel
	err := r.pool.QueryRow(ctx, query, status, completedAt, id).Scan(
		&m.ID,
		&m.FullName,
		&m.Email,
		&m.Address,
		&m.Status,
		&m.TotalPrice,
		&m.CreatedAt,
		&m.CompletedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Order{}, fmt.Errorf("order id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.Order{}, fmt.Errorf("scan error: %w", err)
	}

	m.Items = []domain.BikeOrder{}
	return m.ToDomain(), nil
}