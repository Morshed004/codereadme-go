package parser

import (
	"strings"

	"codereadme/internal/scanner"
)

func BuildPrompt(files []scanner.FileData) string {
	var builder strings.Builder

	builder.WriteString(`You are analyzing a real software project.

Generate a natural, human-written README.md based only on the provided codebase.

Rules:

- Never expose environment variable values, secrets, tokens, API keys, credentials, or private configuration values.
- If environment variables are required, document only their names with placeholder examples.
- Do not copy raw source code into the README unless a very short example is necessary.
- Do not mention internal implementation details unless they are useful for users.
- Do not use emojis.
- Do not use exaggerated marketing language.
- Do not write like an AI assistant.
- Write like an experienced developer documenting their own project.
- Be concise, practical, and clear.
- Do not invent features that do not exist in the codebase.
- Infer project purpose from actual files only.
- If information is missing, omit it instead of guessing.

Write a professional README with these sections when relevant:

# Project Title
A concise title inferred from the codebase.

## Overview
A short explanation of what the project does.

## Features
Only real implemented features.

## Tech Stack
Frameworks, languages, libraries, and services actually used.

## Installation
Real setup instructions.

## Usage
How to run and use the project.

## Project Structure
Brief explanation of key directories.

## Configuration
Document environment variable names only if required.

## Notes
Any relevant implementation details.

## Future Improvements
Only if clearly appropriate.

Write naturally, like documentation written by a human developer for GitHub.

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