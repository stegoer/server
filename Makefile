export GOBIN ?= $(shell pwd)/bin

# Directories containing independent Go modules.
MODULE_DIRS = .

# Many Go tools take file globs or directories as arguments instead of packages.
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: help lint tidy test cover generate dev clean

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo ""
	@echo "  lint        lint all source files"
	@echo "  tidy        check 'go mod tidy' of every module"
	@echo "  test        run all tests of each module directory"
	@echo "  cover   	 run tests with coverage report"
	@echo "  generate    generate ent and gql code"
	@echo "  dev         start development server"
	@echo "  clean       clean object files from package source directories"
	@echo ""
	@echo "Check the Makefile to know exactly what each target is doing."

lint: $(GOLINT) $(STATICCHECK)
	@echo "Running golangci-lint..."
	@golangci-lint run ./... 2>&1
	@echo "Checking 'go mod tidy'..."
	@make tidy

tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true

test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true

cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

generate:
	go generate
	gqlgen

dev:
	@go run cmd/app.go
