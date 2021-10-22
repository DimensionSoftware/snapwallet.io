const Router = require('koa-router'),
  router = new Router(),
  walletsRouter = require('./wallets'),
  transferRouter = require('./transfer'),
  adminRouter = require('./admin'),
  webhooksRouter = require('./webhooks'),
  jobsRouter = require('./jobs'),
  { verifyJWTPlug } = require('../../middleware/auth')

router.use(
  '/webhooks',
  webhooksRouter.routes(),
  webhooksRouter.allowedMethods()
)

router.use(
  '/transfers',
  transferRouter.routes(),
  transferRouter.allowedMethods()
)

router.use('/jobs', jobsRouter.routes(), jobsRouter.allowedMethods())

/**
 * Authenticate client
 */
router.use(verifyJWTPlug)

// TODO: make sure this route is only accessible by admins
router.use('/admin', adminRouter.routes(), adminRouter.allowedMethods())

router.use('/wallets', walletsRouter.routes(), walletsRouter.allowedMethods())

module.exports = router
