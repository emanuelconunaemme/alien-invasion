# Go variables
GO      := go
BINARY  := aliens
GOFMT   := gofmt
GOFILES := $(wildcard *.go)

.PHONY: build run clean fmt test

all: build run

build:
	go build -o $(BINARY)

run: build
	./$(BINARY)

clean:
	go clean
	rm -f $(BINARY)

fmt:
	$(GOFMT) -w $(GOFILES)

test:
	$(GO) test ./...

