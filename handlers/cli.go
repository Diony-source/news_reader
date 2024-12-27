package handlers

import (
	"fmt"
	"log"
	"news_reader/services"
)

// StartCLI initializes the command-line interface for the Mini News Reader
func StartCLI() {
	for {
		fmt.Println("\nMini News Reader")
		fmt.Println("1. Get Latest News")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Fetch and display news
			news, err := services.FetchNews()
			if err != nil {
				log.Printf("Error fetching news: %v\n", err)
				continue
			}
			fmt.Println("\nLatest News:")
			for i, item := range news {
				popularTag := ""
				if item.IsPopular {
					popularTag = "(Most Read)"
				}
				fmt.Printf("%d. %s %s\n", i+1, item.Title, popularTag)
				fmt.Printf("   %s\n", item.Summary)
			}

		case 2:
			// Exit the program
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
