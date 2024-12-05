// Example of using both REST API and SMTP clients
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

	// Example using REST API client
	restClient := shoutbox.NewClient(apiKey)
	restReq := &shoutbox.EmailRequest{
		From:    os.Getenv("SHOUTBOX_FROM"),
		To:      os.Getenv("SHOUTBOX_TO"),
		Subject: "Test from REST API",
		HTML:    "<h1>REST API Test</h1><p>This email was sent using the REST API client.</p>",
	}

	err := restClient.SendEmail(context.Background(), restReq)
	if err != nil {
		log.Printf("REST API error: %v", err)
	} else {
		log.Println("REST API email sent successfully!")
	}

	// Example using SMTP client
	smtpClient := shoutbox.NewSMTPClient(apiKey)
	smtpMsg := &shoutbox.EmailMessage{
		From:    os.Getenv("SHOUTBOX_FROM"),
		To:      []string{os.Getenv("SHOUTBOX_TO")},
		Subject: "Test from SMTP",
		HTML:    "<h1>SMTP Test</h1><p>This email was sent using the SMTP client.</p>",
	}

	err = smtpClient.SendEmail(smtpMsg)
	if err != nil {
		log.Printf("SMTP error: %v", err)
	} else {
		log.Println("SMTP email sent successfully!")
	}
}
