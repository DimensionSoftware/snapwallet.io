#!/bin/bash

# Generate from private key public key
openssl rsa -in bin/private.pem -pubout > bin/public.pem 2> /dev/null

# Base 64 encode public key for .env
base64 -i bin/public.pem
