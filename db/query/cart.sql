-- name: CreateCart :one
INSERT INTO cart (total_price, vat, discount, status)
VALUES ($1, $2, $3, $4)
RETURNING *;