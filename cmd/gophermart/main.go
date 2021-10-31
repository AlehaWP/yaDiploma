package main

import (
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func main() {
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт сервера")
}
