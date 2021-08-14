const Joi = require('joi'),
  FlakeID = require('flake-idgen'),
  Base58 = require('base-58')

const EventSchema = Joi.object({
  id: Joi.string().default(newEventID),
  kind: Joi.string().required(),
  action: Joi.string(),
  data: Joi.any().required(),
  entity: Joi.object({
    kind: Joi.string().valid('BUSINESS', 'USER').required(),
    id: Joi.string().required(),
  }).required(),
  recorded_at: Joi.date().iso().default(new Date().toISOString()),
})

// id can be 10 bits (or datacenter/worker can be used) .. in future -- not needed now
const idGenerator = new FlakeID({ id: 0 })

// DO NOT CHANGE THE BASE ENCODING!!! .. will mess up ordering
function newEventID() {
  return Base58.encode(idGenerator.next())
}

module.exports = { EventSchema, newEventID }
