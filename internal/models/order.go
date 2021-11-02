package models

type Order struct {
	Id     string
	Bonus  float32
	UserID int
}

type OrderRepo interface {
	Get(id string) Order
	Add(Order) error
}
