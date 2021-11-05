package models

import "context"

type UserKeyName string

const UKeyName UserKeyName = "UserID"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string
	ID       int
}

type UsersRepo interface {
	Get(context.Context, *User) (bool, error)
	Add(context.Context, *User) (bool, error)
	Del(context.Context, *User) (bool, error)
}
