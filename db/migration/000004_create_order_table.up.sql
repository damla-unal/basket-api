CREATE TABLE "order"
(
    id          BIGSERIAL PRIMARY KEY,
    total_price BIGINT NOT NULL,
    customer_id BIGINT NOT NULL REFERENCES customer (id),
    created_at  TIMESTAMP DEFAULT now()
);