package storage

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func generateShortCode(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	encodedURL := base64.URLEncoding.EncodeToString(hash[:])
	return encodedURL[:6]
}

func (s *PostgresStorage) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	var existingCode string
	selectQuery, selectArgs, _ := sq.
		Select("short_code").
		From("urlshortener").
		Where(sq.Eq{"original_url": originalURL}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	err := s.db.QueryRow(ctx, selectQuery, selectArgs...).Scan(&existingCode)
	if err == nil {
		return existingCode, nil
	}
	shortCode := generateShortCode(originalURL)
	insertQuery, insertArgs, err := sq.
		Insert("urlshortener").
		Columns("short_code", "original_url").
		Values(shortCode, originalURL).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build query: %w", err)
	}
	_, err = s.db.Exec(ctx, insertQuery, insertArgs...)
	if err != nil {
		return "", fmt.Errorf("failed to create short url: %w", err)
	}
	return shortCode, nil
}

func (s *PostgresStorage) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	var originalURL string
	query, args, err := sq.
		Select("original_url").
		From("urlshortener").
		Where(sq.Eq{"short_code": shortCode}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build query: %w", err)
	}
	err = s.db.QueryRow(ctx, query, args...).Scan(&originalURL)
	if err != nil {
		return "", fmt.Errorf("failed to get original url: %w", err)
	}
	return originalURL, nil
}
