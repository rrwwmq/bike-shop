package bikes_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

func (h *BikesHTTPHandler) DeleteBike(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke DeleteBike handler")

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path")
		return
	}

	if err := h.bikesService.DeleteBike(ctx, id); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete bike")
		return
	}

	responseHandler.NoContentResponse()
}
