package orders_postgres_repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (r *OrdersRepository) GetOrders(ctx context.Context) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, full_name, email, address, status, total_price, created_at, completed_at
		FROM bikeshop.orders;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query orders: %w", err)
	}
	defer rows.Close()

	orders, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Order, error) {
		var m OrderModel
		err := row.Scan(
			&m.ID,
			&m.FullName,
			&m.Email,
			&m.Address,
			&m.Status,
			&m.TotalPrice,
			&m.CreatedAt,
			&m.CompletedAt,
		)
		return m.ToDomain(), err
	})
	if err != nil {
		return nil, fmt.Errorf("collect rows: %w", err)
	}

	return orders, nil
}