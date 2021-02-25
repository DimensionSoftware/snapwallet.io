import { writable } from 'svelte/store'
import {FluxApi, createConfiguration, ServerConfiguration} from 'api-client'

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

async function testyMcAPI(): Promise<any> {
  const api = new FluxApi(createConfiguration({
    baseServer:new ServerConfiguration("http://localhost:5100", {}) ,
  }))

  console.table(await api.fluxPricingData())

}