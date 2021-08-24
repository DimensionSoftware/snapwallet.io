const JwtDecode = require('jsonwebtoken')
const { UnauthenticatedError, UnauthorizedError } = require('../error')
const { createHmac, timingSafeEqual } = require('crypto')

const verifyJWTPlug = async (ctx, next) => {
  try {
    const authHeader = ctx.request.headers['authorization']
    if (!authHeader) throw new UnauthenticatedError()
    const [_header, jwt] = authHeader.split('Bearer ')
    const decodedPubKey = Buffer.from(process.env.JWT_PUBLIC_KEY, 'base64')
    const { aud, sub } = JwtDecode.verify(jwt, decodedPubKey, {
      algorithms: ['RS256'],
    })
    if (aud !== 'ACCESS') throw new UnauthorizedError()
    ctx.user_id = sub
    await next()
  } catch (e) {
    if (e instanceof JwtDecode.JsonWebTokenError) {
      throw new UnauthorizedError()
    }
    throw e
  }
}

const verifyWyreWebhookHmac = async (ctx, next) => {
  const headerSignature = ctx.request.headers['x-api-signature']
  if (!headerSignature)
    throw new UnauthorizedError('Please provide a valid signature.')
  const payload = JSON.stringify(ctx.request.body)
  const signature = createHmac('sha256', process.env.WYRE_SECRET_KEY)
    .update(Buffer.from(payload))
    .digest('hex')
  if (!timingSafeEqual(Buffer.from(headerSignature), Buffer.from(signature))) {
    throw new UnauthorizedError('Please provide a valid signature')
  }
  await next()
}

module.exports = { verifyJWTPlug, verifyWyreWebhookHmac }
