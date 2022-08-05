BEGIN;


CREATE TYPE cart_status AS ENUM ('saved', 'completed');

CREATE TABLE cart
(
    id          BIGSERIAL PRIMARY KEY,
    total_price BIGINT      NOT NULL,
    vat         BIGINT      NOT NULL,
    discount    BIGINT      NOT NULL,
    status      cart_status NOT NULL,
    customer_id BIGINT      NOT NULL REFERENCES customer (id),
    created_at  TIMESTAMP DEFAULT now(),
    updated_at  TIMESTAMP DEFAULT now()
);

COMMIT;

