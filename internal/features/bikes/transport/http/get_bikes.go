package bikes_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type GetBikesResponse []bikeDTOResponse

func (h *BikesHTTPHandler) GetBikes(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke GetBikes handler")

	bikes, err := h.bikesService.GetBikes(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bikes")
		return
	}

	response := GetBikesResponse(bikesDTOFromDomains(bikes))

	responseHandler.JSONResponse(response, http.StatusOK)
}
