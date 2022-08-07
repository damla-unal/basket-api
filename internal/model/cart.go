package model

import (
	"time"
)

type Cart struct {
	ID           int        `json:"id"`
	TotalPrice   int        `json:"total_price"`
	Vat          int        `json:"vat"`
	Discount     int        `json:"discount"`
	CustomerID   int        `json:"customer_id"`
	CustomerName string     `json:"customer_name"`
	Items        []CartItem `json:"cart_items"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
