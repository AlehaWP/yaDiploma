package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/encription"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

type DBUserRepo struct {
	serverDB
}

func (d DBUserRepo) Locate(ctx context.Context, u *models.User) bool {
	db := d.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, user_name, user_token, user_password FROM users WHERE user_token=$1 OR user_name=$2`
	row := db.QueryRowContext(ctx, q, u.Token, u.Login)

	if err := row.Scan(&u.UserID, &u.Login, &u.Token, &u.Password); err != nil {
		logger.Info(err)
		return false
	}
	if u.UserID == 0 {
		return false
	}
	if len(u.Token) == 0 {
		u.Token = encription.EncriptStr(u.Login)
		d.update(ctx, u)
	}
	return true
}

func (d DBUserRepo) Add(ctx context.Context, u *models.User) bool {
	db := d.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	u.Token = encription.EncriptStr(u.Login)

	q := `INSERT INTO users (user_name, user_password, user_token) VALUES ($1,$2,$3)`
	_, err := db.ExecContext(ctx, q, u.Login, u.Password, u.Token)

	if err != nil {
		return false
	}
	return true
}

func (d DBUserRepo) update(ctx context.Context, u *models.User) bool {
	db := d.db
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

func (d DBUserRepo) Del(ctx context.Context, u *models.User) bool {
	return false
}

func NewDBUserRepo() models.UsersRepo {
	ur := new(DBUserRepo)
	ur.serverDB = sr
	return ur
}
