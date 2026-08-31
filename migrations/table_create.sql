CREATE TABLE IF NOT EXISTS restaurants (
    id bigint NOT NULL,
    name VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    website_info VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS dishes (
    id bigint NOT NULL,
    restaurant_id bigint NOT NULL,
    name VARCHAR NOT NULL,
    cuisine_type VARCHAR NOT NULL,
    currency_code VARCHAR NOT NULL,
    price int NOT NULL,
    prep_time INTERVAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS employees (
    id bigint NOT NULL,
    restaurant_id bigint NOT NULL,
    name VARCHAR NOT NULL,
    emp_type VARCHAR NOT NULL,
    free boolean NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customers (
    id bigint NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id bigint NOT NULL,
    restaurant_id bigint NOT NULL,
    food_id bigint ARRAY NOT NULL,
    paid boolean NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE countdown (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    prepare_time INTERVAL NOT NULL
);
