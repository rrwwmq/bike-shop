CREATE SCHEMA IF NOT EXISTS bikeshop;

CREATE TABLE bikeshop.bikes (
    id           SERIAL PRIMARY KEY,
    version      BIGINT NOT NULL DEFAULT 1,
    brand        VARCHAR(255) NOT NULL,
    model        VARCHAR(255) NOT NULL,
    type         VARCHAR(100) NOT NULL,
    price        NUMERIC(10, 2) NOT NULL CHECK(price > 0),
    stock        INTEGER NOT NULL CHECK(stock >= 0),
    description  VARCHAR(1000),

    UNIQUE (brand, model)
);

CREATE TABLE bikeshop.orders (
    id            SERIAL PRIMARY KEY,
    full_name     VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 3 AND 100),
    email         VARCHAR(255) NOT NULL CHECK(email ~ '^[^@]+@[^@]+\.[^@]+$'),
    address       VARCHAR(500) NOT NULL CHECK(char_length(address) BETWEEN 5 AND 500),
    status        VARCHAR(50) NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'completed', 'cancelled')),
    total_price   NUMERIC(10, 2) NOT NULL CHECK(total_price > 0),
    created_at    TIMESTAMPTZ NOT NULL,
    completed_at  TIMESTAMPTZ,

    CHECK (
        (status = 'pending' AND completed_at IS NULL)
        OR
        (status = 'completed' AND completed_at IS NOT NULL AND completed_at >= created_at)
        OR
        (status = 'cancelled' AND completed_at IS NULL)
    )
);

CREATE TABLE bikeshop.bike_order (
    id                  SERIAL PRIMARY KEY,
    order_id            INTEGER NOT NULL REFERENCES bikeshop.orders(id) ON DELETE CASCADE,
    bike_id             INTEGER NOT NULL REFERENCES bikeshop.bikes(id),
    quantity            INTEGER NOT NULL CHECK(quantity > 0),
    price_at_purchase   NUMERIC(10, 2) NOT NULL CHECK(price_at_purchase > 0)
);