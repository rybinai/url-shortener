package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rybinai/url-shortener/internal/server"
	"github.com/rybinai/url-shortener/internal/storage"
	"github.com/rybinai/url-shortener/shortener"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, "postgresql://user:password@localhost:5432/urlshortener")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal("Database ping failed:", err)
	}
	log.Println("Connected to database")

	urlStorage := storage.NewPostgresStorage(db)
	grpcServer := server.NewGRPCServer(urlStorage)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	shortener.RegisterUrlShortenerServer(s, grpcServer)

	log.Printf("gRPC server listening on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
