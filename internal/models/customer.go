package models

type Customer struct {
	User
	Balance float32
}

type CustomersRepo interface {
	Find(string) bool
	Add(Customer) error
	Del(string) bool
	Get(string) Customer
}
