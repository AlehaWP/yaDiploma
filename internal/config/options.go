package config

import (
	"flag"
	"os"
	"sync"

	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/caarlos0/env/v6"
)

var Cfg *Config
var once sync.Once

type Config struct {
	servAddr       string
	dbConnString   string
	accuralAddress string
	appDir         string
}

func (c Config) ServAddr() string {
	return c.servAddr
}

func (c Config) DBConnString() string {
	return c.dbConnString
}

func (c Config) ProgramPath() string {
	return c.appDir
}

func (c Config) AccuralAddress() string {
	return c.accuralAddress
}

type EnvOptions struct {
	ServAddr       string `env:"RUN_ADDRESS"`
	DBConnString   string `env:"DATABASE_URI"`
	AccuralAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

//checkEnv for get options from env to default application options.
func (c *Config) checkEnv() {

	e := &EnvOptions{}
	err := env.Parse(e)
	if err != nil {
		logger.Info("Ошибка чтения конфигурации из переменного окружения", err)
	}
	if len(e.ServAddr) != 0 {
		c.servAddr = e.ServAddr
	}
	if len(e.DBConnString) != 0 {
		c.dbConnString = e.DBConnString
	}
}

//setFlags for get options from console to default application options.
func (c *Config) setFlags() {
	flag.StringVar(&c.servAddr, "a", "localhost:8080", "a server address string")
	flag.StringVar(&c.dbConnString, "d", "user=kseikseich dbname=yad sslmode=disable", "a db connection string")
	flag.StringVar(&c.accuralAddress, "r", "http://localhost:8081", "a accural system address")
	flag.Parse()
}

func createConfig() {
	Cfg = new(Config)

	appDir, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	Cfg.setFlags()
	Cfg.checkEnv()
	Cfg.appDir = appDir
	logger.Info("Создан config")
}

// NewDefOptions return obj like Options interfase.
func NewConfig() {
	once.Do(createConfig)
}
