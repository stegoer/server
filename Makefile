ifneq (,$(wildcard ./.env))
	include .env
endif

GOBIN ?= $(shell pwd)/bin
export

AIR = $(GOBIN)/air
ENT = $(GOBIN)/ent
GQLGEN = $(GOBIN)/gqlgen
GODOC = $(GOBIN)/godoc
GOFUMPT = $(GOBIN)/gofumpt
GOLANGCI_LINT = $(GOBIN)/golangci-lint
GOIMPORTS = $(GOBIN)/goimports-reviser
MIGRATE = $(GOBIN)/migrate

# Directories containing independent Go modules.
MODULE_DIRS = . ./tools

# Many Go tools take file globs or directories as arguments instead of packages.
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

$(AIR):
	cd tools && go install github.com/cosmtrek/air@latest

$(ENT):
	cd tools && go install entgo.io/ent/cmd/ent

$(GQLGEN):
	cd tools && go install github.com/99designs/gqlgen

$(GODOC):
	cd tools && go install golang.org/x/tools/cmd/godoc

$(GOFUMPT):
	cd tools && go install mvdan.cc/gofumpt

$(GOIMPORTS):
	cd tools && go install github.com/incu6us/goimports-reviser

$(GOLANGCI_LINT):
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint

$(MIGRATE):
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz migrate
	mv migrate bin/migrate

.PHONY: describe serve-doc gen fmt lint tidy test cover dev db-init makemigrations migrate-all migrate-up migrate-down build clean help

default: help

describe: $(GQLGEN)  ## Describe ent schema.
	@$(ENT) describe ./ent/schema

serve-doc: $(GODOC)  ## Serve go documentation.
	@$(GODOC) -http=:6060

gen: $(GQLGEN)  ## Generate server files.
	@echo "generating ent files"
	@go generate ./ent
	@echo "generating gqlgen files"
	@$(GQLEN)
	@echo "files generated"

fmt: $(GOFUMPT) $(GOIMPORTS) ## Format source files.
	@echo "formatting via gofumpt and goimports-reviser"
	@$(foreach file,$(GO_FILES),(echo "fmt $(file)" && $(GOFUMPT) -e -w $(file) && $(GOIMPORTS) -project-name github.com/stegoer/server -file-path $(file)) &&) true

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

dev: $(AIR) ## Start development server.
	@echo "starting development server via air..."
	@$(AIR)

db-init:  ## Initialize database.
	@createdb stegoer
	@make migrate-all

makemigrations:  ## Make new migrations.
	@go run cmd/makemigrations/makemigrations.go

migrate-all: $(MIGRATE)  ## Apply all database up migration.
	@$(MIGRATE) -source file://migrations -database $(DATABASE_URL) up

migrate-up: $(MIGRATE)  ## Apply database up migration.
	@$(MIGRATE) -source file://migrations -database $(DATABASE_URL) up 1

migrate-down: $(MIGRATE)  ## Apply database down migration.
	@$(MIGRATE) -source file://migrations -database $(DATABASE_URL) down 1

build:  ## Build server.
	@echo "building server..."
	@go build cmd/server/server.go

clean: ## Cleanup files.
	@echo "cleaning up 'bin'"
	@rm -rf bin/*

help: ## Display help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
