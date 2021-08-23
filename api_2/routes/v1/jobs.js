const Router = require('koa-router')
const router = new Router()
const util = require('util')
const { BadRequestError } = require('../../error')

router.post('/run', async (ctx, _next) => {
  if (!ctx.request.body) {
    const msg = 'No Pub/Sub message received'
    ctx.log.error({ msg })
    throw new BadRequestError(msg)
  }

  if (!ctx.request.body.message) {
    const msg = 'invalid Pub/Sub message format'
    ctx.log.error({ msg })
    throw new BadRequestError(msg)
  }

  const pubSubMessage = ctx.request.body.message
  const job = JSON.parse(Buffer.from(pubSubMessage.data, 'base64'))

  // TODO stuff here

  ctx.log.info({
    msg: `ACKING JOB: ${util.inspect(job, { showHidden: false, depth: null })}`,
  })

  ctx.status = 200
})

module.exports = router
