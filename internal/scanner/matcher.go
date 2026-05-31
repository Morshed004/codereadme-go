package scanner

import (
	"path/filepath"
	"strings"
)

func ShouldIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		// wildcard support
		matched, _ := filepath.Match(pattern, filepath.Base(path))
		if matched {
			return true
		}

		// directory prefix match
		if strings.HasPrefix(path, pattern) {
			return true
		}
	}

	return false
}