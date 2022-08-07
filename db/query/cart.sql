-- name: GetCartByCustomerID :one
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
WHERE c.id = $1;

-- name: UpdateCart :exec
UPDATE cart
SET total_price = sqlc.arg(price),
    vat         = sqlc.arg(vat),
    discount    = sqlc.arg(discount),
    updated_at  = now()
WHERE id = sqlc.arg(id);