package domain

import "time"

type Order struct {
	ID          int
	FullName    string
	Email       string
	Address     string
	Status      string
	TotalPrice  float64
	CreatedAt   time.Time
	CompletedAt *time.Time
	Items       []BikeOrder
}

func NewOrder(id int, fullName string, email string, address string, status string, totalPrice float64, createdAt time.Time, completedAt *time.Time, items []BikeOrder) Order {
	return Order{
		ID:          id,
		FullName:    fullName,
		Email:       email,
		Address:     address,
		Status:      status,
		TotalPrice:  totalPrice,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		Items:       items,
	}
}

func NewOrderUninitialized(fullName string, email string, address string, items []BikeOrder) Order {
	return NewOrder(
		UninitializedID,
		fullName,
		email,
		address,
		"pending",
		0,
		time.Now(),
		nil,
		items,
	)
}
