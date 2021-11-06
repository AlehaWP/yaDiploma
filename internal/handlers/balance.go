package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func HandlerGetUserBalance(br models.BalanceRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Запрос баланса пользователя")
		ctx := r.Context()
		userID := ctx.Value(models.UKeyName).(int)

		cb, err := br.Get(ctx, userID)
		if err != nil {
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(cb)
		if err != nil {
			logger.Info("Ошибка маршализации", res)
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func HandlerGetUserWithdrawals(br models.BalanceRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Запрос списаний пользователя")
		ctx := r.Context()
		userID := ctx.Value(models.UKeyName).(int)

		cb, err := br.GetAll(ctx, userID)
		if err != nil {
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(&cb)
		if err != nil {
			logger.Info("Ошибка маршализации", res)
			logger.Info(http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
