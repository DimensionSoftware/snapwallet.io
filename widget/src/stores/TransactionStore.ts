import type { WyrePaymentMethod } from 'api-client'
import { writable } from 'svelte/store'
import { IAsset, TransactionIntents, TransactionMediums } from '../types'

const initialState = {
  intent: TransactionIntents.BUY,
  inMedium: TransactionMediums.ACH,
  outMedium: TransactionMediums.BLOCKCHAIN,
  sourceId: null,
  sourceCurrency: { name: 'USD', ticker: 'USD', color: '#00D395' },
  destinationCurrency: { name: 'Bitcoin', ticker: 'BTC', color: '#F7931A' },
  sourceAmount: 0.0,
  destinationAmount: 0.0,
  selectedSourcePaymentMethod: null,
  wyrePreview: null,
}

const createStore = () => {
  const { subscribe, update, set } = writable(initialState)

  return {
    subscribe,
    reset: () => set(initialState),
    setSelectedSourcePaymentMethod: (
      selectedSourcePaymentMethod: WyrePaymentMethod,
    ) => {
      update(s => ({ ...s, selectedSourcePaymentMethod }))
    },
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
    // Update both currencies at one time in order
    // to keep svelte + currencies in sync
    setCurrencies: (currencies: {
      sourceCurrency: IAsset
      destinationCurrency: IAsset
    }) => {
      update(s => ({ ...s, ...currencies }))
    },
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
    setWyrePreview: (wyrePreview: object) =>
      update(s => ({ ...s, wyrePreview })),
  }
}

export const transactionStore = createStore()
