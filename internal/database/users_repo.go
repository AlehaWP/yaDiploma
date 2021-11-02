package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/encription"
)

type UserRepo struct {
	serverDB
	models.User
}

func (u UserRepo) Find(ctx context.Context) bool {
	db := u.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id FROM users WHERE user_token=$1 OR user_name=$2`
	var id int
	row := db.QueryRowContext(ctx, q, u.Token, u.Login)

	if err := row.Scan(&id); err != nil {
		return false
	}
	if id == 0 {
		return false
	}
	return true
}

func (u UserRepo) SignIn(ctx context.Context) bool {
	db := u.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, user_token FROM users WHERE user_name=$1 and user_password=$2`
	var (
		token string
		id    int
	)
	row := db.QueryRowContext(ctx, q, u.Login, u.Password)

	if err := row.Scan(&id, &token); err != nil {
		return false
	}
	if id == 0 {
		return false
	}
	if len(token) == 0 {
		u.Token = encription.EncriptStr(u.Login)
	}
	return true
}

func (u UserRepo) Add(ctx context.Context) bool {
	db := u.db
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
