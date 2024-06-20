package util

import (
	"os"
	"path/filepath"
	"strings"
)

func PathSearch() ([]string, error) {
	path := os.Getenv("PATH")
	var filePaths []string

	if path != "" {
		// crawl path for more things to search
		pathFragments := strings.Split(path, ":")
		for _, f := range pathFragments {
			entries, err := os.ReadDir(f)
			if err != nil {
				return filePaths, err
			}
			for _, e := range entries {
				filePaths = append(filePaths, filepath.Join(f, e.Name()))
			}
		}
	}

	return filePaths, nil
}
