package model

import "time"

// Item represents information about an article.
type Item struct {
	ID            string    `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	Link          string    `json:"link,omitempty"`
	Author        string    `json:"author,omitempty"`
	GUID          string    `json:"guid,omitempty"`
	Content       string    `json:"content,omitempty"`
	Hash          string    `json:"hash,omitempty"`
	PublishedDate time.Time `json:"published_date,omitempty"`
}

// Items is a list of Items.
type Items []*Item
