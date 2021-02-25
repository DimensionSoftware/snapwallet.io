GO_LIB_FILES := $(shell find lib -type f -name "*.go" | sort) go.mod go.sum

grpcserver: ${GO_LIB_FILES} $(shell find cmd/grpcserver -type f -name "*.go" | sort)
	go build ./cmd/grpcserver

grpcgateway: ${GO_LIB_FILES} $(shell find cmd/grpcgateway -type f -name "*.go" | sort)
	go build ./cmd/grpcgateway

run: grpcserver grpcgateway
	hivemind

test:
	go test -v ./...