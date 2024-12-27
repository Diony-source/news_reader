package services

import (
	"encoding/json"
	"fmt"
	"news_reader/models"
	"news_reader/utils"
)

// FetchNews fetches the latest news from the API and determines the most popular news
func FetchNews() ([]models.NewsItem, error) {
	apiURL := "https://newsapi.example.com/latest" // Replace with actual API URL
	response, err := utils.MakeAPIRequest(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %v", err)
	}

	var news []models.NewsItem
	err = json.Unmarshal(response, &news)
	if err != nil {
		return nil, fmt.Errorf("error parsing news data: %v", err)
	}

	// Determine the most popular news
	markMostPopular(news)

	return news, nil
}

// markMostPopular marks the most popular news item
func markMostPopular(news []models.NewsItem) {
	if len(news) == 0 {
		return
	}

	mostPopularIndex := 0
	for i, item := range news {
		if item.Popularity > news[mostPopularIndex].Popularity {
			mostPopularIndex = i
		}
	}
	// Mark the most popular news
	news[mostPopularIndex].IsPopular = true
}
