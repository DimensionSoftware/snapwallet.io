import { writable } from 'svelte/store'

function createToast() {
  const { subscribe, set } = writable(null)
  let _timer = null
  return {
    subscribe,
    pop: ({ msg = '', error = false, warning = false, success = false }) => {
      set({ msg, error, warning, success })
      if (_timer) clearTimeout(_timer)
      _timer = setTimeout(() => {
        set(null)
        _timer = null
      }, 4000)
    },
  }
}

export const toaster = createToast()
