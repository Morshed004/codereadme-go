package generator

import (
	"codereadme/internal/llm"
	"codereadme/internal/parser"
	"codereadme/internal/scanner"
	"codereadme/internal/writer"
)

func Generate(files []scanner.FileData, apiKey string) error {
	prompt := parser.BuildPrompt(files)

	readme, err := llm.Generate(prompt, apiKey)
	if err != nil {
		return err
	}

	return writer.SaveREADME(readme)
}