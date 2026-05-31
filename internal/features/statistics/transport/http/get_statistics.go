package statistics_transport_http

import (
	"net/http"

	core_logger "github.com/rrwwmq/bike-shop/internal/core/logger"
	core_http_response "github.com/rrwwmq/bike-shop/internal/core/transport/http/response"
)

type StatisticsResponse struct {
	TotalOrders     int     `json:"total_orders"`
	TotalRevenue    float64 `json:"total_revenue"`
	TotalBikes      int     `json:"total_bikes"`
	MostPopularBike string  `json:"most_popular_bike"`
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke GetStatistics handler")

	stats, err := h.statisticsService.GetStatistics(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}

	responseHandler.JSONResponse(StatisticsResponse{
		TotalOrders:     stats.TotalOrders,
		TotalRevenue:    stats.TotalRevenue,
		TotalBikes:      stats.TotalBikes,
		MostPopularBike: stats.MostPopularBike,
	}, http.StatusOK)
}