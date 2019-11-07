package database

import (
	"database/sql"

	// Postgresql driver import
	_ "github.com/lib/pq"
)

// NewDB returns a new database.
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
