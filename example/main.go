package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/getaudited/audited-go"
)

func run() error {
	client, err := audited.NewClient(audited.Config{
		BaseAPI:  "http://localhost:8080",
		APIToken: "some-token",
	})
	if err != nil {
		return err
	}

	err = client.CreateEvent(context.Background(), audited.Event{})
	if err != nil {
		return err
	}

	return nil
}

type AdminClient struct {
	httpClient *http.Client
	baseAPI    string
}

func login(ctx context.Context, email, password string) (string, error) {
	return "", nil
}

func createSource(ctx context.Context, name string) error {}

func main() {
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
