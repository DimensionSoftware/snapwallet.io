import type { WyreTransfer, WyreTransfers } from 'api-client'
import { writable } from 'svelte/store'

const createStore = ()  => {
  const { subscribe, update } = writable<WyreTransfer[]>([])
  return {
    subscribe,
    fetchUserTransactions: async (): Promise<WyreTransfer[]> => {
      const out = []

      let page = 0
      let pagelen = -1
      while(pagelen !== 0) {
        const { transfers } = await window.API.fluxWyreGetTransfers(page.toString())
        for (const xfer of transfers) {
          out.push(xfer)
        }
        pagelen = transfers.length
        page++
      }

      update(_ => out)

      return out
    },
  }
}

export const transactionsStore = createStore()
