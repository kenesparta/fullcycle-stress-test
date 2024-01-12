SHELL := /bin/bash

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stress-test-cli ./cmd/cli

.PHONY: run
run:
	docker compose -f docker-compose.yaml up -d --build
