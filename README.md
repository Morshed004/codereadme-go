# CodeReadMe Generator

## Description
The **CodeReadMe Generator** is a tool that automates the process of generating a professional README.md file for codebases. By analyzing the structure and content of the existing code files, the generator creates a well-organized README that includes essential sections such as project title, description, features, tech stack, installation instructions, usage information, folder structure, and suggested future improvements.

## Features
- **Automated README Generation**: Generates a README.md file based on the contents of the codebase.
- **Configurable Ignoring Patterns**: Supports ignoring certain files or directories based on user-defined patterns in the `.codereadmeignore` file.
- **OpenRouter AI Integration**: Utilizes OpenRouter API to dynamically generate text based on the analyzed code.
- **Backup Existing README**: Automatically creates a backup of any existing README.md file before updates.

## Tech Stack
- **Go**: The primary programming language used for the project.
- **OpenRouter API**: For generating content using the OpenAI GPT model.
- **Go modules**: For dependency management.

## Installation
Make sure to have Go installed on your machine. Clone the repository and install the necessary dependencies:

```bash
git clone github.com/Morshed004/codereadme-go
cd codereadme-go
go mod tidy
```

### Environment Setup
Create an `.env` file in the root of the project with the following content:

```
OPENROUTER_API_KEY=<your_api_key_here>
```

Replace `<your_api_key_here>` with your actual API key for OpenRouter.

## Usage
Run the application to generate the `README.md` file:
```bash
go run main.go generate
```
The tool will scan the current directory (and subdirectories) for code files and generate a `README.md` file based on its content.

## Folder Structure
```
/Project Root
├── .env                     # Environment variables
├── go.mod                   # Module dependencies
├── go.sum                   # Module checksums
├── main.go                  # Application entry point
├── internal/
│   ├── generator/           # Logic for generating README
│   ├── openrouter/          # Client for OpenRouter API
│   ├── parser/              # Functions to build prompts for the API
│   ├── scanner/             # Scanning logic to read code files
│   └── writer/              # Functions to save README.md
└── .codereadmeignore        # Files and directories to ignore
```