package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rybinai/url-shortener/shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	
	client := shortener.NewUrlShortenerClient(conn)
	command := os.Args[1]

	switch command {
	case "create":
		url := os.Args[2]
		r, err := client.CreateShortUrl(context.Background(), &shortener.CreateRequest{
			OriginalUrl: url,
		})
		if err != nil {
			log.Fatalf("could not create shortener: %v", err)
		}
		fmt.Println("short", r.ShortCode)
	case "get":
		shortCode := os.Args[2]
		r2, err := client.GetOriginalUrl(context.Background(), &shortener.GetOriginalRequest{
			ShortCode: shortCode,
		})
		if err != nil {
			log.Fatalf("could not get original url: %v", err)
		}
		fmt.Println("original url", r2.OriginalUrl)
	}
}
