-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(36),
		order_id VARCHAR(50) UNIQUE,
        bonus NUMERIC,
		date_add timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
