package models

import "context"

type Balance struct {
	BI BalanceIn
	BO BalanceOut
	In bool
}

type BalanceIn struct {
	UserID   int     `json:"-"`
	OrederID string  `json:"order"`
	Sum      float32 `json:"accrual"`
}

type BalanceOut struct {
	UserID    int         `json:"-"`
	OrederID  string      `json:"order"`
	Sum       float32     `json:"sum"`
	Status    OrderStatus `json:"status"`
	Processed string      `json:"processed_at"`
}

type CurrentBalance struct {
	UserID     int     `json:"-"`
	CurBalance float32 `json:"current"`
	Withdrawn  float32 `json:"withdrawn"`
}

type BalanceRepo interface {
	Add(context.Context, *Balance) error
	Get(ctx context.Context, UserId int) (*CurrentBalance, error)
	GetAll(ctx context.Context, UserId int) ([]BalanceOut, error)
}
