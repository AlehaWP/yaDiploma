package database

import (
	"context"

	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func (s *serverDB) createTables(ctx context.Context) {
	tx, err := s.DB.Begin()
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users (
								    id SERIAL PRIMARY KEY,
									user_name VARCHAR(50) UNIQUE,
									user_password VARCHAR(36),
									user_key VARCHAR(36),
									user_token VARCHAR(36),
									date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'))
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}

	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS customers (
									id SERIAL PRIMARY KEY,
									user_id INT UNIQUE,
									balance NUMERIC default 0,
									date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'))
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS orders (
									id SERIAL PRIMARY KEY,
									user_id INT NOT NULL,
									order_id BIGINT UNIQUE,
									accrual NUMERIC default 0,
									order_status VARCHAR(20),
									date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'))
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS balance_log (
									id SERIAL PRIMARY KEY,
									user_id INT NOT NULL,
									order_id BIGINT,
									accrual NUMERIC,
									date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'))
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}

	tx.Commit()
}
