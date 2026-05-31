package orders_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *OrdersService) GetOrders(ctx context.Context) ([]domain.Order, error) {
	orders, err := s.ordersRepository.GetOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}

	return orders, nil
}