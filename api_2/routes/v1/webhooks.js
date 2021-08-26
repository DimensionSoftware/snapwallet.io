const KoaRouter = require('koa-router')
const { getEvent, JOB_PUBLISHER } = require('../../db')
const { verifyWyreWebhookHmac } = require('../../middleware/auth')
const { payoutTask } = require('../../util/get_paid')
const router = new KoaRouter()

router.all(
  '/wyre',
  // TODO: use verifyWyreWebhookHmac
  async (ctx, _next) => {
    ctx.body = {}

    const { dest } = ctx.request.body
    const source = dest.replace('wallet:', '')
    const event = await getEvent(source)

    // Business wallets will land here too
    // Ignore the hook if there's no event in storage.
    if (!event) {
      ctx.log.info({
        msg: 'No event found. Webhook unrelated to wallet_wallet.',
      })
      ctx.status = 200
      return
    }

    const { kind, data } = event

    if (kind.toLowerCase() !== 'transaction') {
      ctx.log.info({
        msg: 'Event is unrelated to transactions',
      })
      ctx.status = 200
      return
    }

    await payoutTask(data)

    ctx.status = 200
    ctx.body = {}
  }
)

module.exports = router
