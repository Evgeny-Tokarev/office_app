package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evgeny-tokarev/office_app/backend/internal/app"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}
	fmt.Println("Config: ", cfg)
	if err := app.Run(cfg); err != nil {
		log.Fatal("error running http server ", err)
	}
}
