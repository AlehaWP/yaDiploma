package database

import (
	"context"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

type DBOrdersRepo struct {
	serverDB
}

func (db *DBOrdersRepo) Get(ctx context.Context, r *models.Order) (bool, error) {
	// ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	// defer cancelfunc()
	// q := `SELECT id, user_name, user_token, user_password FROM users WHERE user_token=$1 OR user_name=$2`
	// row := db.QueryRowContext(ctx, q, u.Token, u.Login)

	// if err := row.Scan(&u.UserID, &u.Login, &u.Token, &u.Password); err != nil {
	// 	logger.Info(err)
	// 	return false, err
	// }
	// if u.UserID == 0 {
	// 	return false, nil
	// }
	// if len(u.Token) == 0 {
	// 	u.Token = encription.EncriptStr(u.Login)
	// 	db.update(ctx, u)
	// }
	return true, nil

}

func (db *DBOrdersRepo) GetAll(ctx context.Context, userID int) ([]models.Order, error) {
	return nil, nil
}

func (db *DBOrdersRepo) Add(ctx context.Context, r *models.Order) (bool, error) {
	return false, nil
}

func (s serverDB) NewDBOrderRepo() models.OrdersRepo {
	ur := new(DBOrdersRepo)
	ur.serverDB = s
	return ur
}
