-- name: ListProducts :many
SELECT *
FROM product;


-- name: GetProductByID :one
SELECT *
FROM product
WHERE id = $1;