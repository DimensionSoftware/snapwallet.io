import { writable } from 'svelte/store'
import type { ProductType, WalletType } from '../types'

type ConfigStoreState = {
  apiKey: string
  appName: string
  sourceAmount: number
  wallets: WalletType[]
  focus: boolean
  intent: 'buy'
  product?: ProductType
  theme: { [cssProp: string]: string }
}

function createStore() {
  const { subscribe, update } = writable<ConfigStoreState>({
    apiKey: '',
    appName: '',
    sourceAmount: 0,
    wallets: [],
    focus: false,
    intent: 'buy',
    product: undefined,
    theme: {},
  })

  return {
    subscribe,
    setInitial: (config: Partial<ConfigStoreState>) =>
      update(s => ({ ...s, ...config })),
  }
}

export const configStore = createStore()
