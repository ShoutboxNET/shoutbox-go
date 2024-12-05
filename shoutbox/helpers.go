package shoutbox

import (
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

// NewAttachmentFromFile creates a new attachment from a file
func NewAttachmentFromFile(filePath string) (Attachment, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Attachment{}, fmt.Errorf("error reading file: %w", err)
	}

	// Detect content type
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return Attachment{
		Filename:    filepath.Base(filePath),
		Content:     content,
		ContentType: contentType,
	}, nil
}

// NewAttachmentFromReader creates a new attachment from an io.Reader
func NewAttachmentFromReader(reader io.Reader, filename string) (Attachment, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return Attachment{}, fmt.Errorf("error reading content: %w", err)
	}

	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return Attachment{
		Filename:    filename,
		Content:     content,
		ContentType: contentType,
	}, nil
}

// ValidateEmail validates an email address format
func ValidateEmail(email string) error {
	// Simple validation for demonstration
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email address: %s", email)
	}
	return nil
}

// ValidateEmailList validates a list of email addresses
func ValidateEmailList(emails []string) error {
	for _, email := range emails {
		if err := ValidateEmail(email); err != nil {
			return err
		}
	}
	return nil
}
