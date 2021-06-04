GO_LIB_FILES := $(shell find lib -type f -name "*.go" | sort) $(shell find cmd -type f -name "*.go" | sort) go.mod go.sum

run: grpcdevserver grpcgateway
	hivemind

test:
	bin/test

coverage:
	bin/coverage

wire:
	bin/generate-wire-code

proto:
	bin/generate-grpc-code
	bin/generate-client

grpcserver: ${GO_LIB_FILES}
	go build ./cmd/grpcserver

grpcdevserver: ${GO_LIB_FILES}
	go build ./cmd/grpcdevserver

grpcgateway: ${GO_LIB_FILES}
	go build ./cmd/grpcgateway
