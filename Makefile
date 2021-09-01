.PHONY: all generate fmt build test

GO ?= go
GOFMT ?= gofmt
GOFMT_FLAGS = -w -l -s

all: generate fmt build

generate:
	$(GO) generate ./...

fmt:
	@find . -name '*.go' | xargs -r $(GOFMT) $(GOFMT_FLAGS)
	$(GO) mod tidy || true

build:
	$(GO) get -v ./...

test:
	$(GO) test -v ./...
