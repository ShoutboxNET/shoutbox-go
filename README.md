# Shoutbox Go Client

A Go client library for the Shoutbox email service, supporting both REST API and SMTP implementations.

## Installation

```bash
go get github.com/shoutboxnet/shoutbox-go
```

## Environment Setup

Create a `.env` file based on the template:

```bash
make env-template   # Creates .env.template
cp .env.template .env
```

Then edit `.env` with your credentials:

```
SHOUTBOX_API_KEY=your_api_key_here
SHOUTBOX_FROM=sender@yourdomain.com
SHOUTBOX_TO=recipient@example.com
```

## Available Make Commands

```bash
make help          # Show available commands
make build         # Build the main program
make run           # Run the main program (requires env vars)
make test          # Run tests (requires env vars)
make run-rest      # Run REST API example
make run-smtp      # Run SMTP example
make clean         # Clean build artifacts
make env-template  # Create .env.template file
```

## Usage

### REST API Client

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/shoutboxnet/shoutbox-go/shoutbox"
)

func main() {
    client := shoutbox.NewClient(os.Getenv("SHOUTBOX_API_KEY"))

    req := &shoutbox.EmailRequest{
        From:    "sender@yourdomain.com",
        To:      "recipient@example.com",
        Subject: "Hello World",
        HTML:    "<h1>Welcome!</h1>",
    }

    err := client.SendEmail(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
}
```

### SMTP Client

```go
package main

import (
    "log"
    "os"

    "github.com/shoutboxnet/shoutbox-go/shoutbox"
)

func main() {
    client := shoutbox.NewSMTPClient(os.Getenv("SHOUTBOX_API_KEY"))

    msg := &shoutbox.EmailMessage{
        From:    "sender@yourdomain.com",
        To:      []string{"recipient@example.com"},
        Subject: "Hello World",
        HTML:    "<h1>Welcome!</h1>",
    }

    err := client.SendEmail(msg)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Attachments

```go
attachment, err := shoutbox.NewAttachmentFromFile("document.pdf")
if err != nil {
    log.Fatal(err)
}

msg := &shoutbox.EmailMessage{
    From:        "sender@yourdomain.com",
    To:          []string{"recipient@example.com"},
    Subject:     "Document Attached",
    HTML:        "<h1>Please find the document attached</h1>",
    Attachments: []shoutbox.Attachment{attachment},
}
```

## Features

- REST API and SMTP support
- File attachments
- Custom headers
- Reply-to address
- Sender name
- Email validation
- Context support (REST API)
- Comprehensive testing

## Testing

Set up environment variables in `.env` file and run:

```bash
make test
```

## Examples

Check the `examples` directory for complete usage examples:

- `examples/rest`: REST API implementation example
- `examples/smtp`: SMTP implementation example with attachments

Run examples using:

```bash
make run-rest    # Run REST API example
make run-smtp    # Run SMTP example
```

## Rate Limits

- 60 requests per minute per API key
- Maximum attachment size: 10MB
- Maximum recipients per email: 50

## License

MIT License
