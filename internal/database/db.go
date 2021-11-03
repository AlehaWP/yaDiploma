package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

var once sync.Once
var sr serverDB

type serverDB struct {
	db *sql.DB
}

//CheckDBConnection trying connect to db.
func (s serverDB) CheckDBConnection(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
	}
}

func connect() {
	ctx := context.Background()
	db, err := sql.Open("postgres", config.Cfg.DBConnString())
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
	}
	sr.db = db
	sr.CheckDBConnection(ctx)
}

func OpenDBConnect() {
	once.Do(connect)
}

func Close() {
	sr.db.Close()
}
