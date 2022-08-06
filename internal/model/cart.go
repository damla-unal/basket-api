package model

import (
	"time"
)

type Cart struct {
	ID           int64      `json:"id"`
	TotalPrice   int64      `json:"total_price"`
	Vat          int64      `json:"vat"`
	Discount     int64      `json:"discount"`
	CustomerID   int64      `json:"customer_id"`
	CustomerName string     `json:"customer_name"`
	Status       CartStatus `json:"status"`
	Items        []CartItem `json:"cart_items"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
