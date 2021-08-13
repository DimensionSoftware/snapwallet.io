const {
  BaseAPIError,
  NotFoundError,
  UnauthorizedError,
} = require('../error/index')

const centralizedErrorPlug = async (ctx, next) => {
  try {
    await next()
    if (ctx.response.status === 404) throw new NotFoundError()
  } catch (e) {
    ctx.log.error({
      stack: e.stack,
      error: e,
      meta: e.originalError,
    })
    if (e instanceof BaseAPIError) {
      ctx.status = e.status
      ctx.body = e
    } else {
      const response = new BaseAPIError()
      ctx.status = response.status
      ctx.body = response
    }
  }
}

module.exports = { centralizedErrorPlug }
