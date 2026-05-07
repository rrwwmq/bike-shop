package domain

import "time"

type Order struct {
	ID         int
	Version    int
	FullName   string
	Email      string
	Status     string
	TotalPrice float64
	CreatedAt  time.Time
	Items      []BikeOrder
}
