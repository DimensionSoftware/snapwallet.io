// Required for using module directly
require('dotenv').config()

const admin = require('firebase-admin'),
  { DatabaseEventsManager } = require('./dbevents')

admin.initializeApp({
  projectId: process.env.FIRESTORE_PROJECT,
})

const collections = {
  users: 'users',
  businesses: 'businesses',
  events: 'events',
}

const db = admin.firestore()

// test
const events = new DatabaseEventsManager(db)
events
  .record({ kind: 'TEST_KIND', data: 123 })
  .then(console.log.bind(0, 'events recorded'))
events
  .record(
    { kind: 'SPOOKY_KIND', data: [1, 2, 3] },
    { kind: 'GOOFY_KIND', data: 'sick' }
  )
  .then(console.log.bind(0, 'events recorded'))

const listUsers = () =>
  db
    .collection(collections.users)
    .get()
    .then((ss) => {
      if (ss.empty) return []
      return ss.map((d) => d.data())
    })

const createBusiness = async ({ name, apiKey, wallet }) => {
  const ref = await db.collection(collections.businesses).add({
    apiKey,
    name,
    wallet,
  })
  const doc = await ref.get()
  return { ...doc.data(), id: ref.id }
}

const createEvent = async ({ type, meta }) => {
  const [event] = await events.record({
    kind: type,
    data: meta,
  })

  return event
}

const getBusinessByAPIKey = async (apiKey) => {
  const result = await db
    .collection(collections.businesses)
    .where('apiKey', '==', apiKey)
    .limit(1)
    .get()

  const doc = result.docs[0]
  if (!doc) return
  return { ...doc.data(), id: doc.id }
}

const getEvent = async (source) => {
  const result = await db
    .collection(collections.events)
    .where('meta.source', '==', source)
    .limit(1)
    .get()
  if (!result.docs[0]) return
  return result.docs[0].data()
}

module.exports = {
  listUsers,
  createBusiness,
  createEvent,
  getBusinessByAPIKey,
  getEvent,
}
