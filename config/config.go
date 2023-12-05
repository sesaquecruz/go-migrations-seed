package config

import (
	"log"
	"os"
)

type Config struct {
	DbDriver       string
	DbDsn          string
	MigrationsPath string
	DataPath       string
}

var cfg *Config

func GetConfig() Config {
	if cfg == nil {
		dbDriver, ok := os.LookupEnv("DB_DRIVER")
		if !ok || dbDriver == "" {
			log.Fatal("env var DB_DRIVER is required")
		}

		dbDsn, ok := os.LookupEnv("DB_DSN")
		if !ok || dbDsn == "" {
			log.Fatal("env var DB_DSN is required")
		}

		migrationsPath, ok := os.LookupEnv("MIGRATIONS_PATH")
		if !ok || migrationsPath == "" {
			log.Fatal("env var MIGRATIONS_PATH is required")
		}

		dataPath, ok := os.LookupEnv("DATA_PATH")
		if !ok || dataPath == "" {
			log.Fatal("env var DATA_PATH is required")
		}

		cfg = &Config{
			DbDriver:       dbDriver,
			DbDsn:          dbDsn,
			MigrationsPath: migrationsPath,
			DataPath:       dataPath,
		}
	}

	return *cfg
}
