package ossignal

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func HandleQuit(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	logger.Info("Получен сигнал на закрытие сервера")
	<-c
	cancel()
}
