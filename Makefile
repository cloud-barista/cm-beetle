# Makefile for CM-Beetle in Cloud-Barista.

MODULE_NAME := cm-beetle
PROJECT_NAME := github.com/cloud-barista/$(MODULE_NAME)
PKG_LIST := $(shell go list $(PROJECT_NAME)/... 2>&1)

GOPROXY_OPTION := GOPROXY=direct # default: GOPROXY=https://proxy.golang.org,direct
GO := $(GOPROXY_OPTION) go
GOPATH := $(shell go env GOPATH)
SWAG := ~/go/bin/swag

.PHONY: all dependency lint update swag swagger build arm prod run stop clean help

all: swag build ## Default target: build the project

dependency: ## Get dependencies
	@echo "Checking dependencies..."
	@$(GO) mod tidy
	@echo "Checked!"

lint: dependency ## Lint the files
	@echo "Running linter..."
	@if [ ! -f "$(GOPATH)/bin/golangci-lint" ] && [ ! -f "$(shell go env GOROOT)/bin/golangci-lint" ]; then \
		$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.2; \
	fi
	@golangci-lint run -E contextcheck -D unused
	@echo "Linter finished!"

update: ## Update all module dependencies
	@echo "Updating dependencies..."
	@cd cmd/$(MODULE_NAME) && $(GO) get -u
	@echo "Checking dependencies..."
	@$(GO) mod tidy
	@echo "Updated!"

swag swagger: ## Generate Swagger API documentation
	@echo ""
	@echo "This commend requires swag binary."
	@echo "If you don't have it, please install swag first with the following commend"
	@echo "- go install github.com/swaggo/swag/cmd/swag@latest"
	@echo ""
	@echo "Generating Swagger API documentation..."
	@ln -sf cmd/$(MODULE_NAME)/main.go ./main.go
	@$(SWAG) i --parseDependency --generalInfo ./main.go --dir ./ --output ./api
	@rm ./main.go
	@echo "Generated Swagger API documentation!"

# build: lint swag ## Build the binary file for amd64
build: ## Build the binary file for amd64
	@echo "Building the binary for amd64..."
	@cd cmd/$(MODULE_NAME) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(MODULE_NAME) main.go
	@echo "Build finished!"

# arm: lint swag ## Build the binary file for ARM
arm: ## Build the binary file for ARM
	@echo "Building the binary for ARM..."
	@cd cmd/$(MODULE_NAME) && CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO) build -o $(MODULE_NAME)-arm main.go
	@echo "Build finished!"

# prod: lint swag ## Build the binary file for production
prod: ## Build the binary file for production
	@echo "Building the binary for amd64 production..."
# Note - Using cgo write normal Go code that imports a pseudo-package "C". I may not need on cross-compiling.
# Note - You can find possible platforms by 'go tool dist list' for GOOS and GOARCH
# Note - Using the -ldflags parameter can help set variable values at compile time.
# Note - Using the -s and -w linker flags can strip the debugging information.	
	@cd cmd/$(MODULE_NAME) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags '-s -w' -tags $(MODULE_NAME) -v -o $(MODULE_NAME) main.go
	@echo "Build finished!"

run: build ## Run the built binary
	@echo "Running the binary..."
	@source conf/setup.env; \
	cd cmd/$(MODULE_NAME) && \
	(./$(MODULE_NAME) || { echo "Trying with sudo..."; sudo ./$(MODULE_NAME); })

stop: ## Stop the built binary
	@echo "Stopping the binary..."
	@sudo killall $(MODULE_NAME) 2>/dev/null || true
	@echo "Stopped!"

clean: ## Remove previous build
	@echo "Cleaning build..."
	@rm -f coverage.out
	@rm -f api/docs.go api/swagger.*
	@cd cmd/$(MODULE_NAME) && $(GO) clean
	@cd cmd/test-cli && $(GO) clean
	@echo "Cleaned!"

test-cli-build: ## Build the test CLI binary
	@echo "Building test CLI binary..."
	@$(GO) build -o cmd/test-cli/test-cli ./cmd/test-cli/main.go
	@echo "Test CLI build finished!"

test-cli: test-cli-build ## Run the test CLI for all CSP-Region pairs
	@echo "Running test CLI for all CSP-Region pairs..."
	@cd cmd/test-cli && ./test-cli -config testdata/config-multi-csp-and-region-pair.json

test-cli-help: ## Show test CLI help
	@echo "Test CLI Help:"
	@cd cmd/test-cli && ./test-cli -h || true

compose: ## Build and up services by docker compose
	@echo "Building and starting services by docker compose..."
	@cd deployments/docker-compose && DOCKER_BUILDKIT=1 docker compose up --build

compose-up: ## Up services by docker compose
	@echo "Starting services by docker compose..."
	@cd deployments/docker-compose && docker compose up

compose-down: ## Down services by docker compose
	@echo "Removing services by docker compose..."
	@cd deployments/docker-compose && docker compose down	

apidiff: ## Compare Swagger API differences (Usage: make apidiff <old_ver> [new_ver] [OPTS="..."])
	@OLD_VER=$(word 2,$(MAKECMDGOALS)); \
	NEW_VER=$(word 3,$(MAKECMDGOALS)); \
	if [ -z "$$OLD_VER" ]; then \
		echo "Error: Please specify at least the old version"; \
		echo "Usage: make apidiff <old_ver> [new_ver] [OPTS=\"...\"]"; \
		echo "Example 1: make apidiff v0.4.0 (Compare v0.4.0 with local)"; \
		echo "Example 2: make apidiff v0.4.0 v0.4.1 (Compare v0.4.0 with v0.4.1)"; \
		echo "Example 3: make apidiff v0.4.0 OPTS=\"-f markdown -o api/diff.md\""; \
		exit 1; \
	fi; \
	OLD_URL="https://raw.githubusercontent.com/cloud-barista/cm-beetle/$$OLD_VER/api/swagger.yaml"; \
	if [ -z "$$NEW_VER" ]; then \
		NEW_SPEC="api/swagger.yaml"; \
		echo "Comparing $$OLD_VER (remote) with local api/swagger.yaml..."; \
	else \
		NEW_SPEC="https://raw.githubusercontent.com/cloud-barista/cm-beetle/$$NEW_VER/api/swagger.yaml"; \
		echo "Comparing $$OLD_VER (remote) with $$NEW_VER (remote)..."; \
	fi; \
	echo "Running deepdiffgo..."; \
	go run deepdiffgo/cmd/deepdiffgo/main.go $$OLD_URL $$NEW_SPEC $(OPTS)

# Dummy targets to prevent make from interpreting branch names as targets
%:
	@:

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
