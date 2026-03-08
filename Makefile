.PHONY: help run build migrate-up migrate-down migrate-create docker-up docker-down docker-logs test clean

include .env

# Variables
BINARY_NAME=gateway
DOCKER_COMPOSE=docker-compose
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

help: ## Show help
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run application
	@go run cmd/gateway/main.go

build: ## Build binary
	@go build -o bin/$(BINARY_NAME) cmd/gateway/main.go

docker-up: ## Start Docker containers
	@$(DOCKER_COMPOSE) up -d
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 3

docker-down: ## Stop Docker containers
	@$(DOCKER_COMPOSE) down

docker-logs: ## Show Docker container logs
	@$(DOCKER_COMPOSE) logs -f

migrate-up: ## Apply migrations (goose)
	@goose -dir migrations postgres "$(DB_URL)" up

migrate-down: ## Rollback migrations (goose)
	@goose -dir migrations postgres "$(DB_URL)" down

migrate-create: ## Create new migration (usage: make migrate-create NAME=create_something)
	@goose -dir migrations create $(NAME) sql

test: ## Run tests
	@go test -v ./...

clean: ## Clean binaries
	@rm -rf bin/
	@go clean

tidy: ## Update dependencies
	@go mod tidy

deps: ## Install dependencies
	@go mod download
