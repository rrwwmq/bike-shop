package auth_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_request "github.com/rrwwmq/bike-shop/internal/core/transport/http/request"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke Login handler")

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}

	token, err := h.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login")
		return
	}

	responseHandler.JSONResponse(LoginResponse{Token: token}, http.StatusOK)
}