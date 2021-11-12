const { getPendingTasks } = require('../db')
const functions = require('firebase-functions')
const { payoutTask } = require('../util/get_paid')

const workers = {
  payoutTask,
}

const taskRunner = functions
  .runWith({ memory: '2GB' })
  .pubsub.schedule('*/10 * * * * *')
  .onRun(async (_ctx) => {
    const tasks = await getPendingTasks()
    const jobs = []

    tasks.forEach((snapshot) => {
      const { worker, options } = snapshot.data()

      const job = workers[worker](options)
        .then(() => snapshot.ref.update({ status: 'completed' }))
        .catch((e) =>
          snapshot.ref.update({ status: 'errored', error: JSON.stringify(e) })
        )

      jobs.push(job)
    })

    return await Promise.allSettled(jobs)
  })

module.exports = { taskRunner }
