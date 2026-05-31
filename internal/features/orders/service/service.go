package orders_service

import (
	"context"
	"time"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type OrdersService struct {
	ordersRepository OrdersRepository
}

func NewOrdersService(ordersRepository OrdersRepository) *OrdersService {
	return &OrdersService{
		ordersRepository: ordersRepository,
	}
}

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int) (domain.Order, error)
	GetOrders(ctx context.Context) ([]domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status string, completedAt *time.Time) (domain.Order, error)
}
