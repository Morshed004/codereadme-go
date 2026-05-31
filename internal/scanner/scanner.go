package scanner

import (
	"os"
	"path/filepath"
)

type FileData struct {
	Path    string
	Content string
}

func Scan(root string) ([]FileData, error) {
	patterns, err := LoadIgnorePatterns()
	if err != nil {
		return nil, err
	}

	var files []FileData

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		cleanPath := filepath.Clean(path)

		if ShouldIgnore(cleanPath, patterns) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		files = append(files, FileData{
			Path:    path,
			Content: string(content),
		})

		return nil
	})

	return files, err
}