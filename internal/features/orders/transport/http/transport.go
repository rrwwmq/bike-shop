package orders_transport_http

import (
	"context"
	"net/http"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_http_middleware "github.com/rrwwmq/bike-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/rrwwmq/bike-shop/internal/core/transport/http/server"
)

type OrdersHTTPHandler struct {
	ordersService OrdersService
	jwtSecret     string
}

type OrdersService interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int) (domain.Order, error)
	GetOrders(ctx context.Context) ([]domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status string) (domain.Order, error)
}

func NewOrdersHTTPHandler(ordersService OrdersService, jwtSecret string) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		ordersService: ordersService,
		jwtSecret:     jwtSecret,
	}
}

func (h *OrdersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/orders",
			Handler: h.CreateOrder,
		},
		{
			Method:  http.MethodGet,
			Path:    "/orders/{id}",
			Handler: h.GetOrder,
		},
		{
			Method:     http.MethodGet,
			Path:       "/orders",
			Handler:    h.GetOrders,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWT(h.jwtSecret)},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/orders/{id}/status",
			Handler:    h.UpdateOrderStatus,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWT(h.jwtSecret)},
		},
	}
}
