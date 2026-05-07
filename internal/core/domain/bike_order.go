package domain

type BikeOrder struct {
	ID              int
	OrderID         int
	BikeID          int
	Quantity        int
	PriceAtPurchase float64
}
