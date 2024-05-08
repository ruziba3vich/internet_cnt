package storage

import (
	"database/sql"

	"github.com/ruziba3vich/shortURL/internal/models"
)

func GenerateShortURL(url string, db *sql.DB) (*models.URL, error) {
	generator := GetURLShortener()
	newURL := generator(url)

	_, err := GetOriginalURL(newURL, db)
	if err != nil {
		if err != sql.ErrNoRows {
			query := `
            INSERT INTO urls (
                original,
                shortened
            )
            VALUES (
                $1, $2
            )
            RETURNING shortened;
        `
			var result models.URL
			row := db.QueryRow(query, url, newURL)
			if err := row.Scan(&result.Url); err != nil {
				return nil, err
			}

			return &result, nil
		}
		return nil, err
	}

	return &models.URL{
		Url: newURL,
	}, nil
}
