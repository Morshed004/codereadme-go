package main

import (
	"fmt"
	"log"
	"os"

	"codereadme/internal/generator"
	"codereadme/internal/scanner"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil{
		fmt.Println("Can not load the env")
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: codereadme <command>")
		fmt.Println("Commands: generate")
		return
	}

	command := os.Args[1]

	switch command {

	case "generate":
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("API_KEY not set")
			return
		}

		files, err := scanner.Scan(".")
		if err != nil {
			panic(err)
		}

		err = generator.Generate(files, apiKey)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("✅ README.md generated")

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available: generate")
	}
}