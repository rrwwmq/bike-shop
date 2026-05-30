package bikes_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

type CreateBikeRequest struct {
	Brand       string  `json:"brand" validate:"required"`
	Model       string  `json:"model" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Description string  `json:"description"`
}

type CreateBikeResponse bikeDTOResponse

func (h *BikesHTTPHandler) CreateBike(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateBike handler")

	var request CreateBikeRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	bikeDomain, err := h.bikesService.CreateBike(ctx, domainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create bike")
		return
	}

	response := CreateBikeResponse(bikeDTOFromDomain(bikeDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateBikeRequest) domain.Bike {
	return domain.NewBikeUninitialized(dto.Brand, dto.Model, dto.Type, dto.Price, dto.Stock, dto.Description)
}