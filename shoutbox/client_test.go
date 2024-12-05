package shoutbox

import (
	"context"
	"os"
	"testing"
)

func TestClient_SendEmail(t *testing.T) {
	apiKey := os.Getenv("SHOUTBOX_API_KEY")
	if apiKey == "" {
		t.Skip("SHOUTBOX_API_KEY not set")
	}

	from := os.Getenv("SHOUTBOX_FROM")
	if from == "" {
		t.Skip("SHOUTBOX_FROM not set")
	}

	to := os.Getenv("SHOUTBOX_TO")
	if to == "" {
		t.Skip("SHOUTBOX_TO not set")
	}

	client := NewClient(apiKey)

	tests := []struct {
		name    string
		req     *EmailRequest
		wantErr bool
	}{
		{
			name: "basic email",
			req: &EmailRequest{
				From:    from,
				To:      to,
				Subject: "Test Email",
				HTML:    "<h1>Test</h1><p>This is a test email from the Shoutbox Go client.</p>",
			},
			wantErr: false,
		},
		{
			name: "email with name and reply-to",
			req: &EmailRequest{
				From:    from,
				To:      to,
				Subject: "Test Email with Name",
				HTML:    "<h1>Test</h1><p>This is a test email with sender name and reply-to.</p>",
				Name:    "Test Sender",
				ReplyTo: from,
			},
			wantErr: false,
		},
		{
			name: "email with custom headers",
			req: &EmailRequest{
				From:    from,
				To:      to,
				Subject: "Test Email with Headers",
				HTML:    "<h1>Test</h1><p>This is a test email with custom headers.</p>",
				Headers: map[string]string{
					"X-Test-Header": "test-value",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.SendEmail(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
