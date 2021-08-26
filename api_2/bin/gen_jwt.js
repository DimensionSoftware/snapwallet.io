#!/usr/local/bin/node

require('dotenv').config()
const JwtSigner = require('jsonwebtoken')

const buffer = Buffer.from(process.env.JWT_PRIVATE_PEM_BASE64, 'base64')

console.log(
  JwtSigner.sign({ aud: 'ACCESS', sub: 'user_123' }, buffer.toString(), {
    algorithm: 'RS256',
  })
)
