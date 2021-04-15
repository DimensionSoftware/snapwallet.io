import type { WyreTransfer, WyreTransfers } from 'api-client'
import { writable } from 'svelte/store'
import { reducePersonalInfoFields } from '../util/profiles'

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
  const { transfers } = await window.API.fluxWyreGetTransfers(
    startingAtPage.toString(),
  )

  if (transfers.length === pagesize) {
    return transfers.concat(await rFetchUserTransfers(startingAtPage + 1))
  }

  return transfers
}

export const transactionsStore = createStore()
