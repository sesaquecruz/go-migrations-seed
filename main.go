package main

import (
	"log"

	"github.com/sesaquecruz/go-sql-seed/config"
	"github.com/sesaquecruz/go-sql-seed/database"
	"github.com/sesaquecruz/go-sql-seed/seed"
)

func main() {
	cfg := config.GetConfig()

	log.Println("connecting to database")
	db, err := database.Connect(cfg.DbDriver, cfg.DbDsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("applying migrations")
	err = seed.ApplyMigrations(cfg.DbDriver, cfg.DbDsn, cfg.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("applying data")
	err = seed.ApplyData(db, cfg.DataPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("all sql seeds were applied")
}
