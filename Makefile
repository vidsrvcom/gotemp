.PHONY: build run test lint clean docker-build docker-run help

# Application name
APP_NAME := gotemp
MAIN_PATH := ./cmd/api

# Go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOVET := $(GOCMD) vet
GOFMT := gofmt
GOMOD := $(GOCMD) mod

# Build directory
BUILD_DIR := bin

# Default target
.DEFAULT_GOAL := help

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## build: Build the application
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

## run: Run the application
run:
	@echo "Running..."
	$(GORUN) $(MAIN_PATH)

## test: Run tests
test:
	@echo "Testing..."
	$(GOTEST) -v -race -cover ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "Testing with coverage..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

## lint: Run linter
lint:
	@echo "Linting..."
	$(GOVET) ./...
	$(GOFMT) -s -l .

## fmt: Format code
fmt:
	@echo "Formatting..."
	$(GOFMT) -s -w .

## tidy: Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

## docker-run: Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME):latest

## all: Build and test
all: lint test build
