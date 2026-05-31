package orders_postgres_repository

import (
	"time"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type OrderModel struct {
	ID          int
	FullName    string
	Email       string
	Address     string
	Status      string
	TotalPrice  float64
	CreatedAt   time.Time
	CompletedAt *time.Time
	Items       []domain.BikeOrder
}

func (m *OrderModel) ToDomain() domain.Order {
	return domain.NewOrder(
		m.ID,
		m.FullName,
		m.Email,
		m.Address,
		m.Status,
		m.TotalPrice,
		m.CreatedAt,
		m.CompletedAt,
		m.Items,
	)
}

type BikeOrderModel struct {
	ID              int
	OrderID         int
	BikeID          int
	Quantity        int
	PriceAtPurchase float64
}

func (m *BikeOrderModel) ToDomain() domain.BikeOrder {
	return domain.NewBikeOrder(
		m.ID,
		m.OrderID,
		m.BikeID,
		m.Quantity,
		m.PriceAtPurchase,
	)
}