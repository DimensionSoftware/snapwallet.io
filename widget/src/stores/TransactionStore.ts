import { writable } from 'svelte/store'
import { IAsset, TransactionMediums } from '../types'

const createStore = () => {
  const { subscribe, update } = writable({
    inMedium: TransactionMediums.ACH,
    outMedium: TransactionMediums.BLOCKCHAIN,
    sourceId: null,
    sourceCurrency: { name: 'USD', ticker: 'USD' },
    destinationCurrency: { name: 'Bitcoin', ticker: 'BTC' },
  })

  return {
    subscribe,
    setSourceCurrency: (sourceCurrency: IAsset) =>
      update(s => ({ ...s, sourceCurrency })),
    setDestinationCurrency: (destinationCurrency: IAsset) =>
      update(s => ({ ...s, destinationCurrency })),
  }
}

export const transactionStore = createStore()
