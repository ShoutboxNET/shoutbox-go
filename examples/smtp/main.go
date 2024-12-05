package main

import (
	"log"
	"os"
	"strings"

	"github.com/shoutboxnet/shoutbox-go/shoutbox"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("SHOUTBOX_API_KEY")
	if apiKey == "" {
		log.Fatal("SHOUTBOX_API_KEY environment variable is not set")
	}

	// Create a new SMTP client
	client := shoutbox.NewSMTPClient(apiKey)

	// Get recipient from environment
	to := os.Getenv("SHOUTBOX_TO")
	if to == "" {
		log.Fatal("SHOUTBOX_TO environment variable is not set")
	}

	// Create a test file for attachment
	testFile := "test.txt"
	err := os.WriteFile(testFile, []byte("This is a test attachment."), 0644)
	if err != nil {
		log.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Create attachment from file
	attachment, err := shoutbox.NewAttachmentFromFile(testFile)
	if err != nil {
		log.Fatalf("Failed to create attachment: %v", err)
	}

	// Create an email message with attachment
	msg := &shoutbox.EmailMessage{
		From:    os.Getenv("SHOUTBOX_FROM"),
		To:      []string{to},
		Subject: "Hello from Shoutbox SMTP",
		HTML: strings.Join([]string{
			"<h1>Hello!</h1>",
			"<p>This email was sent using the Shoutbox SMTP client.</p>",
			"<p>It includes a text file attachment.</p>",
		}, ""),
		Name:    "Shoutbox Test",
		ReplyTo: os.Getenv("SHOUTBOX_FROM"),
		Headers: map[string]string{
			"X-Application": "Shoutbox SMTP Example",
		},
		Attachments: []shoutbox.Attachment{attachment},
	}

	// Send the email
	err = client.SendEmail(msg)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")

	// Example of sending a basic email without attachments
	basicMsg := &shoutbox.EmailMessage{
		From:    os.Getenv("SHOUTBOX_FROM"),
		To:      []string{to},
		Subject: "Basic SMTP Test",
		HTML:    "<h1>Basic Test</h1><p>This is a basic email without attachments.</p>",
	}

	err = client.SendEmail(basicMsg)
	if err != nil {
		log.Fatalf("Failed to send basic email: %v", err)
	}

	log.Println("Basic email sent successfully!")
}
