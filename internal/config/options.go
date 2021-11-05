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
	if len(e.AccuralAddress) != 0 {
		c.accuralAddress = e.AccuralAddress
	}
}

func (c *Config) setDefault() {

	if len(c.servAddr) == 0 {
		c.servAddr = "localhost:8080"
	}
	if len(c.dbConnString) == 0 {
		c.dbConnString = "user=kseikseich dbname=yad sslmode=disable"
	}
	if len(c.accuralAddress) == 0 {
		c.accuralAddress = "http://localhost:8081"
	}
}

//setFlags for get options from console to default application options.
func (c *Config) setFlags() {
	flag.StringVar(&c.servAddr, "a", "", "a server address string")
	flag.StringVar(&c.dbConnString, "d", "", "a db connection string")
	flag.StringVar(&c.accuralAddress, "r", "", "a accural system address")
	flag.Parse()
}

func createConfig() {
	Cfg = new(Config)

	appDir, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	Cfg.checkEnv()
	Cfg.setFlags()
	Cfg.setDefault()
	Cfg.appDir = appDir
	logger.Info("Создан config")
}

// NewDefOptions return obj like Options interfase.
func NewConfig() {
	once.Do(createConfig)
}
