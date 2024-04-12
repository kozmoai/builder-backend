.PHONY: build all test clean

all: build

build: build-http-server build-websocket-server build-http-server-internal

build-http-server:
	go build -o bin/kozmo-builder-backend src/cmd/kozmo-builder-backend/main.go

build-websocket-server:
	go build -o bin/kozmo-builder-backend-websocket src/cmd/kozmo-builder-backend-websocket/main.go

build-http-server-internal:
	go build -o bin/kozmo-builder-backend-internal src/cmd/kozmo-builder-backend-internal/main.go

test:
	PROJECT_PWD=$(shell pwd) go test -race ./...

test-cover:
	go test -cover --count=1 ./...

cover-total:
	go test -cover --count=1 ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total 

cov:
	PROJECT_PWD=$(shell pwd) go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path './*_test.go')

fmt-check:
	@gofmt -l $(shell find . -type f -name '*.go' -not -path './*_test.go')

init-database:
	/bin/bash scripts/postgres-init.sh

clean:
	@ro -fR bin
