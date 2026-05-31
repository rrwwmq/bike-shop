package orders_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (s *OrdersService) GetOrder(ctx context.Context, id int) (domain.Order, error) {
	order, err := s.ordersRepository.GetOrder(ctx, id)
	if err != nil {
		return domain.Order{}, fmt.Errorf("get order: %w", err)
	}

	return order, nil
}