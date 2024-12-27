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
		fmt.Println("1. Get News by Category")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handleCategorySelection()
		case 2:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// handleCategorySelection handles the user's category selection
func handleCategorySelection() {
	fmt.Println("\nSelect a category:")
	fmt.Println("1. Business")
	fmt.Println("2. Entertainment")
	fmt.Println("3. Health")
	fmt.Println("4. Science")
	fmt.Println("5. Sports")
	fmt.Println("6. Technology")
	fmt.Print("Enter your choice: ")

	var categoryChoice int
	fmt.Scanln(&categoryChoice)

	categories := map[int]string{
		1: "business",
		2: "entertainment",
		3: "health",
		4: "science",
		5: "sports",
		6: "technology",
	}

	category, exists := categories[categoryChoice]
	if !exists {
		fmt.Println("Invalid category. Please try again.")
		return
	}

	news, err := services.FetchNews(category)
	if err != nil {
		log.Printf("Error fetching news: %v\n", err)
		return
	}

	fmt.Printf("\nTop news in %s:\n", category)
	for i, item := range news {
		fmt.Printf("%d. %s\n   %s\n   %s\n", i+1, item.Title, item.Summary, item.URL)
	}
}
