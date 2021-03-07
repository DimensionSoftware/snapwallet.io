import App from './App.svelte'
import { FluxApi, createConfiguration, ServerConfiguration } from 'api-client'
import { getFluxSession, setFluxSession } from './util'

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

// TODO: move this to typings/@flux.d.ts
declare global {
  interface Window {
    __api: FluxApi
    Plaid: any
    API: (newToken?: string) => FluxApi
  }
}

function genAPIClient(token?: string): FluxApi {
  return new FluxApi(
    createConfiguration({
      baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
      authMethods: token
        ? {
            Bearer: `Bearer ${token}`,
          }
        : null,
    }),
  )
}

function getAPIClient(newToken?: string): FluxApi {
  if (newToken) {
    setFluxSession(newToken)
  }

  if (newToken || !window.__api) {
    window.__api = genAPIClient(newToken)
  }

  // Remove token when invalid
  console.log('S', getFluxSession())
  if (getFluxSession() !== '') {
    window.__api
      .fluxViewerData()
      .then(resp => {
        console.log('resp', resp)
      })
      .catch(e => {
        setFluxSession('')
      })
  }

  return window.__api
}

window.API = getAPIClient

export default app
