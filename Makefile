# Variables
APP_NAME=test_baltic
DOCKER_COMPOSE=docker-compose
MIGRATE=migrate
DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
MIGRATIONS_PATH=./migrations

# Default target
.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Application build and run
build: ## Build the Go binary
	go build -o $(APP_NAME) ./cmd/main.go

run: ## Run the Go application
	go run ./cmd/main.go

docker-build: ## Build the Docker container
	docker build -t $(APP_NAME):latest .

docker-run: ## Run the application in Docker
	$(DOCKER_COMPOSE) up --build

docker-stop: ## Stop the Docker containers
	$(DOCKER_COMPOSE) down

# Database migrations
migrate-up: ## Run database migrations
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

migrate-down: ## Rollback the last migration
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATIONS_PATH) down 1

migrate-new: ## Create a new migration file
	@read -p "Enter migration name: " name; \
	touch $(MIGRATIONS_PATH)/`date +%Y%m%d%H%M%S`_$$name.up.sql $(MIGRATIONS_PATH)/`date +%Y%m%d%H%M%S`_$$name.down.sql;

# Cleanup
clean: ## Clean up generated files
	rm -f $(APP_NAME)

format: ## Format the code
	gofmt -w .

# Docker and Compose utilities
docker-logs: ## Show logs from the Docker containers
	$(DOCKER_COMPOSE) logs -f

docker-clean: ## Remove stopped containers and dangling images
	docker system prune -f

# Full rebuild
docker-rebuild: docker-stop clean docker-build docker-run ## Clean, rebuild, and run the application in Docker
