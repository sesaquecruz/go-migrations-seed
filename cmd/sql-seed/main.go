package main

import (
	"log"

	"github.com/sesaquecruz/go-sql-seed/config"
	"github.com/sesaquecruz/go-sql-seed/seed"
)

func main() {
	cfg := config.GetConfig()

	err := seed.ApplyAll(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
