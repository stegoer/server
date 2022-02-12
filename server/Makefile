export GOBIN ?= $(shell pwd)/bin

ENT = $(GOBIN)/ent
GQLGEN = $(GOBIN)/gqlgen
GOFUMPT = $(GOBIN)/gofumpt
GOLANGCI_LINT = $(GOBIN)/golangci-lint
GOIMPORTS = $(GOBIN)/goimports-reviser

# Directories containing independent Go modules.
MODULE_DIRS = . ./tools

# Many Go tools take file globs or directories as arguments instead of packages.
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

$(ENT):
	cd tools && go install entgo.io/ent/cmd/ent

$(GQLGEN):
	cd tools && go install github.com/99designs/gqlgen

$(GOFUMPT):
	cd tools && go install mvdan.cc/gofumpt

$(GOIMPORTS):
	cd tools && go install github.com/incu6us/goimports-reviser

$(GOLANGCI_LINT):
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: gen fmt lint tidy test cover ci dev server clean help

default: help

gen: $(ENT) $(GQLGEN)  ## Generate server files.
	@echo "generating ent files"
	@$(ENT) generate ./ent/schema
	@echo "generating gqlgen files"
	@$(GQLEN)

fmt: $(GOFUMPT) $(GOIMPORTS) ## Format source files.
	@echo "formatting via gofumpt and goimports-reviser"
	@$(foreach file,$(GO_FILES),(echo "fmt $(file)" && $(GOFUMPT) -e -w $(file) && $(GOIMPORTS) -project-name github.com/kucera-lukas/stegoer -rm-unused -file-path $(file)) &&) true

lint: $(GOLANGCI_LINT)  ## Lint source files.
	@echo "linting via golangci-lint"
	@$(GOLANGCI_LINT) run --config ./.golangci-lint.yml ./...

tidy:  ## Tidy module dependencies.
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && echo "tidy $(dir)" && go mod tidy -compat=1.17) &&) true

test: ## Run module tests.
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && echo "test $(dir)" && go test -race ./...) &&) true

cover:  ## Run tests with coverage.
	@echo "running tests with coverage"
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

ci:  ## Run continuous integration.
	@make tidy
	@make gen
	@make fmt
	@make lint
	# @make test
	# @make cover

dev:  ## Start development server.
	@echo "starting development server..."
	@go run cmd/stegoer-server/main.go

server:  ## Build and start server.
	@echo "building server..."
	@go build cmd/stegoer-server/main.go
	@echo "starting server..."
	@./main

clean: ## Cleanup files.
	@echo "cleaning up 'bin'"
	@rm bin/*

help: ## Display help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
