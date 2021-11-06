-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_name VARCHAR(50) UNIQUE,
        user_password VARCHAR(36),
        user_key VARCHAR(36),
        user_token VARCHAR(36),
		date_add TIMESTAMPTZ default (NOW() at time zone 'UTC+3'));

CREATE TABLE IF NOT EXISTS customers (
        id SERIAL PRIMARY KEY,
        user_id INT UNIQUE,
        balance NUMERIC,
		date_add TIMESTAMPTZ default (NOW() at time zone 'UTC+3'));

CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NUL,
		order_id BIGINT UNIQUE,
        accrual NUMERIC default 0,
        order_status VARCHAR(20),
		date_add TIMESTAMPTZ default (NOW() at time zone 'UTC+3'));

CREATE TABLE IF NOT EXISTS balance_log (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        order_id VARCHAR(50),
        summ NUMERIC,
        date_add TIMESTAMPTZ default (NOW() at time zone 'UTC+3'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
