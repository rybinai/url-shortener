package server

import (
	"context"
	"time"

	"github.com/rybinai/url-shortener/internal/storage"
	"github.com/rybinai/url-shortener/shortener"
)

type GRPCServer struct {
	shortener.UnimplementedUrlShortenerServer
	storage *storage.PostgresStorage
}

func NewGRPCServer(storage *storage.PostgresStorage) *GRPCServer {
	return &GRPCServer{storage: storage}
}

func (s *GRPCServer) CreateShortUrl(ctx context.Context, req *shortener.CreateRequest) (*shortener.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	shortCode, err := s.storage.CreateShortURL(ctx, req.OriginalUrl)
	if err != nil {
		return nil, err
	}
	return &shortener.CreateResponse{
		ShortCode: shortCode,
	}, nil
}

func (s *GRPCServer) GetOriginalUrl(ctx context.Context, req *shortener.GetOriginalRequest) (*shortener.GetOriginalResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	originalURL, err := s.storage.GetOriginalURL(ctx, req.ShortCode)
	if err != nil {
		return nil, err
	}
	return &shortener.GetOriginalResponse{
		OriginalUrl: originalURL,
	}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedUrlShortenerServer() {}
