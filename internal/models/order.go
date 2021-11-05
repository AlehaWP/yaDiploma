package models

import "context"

type Status string

const (
	Processed  Status = "PROCESSED"
	Processing Status = "PROCESSING"
	Invalid    Status = "INVALID"
	New        Status = "NEW"
)

type Order struct {
	ID      int
	OrderID string
	Status  Status
	Accural float32
	DateAdd string
	UserID  int
}

type OrdersRepo interface {
	Get(context.Context, *Order) (bool, error)
	GetAll(context.Context, int) ([]Order, error)
	Add(context.Context, *Order) error
}
