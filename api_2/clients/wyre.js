const axios = require('axios').default
const { createHmac } = require('crypto')
const { WYRE_API_URL, WYRE_API_SECRET, WYRE_API_KEY, WYRE_WEBHOOK_URL } =
  process.env
const ClientBase = require('./base')
const { v4 } = require('uuid')
const { UnprocessableEntityError } = require('../error')

class Wyre extends ClientBase {
  client = axios.create({ baseURL: WYRE_API_URL })

  getWallet = (params = {}) => {
    return this.get({ url: `/v2/wallet/${params.id}` })
  }

  listWallets = (params = {}) => {
    const { limit = 20, offset = 0 } = params
    return this.get({ url: '/v2/wallets', params: { limit, offset } })
  }

  createWallet = (params = {}) => {
    const {
      type = 'SAVINGS',
      name = `snap_wallet_${v4()}`,
      notes = '',
    } = params
    return this.post({
      url: '/v2/wallets',
      data: {
        name,
        callbackUrl: `${WYRE_WEBHOOK_URL}/v1/webhooks/wyre?wallet_name=${name}`,
        type,
        notes,
      },
    })
  }

  createTransfer = (data) => this.post({ url: '/v3/transfers', data })

  _setSignatureHeaders = (opts = {}) => {
    opts.params = { ...opts.params, timestamp: Date.now() }
    const qs = new URLSearchParams(opts.params).toString()
    const uri = `${WYRE_API_URL}${opts.url}${qs ? '?' + qs : ''}`
    const body = opts.data ? JSON.stringify(opts.data) : ''

    const signature = createHmac('sha256', WYRE_API_SECRET)
      .update(uri + body)
      .digest('hex')

    this.client.defaults.headers.common = {
      ...this.client.defaults.headers.common,
      'X-Api-Key': WYRE_API_KEY,
      'X-Api-Signature': signature,
    }
  }

  errorSwitch = (e) => {
    if (e.response.data) {
      const klass = e.response.data
      if (klass.type === 'InsufficientFundsException')
        return new UnprocessableEntityError(
          'You have insufficient funds for this transaction'
        )
    }
    // TODO: handle Wyre errors here
    return e
  }
}

module.exports = Wyre
