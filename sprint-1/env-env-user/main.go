package main

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	User string `env:"USER"`
}

func main() {
	var config Config

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	log.Println(config)
}
