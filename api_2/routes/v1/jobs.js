const Router = require('koa-router')
const router = new Router()
const util = require('util')
const { BadRequestError } = require('../../error')
const { JobRunSchema } = require('../../schemas/jobs')

router.post('/run', async (ctx, _next) => {
  const vld8n = JobRunSchema.validate(ctx.request.body)
  if (vld8n.error) {
    const msg = 'Invalid Pub/Sub message format'
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
