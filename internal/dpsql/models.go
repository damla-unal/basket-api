// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"fmt"
)

type CartStatus string

const (
	CartStatusSaved     CartStatus = "saved"
	CartStatusCompleted CartStatus = "completed"
)

func (e *CartStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CartStatus(s)
	case string:
		*e = CartStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for CartStatus: %T", src)
	}
	return nil
}

type Cart struct {
	ID         int64        `json:"id"`
	TotalPrice int64        `json:"total_price"`
	Vat        int64        `json:"vat"`
	Discount   int64        `json:"discount"`
	Status     CartStatus   `json:"status"`
	CustomerID int64        `json:"customer_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type CartItem struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
	Discount  int64 `json:"discount"`
}

type Customer struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Order struct {
	ID         int64        `json:"id"`
	TotalPrice int64        `json:"total_price"`
	CustomerID int64        `json:"customer_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
}

type Product struct {
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Price     int64        `json:"price"`
	Vat       int64        `json:"vat"`
	CreatedAt sql.NullTime `json:"created_at"`
}
