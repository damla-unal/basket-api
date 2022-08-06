-- name: CreateCart :one
INSERT INTO cart (total_price, vat, discount, status, customer_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetSavedCartByCustomerID :one
SELECT cart.id,
       total_price,
       vat,
       discount,
       status,
       customer_id,
       c.name as customer_name,
       cart.created_at,
       updated_at
FROM cart
         LEFT JOIN customer c on c.id = cart.customer_id
WHERE c.id = $1
  and cart.status = 'saved';

