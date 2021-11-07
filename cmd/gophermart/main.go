package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/accrual"
	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/AlehaWP/yaDiploma.git/pkg/ossignal"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
	//makeMigrations()
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт сервера")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config.NewConfig()

	sDB := database.OpenDBConnect()
	defer sDB.Close()

	go ossignal.HandleQuit(cancel)

	l := accrual.New(ctx, sDB.NewDBOrdersRepo(), sDB.NewDBBalanceRepo())

	for i := 0; i < 1000000; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			l.Put(strconv.Itoa(i))
		}
	}

	p := <-ctx.Done()
	fmt.Println("Завершено")
	fmt.Println(p)
	time.Sleep(3 * time.Second)

	// s := new(server.Server)
	// s.ServerDB = sDB
	// s.Start(ctx)

	// logger.Info("Сервер остановлен")

}
