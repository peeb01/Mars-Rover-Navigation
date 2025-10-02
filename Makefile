BINARY_NAME=mars
PKG=./...

.PHONY: all prepare build run test cover cover-html clean tidy fmt lint

all: build

prepare:
	@echo "==> Preparing dependencies..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		if command -v snap >/dev/null 2>&1; then \
			sudo snap install golangci-lint --classic || true; \
		else \
			go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		fi \
	fi
	@echo "Dependencies ready."


build:
	@echo "==> Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) main.go

run: build
	@echo "==> Running $(BINARY_NAME)..."
	./bin/$(BINARY_NAME)

test:
	@echo "==> Running tests..."
	go test $(PKG) -cover

cover:
	@echo "==> Generating coverage report (text)..."
	go test $(PKG) -coverprofile=coverage.out
	go tool cover -func=coverage.out

cover-html:
	@echo "==> Generating coverage report (HTML)..."
	go test $(PKG) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html in your browser to view coverage report"

fmt:
	@echo "==> Formatting code..."
	go fmt ./...

tidy:
	@echo "==> Running go mod tidy..."
	go mod tidy

lint:
	@echo "==> Linting code..."
	golangci-lint run ./...

clean:
	@echo "==> Cleaning..."
	rm -rf bin coverage.out coverage.html
