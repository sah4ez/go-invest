package config

import (
	"os"
)

type AppConfig struct {
	PgUser     string
	PgPassword string
	PgDb       string
	ConnStr    string
	LogLevel   string
}

func New() *AppConfig {
	cfg := &AppConfig{
		PgUser:     os.Getenv("POSTGRES_USER"),
		PgPassword: os.Getenv("POSTGRES_PASSWORD"),
		PgDb:       os.Getenv("POSTGRES_DB"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
	}
	cfg.ConnStr = "user=" + cfg.PgUser + " dbname=" + cfg.PgDb + " password=" + cfg.PgPassword + " sslmode=disable"
	return cfg
}
