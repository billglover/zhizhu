package model

// FeedInfo represents information about a feed.
type FeedInfo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
	Language    string `json:"language,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Items       Items  `json:"items,omitempty"`
}
