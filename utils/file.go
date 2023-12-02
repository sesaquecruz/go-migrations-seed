package utils

import (
	"os"
)

type FilterName = func(name string) bool

func ReadDir(path string, filter FilterName) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if filter(entry.Name()) {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}
