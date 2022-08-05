-- name: CreateCart :one
INSERT INTO cart (total_price, vat, discount, status, customer_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;