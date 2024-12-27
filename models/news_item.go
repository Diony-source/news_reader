package models

// NewsItem represents a single news article
type NewsItem struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Popularity int    `json:"popularity"`
	IsPopular bool   // Marks the most popular news
}
