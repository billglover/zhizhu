package storage

import (
	"database/sql"
)

// Storage provides access to a persistant data store.
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new Storage instance backed by the provided sql.DB.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}
