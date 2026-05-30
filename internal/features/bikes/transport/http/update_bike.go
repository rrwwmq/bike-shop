package bikes_transport_http

import (
	"net/http"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type UpdateBikeRequest struct {
	Brand       string  `json:"brand" validate:"required"`
	Model       string  `json:"model" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Description string  `json:"description"`
}

type UpdateBikeResponse bikeDTOResponse

func (h *BikesHTTPHandler) UpdateBike(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke UpdateBike handler")

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path")
		return
	}

	var request UpdateBikeRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	bike, err := h.bikesService.UpdateBike(ctx, id, updateDomainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update bike")
		return
	}

	response := UpdateBikeResponse(bikeDTOFromDomain(bike))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func updateDomainFromDTO(dto UpdateBikeRequest) domain.Bike {
	return domain.NewBikeUninitialized(dto.Brand, dto.Model, dto.Type, dto.Price, dto.Stock, dto.Description)
}
