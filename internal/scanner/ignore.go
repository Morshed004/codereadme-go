package scanner

import (
	"bufio"
	"os"
	"strings"
)

func LoadIgnorePatterns() ([]string, error) {
	file, err := os.Open(".codereadmeignore")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//"){
			continue
		}

		patterns = append(patterns, line)
	}

	return patterns, scanner.Err()
}