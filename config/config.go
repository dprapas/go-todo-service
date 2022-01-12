package config

import (
	"github.com/caarlos0/env/v6"
)

var (
	Configuration Config
)

func Load() {
	Configuration = Config{}
	if err := env.Parse(&Configuration); err != nil {
		panic("Failed to load configuration. Error was: " + err.Error())
	}
}

type Config struct {
	Server
	Application
}

type Server struct {
	Port           string `env:"SERVER_PORT" envDefault:"8080"`
	ReadTimeoutMS  int    `env:"READ_TIMEOUT_MS" envDefault:"15000"`
	WriteTimeoutMS int    `env:"WRITE_TIMEOUT_MS" envDefault:"15000"`
}

type Application struct {
	LogLevel        string `env:"LOG_LEVEL" envDefault:"INFO"`
	ClientTimeoutMS int    `env:"TIMEOUT_MS" envDefault:"30000"`
}
