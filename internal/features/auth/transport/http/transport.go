package auth_transport_http

import (
	"context"
	"net/http"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
	core_http_server "github.com/rrwwmq/bike-shop/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService   AuthService
	adminSecret   string
}

type AuthService interface {
	Register(ctx context.Context, email string, password string) (domain.User, error)
	RegisterAdmin(ctx context.Context, email string, password string) (domain.User, error)
	Login(ctx context.Context, email string, password string) (string, error)
}

func NewAuthHTTPHandler(authService AuthService, adminSecret string) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
		adminSecret: adminSecret,
	}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: h.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/register/admin",
			Handler: h.RegisterAdmin,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.Login,
		},
	}
}