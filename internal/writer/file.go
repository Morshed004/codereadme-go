package writer

import (
	"fmt"
	"os"
)

func SaveREADME(content string) error {
	fileName := "README.md"

	if _, err := os.Stat(fileName); err == nil {
		backupName := "README.old.md"
		input, err := os.ReadFile(fileName)
		if err == nil {
			os.WriteFile(backupName, input, 0644)
			fmt.Println("📦 Backup created: README.old.md")
		}

		fmt.Println("🔄 Updating existing README.md")
	} else {
		fmt.Println("🆕 Creating new README.md")
	}

	return os.WriteFile(fileName, []byte(content), 0644)
}