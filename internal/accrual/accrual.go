package accrual

import (
	"context"
	"encoding/json"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/client"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

// servAddress string

type ordersDB struct {
	or models.OrdersRepo
	br models.BalanceRepo
}

func (o *ordersDB) getOrders(ctx context.Context, st models.OrderStatus) []*models.Order {
	arrOr, err := o.or.GetAllStatus(ctx, st)
	if err != nil {
		logger.Info("Ошибка запроса не обработанных заказов", err)
	}
	return arrOr
}

func (o *ordersDB) updateOrderDB(ctx context.Context, or *models.Order, oOut *models.OrderFromAccrual) {
	if or.Status != oOut.Status {
		or.OrderID = oOut.OrderID
		or.Status = oOut.Status
		or.Accrual = oOut.Accrual

		o.or.Update(ctx, or)

		if oOut.Status == models.OrderStatusProcessed {
			bl := new(models.Balance)
			bl.UserID = or.UserID
			bl.OrderID = or.OrderID
			bl.SumIn = or.Accrual
			err := o.br.Add(ctx, bl)
			if err != nil {
				logger.Info("Ошибка добавления баланса в лог", bl)
			}
		}
	}
}

// default:
// 	l.getNumForSurvey(ctx, models.OrderStatusNew)
// 	l.getNumForSurvey(ctx, models.OrderStatusRegistered)
// 	l.getNumForSurvey(ctx, models.OrderStatusProcessing)
// 	time.Sleep(1 * time.Second)
// }
// }

type Accrual struct {
	servAddress string
	odb         *ordersDB
	l           *WorkersPool
}

func (a *Accrual) includeTrailingBackSlash(st string) string {
	if st[len(st)-1:] != "/" {
		return st + "/"
	}
	return st
}

func (a *Accrual) Put(ctx context.Context, arr []*models.Order) {
	for _, v := range arr {
		a.l.Put(ctx, v)
	}
}

func (a *Accrual) GetOrdersForSurvey(ctx context.Context) {
	arr := a.odb.getOrders(ctx, models.OrderStatusNew)
	a.Put(ctx, arr)
}

func (a *Accrual) getOrderFromAccrual() func(context.Context, interface{}) {
	return func(ctx context.Context, o interface{}) {

		orderForUpdate := o.(*models.Order)
		oIn := new(models.OrderFromAccrual)
		b, ok := client.MakeRequest("GET", a.servAddress+"api/orders/"+orderForUpdate.OrderID, "", "", nil)
		if !ok {
			return
		}
		err := json.Unmarshal(b, oIn)
		if err != nil {
			logger.Info("Error", "Ошибка чтения ответа на запрос", err)
			return
		}

		a.odb.updateOrderDB(ctx, orderForUpdate, oIn)
	}
}

func NewSurveyAccrual(ctx context.Context, or models.OrdersRepo, br models.BalanceRepo, qtyChan int) *Accrual {
	ordersDB := &ordersDB{
		or: or,
		br: br,
	}
	a := new(Accrual)
	a.odb = ordersDB
	a.servAddress = a.includeTrailingBackSlash(config.Cfg.AccrualAddress())
	a.l = NewWorkersPool(ctx, a.getOrderFromAccrual(), qtyChan)

	return a
}
