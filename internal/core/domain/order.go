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
