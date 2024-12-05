package shoutbox

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"strings"
)

// SMTPClient represents a Shoutbox SMTP client
type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Auth     smtp.Auth
}

// NewSMTPClient creates a new Shoutbox SMTP client
func NewSMTPClient(apiKey string) *SMTPClient {
	host := "mail.shoutbox.net"
	return &SMTPClient{
		Host:     host,
		Port:     587,
		Username: "shoutbox",
		Password: apiKey,
		Auth:     smtp.PlainAuth("", "shoutbox", apiKey, host),
	}
}

// Attachment represents an email attachment
type Attachment struct {
	Filename    string
	Content     []byte
	ContentType string
}

// EmailMessage represents an email message for SMTP
type EmailMessage struct {
	From        string
	To          []string
	Subject     string
	HTML        string
	Name        string
	ReplyTo     string
	Attachments []Attachment
	Headers     map[string]string
}

// SendEmail sends an email using SMTP
func (c *SMTPClient) SendEmail(msg *EmailMessage) error {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	// Add headers
	headers := textproto.MIMEHeader{}
	headers.Set("From", formatAddress(msg.From, msg.Name))
	headers.Set("To", strings.Join(msg.To, ", "))
	headers.Set("Subject", msg.Subject)
	headers.Set("MIME-Version", "1.0")
	headers.Set("Content-Type", fmt.Sprintf("multipart/mixed; boundary=%s", writer.Boundary()))

	if msg.ReplyTo != "" {
		headers.Set("Reply-To", msg.ReplyTo)
	}

	// Add custom headers
	for key, value := range msg.Headers {
		headers.Set(key, value)
	}

	// Write headers
	for key, values := range headers {
		for _, value := range values {
			fmt.Fprintf(buffer, "%s: %s\r\n", key, value)
		}
	}
	buffer.WriteString("\r\n")

	// Add HTML part
	htmlPart, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":              {"text/html; charset=UTF-8"},
		"Content-Transfer-Encoding": {"quoted-printable"},
	})
	if err != nil {
		return fmt.Errorf("error creating HTML part: %w", err)
	}
	htmlPart.Write([]byte(msg.HTML))

	// Add attachments
	for _, attachment := range msg.Attachments {
		part, err := writer.CreatePart(textproto.MIMEHeader{
			"Content-Type":              {fmt.Sprintf("%s; name=%q", attachment.ContentType, attachment.Filename)},
			"Content-Disposition":       {fmt.Sprintf("attachment; filename=%q", attachment.Filename)},
			"Content-Transfer-Encoding": {"base64"},
		})
		if err != nil {
			return fmt.Errorf("error creating attachment part: %w", err)
		}

		encoder := base64.NewEncoder(base64.StdEncoding, part)
		encoder.Write(attachment.Content)
		encoder.Close()
	}

	writer.Close()

	// Send email
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		c.Auth,
		msg.From,
		msg.To,
		buffer.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func formatAddress(email, name string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}
