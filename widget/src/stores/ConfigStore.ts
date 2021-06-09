import { writable } from 'svelte/store'
import type { UserIntent, SrcDst, ProductType, WalletType } from '../types'
import { WidgetEnvironments } from '../types'

type ConfigStoreState = {
  apiKey: string
  appName: string
  environment: WidgetEnvironments
  sourceAmount: number
  defaultDestinationAsset?: string
  displayAmount?: SrcDst
  wallets: WalletType[]
  focus: boolean
  intent: UserIntent
  payee: string
  product?: ProductType
  theme: { [cssProp: string]: string }
}

function createStore() {
  const { subscribe, update } = writable<ConfigStoreState>({
    apiKey: '',
    appName: '',
    environment: WidgetEnvironments.PRODUCTION,
    sourceAmount: 0,
    defaultDestinationAsset: undefined,
    displayAmount: 'destination',
    wallets: [],
    focus: false,
    intent: 'buy',
    payee: '',
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
