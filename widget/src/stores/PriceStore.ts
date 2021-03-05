import { writable } from 'svelte/store'
import { supportedCurrencyPairs } from '../constants'

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
    const { rates: prices } = await window.API().fluxPricingData()
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
      return interval
    },
  }
}

export const priceStore = createStore()
