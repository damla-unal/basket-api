CREATE TABLE cart_item
(
    id         BIGSERIAL PRIMARY KEY,
    quantity   BIGINT NOT NULL,
    cart_id    BIGINT NOT NULL REFERENCES cart (id),
    product_id BIGINT NOT NULL REFERENCES product (id)
);