package seed

import (
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sesaquecruz/go-sql-seed/utils"
)

func ApplyData(db *sql.DB, path string) error {
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

	return nil
}
