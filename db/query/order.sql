-- name: CreateOrder :exec
INSERT INTO "order" (total_price, customer_id)
VALUES ($1, $2);

-- name: GetOrdersByCustomerID :many
SELECT *
FROM "order"
WHERE customer_id = $1;