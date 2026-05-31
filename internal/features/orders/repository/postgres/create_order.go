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

func (r *OrdersRepository) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Order{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// 1. получаем цену каждого велосипеда и списываем stock
	var totalPrice float64
	for i, item := range order.Items {
		var price float64
		var stock int

		err := tx.QueryRow(ctx, `
			SELECT price, stock FROM bikeshop.bikes WHERE id = $1
		`, item.BikeID).Scan(&price, &stock)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return domain.Order{}, fmt.Errorf("bike id=%d: %w", item.BikeID, core_errors.ErrNotFound)
			}
			return domain.Order{}, fmt.Errorf("get bike: %w", err)
		}

		if stock < item.Quantity {
			return domain.Order{}, fmt.Errorf("bike id=%d insufficient stock: %w", item.BikeID, core_errors.ErrInvalidArgument)
		}

		_, err = tx.Exec(ctx, `
			UPDATE bikeshop.bikes SET stock = stock - $1, version = version + 1 WHERE id = $2
		`, item.Quantity, item.BikeID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("update bike stock: %w", err)
		}

		order.Items[i].PriceAtPurchase = price
		totalPrice += price * float64(item.Quantity)
	}

	// 2. вставляем order
	var orderID int
	err = tx.QueryRow(ctx, `
		INSERT INTO bikeshop.orders (full_name, email, address, status, total_price, created_at)
		VALUES ($1, $2, $3, 'pending', $4, $5)
		RETURNING id
	`, order.FullName, order.Email, order.Address, totalPrice, time.Now()).Scan(&orderID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("insert order: %w", err)
	}

	// 3. вставляем bike_order позиции
	for i, item := range order.Items {
		var itemID int
		err := tx.QueryRow(ctx, `
			INSERT INTO bikeshop.bike_order (order_id, bike_id, quantity, price_at_purchase)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, orderID, item.BikeID, item.Quantity, item.PriceAtPurchase).Scan(&itemID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("insert bike_order: %w", err)
		}
		order.Items[i].ID = itemID
		order.Items[i].OrderID = orderID
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Order{}, fmt.Errorf("commit transaction: %w", err)
	}

	order.ID = orderID
	order.TotalPrice = totalPrice
	order.Status = "pending"

	return order, nil
}
