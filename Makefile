# Include .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Check for required environment variables
check-env:
	@if [ -z "$(SHOUTBOX_API_KEY)" ]; then \
		echo "Error: SHOUTBOX_API_KEY is not set"; \
		exit 1; \
	fi
	@if [ -z "$(SHOUTBOX_FROM)" ]; then \
		echo "Error: SHOUTBOX_FROM is not set"; \
		exit 1; \
	fi
	@if [ -z "$(SHOUTBOX_TO)" ]; then \
		echo "Error: SHOUTBOX_TO is not set"; \
		exit 1; \
	fi

# Build the main program
build:
	go build -o bin/shoutbox main.go

# Run the main program (requires environment variables)
run: check-env build
	./bin/shoutbox

# Run tests (requires environment variables)
test: check-env
	go test -v ./...

# Run REST API example
run-rest: check-env
	go run examples/rest/main.go

# Run SMTP example
run-smtp: check-env
	go run examples/smtp/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Create .env template file
env-template:
	@echo "SHOUTBOX_API_KEY=" > .env.template
	@echo "SHOUTBOX_FROM=" >> .env.template
	@echo "SHOUTBOX_TO=" >> .env.template
	@echo "Created .env.template file"

# Show help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the main program"
	@echo "  make run         - Run the main program (requires env vars)"
	@echo "  make test        - Run tests (requires env vars)"
	@echo "  make run-rest    - Run REST API example"
	@echo "  make run-smtp    - Run SMTP example"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make env-template - Create .env.template file"
	@echo ""
	@echo "Required environment variables (can be set in .env file):"
	@echo "  SHOUTBOX_API_KEY - Your Shoutbox API key"
	@echo "  SHOUTBOX_FROM    - Sender email address"
	@echo "  SHOUTBOX_TO      - Recipient email address"

.PHONY: check-env build run test run-rest run-smtp clean env-template help
