package main

import (
	"log"
	"news_reader/handlers"
)

func main() {
	log.Println("Starting Mini News Reader...")
	handlers.StartCLI()
}