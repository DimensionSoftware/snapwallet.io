import { writable } from 'svelte/store'
import { IAsset, TransactionIntents, TransactionMediums } from '../types'

const createStore = () => {
  const { subscribe, update } = writable({
    intent: TransactionIntents.BUY,
    inMedium: TransactionMediums.ACH,
    outMedium: TransactionMediums.BLOCKCHAIN,
    sourceId: null,
    sourceCurrency: { name: 'USD', ticker: 'USD' },
    destinationCurrency: { name: 'Bitcoin', ticker: 'BTC' },
    sourceAmount: 0.0,
    destinationAmount: 0.0,
  })

  return {
    subscribe,
    setDestinationAmount: (destinationAmount: number) => {
      update(s => ({ ...s, destinationAmount }))
    },
    setSourceAmount: (sourceAmount: number, destinationPrice: number) => {
      update(s => {
        return {
          ...s,
          sourceAmount,
          // Calculate destination amount for future intent switches
          destinationAmount: destinationPrice * sourceAmount,
        }
      })
    },
    setSourceCurrency: (sourceCurrency: IAsset) =>
      update(s => ({ ...s, sourceCurrency })),
    setDestinationCurrency: (destinationCurrency: IAsset) =>
      update(s => ({ ...s, destinationCurrency })),
    toggleIntent: () =>
      update(s => {
        return {
          ...s,
          intent:
            s.intent === TransactionIntents.BUY
              ? TransactionIntents.SELL
              : TransactionIntents.BUY,
          sourceCurrency: s.destinationCurrency,
          destinationCurrency: s.sourceCurrency,
          sourceAmount: s.destinationAmount,
          destinationAmount: s.sourceAmount,
        }
      }),
  }
}

export const transactionStore = createStore()
