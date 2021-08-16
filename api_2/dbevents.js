// i think this might cause a cycle?
//const { collections: dbCollections } = require('./db')

// TypeError: Cannot read property 'events' of undefined
//     at new DatabaseEventsManager (/Users/dreamcodez/sd/flux/api_2/dbevents.js:8:56)
//     at Object.<anonymous> (/Users/dreamcodez/sd/flux/api_2/db.js:20:16)
//     at Module._compile (internal/modules/cjs/loader.js:1063:30)
//     at Object.Module._extensions..js (internal/modules/cjs/loader.js:1092:10)
//     at Module.load (internal/modules/cjs/loader.js:928:32)
//     at Function.Module._load (internal/modules/cjs/loader.js:769:14)
//     at Module.require (internal/modules/cjs/loader.js:952:19)
//     at require (internal/modules/cjs/helpers.js:88:18)
//     at Object.<anonymous> (/Users/dreamcodez/sd/flux/api_2/routes/v1/transfer.js:3:46)
//     at Module._compile (internal/modules/cjs/loader.js:1063:30)

const { EventSchema } = require('./schemas/event')

class DatabaseEventsManager {
  constructor(db) {
    this.db = db
    this.collection = this.db.collection('events')
  }

  // record n number of events
  async record(...rawEvents) {
    if (rawEvents.length === 0) return []

    // handle first argument as array also
    if (rawEvents.length === 1 && Array.isArray(rawEvents[0]))
      rawEvents = rawEvents[0]

    const events = [],
      batch = this.db.batch()

    for (const rawEvent of rawEvents) {
      const { error: validationError, value: event } =
        EventSchema.validate(rawEvent)

      if (validationError) {
        throw validationError
      }

      const ref = this.collection.doc(event.id)
      batch.set(ref, event)

      events.push(event)
    }

    await batch.commit()

    return events
  }

  // get by id
  async get(id) {
    const ref = db.collection('events').doc(id)

    const doc = await ref.get()
    if (doc.exists) {
      return doc.data()
    } else {
      return null
    }
  }

  // TODO:

  // scan by date
  async scan(from, to) {
  }
}


module.exports = { DatabaseEventsManager }
