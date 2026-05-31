package auth_transport_http

import (
	"fmt"
	"net/http"

	core_errors "github.com/rrwwmq/bike-shop/internal/core/errors"
	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type RegisterAdminRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	AdminSecret string `json:"admin_secret" validate:"required"`
}

func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke Register handler")

	var request RegisterRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	user, err := h.authService.Register(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to register user")
		return
	}

	responseHandler.JSONResponse(RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, http.StatusCreated)
}

func (h *AuthHTTPHandler) RegisterAdmin(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke RegisterAdmin handler")

	var request RegisterAdminRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	if request.AdminSecret != h.adminSecret {
		responseHandler.ErrorResponse(
			fmt.Errorf("invalid admin secret: %w", core_errors.ErrInvalidArgument),
			"forbidden",
		)
		return
	}

	user, err := h.authService.RegisterAdmin(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to register admin")
		return
	}

	responseHandler.JSONResponse(RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, http.StatusCreated)
}
