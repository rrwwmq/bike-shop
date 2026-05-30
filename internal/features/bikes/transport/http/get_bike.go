package bikes_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type GetBikeResponse bikeDTOResponse

func (h *BikesHTTPHandler) GetBike(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke GetBike handler")

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path")
		return
	}

	bike, err := h.bikesService.GetBike(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bike")
		return
	}

	response := GetBikeResponse(bikeDTOFromDomain(bike))
	responseHandler.JSONResponse(response, http.StatusOK)
}
