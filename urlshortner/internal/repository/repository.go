package repository

import (
	"database/sql"
	"time"
)

// expiry_time = null, then never expires
func CreateTable(db *sql.DB) error {
	sql := `
		CREATE TABLE IF NOT EXISTS urls (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    short_url TEXT NOT NULL UNIQUE,
		    original_url TEXT NOT NULL,
			expiry_time DATETIME
		)
	`
	_, err := db.Exec(sql)
	return err
}

// if duplicate short_url, then ignore
func StoreURL(db *sql.DB, shortURL, originalURL string, expiryTime time.Time) error {
	// format to sqlite datetime
	expiryTimeFormatted := expiryTime.Format("2006-01-02 15:04:05")
	sql := `
		INSERT INTO urls (short_url, original_url, expiry_time)
		VALUES (?, ?, ?)
		ON CONFLICT(short_url) DO NOTHING
	`
	_, err := db.Exec(sql, shortURL, originalURL, expiryTimeFormatted)
	return err
}

func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	sql := `
		SELECT original_url
		FROM urls
		WHERE short_url = ? LIMIT 1
	`
	err := db.QueryRow(sql, shortURL).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil

}

func SetTTL(db *sql.DB, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	sql := `
		DELETE FROM urls
		WHERE expiry_time < DATETIME('now')
	`
	for range ticker.C {
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}
