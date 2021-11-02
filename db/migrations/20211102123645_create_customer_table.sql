-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
        id SERIAL,
		user_id INT PRIMARY KEY,
        order_id INT,
        balance NUMERIC,
		date_add timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
