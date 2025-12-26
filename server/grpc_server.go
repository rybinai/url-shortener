package server

import (
	"context"
	"time"

	"github.com/rybinai/url-shortener/internal/storage"
	"github.com/rybinai/url-shortener/shortener"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	shortener.UnimplementedUrlShortenerServer
	storage *storage.PostgresStorage
}

func NewGRPCServer(storage *storage.PostgresStorage) *GRPCServer {
	return &GRPCServer{storage: storage}
}

func (s *GRPCServer) CreateShortUrl(ctx context.Context, req *shortener.CreateShortUrlRequest) (*shortener.CreateShortUrlResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	shortCode, err := s.storage.CreateShortURL(ctx, req.OriginalUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create short url: %v", err)
	}
	return &shortener.CreateShortUrlResponse{
		ShortCode: shortCode,
	}, nil
}

func (s *GRPCServer) GetOriginalUrl(ctx context.Context, req *shortener.GetOriginalUrlRequest) (*shortener.GetOriginalUrlResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	originalURL, err := s.storage.GetOriginalURL(ctx, req.ShortCode)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get original url: %v", err)
	}
	return &shortener.GetOriginalUrlResponse{
		OriginalUrl: originalURL,
	}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedUrlShortenerServer() {}
