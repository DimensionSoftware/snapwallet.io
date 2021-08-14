require('dotenv').config()

const Koa = require('koa')
const koaRouter = require('koa-router')
const app = new Koa()
const router = new koaRouter()
const bodyParser = require('koa-bodyparser')
const v1Router = require('./routes/v1')
const { centralizedErrorPlug } = require('./middleware/error')
const { loggerPlug } = require('./middleware/logging')
const cors = require('@koa/cors')

/**
 * Common Middleware
 */
app.use(loggerPlug)
app.use(centralizedErrorPlug)
app.use(bodyParser())
app.use(cors())

/**
 * Routing
 */
router.use('/v1', v1Router.routes(), v1Router.allowedMethods())
app.use(router.routes(), router.allowedMethods())

app.listen(process.env.PORT || 3000)
