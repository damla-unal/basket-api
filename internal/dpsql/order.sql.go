// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO "order" (total_price, customer_id)
VALUES ($1, $2)
`

type CreateOrderParams struct {
	TotalPrice int64 `json:"total_price"`
	CustomerID int64 `json:"customer_id"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.Exec(ctx, createOrder, arg.TotalPrice, arg.CustomerID)
	return err
}

const getOrdersByCustomerID = `-- name: GetOrdersByCustomerID :many
SELECT id, total_price, customer_id, created_at
FROM "order"
WHERE customer_id = $1
`

func (q *Queries) GetOrdersByCustomerID(ctx context.Context, customerID int64) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersByCustomerID, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.TotalPrice,
			&i.CustomerID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
