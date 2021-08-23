const Joi = require('joi')

const JobRunSchema = Joi.object({
  message: Joi.object({
    data: Joi.string().base64().required(),
  }).required(),
})

module.exports = { JobRunSchema }
