const logger = require('koa-pino-logger')
const { v4: uuid } = require('uuid')

const loggerPlug = logger({
  redact: ['req.headers.authorization'],
  genReqId: uuid,
  prettyPrint: process.env.NODE_ENV === 'local',
  level: process.env.LOG_LEVEL || 'info',
})

module.exports = { loggerPlug }
