package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/AlehaWP/yaDiploma.git/internal/accrual"
	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/internal/server"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/AlehaWP/yaDiploma.git/pkg/ossignal"
	"github.com/AlehaWP/yaDiploma.git/pkg/workers"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
	//makeMigrations()
	var wg sync.WaitGroup
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт сервера")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config.NewConfig()

	sDB := database.OpenDBConnect()
	defer sDB.Close()

	wg.Add(2)
	go func() {
		ossignal.HandleQuit(cancel)
		wg.Done()
	}()

	w := workers.NewWorkersPool(10)
	defer w.Close()

	l := accrual.NewSurveyAccrual(sDB.NewDBOrdersRepo(), sDB.NewDBBalanceRepo(), w)
	go func() {
		l.GetOrdersForSurveyFromDB(ctx)
		wg.Done()
	}()

	s := new(server.Server)
	s.ServerDB = sDB
	s.Start(ctx)
	wg.Wait()
	logger.Info("Сервер остановлен")

}
