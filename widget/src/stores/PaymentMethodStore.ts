import { writable } from 'svelte/store'

const createStore = () => {
  const store = writable({
    wyrePaymentMethods: [],
  })

  const { subscribe, update, set } = store

  const fetchWyrePaymentMethods = async () => {
    const {
      paymentMethods: wyrePaymentMethods,
    } = await window.API.fluxWyreGetPaymentMethods()
    update(s => ({ ...s, wyrePaymentMethods }))
  }

  return {
    subscribe,
    reset: () => set({ wyrePaymentMethods: [] }),
    fetchWyrePaymentMethods,
  }
}

export const paymentMethodStore = createStore()
