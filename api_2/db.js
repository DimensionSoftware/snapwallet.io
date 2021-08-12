// Required for using module directly
require('dotenv').config()

const admin = require('firebase-admin')

admin.initializeApp({
  projectId: process.env.FIRESTORE_PROJECT,
})

const collections = {
  users: 'users',
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

module.exports = { listUsers }
