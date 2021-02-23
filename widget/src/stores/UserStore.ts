import { writable } from 'svelte/store'

function createStore() {
  const { subscribe, update } = writable({ intent: 'buy' })

  return {
    subscribe,
    setIntent: (intent: string) => update(s => ({ ...s, intent })),
  }
}

export const userStore = createStore()
