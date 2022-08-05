-- name: CreateCustomer :one
INSERT INTO customer (name, email)
VALUES ($1, $2)
RETURNING *;