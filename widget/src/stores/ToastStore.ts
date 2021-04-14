import { writable } from 'svelte/store'

function createToast() {
  const { subscribe, set } = writable(null)
  let _timer = null

  const clearTimer = () => {
      if (_timer) clearTimeout(_timer)
    },
    dismiss = () => {
      // dismiss toast early
      window.dispatchEvent(new Event('unblurryHeader'))
      clearTimer()
      set(null)
      _timer = null
    }

  return {
    subscribe,
    dismiss,
    pop: ({ msg = '', error = false, warning = false, success = false }) => {
      set({ msg, error, warning, success })
      clearTimer()
      window.dispatchEvent(new Event('blurryHeader'))
      _timer = setTimeout(() => {
        dismiss()
      }, 4000)
    },
  }
}

export const toaster = createToast()
