.PHONY: help build dev test clean install-deps docker-up docker-down

.DEFAULT_GOAL := help

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-deps: ## Install dependencies for both backend and frontend
	cd backend && go mod download
	cd frontend && npm install

build: ## Build both backend and frontend
	cd backend && go build -o bin/api cmd/api/main.go
	cd backend && go build -o bin/agent cmd/agent/main.go
	cd frontend && npm run build

dev: ## Run development environment with hot reload
	docker-compose up -d

dev-backend: ## Run backend with hot reload (requires air)
	cd backend && air

dev-frontend: ## Run frontend with hot reload
	cd frontend && npm run dev

test: ## Run all tests
	cd backend && go test -v ./...

test-backend: ## Run backend tests
	cd backend && go test -v ./...

test-frontend: ## Run frontend tests
	cd frontend && npm run test

db-migrate: ## Run database migrations
	cd backend && go run cmd/migrate/main.go

db-seed: ## Seed database with test data
	cd backend && go run cmd/seed/main.go

clean: ## Clean build artifacts
	rm -rf backend/bin
	rm -rf frontend/dist

docker-build: ## Build Docker images
	docker-compose build

docker-up: ## Start Docker containers
	docker-compose up -d

docker-down: ## Stop Docker containers
	docker-compose down

docker-logs: ## View Docker logs
	docker-compose logs -f

db-shell: ## Connect to PostgreSQL database
	docker-compose exec postgres psql -U vps_user -d vps_security

api-docs: ## Open API documentation
	open http://localhost:8080/docs
