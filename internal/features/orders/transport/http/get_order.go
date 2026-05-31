package orders_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type GetOrderResponse orderDTOResponse

func (h *OrdersHTTPHandler) GetOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke GetOrder handler")

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path")
		return
	}

	order, err := h.ordersService.GetOrder(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get order")
		return
	}

	response := GetOrderResponse(orderDTOFromDomain(order))
	responseHandler.JSONResponse(response, http.StatusOK)
}