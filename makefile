SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} 
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: 
	goreload --d ./cmd/api

test:
	$(GOCMD) test ./...

deps:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache:
	$(GOCMD) clean -modcache