// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: cart.sql

package db

import (
	"context"
	"database/sql"
)

const getCartByCustomerID = `-- name: GetCartByCustomerID :one
SELECT cart.id,
       total_price,
       vat,
       discount,
       customer_id,
       c.name as customer_name,
       cart.created_at,
       updated_at
FROM cart
         LEFT JOIN customer c on c.id = cart.customer_id
WHERE c.id = $1
`

type GetCartByCustomerIDRow struct {
	ID           int64          `json:"id"`
	TotalPrice   int64          `json:"total_price"`
	Vat          int64          `json:"vat"`
	Discount     int64          `json:"discount"`
	CustomerID   int64          `json:"customer_id"`
	CustomerName sql.NullString `json:"customer_name"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

func (q *Queries) GetCartByCustomerID(ctx context.Context, id int64) (GetCartByCustomerIDRow, error) {
	row := q.db.QueryRow(ctx, getCartByCustomerID, id)
	var i GetCartByCustomerIDRow
	err := row.Scan(
		&i.ID,
		&i.TotalPrice,
		&i.Vat,
		&i.Discount,
		&i.CustomerID,
		&i.CustomerName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCart = `-- name: UpdateCart :exec
UPDATE cart
SET total_price = $1,
    vat         = $2,
    discount    = $3,
    updated_at  = now()
WHERE id = $4
`

type UpdateCartParams struct {
	Price    int64 `json:"price"`
	Vat      int64 `json:"vat"`
	Discount int64 `json:"discount"`
	ID       int64 `json:"id"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) error {
	_, err := q.db.Exec(ctx, updateCart,
		arg.Price,
		arg.Vat,
		arg.Discount,
		arg.ID,
	)
	return err
}
