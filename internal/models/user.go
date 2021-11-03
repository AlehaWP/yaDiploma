package models

import "context"

type UserKeyName string

const UKeyName UserKeyName = "UserUID"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string
	UserID   int
}

type UsersRepo interface {
	Locate(context.Context, *User) bool
	Add(context.Context, *User) bool
	Del(context.Context, *User) bool
}
