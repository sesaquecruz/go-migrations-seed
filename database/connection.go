package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Connect(driver string, dsn string) (*sql.DB, error) {
	log.Println("connecting to database")

	switch driver {
	case "postgres":
		dsn = fmt.Sprintf("%s://%s", driver, dsn)
		break
	case "mysql":
		break
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
