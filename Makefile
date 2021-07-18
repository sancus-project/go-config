.PHONY: all fmt build test

GO ?= go

all: generate fmt build

generate:
	$(GO) generate ./...

fmt:
	$(GO) fmt ./...
	$(GO) mod tidy || true

build:
	$(GO) get -v ./...

test:
	$(GO) test -v ./...
