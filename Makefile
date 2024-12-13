.PHONY: dev tidy lint build

# Command to run the dev server
dev:
	@echo "Starting development server with hot reload..."
	air

# Clean up Go module dependencies
tidy:
	@echo "Tidying Go modules..."
	go mod tidy

# Run linter
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Build the project
build:
	@echo "Building the Go project..."
	go build -o bin/app .


## Purpose Step CI
#steps:
# - name: Run Makefile commands
#  run: |
#      make tidy
#      make lint