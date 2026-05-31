package orders_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

func (h *OrdersHTTPHandler) UpdateOrderStatus(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke UpdateOrderStatus handler")

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path")
		return
	}

	var request UpdateOrderStatusRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	order, err := h.ordersService.UpdateOrderStatus(ctx, id, request.Status)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update order status")
		return
	}

	responseHandler.JSONResponse(orderDTOFromDomain(order), http.StatusOK)
}