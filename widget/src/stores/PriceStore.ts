import { writable } from 'svelte/store'
import { FluxApi, createConfiguration, ServerConfiguration } from 'api-client'
import { supportedCurrencyPairs } from '../constants'

// FIXME: example of how to configure client for bearer, it will automatically send to secure routes while omitting for public routes
const tokenHere = 'not-the-right-thing'
function genAPIClient(token?: string): FluxApi {
  return new FluxApi(
    createConfiguration({
      baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
      authMethods: token
        ? {
            Bearer: `Bearer ${token}`,
          }
        : null,
    }),
  )
}
new FluxApi(
  createConfiguration({
    baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
    authMethods: {
      Bearer: `Bearer ${tokenHere}`,
    },
  }),
)

// for testing, when needed, uncomment :D
;(window as any).api = genAPIClient()

const stubPrices = () => {
  return supportedCurrencyPairs.wyre.reduce((acc, key) => {
    const [source, destination] = key.split('_')
    acc[key] = {
      [source]: 0.0,
      [destination]: 0.0,
    }
    return acc
  }, {})
}

const createStore = () => {
  const store = writable({
    prices: stubPrices(),
  })

  const { subscribe, update } = store

  const fetchPrices = async () => {
    const { rates: prices } = await api.fluxPricingData()
    const result = {}
    await Promise.all(
      supportedCurrencyPairs.wyre.map(s => {
        if (!prices[s]) return
        result[s] = prices[s]['rate']
      }),
    )
    update(s => ({ ...s, prices: result }))
  }

  return {
    subscribe,
    fetchPrices,
    pollPrices: () => {
      const interval = setInterval(async () => {
        try {
          await fetchPrices()
        } catch (e) {
          clearInterval(interval)
        }
      }, 15_000)
    },
  }
}

export const priceStore = createStore()
