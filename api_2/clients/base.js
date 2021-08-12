const axios = require('axios').default

class ClientBase {
  client = axios.create()
  errorSwitch = (e) => e

  get = (opts = {}) => {
    return this.request({ ...opts, method: 'GET' })
  }

  post = (opts = {}) => {
    return this.request({ ...opts, method: 'POST' })
  }

  put = (opts = {}) => {
    return this.request({ ...opts, method: 'PUT' })
  }

  patch = (opts = {}) => {
    return this.request({ ...opts, method: 'PATCH' })
  }

  delete = (opts = {}) => {
    return this.request({ ...opts, method: 'DELETE' })
  }

  request = async (opts = {}) => {
    try {
      this._setSignatureHeaders(opts)
      return await this.client.request(opts)
    } catch (e) {
      throw this.errorSwitch(e)
    }
  }

  _setSignatureHeaders = (_opts = {}) => {}
}

module.exports = ClientBase
