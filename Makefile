# Go parameters
BINARY_NAME=mars
PKG=./pkg

.PHONY: all build run test cover cover-html clean tidy fmt

# Build binary
build:
	@echo "==> Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) main.go

# Run program
run: build
	@echo "==> Running $(BINARY_NAME)..."
	./bin/$(BINARY_NAME)

# Run tests in pkg/ with coverage summary
test:
	@echo "==> Running tests in pkg/..."
	go test $(PKG) -cover

# Generate coverage report (text) for pkg/
cover:
	@echo "==> Generating coverage report (text)..."
	go test $(PKG) -coverprofile=coverage.out
	go tool cover -func=coverage.out

# Go linting
lint:
	@echo "==> Linting code..."
	golangci-lint run ./...

# Generate coverage report (HTML) for pkg/
cover-html:
	@echo "==> Generating coverage report (HTML)..."
	go test $(PKG) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html in your browser to view coverage report"

# Format code
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# Tidy modules
tidy:
	@echo "==> Running go mod tidy..."
	go mod tidy

# Clean up build artifacts
clean:
	@echo "==> Cleaning..."
	rm -rf bin coverage.out coverage.html
