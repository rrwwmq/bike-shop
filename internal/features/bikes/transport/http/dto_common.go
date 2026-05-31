package bikes_transport_http

import "github.com/rrwwmq/bike-shop/internal/core/domain"

type bikeDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func bikeDTOFromDomain(bike domain.Bike) bikeDTOResponse {
	return bikeDTOResponse{
		ID:          bike.ID,
		Version:     bike.Version,
		Brand:       bike.Brand,
		Model:       bike.Model,
		Type:        bike.Type,
		Price:       bike.Price,
		Stock:       bike.Stock,
		Description: bike.Description,
	}
}

func bikesDTOFromDomains(bikes []domain.Bike) []bikeDTOResponse {
	bikeDTO := make([]bikeDTOResponse, len(bikes))

	for i, bike := range bikes {
		bikeDTO[i] = bikeDTOFromDomain(bike)
	}

	return bikeDTO
}
