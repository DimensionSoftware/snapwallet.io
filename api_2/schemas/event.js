const EventSchema = Joi.object({
  id: Joi.string().default(() => v4()),
  kind: Joi.string().required(),
  action: Joi.string(),
  data: Joi.any().required(),
  entity: {
    kind: Joi.string().valid('business', 'user').required(),
    id: Joi.string().required(),
  },
  recorded_at: Joi.date().iso().default(new Date().toISOString()),
})

module.exports = { EventSchema }
