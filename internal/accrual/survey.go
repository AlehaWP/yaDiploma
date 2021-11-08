package accrual

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

type listOrders struct {
	or    models.OrdersRepo
	br    models.BalanceRepo
	chPut chan *models.Order
}

func (l *listOrders) Put(ctx context.Context, o *models.Order) {
	select {
	case <-ctx.Done():
		break
	default:
		l.chPut <- o
	}
}

func (l *listOrders) sendData(ctx context.Context, o *models.Order) {
	resp, err := http.Get("http://localhost:8082/api/orders/" + o.OrderID)
	if err != nil {
		logger.Info("Error", "Ошибка выполнения запроса", err)
		return
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		logger.Info(http.StatusTooManyRequests, "Приостановка запросов к серверу расчета начислений")
		time.Sleep(time.Second)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		logger.Info(http.StatusInternalServerError, "Ошибка обработки запроса сервером начислений", o)
		return
	}

	if resp.StatusCode == http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Info("Error", "Ошибка чтения ответа на запрос", err)
			return
		}
		defer resp.Body.Close()

		oOut := new(models.OrderFromAccuyral)
		err = json.Unmarshal(b, oOut)
		if err != nil {
			logger.Info("Error", "Ошибка чтения ответа на запрос", err)
			return
		}
		if o.Status != oOut.Status {
			o.OrderID = oOut.OrderID
			o.Status = oOut.Status
			o.Accural = oOut.Accural

			l.or.Update(ctx, o)

			if oOut.Status == models.OrderStatusProcessed {
				bl := new(models.Balance)
				bl.UserID = o.UserID
				bl.OrderID = o.OrderID
				bl.SumIn = o.Accural
				err = l.br.Add(ctx, bl)
				if err != nil {
					logger.Info("Ошибка добавления баланса в лог", bl)
				}
			}
		}
	}
}

func BeginSurvey(ctx context.Context, a string, o models.OrdersRepo, b models.BalanceRepo, numOfWorkers int) {
	var wg sync.WaitGroup

	jobCh := make(chan *models.Order, 100)
	l := &listOrders{
		or:    o,
		br:    b,
		chPut: jobCh,
	}

	wg.Add(numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func() {
			for job := range jobCh {
				l.sendData(ctx, job)
			}
			wg.Done()
		}()
	}

	for {
		select {
		case <-ctx.Done():
			close(jobCh)
			wg.Wait()
			return
		default:
			l.getNumForSurvey(ctx, models.OrderStatusNew)
			l.getNumForSurvey(ctx, models.OrderStatusProcessing)
			time.Sleep(2 * time.Second)
		}
	}

}

func (l *listOrders) getNumForSurvey(ctx context.Context, st models.OrderStatus) {
	arrOr, err := l.or.GetAllStatus(ctx, st)
	if err != nil {
		logger.Info("Ошибка запроса не обработанных заказов", err)
	}

	for _, v := range arrOr {
		l.chPut <- v
	}
}
