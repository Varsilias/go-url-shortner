package db

import (
	"database/sql"
	"errors"
)

// CreateTable ensures the urls table exists.
func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			short_url TEXT NOT NULL,
			original_url TEXT NOT NULL
		);
	`

	_, err := db.Exec(query)
	return err
}

// StoreURL inserts new short URL and original URL into the database
func StoreURL(db *sql.DB, shortURL string, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

// GetOriginalURL fetches the original URL by the short URL
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ?`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}

func GetOriginalURLExists(db *sql.DB, originalURL string) (bool, error) {
	var o string
	query := `SELECT original_url FROM urls WHERE original_url = ?`
	err := db.QueryRow(query, originalURL).Scan(&o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func GetShortURLExists(db *sql.DB, shortURL string) (bool, error) {
	var s string
	query := `SELECT short_url FROM urls WHERE short_url = ?`
	err := db.QueryRow(query, shortURL).Scan(&s)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

// TODO: fix duplicate original url - done
// handle possible collision
