package handlers

import (
	"io"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/AlehaWP/yaDiploma.git/pkg/luhn"
)

func HandlersNewOrder(or models.OrdersRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Обработка нового заказа")
		ctx := r.Context()
		userID := ctx.Value(models.UKeyName).(int)
		order := new(models.Order)
		order.UserID = userID

		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ordNum := string(b)
		if ok := luhn.CheckString(ordNum); !ok {
			logger.Info(ordNum, http.StatusUnprocessableEntity)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		order.OrderID = string(b)

		finded, err := or.Get(ctx, order)
		if err != nil {
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if finded {
			status := http.StatusOK
			if order.UserID != userID {
				status = http.StatusConflict
			}
			logger.Info(status)
			w.WriteHeader(status)
			return
		}
		if ok, err := or.Add(ctx, order); !ok || err != nil {
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Info(http.StatusAccepted)
		w.WriteHeader(http.StatusAccepted)
	}
}
