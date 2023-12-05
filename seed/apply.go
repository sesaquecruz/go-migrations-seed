package seed

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sesaquecruz/go-sql-seed/config"
	"github.com/sesaquecruz/go-sql-seed/database"
	"github.com/sesaquecruz/go-sql-seed/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplyAll(cfg *config.Config) error {
	db, err := database.Connect(cfg.DbDriver, cfg.DbDsn)
	if err != nil {
		return err
	}

	err = ApplyMigrations(cfg.DbDriver, cfg.DbDsn, cfg.MigrationsPath)
	if err != nil {
		return err
	}

	err = ApplyData(db, cfg.DataPath)
	if err != nil {
		return err
	}

	return nil
}

func ApplyMigrations(driver string, dsn string, path string) error {
	log.Println("applying migrations")

	migrations, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		fmt.Sprintf("%s://%s", driver, dsn),
	)
	if err != nil {
		return err
	}
	defer migrations.Close()

	err = migrations.Up()
	if err != nil {
		return err
	}

	log.Println("migrations applied")
	return nil
}

func ApplyData(db *sql.DB, path string) error {
	log.Println("applying data")

	filter := func(name string) bool {
		return strings.HasSuffix(name, ".sql")
	}

	files, err := utils.ReadDir(path, filter)
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, file := range files {
		data, err := os.ReadFile(filepath.Join(path, file))
		if err != nil {
			return err
		}

		_, err = db.Exec(string(data))
		if err != nil {
			return err
		}
	}

	log.Println("data applied")
	return nil
}
