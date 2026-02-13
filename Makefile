.PHONY: build run test lint clean

# Build the application
build:
	@echo "Building todotxt-tui..."
	@go build -o bin/todotxt-tui ./cmd/todotxt-tui

# Run the application
run: build
	@./bin/todotxt-tui

# Run all tests with Ginkgo
test:
	@echo "Running tests..."
	@ginkgo -r --randomize-all --randomize-suites --fail-on-pending --cover

# Run linting and formatting
lint:
	@echo "Running linters..."
	@go mod tidy
	@goimports -w .
	@golangci-lint run ./...
	@go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out
	@find . -name "*.test" -delete
	@find . -name "*.out" -delete

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/onsi/ginkgo/v2/ginkgo@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run tests with coverage report
coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
