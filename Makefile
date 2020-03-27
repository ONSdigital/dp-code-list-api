SHELL=bash

BUILD=build
BUILD_ARCH=$(BUILD)/$(GOOS)-$(GOARCH)
BIN_DIR?=.

export GOOS?=$(shell go env GOOS)
export GOARCH?=$(shell go env GOARCH)

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)
LDFLAGS=-ldflags "-w -s -X 'main.Version=${VERSION}' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'"

build:
	@mkdir -p $(BUILD_ARCH)/$(BIN_DIR)
	go build $(LDFLAGS) -o $(BUILD_ARCH)/$(BIN_DIR)/dp-code-list-api cmd/dp-code-list-api/main.go
test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

debug:
	HUMAN_LOG=1 GRAPH_DRIVER_TYPE=neo4j GRAPH_ADDR="bolt://localhost:7687" go run $(LDFLAGS) -race cmd/dp-code-list-api/main.go

acceptance:
	HUMAN_LOG=1 go run $(LDFLAGS) -race cmd/dp-code-list-api/main.go

.PHONY: test build debug acceptance
