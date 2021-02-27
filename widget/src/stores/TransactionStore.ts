import { writable } from 'svelte/store'
import { IAsset, TransactionMediums } from '../types'

const createStore = () => {
  const { subscribe, update } = writable({
    inMedium: TransactionMediums.ACH,
    outMedium: TransactionMediums.BLOCKCHAIN,
    sourceId: null,
    sourceCurrency: { name: 'USD', ticker: 'USD' },
    destinationCurrency: { name: 'Bitcoin', ticker: 'BTC' },
    sourceAmount: 0.0,
  })

  return {
    subscribe,
    setSourceAmount: (sourceAmount: number) => {
      update(s => ({ ...s, sourceAmount }))
    },
    setSourceCurrency: (sourceCurrency: IAsset) =>
      update(s => ({ ...s, sourceCurrency })),
    setDestinationCurrency: (destinationCurrency: IAsset) =>
      update(s => ({ ...s, destinationCurrency })),
  }
}

export const transactionStore = createStore()
