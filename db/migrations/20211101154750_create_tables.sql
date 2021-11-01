-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_test (
		id SERIAL PRIMARY KEY,
		user_uuid VARCHAR(36),
		user_name VARCHAR(50) UNIQUE,
        user_password VARCHAR(36),
        user_key VARCHAR(36),
        user_token VARCHAR(36),
        user_role VARCHAR(50),
		date_add timestamp);

CREATE TABLE IF NOT EXISTS users_test2 (
		id SERIAL PRIMARY KEY,
		user_uuid VARCHAR(36),
		user_name VARCHAR(50) UNIQUE,
        user_password VARCHAR(36),
        user_key VARCHAR(36),
        user_token VARCHAR(36),
        user_role VARCHAR(50),
		date_add timestamp)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
