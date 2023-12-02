package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Connect(driver string, dsn string) (*sql.DB, error) {
	switch driver {
	case "postgres":
		break
	case "mysql":
		break
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}

	db, err := sql.Open(driver, fmt.Sprintf("%s://%s", driver, dsn))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
