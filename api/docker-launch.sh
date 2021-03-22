#!/bin/bash -xeu

./grpcserver &
while ! echo exit | nc localhost 50051; do sleep 0.1; done
./grpcgateway &

wait -n
