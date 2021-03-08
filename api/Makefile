GO_LIB_FILES := $(shell find lib -type f -name "*.go" | sort) $(shell find cmd -type f -name "*.go" | sort) go.mod go.sum

run: grpcserver  grpcgateway
	hivemind

test:
	go test -coverprofile cover.out -v ./...

coverage: test
	go tool cover -html cover.out

wire:
	bin/generate-wire-code

proto:
	bin/generate-grpc-code
	bin/generate-client

grpcserver: ${GO_LIB_FILES}
	go build ./cmd/grpcserver

grpcgateway: ${GO_LIB_FILES}
	go build ./cmd/grpcgateway