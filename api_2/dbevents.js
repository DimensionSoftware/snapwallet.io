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
    const ref = this.db.collection('events').doc(id)

    const doc = await ref.get()
    if (doc.exists) {
      return doc.data()
    } else {
      return null
    }
  }

  // list by entity id / no guarantee of ordering to avoid complex index ; can be done by app server
  async listByEntityID(entityID) {
    const ref = this.db.collection('events')

    const snapshot = await ref.where('entity.id', '==', entityID).get()

    const events = []
    snapshot.forEach((doc) => {
      events.push(doc.data())
    })

    return events
  }

  async listEventsBySource(source) {
    const ref = this.db.collection('events')

    const snapshot = await ref.where('source', '==', source).get()

    const events = []
    snapshot.forEach((doc) => {
      events.push(doc.data())
    })

    return events
  }
}

module.exports = { DatabaseEventsManager }
