package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server    Server
	Database  Database
	Cache     Cache
	Messaging Messaging
}

type Server struct {
	Host string `envconfig:"DEV_SERVER_HOST" default:"0.0.0.0"`
	Port int    `envconfig:"DEV_SERVER_PORT" default:"8080"`
}

type Database struct {
	Postgres   Postgres
	ClickHouse ClickHouse
}

type Postgres struct {
	Username string `envconfig:"DEV_DATABASE_USERNAME"`
	Password string `envconfig:"DEV_DATABASE_PASSWORD"`
	Host     string `envconfig:"DEV_DATABASE_HOST"     default:"0.0.0.0"`
	Port     int    `envconfig:"DEV_DATABASE_PORT"     default:"5432"`
	DBName   string `envconfig:"DEV_DATABASE_DBNAME"`
}

type ClickHouse struct {
	Host     string `envconfig:"DEV_CLICKHOUSE_HOST"     default:"localhost"`
	Port     int    `envconfig:"DEV_CLICKHOUSE_PORT"     default:"8123"`
	Database string `envconfig:"DEV_CLICKHOUSE_DATABASE"`
	User     string `envconfig:"DEV_CLICKHOUSE_USER"`
	Password string `envconfig:"DEV_CLICKHOUSE_PASSWORD"`
}

type Cache struct {
	Host     string `envconfig:"DEV_REDIS_HOST"     default:"localhost"`
	Port     int    `envconfig:"DEV_REDIS_PORT"     default:"6379"`
	Password string `envconfig:"DEV_REDIS_PASSWORD"`
	DB       int    `envconfig:"DEV_REDIS_DB"       default:"0"`
}

type Messaging struct {
	URL string `envconfig:"DEV_NATS_URL" default:"nats://localhost:4222"`
}

// LoadConfig reads config from environment variables.
func LoadConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("dev", &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshall config: %w", err)
	}

	return cfg, nil
}
