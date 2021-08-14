const { EventSchema } = require('./schemas/event')

const collectionName = 'events'

class DatabaseEventsManager {
  constructor(db) {
    this.db = db
    this.collection = this.db.collection(collectionName)
  }

  async record(...rawEvents) {
    if (!rawEvents.length) return []

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
}

module.exports = { DatabaseEventsManager }
