import { writable } from 'svelte/store'

const createStore = () => {
  const { subscribe, update } = writable([])
  return {
    subscribe,
    fetchUserTransactions: async () => {
      const { transfers } = await window.API.fluxWyreGetTransfers()
      return update(_ => transfers)
    },
  }
}

export const transactionsStore = createStore()
