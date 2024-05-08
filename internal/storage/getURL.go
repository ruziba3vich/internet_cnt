package storage

import (
	"database/sql"
	"errors"
)

func GetOriginalURL(shortURL string, db *sql.DB) (string, error) {
	query := "SELECT original FROM Urls WHERE shortened = $1;"
	var result string
	err := db.QueryRow(query, shortURL).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no URL found")
		}
		return "", err
	}

	return result, nil
}
