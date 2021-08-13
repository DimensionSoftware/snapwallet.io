const { BaseAPIError, UnprocessableEntityError } = require('.')

class WyreError extends BaseAPIError {
  constructor(e) {
    super()
    this.originalError = e.response ? e.response.data : e
    this.mapError()
  }

  mapError = () => {
    switch (this.originalError.type) {
      case 'InsufficientFundsException': {
        const e = new UnprocessableEntityError()
        this.message = 'There are insufficient funds for this transaction.'
        this.status = e.status
        this.error = e.error
      }
    }
  }
}

module.exports = { WyreError }
