package models

// NewsItem represents a single news article
type NewsItem struct {
	Title   string `json:"title"`
	Summary string `json:"description"`
	URL     string `json:"url"`
}
