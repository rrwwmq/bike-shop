package statistics_postgres_repository

import (
	"context"
	"fmt"

	"github.com/rrwwmq/bike-shop/internal/core/domain"
)

func (r *StatisticsRepository) GetStatistics(ctx context.Context) (domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			COUNT(o.id) AS total_orders,
			COALESCE(SUM(o.total_price), 0) AS total_revenue,
			(SELECT COUNT(*) FROM bikeshop.bikes) AS total_bikes,
			COALESCE(
				(SELECT b.model || ' ' || b.brand
				FROM bikeshop.bike_order bo
				JOIN bikeshop.bikes b ON b.id = bo.bike_id
				GROUP BY b.id, b.model, b.brand
				ORDER BY SUM(bo.quantity) DESC
				LIMIT 1),
			'Нет данных') AS most_popular_bike
		FROM bikeshop.orders o;
	`

	var stats domain.Statistics
	err := r.pool.QueryRow(ctx, query).Scan(
		&stats.TotalOrders,
		&stats.TotalRevenue,
		&stats.TotalBikes,
		&stats.MostPopularBike,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("scan error: %w", err)
	}

	return stats, nil
}