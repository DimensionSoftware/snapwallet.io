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
  .record({
    entity: { kind: 'BUSINESS', id: 'biz123' },
    kind: 'TEST_KIND',
    data: 123,
  })
  .then(console.log.bind(0, 'events recorded'))
events
  .record(
    {
      entity: { kind: 'USER', id: '123' },
      kind: 'SPOOKY_KIND',
      data: [1, 2, 3],
    },
    { entity: { kind: 'USER', id: '321' }, kind: 'GOOFY_KIND', data: 'sick' }
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

const createEvent = async ({ type, meta, entity }) => {
  const [event] = await events.record({
    kind: type,
    data: meta,
    entity,
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
    .where('data.source', '==', source)
    .limit(1)
    .get()
  if (!result.docs[0]) return
  return result.docs[0].data()
}

const insertTask = async (config = {}) => {
  const { worker, options } = config
  const now = new Date().toISOString()
  return db.collection('tasks').add({
    perform_at: now,
    status: 'pending',
    worker,
    options,
  })
}

const getPendingTasks = async (limit = 25) => {
  const now = new Date().toISOString()
  return db
    .collection('tasks')
    .where('perform_at', '<=', now)
    .where('status', '==', 'pending')
    .limit(limit)
    .get()
}

module.exports = {
  listUsers,
  createBusiness,
  createEvent,
  getBusinessByAPIKey,
  getEvent,
  collections,
  insertTask,
  getPendingTasks,
}
