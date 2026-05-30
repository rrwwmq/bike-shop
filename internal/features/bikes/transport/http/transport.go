package bikes_transport_http

import (
	"context"
	"net/http"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_http_server "github.com/rrwwmq/bike-shop/internal/core/transport/http/server"
)

type BikesHTTPHandler struct {
	bikesService BikesService
}

type BikesService interface {
	CreateBike(ctx context.Context, bike domain.Bike) (domain.Bike, error)
	GetBikes(ctx context.Context) ([]domain.Bike, error)
	GetBike(ctx context.Context, id int) (domain.Bike, error)
	UpdateBike(ctx context.Context, id int, bike domain.Bike) (domain.Bike, error)
	DeleteBike(ctx context.Context, id int) error
}

func NewBikesHTTPHandler(bikesService BikesService) *BikesHTTPHandler {
	return &BikesHTTPHandler{bikesService: bikesService}
}

func (h *BikesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/bikes",
			Handler: h.CreateBike,
		},

		{
			Method:  http.MethodGet,
			Path:    "/bikes",
			Handler: h.GetBikes,
		},

		{
			Method:  http.MethodGet,
			Path:    "/bikes/{id}",
			Handler: h.GetBike,
		},

		{
			Method:  http.MethodPut,
			Path:    "/bikes/{id}",
			Handler: h.UpdateBike,
		},

		{
			Method:  http.MethodDelete,
			Path:    "/bikes/{id}",
			Handler: h.DeleteBike,
		},
	}
}
