package orders_service

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

func (s *OrdersService) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	if len(order.Items) == 0 {
		return domain.Order{}, fmt.Errorf("order must have at least one item: %w", core_errors.ErrInvalidArgument)
	}

	order, err := s.ordersRepository.CreateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, fmt.Errorf("create order: %w", err)
	}

	return order, nil
}