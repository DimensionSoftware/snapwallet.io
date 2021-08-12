const Router = require('koa-router')
const Wyre = require('../../clients/wyre')
const router = new Router()

router.post('/', async (ctx, _next) => {
  const {
    source,
    sourceCurrency,
    sourceAmount,
    destCurrency,
    preview = true,
    dest,
  } = ctx.request.body

  const params = {
    source,
    sourceCurrency,
    sourceAmount,
    dest,
    destCurrency,
    message: '',
    preview,
    autoConfirm: false,
  }

  const wyre = new Wyre()
  const { data } = await wyre.createTransfer(params)
  ctx.body = { data }
})

module.exports = router
