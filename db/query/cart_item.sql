-- name: CreateCartItem :one
INSERT INTO cart_item (quantity,
                       cart_id,
                       product_id)
VALUES ($1, $2, $3)
RETURNING *;