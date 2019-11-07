package storage

import (
	"fmt"

	"github.com/billglover/zhizhu/model"
)

// CreateFeed persists a new feed to the underlying data store.
func (s *Storage) CreateFeed(f *model.FeedInfo) error {
	stmt, err := s.db.Prepare(`INSERT INTO feeds
	(title, link, description, language, copyright)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(f.Title, f.Link, f.Description, f.Language, f.Copyright).Scan(&f.ID)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
