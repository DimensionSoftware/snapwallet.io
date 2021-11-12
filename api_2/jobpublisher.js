class JobPublisher {
  constructor(dbevents, topic) {
    this.dbevents = dbevents
    this.topic = topic
  }

  async publish(job) {
    const payload = Buffer.from(
      JSON.stringify({ config: job.config, worker: job.worker })
    )
    this.topic.publish(payload)
    // todo: this.events.record({ /* some event related to job starting */ })
  }
}

module.exports = { JobPublisher }
