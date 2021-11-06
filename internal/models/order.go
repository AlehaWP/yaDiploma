package models

import "context"

type Status string

const (
	OrderStatusProcessed  Status = "PROCESSED"
	OrderStatusProcessing Status = "PROCESSING"
	OrderStatusInvalid    Status = "INVALID"
	OrderStatusNew        Status = "NEW"
)

type Order struct {
	ID      int     `json:"-"`
	OrderID int     `json:"number"`
	Status  Status  `json:"status"`
	Accural float32 `json:"accrual"`
	DateAdd string  `json:"uploaded_at"`
	UserID  int     `json:"-"`
}

type OrdersRepo interface {
	Get(context.Context, *Order) (bool, error)
	GetAll(context.Context, int) ([]Order, error)
	Add(context.Context, *Order) error
}
