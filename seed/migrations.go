package seed

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func ApplyMigrations(driver string, dsn string, path string) error {
	migrations, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		fmt.Sprintf("%s://%s", driver, dsn),
	)
	if err != nil {
		return err
	}
	defer migrations.Close()

	err = migrations.Up()
	return err
}
