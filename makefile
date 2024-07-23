.PHONY: all build up down restart-db check-db

# Application and database container names
APP_NAME=app
DB_CONTAINER_NAME=mariadb

# Default target: build and start the application
all: build up

# Build the application
build:
	@echo "Building the Go application..."
	docker-compose build $(APP_NAME)

# Start the containers
up: check-db
	@echo "Starting the application..."
	docker-compose up -d $(APP_NAME)

# Stop the containers
down:
	@echo "Stopping all containers..."
	docker-compose down

# Check if the database container is already running
check-db:
	@echo "Checking if the database container is already running..."
	@if [ $$(docker ps -q -f name=$(DB_CONTAINER_NAME)) ]; then \
		echo "Database container is already running."; \
	else \
		echo "Starting the database container..."; \
		docker-compose up -d $(DB_CONTAINER_NAME); \
	fi

# Force restart the database container
restart-db:
	@echo "Restarting the database container..."
	docker-compose restart $(DB_CONTAINER_NAME)