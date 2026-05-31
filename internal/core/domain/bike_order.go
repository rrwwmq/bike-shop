package domain

type BikeOrder struct {
	ID              int
	OrderID         int
	BikeID          int
	Quantity        int
	PriceAtPurchase float64
}

func NewBikeOrder(id int, orderID int, bikeID int, quantity int, priceAtPurchase float64) BikeOrder {
	return BikeOrder{
		ID:              id,
		OrderID:         orderID,
		BikeID:          bikeID,
		Quantity:        quantity,
		PriceAtPurchase: priceAtPurchase,
	}
}

func NewBikeOrderUninitialized(bikeID int, quantity int) BikeOrder {
	return NewBikeOrder(
		UninitializedID,
		UninitializedID,
		bikeID,
		quantity,
		0,
	)
}
