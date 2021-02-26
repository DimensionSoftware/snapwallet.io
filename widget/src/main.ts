import App from './App.svelte'

import {FluxApi, createConfiguration, ServerConfiguration} from 'api-client'


async function testyMcAPI(): Promise<void> {
  const api = new FluxApi(createConfiguration({
    baseServer:new ServerConfiguration("http://localhost:5100", {}) ,
  }))

  console.table(await api.fluxPricingData())
}
testyMcAPI()

const queryParams = new URLSearchParams(window.location.search)

const app = new App({
  target: document.body,
  props: {
    apiKey: queryParams.get('apiKey') || '',
    appName: queryParams.get('appName') || 'Flux',
    intent: queryParams.get('intent') || 'buy',
  },
})

export default app
