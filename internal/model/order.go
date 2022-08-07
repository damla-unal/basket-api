package model

import "time"

type Order struct {
	ID         int       `json:"id"`
	TotalPrice int       `json:"total_price"`
	CustomerID int       `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
