package services

import (
	"encoding/json"
	"fmt"
	"news_reader/models"
	"news_reader/utils"
	"os"
)

// FetchNews fetches news based on the selected category
func FetchNews(category string) ([]models.NewsItem, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("NEWS_API_KEY is not set in environment variables")
	}

	// Build the API URL for the selected category
	apiURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?category=%s&country=us&apiKey=%s", category, apiKey)
	response, err := utils.MakeAPIRequest(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %v", err)
	}

	var result struct {
		Articles []models.NewsItem `json:"articles"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing news data: %v", err)
	}

	return result.Articles, nil
}
