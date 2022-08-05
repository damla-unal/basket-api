// Code generated by sqlc. DO NOT EDIT.
// source: product.sql

package db

import (
	"context"
)

const listProducts = `-- name: ListProducts :many
SELECT id, title, price, vat, created_at
FROM product
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Price,
			&i.Vat,
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