BINARY_NAME=depgraph
VERSION?=dev
BUILD_DIR=bin
GO=go
GOFLAGS=-trimpath
LDFLAGS=-ldflags "-s -w -X github.com/alokshukla631/depgraph/internal/cli.Version=$(VERSION)"

.PHONY: all build clean test lint fmt vet help

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':'

## build: Build the binary
build:
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/depgraph

## test: Run all unit tests
test:
	$(GO) test -race -count=1 ./...

## test-verbose: Run all unit tests with verbose output
test-verbose:
	$(GO) test -race -count=1 -v ./...

## test-coverage: Run tests with coverage report
test-coverage:
	$(GO) test -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## lint: Run golangci-lint
lint:
	golangci-lint run ./...

## fmt: Format all Go files
fmt:
	$(GO) fmt ./...
	goimports -w .

## vet: Run go vet
vet:
	$(GO) vet ./...

## clean: Remove build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

## run: Build and run with sample compose file
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) analyze testdata/compose/simple.yml

## install: Install the binary to $GOPATH/bin
install:
	$(GO) install $(GOFLAGS) $(LDFLAGS) ./cmd/depgraph

all: fmt vet lint test build
