# Makefile for CM-Beetle in Cloud-Barista.

MODULE_NAME := cm-beetle
PROJECT_NAME := github.com/cloud-barista/$(MODULE_NAME)
PKG_LIST := $(shell go list $(PROJECT_NAME)/... 2>&1)

GOPROXY_OPTION := GOPROXY=direct # default: GOPROXY=https://proxy.golang.org,direct
GO := $(GOPROXY_OPTION) go
GOPATH := $(shell go env GOPATH)
SWAG := ~/go/bin/swag

.PHONY: all dependency tidy lint update swag swagger build arm prod run stop clean help

all: swag build ## Default target: build the project

dependency: ## Get dependencies
	@echo "Checking dependencies..."
	@$(GO) mod tidy
	@echo "Checked!"

tidy: ## Run go mod tidy for all modules (root, transx, analyzer, deepdiffgo)
	@echo "Running go mod tidy for all modules..."
	@echo "  - Root module..."
	@$(GO) mod tidy
	@echo "  - transx module..."
	@cd transx && $(GO) mod tidy
	@echo "  - analyzer module..."
	@cd analyzer && $(GO) mod tidy
	@echo "  - deepdiffgo module..."
	@cd deepdiffgo && $(GO) mod tidy
	@echo "All modules tidied!"

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
	@$(SWAG) i --parseDependency --generalInfo ./main.go --dir ./ --exclude deepdiffgo,transx/examples --output ./api
	@rm ./main.go
	@echo "Generated Swagger API documentation!"

api-guide: ## Generate API guide using AI (Usage: make api-guide API_PATH=/migration/ns/{nsId}/mci)
	@./scripts/ask-ai-api-guide.sh "$(API_PATH)"

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

prepare-volumes: ## Create bind-mount directories with correct ownership
	@echo "Preparing container-volume directories in deployments/docker-compose/data/..."
	@mkdir -p \
		deployments/docker-compose/data/cb-tumblebug-container/meta_db \
		deployments/docker-compose/data/cb-tumblebug-container/log \
		deployments/docker-compose/data/cb-spider-container/meta_db \
		deployments/docker-compose/data/cb-spider-container/log \
		deployments/docker-compose/data/etcd/data \
		deployments/docker-compose/data/openbao-data \
		deployments/docker-compose/data/mc-terrarium-container/.terrarium \
		deployments/docker-compose/data/cm-beetle-container/db \
		deployments/docker-compose/data/cm-beetle-container/log \
		2>/dev/null || \
	sudo mkdir -p \
		deployments/docker-compose/data/cb-tumblebug-container/meta_db \
		deployments/docker-compose/data/cb-tumblebug-container/log \
		deployments/docker-compose/data/cb-spider-container/meta_db \
		deployments/docker-compose/data/cb-spider-container/log \
		deployments/docker-compose/data/etcd/data \
		deployments/docker-compose/data/openbao-data \
		deployments/docker-compose/data/mc-terrarium-container/.terrarium \
		deployments/docker-compose/data/cm-beetle-container/db \
		deployments/docker-compose/data/cm-beetle-container/log
	@# Fix ownership for mc-terrarium volume (container runs as appuser, uid 1000)
	@if [ "$$(stat -c '%u' deployments/docker-compose/data/mc-terrarium-container/.terrarium 2>/dev/null)" != "$$(id -u)" ]; then \
		echo "Fixing ownership of mc-terrarium volume..."; \
		sudo chown -R $$(id -u):$$(id -g) deployments/docker-compose/data/mc-terrarium-container/.terrarium; \
	fi
	@echo "Prepared!"

up: compose # Build and up services by docker compose

down: compose-down # Down services by docker compose

# ===== Initialization =====
SHELL := /bin/bash

init: ## Run initialization sequence (credential registration for OpenBao and Tumblebug)
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "CM-Beetle (with CB-Tumblebug) Initialization"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@IS_TMP_KEY=0; \
	cleanup_tmp_key() { \
		if [ "$$IS_TMP_KEY" = "1" ] && [ -f ~/.cloud-barista/.tmp_enc_key ]; then \
			rm -f ~/.cloud-barista/.tmp_enc_key; \
		fi; \
	}; \
	trap cleanup_tmp_key EXIT INT TERM HUP; \
	if [ ! -f ~/.cloud-barista/.tmp_enc_key ]; then \
		echo "Notice: A temporary key file will be created for initialization."; \
		echo "        It will be removed automatically after initialization."; \
		printf "Enter the password for credentials.yaml.enc: "; \
		read -s PASS; \
		echo ""; \
		printf "%s" "$$PASS" > ~/.cloud-barista/.tmp_enc_key; \
		chmod 600 ~/.cloud-barista/.tmp_enc_key; \
		IS_TMP_KEY=1; \
	fi; \
	( \
		echo "1. Registering credentials to OpenBao..."; \
		chmod +x ./deployments/docker-compose/openbao/openbao-register-creds.sh 2>/dev/null || true; \
		./deployments/docker-compose/openbao/openbao-register-creds.sh -y && \
		echo "" && \
		echo "2. Registering credentials to Tumblebug..." && \
		chmod +x ./deployments/docker-compose/cb-tumblebug/init/init.sh 2>/dev/null || true; \
		./deployments/docker-compose/cb-tumblebug/init/init.sh; \
	); \
	EXIT_CODE=$$?; \
	if [ "$$EXIT_CODE" -ne 0 ]; then \
		echo "Initialization failed."; \
	fi; \
	exit $$EXIT_CODE
	@echo "Initialization complete!"

init-openbao: ## Initialize OpenBao (one-time setup: generate unseal key + root token)
	@echo "Initializing OpenBao..."
	@chmod +x ./deployments/docker-compose/openbao/openbao-init.sh 2>/dev/null || true
	@./deployments/docker-compose/openbao/openbao-init.sh

unseal: ## Unseal OpenBao (needed after every container restart)
	@echo "Trying to unseal OpenBao (if not already unsealed)..."
	@chmod +x ./deployments/docker-compose/openbao/openbao-unseal.sh 2>/dev/null || true
	@./deployments/docker-compose/openbao/openbao-unseal.sh || true

logs: ## Follow Docker Compose logs (docker compose logs -f)
	@cd deployments/docker-compose && docker compose logs -f

status: ## Show status of Docker Compose services (docker compose ps)
	@cd deployments/docker-compose && docker compose ps --format "table {{.Name}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}"

ps: status ## Alias for status

compose: prepare-volumes ## Build and up services by docker compose
	@echo "Starting OpenBao..."
	@cd deployments/docker-compose && docker compose up -d openbao
	@if [ ! -f deployments/docker-compose/.env ] || ! grep -q '^VAULT_TOKEN=.+' deployments/docker-compose/.env 2>/dev/null; then \
		echo "VAULT_TOKEN not found — running first-time OpenBao initialization..."; \
		bash deployments/docker-compose/openbao/openbao-init.sh; \
	fi
	@$(MAKE) unseal
	@echo "Building and starting all services by docker compose..."
	@cd deployments/docker-compose && DOCKER_BUILDKIT=1 docker compose up --build

compose-down: ## Down services by docker compose
	@echo "Removing services by docker compose..."
	@cd deployments/docker-compose && docker compose down	

clean-db: compose-down ## Clean all database metadata and persistent data (excluding OpenBao)
	@echo "Running cleanDB script..."
	@chmod +x ./deployments/docker-compose/scripts/cleanDB.sh 2>/dev/null || true
	@./deployments/docker-compose/scripts/cleanDB.sh

clean-all: compose-down clean-db ## Full reset including OpenBao (requires re-init)
	@echo "Cleaning OpenBao configuration and secrets..."
	@sudo rm -rf deployments/docker-compose/data/openbao-data/
	@find deployments/docker-compose/openbao/secrets -type f ! -name ".gitkeep" -delete
	@sed -i 's/^VAULT_TOKEN=.*/VAULT_TOKEN=/' deployments/docker-compose/.env 2>/dev/null || true
	@echo "Cleaned! Run 'make up' to re-initialize."

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
