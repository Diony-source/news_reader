package main

import (
	"log"
	"news_reader/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Starting Mini News Reader...")
	handlers.StartCLI()
}
