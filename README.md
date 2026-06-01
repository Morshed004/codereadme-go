# CodeREADME

**CodeREADME** is a CLI tool written in Go that automatically generates professional `README.md` files for your projects. It scans your local codebase, compiles the relevant source files, and leverages the power of Large Language Models (LLMs) via the OpenRouter API to document your project structure, features, and usage.

## Features

- **Automated Codebase Scanning**: Recursively walks through your directory to collect source code.
- **Smart Ignoring**: Supports a `.codereadmeignore` file to exclude binaries, dependencies, and sensitive files from being sent to the LLM.
- **LLM Integration**: Uses OpenRouter (specifically configured for `google/gemma-4-31b-it:free`) to generate high-quality, context-aware documentation.
- **Safety First**: Automatically creates a `README.old.md` backup if an existing README is detected before updating.
- **Clean Architecture**: Organized into internal packages for scanning, parsing, generation, and writing.

## 🛠 Tech Stack

- **Language**: [Go](https://go.dev/) (v1.26.3)
- **API**: [OpenRouter](https://openrouter.ai/)
- **Environment Management**: `godotenv`

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Morshed004/codereadme-go.git
   cd codereadme-go
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure Environment**:
   Create a `.env` file in the root directory and add your OpenRouter API key:
   ```env
   OPENROUTER_API_KEY=your_api_key_here
   ```

4. **Build the application**:
   ```bash
   go build -o codereadme main.go
   ```

## Usage

### 1. Set up Ignore List
Create a `.codereadmeignore` file to prevent unnecessary files (like `.git`, `node_modules`, or `.env`) from being uploaded to the API:
```text
.git
.env
mod.sum
*.exe
```

### 2. Generate the README
Run the following command in the root of the project you wish to document:
```bash
./codereadme generate
```

**What happens next?**
- The tool scans your directory.
- It constructs a prompt containing your code.
- It calls the OpenRouter API.
- It saves the result as `README.md` (and backups the old one if it exists).

## Folder Structure

```text
.
├── internal/
│   ├── generator/    # Orchestrates the flow between scanner, parser, and writer
│   │   └── readme.go
│   ├── openrouter/    # Handles HTTP communication with the OpenRouter API
│   │   └── client.go
│   ├── parser/        # Formats the scanned files into a structured LLM prompt
│   │   └── parser.go
│   ├── scanner/       # Logic for walking the file system and handling ignore patterns
│   │   ├── ignore.go
│   │   ├── matcher.go
│   │   └── scanner.go
│   └── writer/        # Handles file writing and backup logic
│       └── file.go
├── .env               # Environment variables (API Keys)
├── .codereadmeignore  # Files to exclude from scanning
├── go.mod             # Go module definition
└── main.go            # CLI Entry point
```