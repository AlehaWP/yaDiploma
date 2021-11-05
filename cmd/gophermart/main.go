package main

import (
	"context"
	"database/sql"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/internal/server"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/AlehaWP/yaDiploma.git/pkg/ossignal"
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

	<-ctx.Done()
	logger.Info("Сервер остановлен")

}
