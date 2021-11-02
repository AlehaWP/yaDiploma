-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_name VARCHAR(50) UNIQUE,
        user_password VARCHAR(36),
        user_key VARCHAR(36),
        user_token VARCHAR(36),
		date_add timestamp);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
