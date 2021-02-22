import { writable } from 'svelte/store'

function createToast() {
  const { subscribe, set } = writable(null)

  return {
    subscribe,
    pop: ({ msg = '', error = false, warning = false, success = false }) => {
      set({ msg, error, warning, success })
      setTimeout(() => {
        set(null)
      }, 4000)
    },
  }
}

export const toaster = createToast()
