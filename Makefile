# Variables for easy updates
BINARY_NAME=portfolio
GO_FILES=server.go

# The 'all' target is usually the default
all: build test

build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) $(GO_FILES)

run: 
	go run $(GO_FILES)

test:
	go test -v ./...

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

# .PHONY tells Make that these are commands, not actual files on disk
.PHONY: all build run test clean