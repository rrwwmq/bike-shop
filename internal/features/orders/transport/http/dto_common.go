package orders_transport_http

import (
	"time"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type orderDTOResponse struct {
	ID          int                    `json:"id"`
	FullName    string                 `json:"full_name"`
	Email       string                 `json:"email"`
	Address     string                 `json:"address"`
	Status      string                 `json:"status"`
	TotalPrice  float64                `json:"total_price"`
	CreatedAt   time.Time              `json:"created_at"`
	CompletedAt *time.Time             `json:"completed_at"`
	Items       []bikeOrderDTOResponse `json:"items"`
}

type bikeOrderDTOResponse struct {
	ID              int     `json:"id"`
	BikeID          int     `json:"bike_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

func orderDTOFromDomain(order domain.Order) orderDTOResponse {
	items := make([]bikeOrderDTOResponse, len(order.Items))
	for i, item := range order.Items {
		items[i] = bikeOrderDTOResponse{
			ID:              item.ID,
			BikeID:          item.BikeID,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.PriceAtPurchase,
		}
	}

	return orderDTOResponse{
		ID:          order.ID,
		FullName:    order.FullName,
		Email:       order.Email,
		Address:     order.Address,
		Status:      order.Status,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   order.CreatedAt,
		CompletedAt: order.CompletedAt,
		Items:       items,
	}
}
