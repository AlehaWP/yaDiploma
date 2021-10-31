package models

type Customer struct {
	User
	Balance float32
}

type CustomerRepo interface {
	Find(string) bool
	Add(Customer) error
	Del(string) bool
	Get(string) Customer
}
