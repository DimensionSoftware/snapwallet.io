#!/bin/bash -xeu

API_KEY=$(date +%s | sha256sum | base64 | head -c 35)
echo "API KEY: $API_KEY"

curl -XPOST \
  -H "Content-Type: application/json" \
  -d "{\"secretKey\": \"$API_KEY\"}" \
  https://api.testwyre.com/v2/sessions/auth/key
