package models

type UserKeyName string

const UKeyName UserKeyName = "UserID"

type User struct {
	Login    string
	Password string
	Token    string
}

type UserRepo interface {
	Find(string) bool
	Add(User) string
	Del(string) bool
	Get(string) User
}
