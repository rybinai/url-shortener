package storage

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func generateShortCode() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.IntN(len(letters))]
	}
	return string(code)
}

func (s *PostgresStorage) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	shortCode := generateShortCode()
	query := `INSERT INTO urlshortener (short_code, original_url) VALUES ($1, $2)`
	_, err := s.db.Exec(ctx, query, shortCode, originalURL)
	if err != nil {
		return "", fmt.Errorf("failed to create short url: %w", err)
	}
	return shortCode, nil
}

func (s *PostgresStorage) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urlshortener WHERE short_code = $1`
	err := s.db.QueryRow(ctx, query, shortCode).Scan(&originalURL)
	if err != nil {
		return "", fmt.Errorf("url not found: %w", err)
	}
	return originalURL, nil
}
