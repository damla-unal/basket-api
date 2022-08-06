-- name: CreateCartItem :one
INSERT INTO cart_item (quantity,
                       cart_id,
                       product_id)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetCartItemsByCartID :many
SELECT p.title as product_title, cart_id, quantity, cart_item.discount, product_id
FROM cart_item
         LEFT JOIN cart c on c.id = cart_item.cart_id
         LEFT JOIN product p on p.id = cart_item.product_id
WHERE cart_id = $1;

-- name: UpsertCartItem :exec
INSERT INTO cart_item (quantity, cart_id, product_id)
VALUES (sqlc.arg(quantity), sqlc.arg(cart_id), sqlc.arg(product_id))
ON CONFLICT (cart_id, product_id)
    DO UPDATE
    SET quantity = cart_item.quantity + excluded.quantity;