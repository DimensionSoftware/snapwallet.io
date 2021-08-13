class BaseAPIError extends Error {
  status = 500
  error = 'internal_server'
  message = 'Oops, an unexpected error occurred.'
  constructor(msg) {
    super()
    this.message = msg || this.message
    this.upperCaseMessage()
  }

  upperCaseMessage = () => {
    if (this.message) {
      this.message =
        this.message.charAt(0).toUpperCase() + this.message.slice(1)
    }
  }

  toJSON = () => {
    return {
      error: this.error,
      description: this.message,
    }
  }
}

class UnauthorizedError extends BaseAPIError {
  status = 403
  error = 'unauthorized'
  message = 'You are not authorized to perform the requested action.'
}

class UnauthenticatedError extends BaseAPIError {
  status = 401
  error = 'unauthenticated'
  constructor(msg) {
    super(msg || 'Please sign in and try again.')
  }
}

class UnprocessableEntityError extends BaseAPIError {
  status = 422
  error = 'validation_error'
  constructor(msg) {
    super(msg || 'One or more provided values could not be understood.')
  }
}

class NotFoundError extends BaseAPIError {
  status = 404
  error = 'not_found'
  constructor(msg) {
    super(msg || 'The requested resource could not be located.')
  }
}

module.exports = {
  BaseAPIError,
  UnauthorizedError,
  UnauthenticatedError,
  UnprocessableEntityError,
  NotFoundError,
}