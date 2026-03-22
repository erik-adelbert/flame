.PHONY: build run test bench fmt clean

BIN_DIR := bin
BIN := $(BIN_DIR)/flame

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN) ./cmd/flame/main.go

run:
	go run ./cmd/flame/main.go

test:
	go test ./...

bench:
	go test ./flame -run '^$$' -bench . -benchmem

fmt:
	gofmt -w ./cmd ./flame

clean:
	rm -rf $(BIN_DIR)
