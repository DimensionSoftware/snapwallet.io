const KoaRouter = require('koa-router')
const { v4 } = require('uuid')
const Wyre = require('../../clients/wyre')
const router = new KoaRouter()
const { createBusiness } = require('../../db')

router.post('/business', async (ctx, _next) => {
  const { name } = ctx.request.body
  const wyre = new Wyre()
  const { data: wallet } = await wyre.createWallet()
  const business = await createBusiness({
    name,
    wallet: {
      depositAddresses: wallet.depositAddresses,
      id: wallet.id,
      name: wallet.name,
      processor: 'wyre',
      type: wallet.type,
    },
    apiKey: v4(),
  })
  ctx.body = business
})

module.exports = router
