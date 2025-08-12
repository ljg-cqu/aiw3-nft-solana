# AIW3 NFT API Makefile

.PHONY: help build run test clean docker-build docker-run docker-stop deps fmt lint

# Default target
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker Compose"
	@echo "  deps         - Download dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"

# Build the application
build:
	@echo "Building AIW3 NFT API..."
	go build -o bin/aiw3-nft-api main.go

# Run the application
run:
	@echo "Starting AIW3 NFT API..."
	go run main.go

# Run with custom port
run-port:
	@echo "Starting AIW3 NFT API on port $(PORT)..."
	go run main.go --port $(PORT)

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	golangci-lint run

# Install linter
install-lint:
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t aiw3-nft-api .

# Run with Docker Compose
docker-run:
	@echo "Starting with Docker Compose..."
	docker-compose up -d

# Run with Docker Compose (with logs)
docker-run-logs:
	@echo "Starting with Docker Compose (with logs)..."
	docker-compose up

# Stop Docker Compose
docker-stop:
	@echo "Stopping Docker Compose..."
	docker-compose down

# View Docker logs
docker-logs:
	@echo "Viewing Docker logs..."
	docker-compose logs -f

# Development setup
dev-setup: deps install-lint
	@echo "Development environment setup complete!"

# Production build
prod-build:
	@echo "Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o bin/aiw3-nft-api main.go

# Check if server is running
health-check:
	@echo "Checking server health..."
	curl -f http://localhost:3000/health || echo "Server is not running"

# Generate API documentation (if using additional tools)
docs:
	@echo "API documentation is auto-generated and available at /docs endpoint"

# Database migration (placeholder for future use)
migrate:
	@echo "Database migrations would go here..."

# Seed data (placeholder for future use)
seed:
	@echo "Data seeding would go here..."

# Full development cycle
dev: deps fmt lint test build
	@echo "Development cycle complete!"

# CI/CD pipeline simulation
ci: deps fmt lint test prod-build
	@echo "CI pipeline complete!"