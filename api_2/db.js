// Required for using module directly
require('dotenv').config()

const admin = require('firebase-admin')

admin.initializeApp({
  projectId: process.env.FIRESTORE_PROJECT,
})

const collections = {
  users: 'users',
  businesses: 'businesses',
  events: 'events',
}

const db = admin.firestore()

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
  const ref = await db.collection(collections.events).add({
    type,
    meta,
  })
  const doc = await ref.get()
  return { ...doc.data(), id: ref.id }
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
