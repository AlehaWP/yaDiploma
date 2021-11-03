package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func getUserFromRequest(r *http.Request) (*models.User, bool) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Info("Ошибка обработки запроса", err)
		return nil, false
	}
	user := new(models.User)

	err = json.Unmarshal(b, user)
	if err != nil {
		logger.Info("Ошибка Unmarshal", err)
		return nil, false
	}
	return user, true
}

func HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	logger.Info("Обработка запроса регистрации")
	ctx := r.Context()

	user, ok := getUserFromRequest(r)
	if !ok {
		logger.Info(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userRepo := database.NewDBUserRepo()

	if finded := userRepo.Locate(ctx, user); finded {
		logger.Info(http.StatusConflict)
		w.WriteHeader(http.StatusConflict)
		return
	}

	if ok := userRepo.Add(ctx, user); !ok {
		logger.Info(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Authorization", "Bearer "+user.Token)
	w.WriteHeader(http.StatusCreated)
}
