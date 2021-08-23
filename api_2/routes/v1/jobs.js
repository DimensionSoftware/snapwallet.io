const Router = require('koa-router')
const router = new Router()
const util = require('util')

router.post('/run', async (ctx, _next) => {
  if (!ctx.request.body) {
    const msg = 'no Pub/Sub message received'
    console.error(`error: ${msg}`)
    res.status(400).send(`Bad Request: ${msg}`)
    return;
  }
  if (!ctx.request.body.message) {
    const msg = 'invalid Pub/Sub message format'
    console.error(`error: ${msg}`)
    res.status(400).send(`Bad Request: ${msg}`)
    return
  }

  const pubSubMessage = ctx.request.body.message
  const job = JSON.parse(Buffer.from(pubSubMessage.data, 'base64'))

  // TODO stuff here

  console.log('ACKING JOB:', util.inspect(job, {showHidden: false, depth: null}))

  res.status(200).send()
})

module.exports = router
