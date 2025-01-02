package config

type Config struct {
	Port        int    `env:"SERVER_PORT" envDefault:"13005"`
	PgPort      string `env:"PG_PORT" envDefault:"5458"`
	PgHost      string `env:"PG_HOST" envDefault:"0.0.0.0"`
	PgDBName    string `env:"PG_DB_NAME" envDefault:"db"`
	PgUser      string `env:"PG_USER" envDefault:"db"`
	PgPwd       string `env:"PG_PWD" envDefault:"no-db"`
	LogLvl      string `env:"LOG_LEVEL" envDefault:"DebugLevel"`
	TokenSecret string `env:"TOKEN_SECRET" envDefault:"secret"`
	TokenType   string `env:"TOKEN_TYPE" envDefault:"jwt"`
	GeoApiToken string `env:"GOOGLE_MAPS_API_TOKEN"`
}
