FROM golang:1.16.0-alpine3.13

WORKDIR /tmp/build

RUN apk add --no-cache \
  bash

COPY go.mod go.sum .
COPY bin bin
COPY cmd cmd
COPY lib lib

RUN bin/build

FROM alpine:3.13

WORKDIR /usr/local/flux

COPY --from=0 \
  /tmp/build/grpcserver \
  /tmp/build/grpcgateway \
  .

COPY lib/swagger lib/swagger

