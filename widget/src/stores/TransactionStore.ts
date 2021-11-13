import type { WyrePaymentMethod } from 'api-client'
import { writable } from 'svelte/store'
import { IAsset, TransactionIntents, TransactionMediums } from '../types'
import { computeTransactionExpiration } from '../util/transactions'

type StoreState = {
  intent: TransactionIntents
  inMedium: TransactionMediums
  outMedium: TransactionMediums
  sourceId: null
  sourceCurrency: IAsset
  destinationCurrency: IAsset
  sourceAmount: number
  destinationAmount: number
  selectedSourcePaymentMethod: WyrePaymentMethod | null
  wyrePreview: {
    sourceAmount: number
    id: string
    dest: string
    sourceCurrency: string
    destAmount: number
    destCurrency: string
    exchangeRate: number
    fees: {
      [ticker: string]: number
    }
  }
  transactionExpirationTimer?: number
  transactionExpirationSeconds: number
}

const initialState: StoreState = {
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
  transactionExpirationTimer: null,
  transactionExpirationSeconds: null,
}

const createStore = () => {
  const { subscribe, update } = writable<StoreState>(initialState)

  const updateTxnSeconds = (expiresAt: string) => {
    const transactionExpirationSeconds = computeTransactionExpiration(expiresAt)
    update(s => {
      if (transactionExpirationSeconds <= 0)
        clearInterval(s.transactionExpirationTimer)
      return { ...s, transactionExpirationSeconds }
    })
  }

  const beginExpirationTimer = expiresAt => {
    const transactionExpirationTimer = setInterval(() => {
      try {
        updateTxnSeconds(expiresAt)
      } catch {
        clearInterval(transactionExpirationTimer)
        update(s => ({ ...s, transactionExpirationSeconds: null }))
      }
    }, 1000)
    return transactionExpirationTimer
  }

  return {
    subscribe,
    reset: () =>
      update(s => {
        clearInterval(s.transactionExpirationTimer)
        return initialState
      }),
    update: (fields: Partial<StoreState>) => update(s => ({ ...s, ...fields })),
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
    setWyrePreview: (wyrePreview: object & { expiresAt: string }) => {
      update(s => {
        // Clear existing expiration timer
        clearInterval(s.transactionExpirationTimer)
        const transactionExpirationSeconds = computeTransactionExpiration(
          wyrePreview.expiresAt,
        )
        const transactionExpirationTimer = beginExpirationTimer(
          wyrePreview.expiresAt,
        )
        return {
          ...s,
          wyrePreview,
          transactionExpirationTimer,
          transactionExpirationSeconds,
        }
      })
    },
  }
}

export const transactionStore = createStore()
