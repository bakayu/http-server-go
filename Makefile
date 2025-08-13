# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_CLEAN=$(GO_CMD) clean

# Binaries
CLI_BINARY=cli
WEBSERVER_BINARY=webserver
BUILD_DIR=build

# Default target executed when just `make` is run
all: test build

# Build the applications
build: cli webserver

cli:
	@echo "Building CLI binary..."
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(CLI_BINARY) ./cmd/cli

webserver:
	@echo "Building webserver binary..."
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(WEBSERVER_BINARY) ./cmd/webserver

# Run the applications
run-cli: cli
	@echo "Starting the CLI..."
	./$(BUILD_DIR)/$(CLI_BINARY)

run-webserver: webserver
	@echo "Starting the webserver on http://localhost:5000..."
	./$(BUILD_DIR)/$(WEBSERVER_BINARY)

# Run the tests
test:
	@echo "Running tests..."
	$(GO_TEST) -v ./...

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	$(GO_CLEAN)

# Declare targets that are not files
.PHONY: all build cli webserver run-cli run-webserver test clean