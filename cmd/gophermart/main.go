package main

import (
	"database/sql"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/pressly/goose/v3"
)

// makeMigrations start here for autotests
func makeMigrations() {
	logger.NewLogs()
	p := "Миграции базы данных:"
	logger.Info(p, "Старт")
	config.NewConfig()
	logger.Info(p, "Подключение к БД")
	db, err := sql.Open("postgres", config.Cfg.DBConnString())
	if err != nil {
		logger.Error(p, err)
	}

	defer db.Close()
	// setup database
	logger.Info(p, "Применение миграций")
	if err := goose.Up(db, "../../db/migrations"); err != nil {
		logger.Error(p, err)
	}
	logger.Info(p, "Завершение") // run app
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
	makeMigrations()
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт сервера")

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	config.NewConfig()

	// sDB := database.OpenDBConnect()
	// defer sDB.Close()

	// // go ossignal.HandleQuit(cancel)

	// s := new(server.Server)
	// s.ServerDB = sDB
	// s.Start(ctx)

	http.HandleFunc("/", HelloWorld)

	http.HandleFunc("/api/user/register", HelloWorld)
	http.HandleFunc("/api/user/login", HelloWorld)
	// запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)
	logger.Info("Сервер остановлен")

}
