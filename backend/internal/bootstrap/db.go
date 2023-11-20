package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	_ "github.com/jackc/pgx/stdlib"
)

var TestDB *sql.DB

func InitSqlDB(cfg config.Config) (*sql.DB, error) {

	db, err := sql.Open("pgx", FormatConnect(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func FormatConnect(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPwd, cfg.PgHost, cfg.PgPort, cfg.PgDBName,
	)
}
