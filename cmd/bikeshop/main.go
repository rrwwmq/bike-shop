package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_repository_postgres_pool "github.com/rrwwmq/bike-shop/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/rrwwmq/bike-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/rrwwmq/bike-shop/internal/core/transport/http/server"
	bikes_postgres_repository "github.com/rrwwmq/bike-shop/internal/features/bikes/repository/postgres"
	bikes_service "github.com/rrwwmq/bike-shop/internal/features/bikes/service"
	bikes_transport_http "github.com/rrwwmq/bike-shop/internal/features/bikes/transport/http"
	orders_postgres_repository "github.com/rrwwmq/bike-shop/internal/features/orders/repository/postgres"
	orders_service "github.com/rrwwmq/bike-shop/internal/features/orders/service"
	orders_transport_http "github.com/rrwwmq/bike-shop/internal/features/orders/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := core_repository_postgres_pool.NewConnectionPool(ctx, core_repository_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "bikes"))
	bikesRepository := bikes_postgres_repository.NewBikesRepository(pool)
	bikesService := bikes_service.NewBikesService(bikesRepository)
	bikesTransportHTTP := bikes_transport_http.NewBikesHTTPHandler(bikesService)

	logger.Debug("initializing feature", zap.String("feature", "orders"))
	ordersRepository := orders_postgres_repository.NewOrdersRepository(pool)
	ordersService := orders_service.NewOrdersService(ordersRepository)
	ordersTransportHTTP := orders_transport_http.NewOrdersHTTPHandler(ordersService, core_http_server.NewConfigMust().AdminKey)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	routes := append(bikesTransportHTTP.Routes(), ordersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRouters(routes...)

	httpServer.RegisterAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
