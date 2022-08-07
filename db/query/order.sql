-- name: CreateOrder :exec
INSERT INTO "order" (total_price, customer_id)
VALUES ($1, $2);