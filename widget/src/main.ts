import App from './App.svelte'

import { FluxApi, createConfiguration, ServerConfiguration } from 'api-client'

async function testyMcAPI(): Promise<void> {
  const api = new FluxApi(
    createConfiguration({
      baseServer: new ServerConfiguration('http://localhost:5100', {}),
    }),
  )

  console.log(await api.fluxPricingData())
}
testyMcAPI()

const queryParams = new URLSearchParams(window.location.search)

const app = new App({
  target: document.body,
  props: {
    apiKey: queryParams.get('apiKey') || '',
    appName: queryParams.get('appName') || 'Flux',
    intent: queryParams.get('intent') || 'buy',
    theme: {
      /* NOTE: each attribute maps to a css variable which is prefixed by --theme- */
      // color: '#7f8c8d',
      // colorDarkened: '#2c3e50',
      // colorLightened: '#95a5a6',
    },
  },
})

export default app
