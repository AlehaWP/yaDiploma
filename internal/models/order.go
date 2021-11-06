package models

import "context"

type OrderStatus string

const (
	OrderStatusProcessed  OrderStatus = "PROCESSED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusNew        OrderStatus = "NEW"
)

type Order struct {
	ID      int         `json:"-"`
	OrderID string      `json:"number"`
	Status  OrderStatus `json:"status"`
	Accural float32     `json:"accrual"`
	DateAdd string      `json:"uploaded_at"`
	UserID  int         `json:"-"`
}

type OrdersRepo interface {
	Get(context.Context, *Order) (bool, error)
	GetAll(ctx context.Context, UserID int) ([]Order, error)
	Add(context.Context, *Order) error
}
