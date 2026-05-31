package bikes_transport_http

import (
	"context"
	"net/http"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_http_middleware "github.com/rrwwmq/bike-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/rrwwmq/bike-shop/internal/core/transport/http/server"
)

type BikesHTTPHandler struct {
	bikesService BikesService
	jwtSecret    string
}

type BikesService interface {
	CreateBike(ctx context.Context, bike domain.Bike) (domain.Bike, error)
	GetBikes(ctx context.Context) ([]domain.Bike, error)
	GetBike(ctx context.Context, id int) (domain.Bike, error)
	UpdateBike(ctx context.Context, id int, bike domain.Bike) (domain.Bike, error)
	DeleteBike(ctx context.Context, id int) error
}

func NewBikesHTTPHandler(bikesService BikesService, jwtSecret string) *BikesHTTPHandler {
	return &BikesHTTPHandler{
		bikesService: bikesService,
		jwtSecret:    jwtSecret,
	}
}

func (h *BikesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/bikes",
			Handler:    h.CreateBike,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWT(h.jwtSecret)},
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
			Method:     http.MethodPut,
			Path:       "/bikes/{id}",
			Handler:    h.UpdateBike,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWT(h.jwtSecret)},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/bikes/{id}",
			Handler:    h.DeleteBike,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWT(h.jwtSecret)},
		},
	}
}
