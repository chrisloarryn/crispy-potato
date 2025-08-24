.PHONY: build run test clean docker-build docker-run docker-down deps lint fmt vet

# Build the application
build:
	go build -o bin/crispy-potato cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf uploads/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run all checks
check: fmt vet lint test

# Docker build
docker-build:
	docker build -t crispy-potato .

# Docker run
docker-run:
	docker-compose up -d

# Docker stop
docker-down:
	docker-compose down

# Docker logs
docker-logs:
	docker-compose logs -f

# Create uploads directories
setup-dirs:
	mkdir -p uploads/avatars uploads/banners

# Development setup
dev-setup: deps setup-dirs
	@echo "Development environment setup complete!"

# Production build
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/crispy-potato cmd/api/main.go

# Help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  check         - Run all checks"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run with Docker Compose"
	@echo "  docker-down   - Stop Docker Compose"
	@echo "  docker-logs   - View Docker logs"
	@echo "  setup-dirs    - Create uploads directories"
	@echo "  dev-setup     - Setup development environment"
	@echo "  prod-build    - Build for production"
	@echo "  help          - Show this help"
