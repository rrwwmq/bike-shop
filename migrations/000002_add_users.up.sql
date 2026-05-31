CREATE TABLE bikeshop.users (
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL CHECK(email ~ '^[^@]+@[^@]+\.[^@]+$'),
    password_hash VARCHAR(255) NOT NULL,
    role          VARCHAR(50) NOT NULL DEFAULT 'customer' CHECK(role IN ('admin', 'customer')),
    created_at    TIMESTAMPTZ NOT NULL,

    UNIQUE(email)
);

ALTER TABLE bikeshop.orders
    ADD COLUMN user_id INTEGER REFERENCES bikeshop.users(id);