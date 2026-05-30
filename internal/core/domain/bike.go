package domain

type Bike struct {
	ID          int
	Version     int
	Brand       string
	Model       string
	Type        string
	Price       float64
	Stock       int
	Description string
}

func NewBike(id int, version int, brand string, model string, bikeType string, price float64, stock int, description string) Bike {
	return Bike{
		ID:          id,
		Version:     version,
		Brand:       brand,
		Model:       model,
		Type:        bikeType,
		Price:       price,
		Stock:       stock,
		Description: description,
	}
}

func NewBikeUninitialized(brand string, model string, bikeType string, price float64, stock int, description string) Bike {
	return NewBike(
		UninitializedID,
		UninitializedVersion,
		brand,
		model,
		bikeType,
		price,
		stock,
		description,
	)
}
