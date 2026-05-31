package orders_service

import (
	"context"
	"fmt"
	"time"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
)

var validStatuses = map[string]bool{
	"pending":   true,
	"completed": true,
	"cancelled": true,
}

func (s *OrdersService) UpdateOrderStatus(ctx context.Context, id int, status string) (domain.Order, error) {
	if !validStatuses[status] {
		return domain.Order{}, fmt.Errorf("invalid status '%s': %w", status, core_errors.ErrInvalidArgument)
	}

	var completedAt *time.Time
	if status == "completed" {
		t := time.Now()
		completedAt = &t
	}

	order, err := s.ordersRepository.UpdateOrderStatus(ctx, id, status, completedAt)
	if err != nil {
		return domain.Order{}, fmt.Errorf("update order status: %w", err)
	}

	return order, nil
}
