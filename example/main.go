package main

import (
	"context"
	"log"
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

func login() (string, error) {
	return "", nil
}

func createSource() {}

func main() {
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
