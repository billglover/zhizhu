package storage

import (
	"fmt"

	"github.com/billglover/zhizhu/model"
)

// CreateItem persists a new feed to the underlying data store.
func (s *Storage) CreateItem(feedID string, i *model.Item) error {
	stmt, err := s.db.Prepare(`INSERT INTO items
	(feed_id, title, link, hash, description, guid, content, published_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(feedID, i.Title, i.Link, i.Hash, i.Description, i.GUID, i.Content, i.PublishedDate).Scan(&i.ID)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
