const KoaRouter = require('koa-router')
const Wyre = require('../../clients/wyre')
const { getEvent } = require('../../db')
const { verifyWyreWebhookHmac } = require('../../middleware/auth')
const router = new KoaRouter()

router.all(
  '/wyre',
  // TODO: use verifyWyreWebhookHmac
  async (ctx, _next) => {
    ctx.body = {}

    const wyre = new Wyre()
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

    const { type, meta } = event

    if (type !== 'transaction') {
      ctx.log.info({
        msg: 'Event is unrelated to transactions',
      })
      ctx.status = 200
      return
    }

    // TODO: figure out what fee should be and take from environment.
    const swAmount = 0.0025 * meta.sourceAmount
    const remainingAmount = meta.sourceAmount - swAmount
    const baseParams = {
      preview: false,
      autoConfirm: true,
      sourceCurrency: meta.sourceCurrency,
      destCurrency: meta.destCurrency,
    }

    await Promise.all([
      wyre.createTransfer({
        ...baseParams,
        source: `wallet:${meta.source}`,
        sourceAmount: swAmount,
        dest: `wallet:${process.env.SNAP_WALLET_WYRE_SAVINGS_WALLET}`,
      }),
      wyre.createTransfer({
        ...baseParams,
        source: `wallet:${meta.source}`,
        sourceAmount: remainingAmount,
        dest: `wallet:${meta.destination}`,
      }),
    ])

    ctx.status = 200
    ctx.body = {}
  }
)

module.exports = router
