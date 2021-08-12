const Router = require('koa-router')
const Wyre = require('../../clients/wyre')
const router = new Router()

router.post('/', async (ctx, _next) => {
  const wyre = new Wyre()
  const { data } = await wyre.createWallet({ type: 'SAVINGS' })
  ctx.body = { data }
})

router.get('/:walletId', async (ctx, _next) => {
  const { walletId } = ctx.params
  const wyre = new Wyre()
  const { data } = await wyre.getWallet({ id: walletId })
  ctx.body = { data }
})

module.exports = router
