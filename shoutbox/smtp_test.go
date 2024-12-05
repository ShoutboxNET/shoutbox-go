package shoutbox

import (
	"os"
	"testing"
)

func TestSMTPClient_SendEmail(t *testing.T) {
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

	client := NewSMTPClient(apiKey)

	tests := []struct {
		name    string
		msg     *EmailMessage
		wantErr bool
	}{
		{
			name: "basic email",
			msg: &EmailMessage{
				From:    from,
				To:      []string{to},
				Subject: "SMTP Test Email",
				HTML:    "<h1>Test</h1><p>This is a test email from the Shoutbox SMTP client.</p>",
			},
			wantErr: false,
		},
		{
			name: "email with name and reply-to",
			msg: &EmailMessage{
				From:    from,
				To:      []string{to},
				Subject: "SMTP Test Email with Name",
				HTML:    "<h1>Test</h1><p>This is a test email with sender name and reply-to.</p>",
				Name:    "Test Sender",
				ReplyTo: from,
			},
			wantErr: false,
		},
		{
			name: "email with custom headers",
			msg: &EmailMessage{
				From:    from,
				To:      []string{to},
				Subject: "SMTP Test Email with Headers",
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
			err := client.SendEmail(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "invalid email",
			email:   "invalid-email",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEmailList(t *testing.T) {
	tests := []struct {
		name    string
		emails  []string
		wantErr bool
	}{
		{
			name:    "valid emails",
			emails:  []string{"test1@example.com", "test2@example.com"},
			wantErr: false,
		},
		{
			name:    "invalid email in list",
			emails:  []string{"test1@example.com", "invalid-email"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmailList(tt.emails)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmailList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
