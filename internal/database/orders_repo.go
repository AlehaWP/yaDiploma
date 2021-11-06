package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

type DBOrdersRepo struct {
	serverDB
}

func (db *DBOrdersRepo) Get(ctx context.Context, o *models.Order) (bool, error) {
	logger.Info("Проверка наличия заказа")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, user_id FROM orders WHERE order_id=$1`
	row := db.QueryRowContext(ctx, q, o.OrderID)

	if err := row.Scan(&o.ID, &o.UserID); err != nil && err != sql.ErrNoRows {
		logger.Info(err)
		return false, err
	}
	if o.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (db *DBOrdersRepo) GetAll(ctx context.Context, userID int) ([]models.Order, error) {
	logger.Info("Запрос заказов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, order_id, accrual, order_status, date_add FROM orders WHERE user_id=$1`
	rows, err := db.QueryContext(ctx, q, userID)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var aOrders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.OrderID, &o.Accural, &o.Status, &o.DateAdd); err != nil {
			logger.Info(err)
			return nil, err
		}
		aOrders = append(aOrders, o)
	}
	err = rows.Err()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return aOrders, nil
}

func (db *DBOrdersRepo) Add(ctx context.Context, o *models.Order) error {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	q := `INSERT INTO orders (order_id, user_id, order_status) VALUES ($1,$2, $3) RETURNING ID`
	row := db.QueryRowContext(ctx, q, o.OrderID, o.UserID, models.OrderStatusNew)

	if err := row.Scan(&o.ID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func (s serverDB) NewDBOrdersRepo() models.OrdersRepo {
	or := new(DBOrdersRepo)
	or.serverDB = s
	return or
}
