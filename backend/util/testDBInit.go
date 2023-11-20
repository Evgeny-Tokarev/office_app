package util

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evgeny-tokarev/office_app/backend/internal/bootstrap"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"time"
)

type CreateOfficeParams struct {
	Name    string `db:"name"`
	Address string `db:"address"`
}

type Office struct {
	ID        int64          `db:"id"`
	Name      string         `db:"name"`
	Address   string         `db:"address"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	ImgFile   sql.NullString `db:"img_file"`
}

func InitTestDB() (*sql.DB, error) {
	cfg := config.Config{}
	var testDB *sql.DB

	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("unable to load .env file for test: %e", err)
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to retrieve env variables for test, %v", err)
	}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("failed to retrieve env variables, %v", err)
	}
	fmt.Println("Config for test DB: ", cfg)
	testDB, err = sql.Open("pgx", bootstrap.FormatConnect(cfg))
	if err != nil {
		return nil, err
	}
	return testDB, nil
}
