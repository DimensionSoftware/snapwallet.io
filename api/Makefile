build_gateway: cmd lib go.mod go.sum
	go build ./cmd/grpcgateway

build_server: cmd lib go.mod go.sum
	go build ./cmd/grpcserver

gateway: build_gateway
	./grpcgateway

server: build_server
	./grpcserver

run: build_gateway build_server
	hivemind

test:
	go test -v ./...