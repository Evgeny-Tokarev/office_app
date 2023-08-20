package bootstrap

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lesson4/internal/config"
)

func InitSqlDB(cfg config.Config) (*sql.DB, error) {

	db, err := sql.Open("pgx", formatConnect(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitSqlxDB(cfg config.Config) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", formatConnect(cfg))
}

func InitGormDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PgHost, cfg.PgUser, cfg.PgPwd, cfg.PgDBName, cfg.PgPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func formatConnect(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPwd, cfg.PgHost, cfg.PgPort, cfg.PgDBName,
	)
}
