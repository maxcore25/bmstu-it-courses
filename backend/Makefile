.PHONY: dev build run swag help

dev:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Starting dev mode (Air hot reload)..."
	air

build:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Building Go application..."
	@mkdir -p bin
	go build -o bin/app .
	@echo "Build complete: bin/app"

run:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Running Go application..."
	go run .

swag:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"

help:
	@echo "Available commands:"
	@echo "  dev    - Start Air hot reload (with Swagger hot reload)"
	@echo "  build  - Build Go app into ./bin/app (Swagger updated first)"
	@echo "  run    - Run Go app (no reload)"
	@echo "  swag   - Generate Swagger docs"
	@echo "  help   - Show this help"
