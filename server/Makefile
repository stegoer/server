export GOBIN ?= $(shell pwd)/bin

# Directories containing independent Go modules.
MODULE_DIRS = .

.PHONY: help lint tidy test cover codegen dev

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo ""
	@echo "  lint        lint all source files"
	@echo "  tidy        check 'go mod tidy' of every module"
	@echo "  test        run all tests of each module directory"
	@echo "  cover   	 run tests with coverage report"
	@echo "  codegen     generate ent and graphql code"
	@echo "  dev         start development server"
	@echo ""
	@echo "Check the Makefile to know exactly what each target is doing."

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./... 2>&1
	@echo "Checking 'go mod tidy'..."
	@make tidy

tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy -compat=1.17) &&) true

test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true

cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

codegen:
	@echo "Generating ent..."
	@ent generate ./ent/schema
	@echo "Generating gqlgen..."
	@gqlgen

dev:
	@echo "Starting server..."
	@go run cmd/stegoer-server/main.go
