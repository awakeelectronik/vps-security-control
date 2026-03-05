.PHONY: help run build test migrate seed

help:
	@echo "VPS Security Control - Backend Commands"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the API server
	go run cmd/api/main.go

build: ## Build the API binary
	go build -o bin/api cmd/api/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	@echo "\nCoverage report generated: coverage.out"

migrate: ## Run database migrations
	go run cmd/migrate/main.go

seed: ## Seed database with test data
	go run cmd/seed/main.go

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run ./...

dev: ## Run with hot reload (requires air)
	air
