const Router = require('koa-router'),
  router = new Router(),
  walletsRouter = require('./wallets'),
  transferRouter = require('./transfer')

router.use('/wallets', walletsRouter.routes(), walletsRouter.allowedMethods())

router.use(
  '/transfers',
  transferRouter.routes(),
  transferRouter.allowedMethods()
)

module.exports = router
