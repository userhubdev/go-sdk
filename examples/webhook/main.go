package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/userhubdev/go-sdk/eventsv1"
	"github.com/userhubdev/go-sdk/webhook"
)

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	signingSecret := os.Getenv("SIGNING_SECRET")
	if signingSecret == "" {
		return fmt.Errorf("SIGNING_SECRET required")
	}

	wh := webhook.New(signingSecret).
		OnEvent(func(ctx context.Context, event *eventsv1.Event) error {
			fmt.Println("Event:", event.Type)

			switch event.Type {
			case "organizations.changed":
				organization := event.OrganizationsChanged.Organization
				fmt.Println(" - Organization:", organization.Id, organization.DisplayName)
			case "users.changed":
				user := event.UsersChanged.User
				fmt.Println(" - User:", user.Id, user.DisplayName)
			}

			return nil
		})

	return http.ListenAndServe(":"+port, wh)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
