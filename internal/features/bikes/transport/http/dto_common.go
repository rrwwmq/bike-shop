package bikes_transport_http

import "github.com/rrwwmq/bike-shop/internal/core/domain"

type bikeDTOResponse struct {
	ID          int
	Version     int
	Brand       string
	Model       string
	Type        string
	Stock       int
	Description string
}

func bikeDTOFromDomain(bike domain.Bike) bikeDTOResponse {
	return bikeDTOResponse{
		ID:          bike.ID,
		Version:     bike.Version,
		Brand:       bike.Brand,
		Model:       bike.Model,
		Type:        bike.Type,
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
