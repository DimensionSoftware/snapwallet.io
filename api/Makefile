GO_LIB_FILES := $(shell find lib -type f -name "*.go" | sort) $(shell find cmd -type f -name "*.go" | sort) go.mod go.sum

run: build_servers
	hivemind

build_servers: grpcserver  grpcgateway


test:
	go test -v ./...

wire:
	bin/generate-wire-code

proto:
	bin/generate-grpc-code
	bin/generate-client

grpcserver: ${GO_LIB_FILES}
	go build ./cmd/grpcserver

grpcgateway: ${GO_LIB_FILES}
	go build ./cmd/grpcgateway