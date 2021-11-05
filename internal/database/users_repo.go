package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/encription"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

type DBUserRepo struct {
	serverDB
}

func (db DBUserRepo) Get(ctx context.Context, u *models.User) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT COALESCE(id, 0), user_name, user_token, user_password FROM users WHERE user_token=$1 OR user_name=$2`
	row := db.QueryRowContext(ctx, q, u.Token, u.Login)
	if err := row.Scan(&u.UserID, &u.Login, &u.Token, &u.Password); err != nil && err != sql.ErrNoRows {
		logger.Info(q, err)
		return false, err
	}
	if u.UserID == 0 {
		return false, nil
	}
	if len(u.Token) == 0 {
		u.Token = encription.EncriptStr(u.Login)
		db.update(ctx, u)
	}
	return true, nil
}

func (db DBUserRepo) Add(ctx context.Context, u *models.User) (bool, error) {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	u.Token = encription.EncriptStr(u.Login)

	q := `INSERT INTO users (user_name, user_password, user_token) VALUES ($1,$2,$3)`
	_, err := db.ExecContext(ctx, q, u.Login, u.Password, u.Token)

	if err != nil {
		logger.Info(q, err)
		return false, err
	}
	return true, nil
}

func (db DBUserRepo) update(ctx context.Context, u *models.User) bool {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	u.Token = encription.EncriptStr(u.Login)

	q := `UPDATE users SET user_name=$2, user_password=$3, user_token=$4 WHERE ID=$4`
	_, err := db.ExecContext(ctx, q, u.UserID, u.Login, u.Password, u.Token)

	if err != nil {
		return false
	}
	return true
}

func (db DBUserRepo) Del(ctx context.Context, u *models.User) (bool, error) {
	return false, nil
}

func (s serverDB) NewDBUserRepo() models.UsersRepo {
	ur := new(DBUserRepo)
	ur.serverDB = s
	return ur
}
