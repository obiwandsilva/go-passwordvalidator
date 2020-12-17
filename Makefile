.PHONY: test/unity test/integration test/all

GOIMPORTS ?= golang.org/x/tools/cmd/goimports
GOLINT ?= github.com/golangci/golangci-lint/cmd/golangci-lint
SOURCES := $(shell go run $(GOIMPORTS) -l **/**.go | xargs)

goimports:
	go run $(GOIMPORTS) -w $(SOURCES)

lint:
	go run $(GOLINT) run

test/unit:
	go test -v -race ./domain/...

test/integration:
	docker-compose up --build -d
	(go test -v -race ./it/... && docker-compose down) || (docker-compose down && exit 1)

test/all: test/unit test/integration
