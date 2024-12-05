package main

import (
	"context"
	"log"
	"os"

	"github.com/shoutboxnet/shoutbox-go/shoutbox"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("SHOUTBOX_API_KEY")
	if apiKey == "" {
		log.Fatal("SHOUTBOX_API_KEY environment variable is not set")
	}

	// Create a new client
	client := shoutbox.NewClient(apiKey)

	// Create an email request
	req := &shoutbox.EmailRequest{
		From:    os.Getenv("SHOUTBOX_FROM"),
		To:      os.Getenv("SHOUTBOX_TO"),
		Subject: "Hello from Shoutbox REST API",
		HTML:    "<h1>Hello!</h1><p>This email was sent using the Shoutbox REST API client.</p>",
		Name:    "Shoutbox Test",
		ReplyTo: os.Getenv("SHOUTBOX_FROM"),
		Headers: map[string]string{
			"X-Application": "Shoutbox Example",
		},
	}

	// Send the email
	err := client.SendEmail(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
