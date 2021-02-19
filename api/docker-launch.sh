#!/bin/bash -xeu

./grpcserver &
./grpcgateway &

wait -n
