import type { WyreTransfer } from 'api-client'
import { writable } from 'svelte/store'

const createStore = () => {
  const { subscribe, update } = writable<WyreTransfer[]>([])
  return {
    subscribe,
    fetchUserTransactions: async (): Promise<WyreTransfer[]> => {
      const txns = await rFetchUserTransfers()
      update(_ => txns)
      return txns
    },
  }
}

const pagesize = 30

async function rFetchUserTransfers(
  startingAtPage = 0,
): Promise<WyreTransfer[]> {
  let { transactions } = await window.API.fluxGetTransactions(
    startingAtPage.toString(),
  )

  transactions = transactions.filter(t => t.status.toUpperCase() !== 'QUOTED')

  if (transactions.length === pagesize) {
    return transactions.concat(await rFetchUserTransfers(startingAtPage + 1))
  }

  return transactions
}

export const transactionsStore = createStore()

const createTransactionDetailsStore = () => {
  const { subscribe, set } = writable<{ transaction: WyreTransfer }>({
    transaction: null,
  })
  return {
    subscribe,
    set,
    reset: () => set(null),
  }
}

export const transactionDetailsStore = createTransactionDetailsStore()
