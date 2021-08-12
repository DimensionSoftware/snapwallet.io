const JwtDecode = require('jsonwebtoken')
const { UnauthenticatedError, UnauthorizedError } = require('../error')

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

module.exports = { verifyJWTPlug }
