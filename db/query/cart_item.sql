-- name: GetCartItemsByCartID :many
SELECT p.title as product_title, cart_id, cart_item.id, quantity, cart_item.discount, cart_item.price, product_id
FROM cart_item
         LEFT JOIN cart c on c.id = cart_item.cart_id
         LEFT JOIN product p on p.id = cart_item.product_id
WHERE cart_id = $1;

-- name: UpsertCartItem :exec
INSERT INTO cart_item (quantity, cart_id, product_id, price)
VALUES (sqlc.arg(quantity), sqlc.arg(cart_id), sqlc.arg(product_id), sqlc.arg(price))
ON CONFLICT (cart_id, product_id)
    DO UPDATE
    SET quantity = cart_item.quantity + excluded.quantity,
        price    = cart_item.price + excluded.price;

-- name: UpdateCartItem :exec
UPDATE cart_item
SET quantity = sqlc.arg(quantity),
    discount = sqlc.arg(discount),
    price    = sqlc.arg(price)
WHERE id = sqlc.arg(id);


-- name: DeleteCartItem :exec
DELETE
FROM cart_item
WHERE id = $1;

-- name: GetCartItemByID :one
SELECT *
FROM cart_item
WHERE id = $1;