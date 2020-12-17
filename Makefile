.PHONY: test/unity test/integration test/all

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint --timeout 10m

goimports:
	go run golang.org/x/tools/cmd/goimports -w ./...

test/unit:
	go test -v -race ./domain/...

test/integration:
	docker-compose up --build -d
	(go test -v -race ./it/... && docker-compose down) || (docker-compose down && exit 1)

test/all: test/unit test/integration
