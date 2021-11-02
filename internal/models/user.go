package models

import "context"

type UserKeyName string

const UKeyName UserKeyName = "UserUID"

type User struct {
	Login    string
	Password string
	Token    string
	UserID   int
}

type UsersRepo interface {
	Find(context.Context) bool
	SignIn(context.Context) bool
	Add(context.Context) bool
	Del(string) bool
	Get(string) User
}
