package repository

import (
	"context"

	"github.com/MitSonani/go-url-shortner/internal/db"
)

type URLRepository struct {
}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (r *URLRepository) FindByURL(originalURL string) (string, error) {
	query := `SELECT short_code FROM urls WHERE original_url = $1`

	var code string
	err := db.Conn.QueryRow(context.Background(), query, originalURL).Scan(&code)
	if err != nil {
		return "", err
	}

	return code, nil
}

// ➕ Create new URL
func (r *URLRepository) Create(shortCode string, originalURL string) (string, error) {
	query := `
	INSERT INTO urls (original_url, short_code)
	VALUES ($1, $2)
	RETURNING short_code;
	`

	var savedCode string
	err := db.Conn.QueryRow(context.Background(), query, originalURL, shortCode).Scan(&savedCode)

	return savedCode, err
}

// 🔁 Get original URL
func (r *URLRepository) Get(shortCode string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_code = $1`

	var url string
	err := db.Conn.QueryRow(context.Background(), query, shortCode).Scan(&url)

	return url, err
}
