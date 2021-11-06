package main

<<<<<<< HEAD
import "fmt"

func main() {
	fmt.Println("Поехали")
=======
import (
	"context"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/internal/server"
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

	s := new(server.Server)
	s.ServerDB = sDB
	s.Start(ctx)

	logger.Info("Сервер остановлен")

>>>>>>> balance
}
