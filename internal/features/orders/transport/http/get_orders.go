package orders_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type GetOrdersResponse []orderDTOResponse

func (h *OrdersHTTPHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke GetOrders handler")

	orders, err := h.ordersService.GetOrders(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get orders")
		return
	}

	response := make(GetOrdersResponse, len(orders))
	for i, order := range orders {
		response[i] = orderDTOFromDomain(order)
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}