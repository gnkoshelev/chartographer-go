package config

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
}

var Cfg Config

func New() {
	var serverAddress string
	var once sync.Once

	once.Do(func() {
		if err := env.Parse(&Cfg); err != nil {
			log.Fatal(err)
		}

		flag.StringVar(&serverAddress, "a", Cfg.ServerAddress, "An address and a port of service.")

		flag.Parse()

		Cfg.ServerAddress = serverAddress
	})
}
