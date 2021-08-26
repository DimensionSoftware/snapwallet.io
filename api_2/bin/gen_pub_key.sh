#!/bin/bash

# Generate from private key public key
openssl rsa -in private.pem -pubout > public.pem 2> /dev/null

# Base 64 encode public key for .env
base64 -i bin/public.pem
