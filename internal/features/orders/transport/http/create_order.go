package orders_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type CreateOrderItemRequest struct {
	BikeID   int `json:"bike_id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}

type CreateOrderRequest struct {
	FullName string                   `json:"full_name" validate:"required"`
	Email    string                   `json:"email" validate:"required,email"`
	Address  string                   `json:"address" validate:"required"`
	Items    []CreateOrderItemRequest  `json:"items" validate:"required,min=1"`
}

type CreateOrderResponse orderDTOResponse

func (h *OrdersHTTPHandler) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateOrder handler")

	var request CreateOrderRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	order, err := h.ordersService.CreateOrder(ctx, orderDomainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create order")
		return
	}

	response := CreateOrderResponse(orderDTOFromDomain(order))
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func orderDomainFromDTO(dto CreateOrderRequest) domain.Order {
	items := make([]domain.BikeOrder, len(dto.Items))
	for i, item := range dto.Items {
		items[i] = domain.NewBikeOrderUninitialized(item.BikeID, item.Quantity)
	}

	return domain.NewOrderUninitialized(dto.FullName, dto.Email, dto.Address, items)
}