package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rybinai/url-shortener/shortener"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddress = "localhost:8080"
)

var rootCmd = &cobra.Command{
	Use:   "url-shortener",
	Short: "url-shortener CLI application",
}

var createCmd = &cobra.Command{
	Use:   "create [URL]",
	Short: "create a url",
	Args:  cobra.ExactArgs(1),
	Run:   runCreate,
}

var getCmd = &cobra.Command{
	Use:   "get [shortCode]",
	Short: "get original url by short code",
	Args:  cobra.ExactArgs(1),
	Run:   runGet,
}

func main() {
	rootCmd.AddCommand(createCmd, getCmd)
	rootCmd.Execute()
}

func createClient() shortener.UrlShortenerClient {
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not connect to server: %v\n", err)
		os.Exit(1)
	}
	return shortener.NewUrlShortenerClient(conn)
}

func runCreate(_ *cobra.Command, args []string) {
	url := args[0]
	client := createClient()

	r, err := client.CreateShortUrl(context.Background(), &shortener.CreateShortUrlRequest{
		OriginalUrl: url,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create short URL: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("short", r.ShortCode)
}

func runGet(_ *cobra.Command, args []string) {
	shortCode := args[0]
	client := createClient()
	r2, err := client.GetOriginalUrl(context.Background(), &shortener.GetOriginalUrlRequest{
		ShortCode: shortCode,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not get original URL: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("original url", r2.OriginalUrl)
}
