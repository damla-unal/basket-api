BEGIN;

CREATE TABLE customer
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    email      TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE product
(
    id         BIGSERIAL PRIMARY KEY,
    title      TEXT   NOT NULL,
    price      BIGINT NOT NULL,
    vat        BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE "order"
(
    id          BIGSERIAL PRIMARY KEY,
    total_price BIGINT NOT NULL,
    created_at  TIMESTAMP DEFAULT now()
);

CREATE TABLE cart
(
    id          BIGSERIAL PRIMARY KEY,
    total_price BIGINT NOT NULL,
    vat         BIGINT NOT NULL,
    discount    BIGINT NOT NULL,
    created_at  TIMESTAMP DEFAULT now(),
    updated_at  TIMESTAMP DEFAULT now()
);

COMMIT;