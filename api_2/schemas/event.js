const Joi = require('joi'),
  { nanoid } = require('nanoid')

const EventSchema = Joi.object({
  id: Joi.string().default(nanoid),
  kind: Joi.string().required(),
  action: Joi.string(),
  data: Joi.any().required(),
  entity: Joi.object({
    kind: Joi.string().valid('BUSINESS', 'USER').required(),
    id: Joi.string().required(),
  }).required(),
  recorded_at: Joi.date().iso().default(new Date().toISOString()),
})

module.exports = { EventSchema }
