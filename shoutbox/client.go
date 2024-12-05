package shoutbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client represents a Shoutbox API client
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// EmailRequest represents an email request to the Shoutbox API
type EmailRequest struct {
	From    string   `json:"from"`
	To      string   `json:"to"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
	Name    string   `json:"name,omitempty"`
	ReplyTo string   `json:"reply_to,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// NewClient creates a new Shoutbox API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		baseURL:    "https://api.shoutbox.net",
	}
}

// SendEmail sends an email using the Shoutbox API
func (c *Client) SendEmail(ctx context.Context, req *EmailRequest) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/send",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("error response with status %d", resp.StatusCode)
		}
		return fmt.Errorf("api error: %s", errResp.Error)
	}

	return nil
}
