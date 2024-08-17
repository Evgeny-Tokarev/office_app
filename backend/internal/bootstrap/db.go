package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	_ "github.com/jackc/pgx/stdlib"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *sql.DB

func InitSqlDB(cfg config.Config) (*sql.DB, error) {
	connectionString := FormatConnectForSql(cfg)
	log.Info("Connecting to database with connection string:", connectionString)
	db, err := sql.Open("pgx", FormatConnectForSql(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func FormatConnectForSql(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPwd, cfg.PgHost, cfg.PgPort, cfg.PgDBName,
	)
}

func InitGormDB(cfg config.Config) (*gorm.DB, error) {
	connectionString := FormatConnectForSql(cfg)
	log.Info("Connecting to database with connection string:", connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
