GO_LIB_FILES := $(shell find lib -type f -name "*.go" | sort) go.mod go.sum
GO_GENERATED_PROTOCOL_FILES := $(shell find lib/protocol -type f -name "*.go" | sort)
GO_GENERATED_SWAGGER_FILE := lib/swagger/swagger.json

build: grpcserver grpcgateway

build_client: ${GO_GENERATED_SWAGGER_FILE}
	rm -rf client
	openapi-generator-cli generate -g typescript -i ${GO_GENERATED_SWAGGER_FILE} -o client

test: ${GO_LIB_FILES}
	go test -v ./...

run: build
	hivemind

clean:
	rm grpcserver grpcgateway

lib/wire/wire_gen.go: wire.go
	bin/generate-wire-code

${GO_GENERATED_PROTOCOL_FILES} ${GO_GENERATED_SWAGGER_FILE}: lib/protocol/api.proto
	bin/generate-grpc-code

grpcserver: ${GO_LIB_FILES} $(shell find cmd/grpcserver -type f -name "*.go" | sort)
	go build ./cmd/grpcserver

grpcgateway: ${GO_LIB_FILES} $(shell find cmd/grpcgateway -type f -name "*.go" | sort)
	go build ./cmd/grpcgateway

