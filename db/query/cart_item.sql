-- name: GetCartItemsByCartID :many
SELECT p.title as product_title,
       cart_id,
       cart_item.id,
       quantity,
       cart_item.discount,
       cart_item.price,
       product_id,
       p.price as qty_price,
       p.vat   as product_vat
FROM cart_item
         LEFT JOIN cart c on c.id = cart_item.cart_id
         LEFT JOIN product p on p.id = cart_item.product_id
WHERE cart_id = $1;

-- name: UpsertCartItem :exec
INSERT INTO cart_item (quantity, discount, cart_id, product_id, price)
VALUES (sqlc.arg(quantity), sqlc.arg(discount), sqlc.arg(cart_id), sqlc.arg(product_id), sqlc.arg(price))
ON CONFLICT (cart_id, product_id)
    DO UPDATE
    SET quantity = cart_item.quantity + excluded.quantity,
        price    = cart_item.price + excluded.price,
        discount = excluded.discount;

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

-- name: GetCartItemDetailsByID :one
SELECT cart_item.id,
       quantity,
       cart_id,
       product_id,
       discount,
       cart_item.price,
       title   as prodcut_title,
       p.price as qty_price,
       vat     as product_vat
FROM cart_item
         LEFT JOIN product p on p.id = cart_item.product_id
WHERE cart_item.id = $1;

-- name: DeleteAllCartItem :exec
DELETE
FROM cart_item
WHERE cart_id = $1;