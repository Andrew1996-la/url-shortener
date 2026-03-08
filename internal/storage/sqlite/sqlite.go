package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	operation := "storage.sqlite.NewStorage"

	db, err := sql.Open("sqlite", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	db.Prepare(`
	CREATE TABLE IF NOT EXISTS urls (
	    id INTEGER PRIMARY KEY,
	    alias TEXT NOT NULL,
	    url TEXT NOT NULL,
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	return &Storage{db: db}, nil
}
