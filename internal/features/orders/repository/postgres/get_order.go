package orders_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (r *OrdersRepository) GetOrder(ctx context.Context, id int) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	orderQuery := `
		SELECT id, full_name, email, address, status, total_price, created_at, completed_at
		FROM bikeshop.orders
		WHERE id = $1;
	`

	var m OrderModel
	err := r.pool.QueryRow(ctx, orderQuery, id).Scan(
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
		return domain.Order{}, fmt.Errorf("scan order: %w", err)
	}

	itemsQuery := `
		SELECT id, order_id, bike_id, quantity, price_at_purchase
		FROM bikeshop.bike_order
		WHERE order_id = $1;
	`

	rows, err := r.pool.Query(ctx, itemsQuery, id)
	if err != nil {
		return domain.Order{}, fmt.Errorf("query bike_order: %w", err)
	}
	defer rows.Close()

	items, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.BikeOrder, error) {
		var item BikeOrderModel
		err := row.Scan(
			&item.ID,
			&item.OrderID,
			&item.BikeID,
			&item.Quantity,
			&item.PriceAtPurchase,
		)
		return item.ToDomain(), err
	})
	if err != nil {
		return domain.Order{}, fmt.Errorf("collect bike_order rows: %w", err)
	}

	m.Items = items
	return m.ToDomain(), nil
}
