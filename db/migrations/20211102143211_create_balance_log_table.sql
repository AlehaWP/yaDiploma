-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS balance_log (
		id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
		summ NUMERIC,
		date_add timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
