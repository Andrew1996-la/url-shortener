package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Andrew1996-la/url-shortenerr/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	operation := "storage.sqlite.NewStorage"

	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
	    id INTEGER PRIMARY KEY,
	    alias TEXT NOT NULL UNIQUE,
	    url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	operation := "storage.sqlite.SaveURL"

	res, err := s.db.Exec(`
		INSERT INTO urls (alias, url)
		VALUES (?, ?)`,
		alias, urlToSave,
	)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	operation := "storage.sqlite.GetURL"

	var url string

	err := s.db.QueryRow(`
		SELECT url FROM urls WHERE alias = ?`,
		alias,
	).Scan(&url)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrUrlNotFound
		}
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	return url, nil
}

func (s *Storage) DeleteURL(alias string) error {
	operation := "storage.sqlite.DeleteURL"

	_, err := s.db.Exec(`DELETE FROM urls WHERE alias = ?`, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
