package accrual

import (
	"container/list"
	"context"
	"fmt"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

type listOrders struct {
	*list.List
	or    models.OrdersRepo
	br    models.BalanceRepo
	chPut chan string
}

func (l *listOrders) Put(o string) {
	l.chPut <- o
}

func SendData(ctx context.Context, l *listOrders) {
	for {
		select {
		case o := <-l.chPut:
			l.PushBack(o)
		case p := <-ctx.Done():
			fmt.Println("Завершено")
			fmt.Println(p)
			close(l.chPut)
			return
		default:
			if l.Len() > 0 {
				fmt.Println("len: ", l.Len())
				e := l.Front()
				num := e.Value.(string)
				fmt.Println("num :", num)
				l.Remove(e)
			}
		}
	}
}

func New(ctx context.Context, o models.OrdersRepo, b models.BalanceRepo) *listOrders {
	l := &listOrders{
		List:  list.New(),
		or:    o,
		br:    b,
		chPut: make(chan string),
	}
	go SendData(ctx, l)
	return l

}
