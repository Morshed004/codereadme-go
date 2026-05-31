package parser

import (
	"strings"

	"codereadme/internal/scanner"
)

func BuildPrompt(files []scanner.FileData) string {
	var builder strings.Builder

	builder.WriteString(`Analyze this codebase and generate a professional README.md.

Include:
- Project title
- Description
- Features
- Tech stack
- Installation
- Usage
- Folder structure
- Future improvements

Codebase:
`)

	for _, file := range files {
		builder.WriteString("\n--- FILE: ")
		builder.WriteString(file.Path)
		builder.WriteString(" ---\n")
		builder.WriteString(file.Content)
		builder.WriteString("\n")
	}

	return builder.String()
}
